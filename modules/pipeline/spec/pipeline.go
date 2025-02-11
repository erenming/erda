// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package spec

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/conf"
	"github.com/erda-project/erda/pkg/strutil"
)

// Pipeline
type Pipeline struct {
	PipelineBase
	PipelineExtra
	Labels map[string]string
}

type PipelineWithStage struct {
	Pipeline
	PipelineStages []*PipelineStageWithTask
}

type PipelineStageWithTask struct {
	PipelineStage
	PipelineTasks []*PipelineTask
}

type PipelineCombosReq struct {
	Branches []string `json:"branches"`
	Sources  []string `json:"sources"`
	YmlNames []string `json:"ymlNames"`
}

type PipelineWithTasks struct {
	Pipeline *Pipeline
	Tasks    []*PipelineTask
}

func (p *PipelineWithTasks) DoneTasks() []string {
	var dones []string
	for _, task := range p.Tasks {
		if task.Status.IsEndStatus() || task.Status == apistructs.PipelineStatusDisabled {
			dones = append(dones, task.Name)
		}
	}
	return dones
}

func (p *Pipeline) GenIdentityInfo() apistructs.IdentityInfo {
	// 默认从快照获取
	if !p.Snapshot.IdentityInfo.Empty() {
		return p.Snapshot.IdentityInfo
	}
	// 老数据从 extra 里获取
	var userID string
	if p.Extra.SubmitUser != nil && p.Extra.SubmitUser.ID != nil {
		userID = fmt.Sprintf("%v", p.Extra.SubmitUser.ID)
	}
	return apistructs.IdentityInfo{
		UserID:         userID,
		InternalClient: p.Extra.InternalClient,
	}
}

func (p *Pipeline) GetSubmitUserID() string {
	var userID string
	if p.Extra.SubmitUser != nil && p.Extra.SubmitUser.ID != nil {
		userID = fmt.Sprintf("%v", p.Extra.SubmitUser.ID)
	}
	return mustUserID(userID)
}

func (p *Pipeline) GetRunUserID() string {
	var userID string
	if p.Extra.RunUser != nil && p.Extra.RunUser.ID != nil {
		userID = fmt.Sprintf("%v", p.Extra.RunUser.ID)
	}
	return mustUserID(userID)
}

func (p *Pipeline) GetCancelUserID() string {
	var userID string
	if p.Extra.CancelUser != nil && p.Extra.CancelUser.ID != nil {
		userID = fmt.Sprintf("%v", p.Extra.CancelUser.ID)
	}
	return mustUserID(userID)
}

func (p *Pipeline) GetLabel(labelKey string) string {
	return p.MergeLabels()[labelKey]
}

func mustUserID(userID string) string {
	if userID != "" {
		return userID
	}
	return conf.InternalUserID()
}

func (p *Pipeline) MergeLabels() map[string]string {
	mergeLabels := make(map[string]string)
	for k, v := range p.NormalLabels {
		mergeLabels[k] = v
	}
	for k, v := range p.Labels {
		mergeLabels[k] = v
	}
	return mergeLabels
}

// GenerateNormalLabelsForCreateV2
// pipeline.createV2 有一些字段通过标签来传递，例如 commit
func (p *Pipeline) GenerateNormalLabelsForCreateV2() map[string]string {
	labels := p.MergeLabels()

	// org
	labels[apistructs.LabelOrgName] = p.GetOrgName()
	// workspace
	labels[apistructs.LabelDiceWorkspace] = p.Extra.DiceWorkspace.String()
	// commit
	labels[apistructs.LabelCommit] = p.GetCommitID()
	commitDetailStr, err := json.Marshal(p.CommitDetail)
	if err == nil && string(commitDetailStr) != "{}" {
		labels[apistructs.LabelCommitDetail] = string(commitDetailStr)
	}
	// userID
	labels[apistructs.LabelUserID] = p.GetSubmitUserID()

	labels[apistructs.LabelPipelineType] = p.Type.String()
	labels[apistructs.LabelPipelineYmlSource] = p.Extra.PipelineYmlSource.String()
	labels[apistructs.LabelPipelineTriggerMode] = p.TriggerMode.String()
	var cronTriggerTimeStr string
	if p.Extra.CronTriggerTime != nil {
		cronTriggerTimeStr = strconv.FormatInt(p.Extra.CronTriggerTime.UnixNano(), 10)
	}
	labels[apistructs.LabelPipelineCronTriggerTime] = cronTriggerTimeStr
	if p.CronID != nil {
		labels[apistructs.LabelPipelineCronID] = strconv.FormatUint(*p.CronID, 10)
	}

	for k, v := range labels {
		if v == "" {
			delete(labels, k)
		}
	}
	return labels
}

// GenerateV1UniquePipelineYmlName 为 v1 pipeline 返回 pipelineYmlName，该 name 在 source 下唯一
// 生成规则: AppID/DiceWorkspace/Branch/PipelineYmlPath
// 1) 100/PROD/master/ec/dws/itm/workflow/item_1d_df_process.workflow
// 2) 200/DEV/feature/dice/pipeline.yml
func (p *Pipeline) GenerateV1UniquePipelineYmlName(originPipelineYmlPath string) string {
	// source != (dice || bigdata) 时无需转换
	if !(p.PipelineSource == apistructs.PipelineSourceDice || p.PipelineSource == apistructs.PipelineSourceBigData) {
		return originPipelineYmlPath
	}
	// 若 originPipelineYmlPath 已经符合生成规则，则直接返回
	ss := strutil.Split(originPipelineYmlPath, "/", true)
	if len(ss) > 3 {
		appID, _ := strconv.ParseUint(ss[0], 10, 64)
		workspace := ss[1]
		branchWithYmlName := strutil.Join(ss[2:], "/", true)
		branchPrefix := strutil.Concat(p.Labels[apistructs.LabelBranch], "/")
		if strconv.FormatUint(appID, 10) == p.Labels[apistructs.LabelAppID] &&
			workspace == p.Extra.DiceWorkspace.String() &&
			strutil.HasPrefixes(branchWithYmlName, branchPrefix) &&
			len(branchWithYmlName) > len(branchPrefix) {
			return originPipelineYmlPath
		}
	}
	return fmt.Sprintf("%s/%s/%s/%s", p.Labels[apistructs.LabelAppID], p.Extra.DiceWorkspace.String(),
		p.Labels[apistructs.LabelBranch], originPipelineYmlPath)
}

// DecodeV1UniquePipelineYmlName 根据 GenerateV1UniquePipelineYmlName 生成规则，反解析得到 originName
func (p *Pipeline) DecodeV1UniquePipelineYmlName(name string) string {
	prefix := fmt.Sprintf("%s/%s/%s/", p.Labels[apistructs.LabelAppID], p.Extra.DiceWorkspace.String(), p.Labels[apistructs.LabelBranch])
	return strutil.TrimPrefixes(name, prefix)
}

func (p *Pipeline) GetConfigManageNamespaces() []string {
	return strutil.DedupSlice(
		append([]string{p.Extra.ConfigManageNamespaceOfSecretsDefault, p.Extra.ConfigManageNamespaceOfSecrets},
			p.Extra.ConfigManageNamespaces...), true)
}

// EnsureGC without nil field
func (p *Pipeline) EnsureGC() {
	var halfHourSec uint64 = 60 * 30
	gc := &p.Extra.GC
	// resource
	if gc.ResourceGC.SuccessTTLSecond == nil {
		gc.ResourceGC.SuccessTTLSecond = &halfHourSec
	}
	if gc.ResourceGC.FailedTTLSecond == nil {
		gc.ResourceGC.FailedTTLSecond = &halfHourSec
	}
	// database
	if gc.DatabaseGC.Analyzed.NeedArchive == nil {
		gc.DatabaseGC.Analyzed.NeedArchive = &[]bool{false}[0]
	}
	if gc.DatabaseGC.Analyzed.TTLSecond == nil {
		gc.DatabaseGC.Analyzed.TTLSecond = &[]uint64{60 * 60 * 24}[0]
	}
	if gc.DatabaseGC.Finished.NeedArchive == nil {
		gc.DatabaseGC.Finished.NeedArchive = &[]bool{true}[0]
	}
	if gc.DatabaseGC.Finished.TTLSecond == nil {
		gc.DatabaseGC.Finished.TTLSecond = &[]uint64{60 * 60 * 24 * 30 * 2}[0]
	}
}

func (p *Pipeline) GetResourceGCTTL() uint64 {
	p.EnsureGC()
	resourceGCTTL := *p.Extra.GC.ResourceGC.FailedTTLSecond
	if p.Status.IsSuccessStatus() {
		resourceGCTTL = *p.Extra.GC.ResourceGC.SuccessTTLSecond
	}
	return resourceGCTTL
}
