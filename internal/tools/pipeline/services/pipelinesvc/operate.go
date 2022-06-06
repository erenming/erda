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

package pipelinesvc

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/internal/tools/pipeline/services/apierrors"
)

type OperateAction string

var (
	OpDisableTask OperateAction = "DISABLE-TASK"
	OpEnableTask  OperateAction = "ENABLE-TASK"
	OpPauseTask   OperateAction = "PAUSE-TASK"
	OpUnpauseTask OperateAction = "UNPAUSE-TASK"
)

func (s *PipelineSvc) Operate(pipelineID uint64, req *apistructs.PipelineOperateRequest) error {
	if req == nil || len(req.TaskOperates) <= 0 {
		return nil
	}

	p, err := s.dbClient.GetPipeline(pipelineID)
	if err != nil {
		return apierrors.ErrOperatePipeline.InternalError(err)
	}

	extra := p.PipelineExtra
	extra.Extra.TaskOperates = append(extra.Extra.TaskOperates, req.TaskOperates...)

	return s.dbClient.UpdatePipelineExtraExtraInfoByPipelineID(p.ID, extra.Extra)
}
