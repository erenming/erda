// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tabExecuteButton

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister/base"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/internal/apps/dop/component-protocol/components/auto-test-scenes/common/gshelper"
	"github.com/erda-project/erda/internal/apps/dop/component-protocol/types"
	autotestv2 "github.com/erda-project/erda/internal/apps/dop/services/autotest_v2"
	protocol "github.com/erda-project/erda/internal/core/openapi/legacy/component-protocol"
)

type ComponentAction struct {
	sdk          *cptype.SDK
	bdl          *bundle.Bundle
	Type         string `json:"type"`
	Props        Props  `json:"props"`
	State        State  `json:"state"`
	gsHelper     *gshelper.GSHelper
	AutoTestSvc  *autotestv2.Service
	ActiveConfig string
}

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "tabExecuteButton",
		func() servicehub.Provider { return &ComponentAction{} })
}

func (a *ComponentAction) marshal(c *cptype.Component) error {
	stateValue, err := json.Marshal(a.State)
	if err != nil {
		return err
	}
	var state map[string]interface{}
	err = json.Unmarshal(stateValue, &state)
	if err != nil {
		return err
	}

	propValue, err := json.Marshal(a.Props)
	if err != nil {
		return err
	}
	var prop map[string]interface{}
	err = json.Unmarshal(propValue, &prop)
	if err != nil {
		return err
	}

	c.Props = prop
	c.State = state
	c.Type = a.Type
	return nil
}

func (a *ComponentAction) unmarshal(c *cptype.Component) error {
	stateValue, err := json.Marshal(c.State)
	if err != nil {
		return err
	}
	var state State
	err = json.Unmarshal(stateValue, &state)
	if err != nil {
		return err
	}

	propValue, err := json.Marshal(c.Props)
	if err != nil {
		return err
	}
	var props Props
	err = json.Unmarshal(propValue, &props)
	if err != nil {
		return err
	}

	a.State = state
	a.Type = c.Type
	a.Props = props
	return nil
}

type State struct {
	Env      string `json:"env"`
	ScenesID uint64 `json:"scenesID"`
}

type ClientMetaData struct {
	Env       string `json:"env"`
	ScenesID  uint64 `json:"scenesID"`
	ConfigEnv string `json:"configEnv"`
}

type Props struct {
	Test  string `json:"text"`
	Type  string `json:"type"`
	Menus []Menu `json:"menu"`
}

type Menu struct {
	Text       string                 `json:"text"`
	Key        string                 `json:"key"`
	Operations map[string]interface{} `json:"operations"`
}

type ClickOperation struct {
	Key    string      `json:"key"`
	Reload bool        `json:"reload"`
	Meta   interface{} `json:"meta"`
}

func (ca *ComponentAction) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	gh := gshelper.NewGSHelper(gs)

	ca.sdk = cputil.SDK(ctx)
	ca.bdl = ctx.Value(types.GlobalCtxKeyBundle).(*bundle.Bundle)
	ca.gsHelper = gshelper.NewGSHelper(gs)

	err := ca.unmarshal(c)
	if err != nil {
		return err
	}

	sceneID := gh.GetFileTreeSceneID()
	ca.ActiveConfig = ca.gsHelper.GetGlobalActiveConfig()
	if ca.ActiveConfig == gshelper.SceneConfigKey && sceneID == 0 {
		return nil
	}
	ca.State.ScenesID = sceneID

	defer func() {
		fail := ca.marshal(c)
		if err == nil && fail != nil {
			err = fail
		}
	}()

	ca.AutoTestSvc = ctx.Value(types.AutoTestPlanService).(*autotestv2.Service)
	switch event.Operation {
	case cptype.InitializeOperation, cptype.RenderingOperation:
		if ca.ActiveConfig == gshelper.SceneConfigKey {
			if err := ca.handleSceneDefault(); err != nil {
				return err
			}
		} else {
			if err := ca.handleSceneSetDefault(); err != nil {
				return err
			}
		}
	case "execute":
		err := ca.handleClick(event, gs)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *ComponentAction) handleSceneSetDefault() error {
	setID := a.gsHelper.GetGlobalSelectedSetID()
	if setID == 0 {
		return nil
	}
	sceneSet, err := a.AutoTestSvc.GetSceneSet(setID)
	if err != nil {
		return err
	}
	spaces, err := a.AutoTestSvc.GetSpace(sceneSet.SpaceID)
	if err != nil {
		return err
	}

	project, err := a.bdl.GetProject(uint64(spaces.ProjectID))
	if err != nil {
		return err
	}
	testClusterName, ok := project.ClusterConfig[string(apistructs.TestWorkspace)]
	if !ok {
		return fmt.Errorf("not found cluster")
	}
	var autoTestGlobalConfigListRequest apistructs.AutoTestGlobalConfigListRequest
	autoTestGlobalConfigListRequest.ScopeID = strconv.Itoa(int(project.ID))
	autoTestGlobalConfigListRequest.Scope = "project-autotest-testcase"
	autoTestGlobalConfigListRequest.UserID = a.sdk.Identity.UserID
	configs, err := a.bdl.ListAutoTestGlobalConfig(autoTestGlobalConfigListRequest)
	if err != nil {
		return err
	}

	a.Props.Menus = []Menu{
		{
			Text: a.sdk.I18n("empty"),
			Key:  "无",
			Operations: map[string]interface{}{
				apistructs.ClickOperation.String(): ClickOperation{
					Key:    "execute",
					Reload: true,
					Meta: ClientMetaData{
						Env:       testClusterName,
						ScenesID:  setID,
						ConfigEnv: "",
					},
				},
			},
		},
	}

	for _, v := range configs {
		a.Props.Menus = append(a.Props.Menus, Menu{
			Text: v.DisplayName,
			Key:  v.Ns,
			Operations: map[string]interface{}{
				apistructs.ClickOperation.String(): ClickOperation{
					Key:    "execute",
					Reload: true,
					Meta: ClientMetaData{
						Env:       testClusterName,
						ScenesID:  setID,
						ConfigEnv: v.Ns,
					},
				},
			},
		})
	}

	a.Props.Type = "primary"
	a.Props.Test = a.sdk.I18n("execute")
	return nil
}

func (a *ComponentAction) handleSceneDefault() error {
	scenesID := a.State.ScenesID
	var req apistructs.AutotestSceneRequest
	req.SceneID = scenesID
	req.UserID = a.sdk.Identity.UserID
	scene, err := a.bdl.GetAutoTestScene(req)
	if err != nil {
		return err
	}

	spaces, err := a.bdl.GetTestSpace(scene.SpaceID)
	if err != nil {
		return err
	}

	project, err := a.bdl.GetProject(uint64(spaces.ProjectID))
	if err != nil {
		return err
	}
	testClusterName, ok := project.ClusterConfig[string(apistructs.TestWorkspace)]
	if !ok {
		return fmt.Errorf("not found cluster")
	}
	var autoTestGlobalConfigListRequest apistructs.AutoTestGlobalConfigListRequest
	autoTestGlobalConfigListRequest.ScopeID = strconv.Itoa(int(project.ID))
	autoTestGlobalConfigListRequest.Scope = "project-autotest-testcase"
	autoTestGlobalConfigListRequest.UserID = a.sdk.Identity.UserID
	configs, err := a.bdl.ListAutoTestGlobalConfig(autoTestGlobalConfigListRequest)
	if err != nil {
		return err
	}

	a.Props.Menus = []Menu{
		{
			Text: a.sdk.I18n("empty"),
			Key:  "无",
			Operations: map[string]interface{}{
				apistructs.ClickOperation.String(): ClickOperation{
					Key:    "execute",
					Reload: true,
					Meta: ClientMetaData{
						Env:       testClusterName,
						ScenesID:  a.State.ScenesID,
						ConfigEnv: "",
					},
				},
			},
		},
	}

	for _, v := range configs {
		a.Props.Menus = append(a.Props.Menus, Menu{
			Text: v.DisplayName,
			Key:  v.Ns,
			Operations: map[string]interface{}{
				apistructs.ClickOperation.String(): ClickOperation{
					Key:    "execute",
					Reload: true,
					Meta: ClientMetaData{
						Env:       testClusterName,
						ScenesID:  a.State.ScenesID,
						ConfigEnv: v.Ns,
					},
				},
			},
		})
	}

	a.Props.Type = "primary"
	a.Props.Test = a.sdk.I18n("execute")
	return nil
}

func (a *ComponentAction) handleClick(event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	gh := gshelper.NewGSHelper(gs)
	metaJson, err := json.Marshal(event.OperationData["meta"])
	if err != nil {
		return err
	}
	var metaData ClientMetaData
	err = json.Unmarshal(metaJson, &metaData)
	if err != nil {
		return err
	}
	gh.SetExecuteButtonActiveKey("fileExecute")
	if a.ActiveConfig == gshelper.SceneConfigKey {
		var req apistructs.AutotestExecuteSceneRequest
		req.AutoTestScene.ID = metaData.ScenesID
		req.ClusterName = metaData.Env
		req.ConfigManageNamespaces = metaData.ConfigEnv
		req.UserID = a.sdk.Identity.UserID
		pipeline, err := a.bdl.ExecuteDiceAutotestScene(req)
		if err != nil {
			(*gs)[protocol.GlobalInnerKeyError.String()] = err.Error()
		} else {
			logrus.Infof("run scene pipeline success scene id: %v, pipelineID %v", req.AutoTestScene.ID, pipeline.Data.ID)
		}
	} else {
		var req apistructs.AutotestExecuteSceneSetRequest
		req.AutoTestSceneSet.ID = metaData.ScenesID
		req.ClusterName = metaData.Env
		req.ConfigManageNamespaces = metaData.ConfigEnv
		req.UserID = a.sdk.Identity.UserID
		pipeline, err := a.AutoTestSvc.ExecuteAutotestSceneSet(req)
		if err != nil {
			(*gs)[protocol.GlobalInnerKeyError.String()] = err.Error()
		} else {
			logrus.Infof("run scene set pipeline success scene set id: %v, pipelineID %v", req.AutoTestSceneSet.ID, pipeline.ID)
		}
	}
	return nil
}
