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

package common

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/pkg/http/httpclient"
	"github.com/erda-project/erda/tools/cli/command"
	"github.com/erda-project/erda/tools/cli/format"
)

func GetClusters(ctx *command.Context, orgid string) (apistructs.ClusterListResponse, error) {
	var resp apistructs.ClusterListResponse
	var b bytes.Buffer
	var response *httpclient.Response
	var err error

	if orgid == "" {
		response, err = ctx.Get().Path("/api/clusters").Do().Body(&b)
	} else {
		response, err = ctx.Get().Path("/api/clusters").Param("orgID", orgid).Do().Body(&b)
	}
	if err != nil {
		return apistructs.ClusterListResponse{}, fmt.Errorf(
			format.FormatErrMsg("get clusters", "failed to request ("+err.Error()+")", false))
	}

	if !response.IsOK() {
		return apistructs.ClusterListResponse{}, fmt.Errorf(format.FormatErrMsg("get clusters",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				response.StatusCode(), response.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &resp); err != nil {
		return apistructs.ClusterListResponse{}, fmt.Errorf(format.FormatErrMsg("get clusters",
			fmt.Sprintf("failed to unmarshal build detail response ("+err.Error()+")"), false))
	}

	if !resp.Success {
		return apistructs.ClusterListResponse{}, fmt.Errorf(format.FormatErrMsg("get clusters",
			fmt.Sprintf("failed to request, error code: %s, error message: %s",
				resp.Error.Code, resp.Error.Msg), false))
	}

	return resp, nil
}

func GetClusterDetail(ctx *command.Context, cluster string) (apistructs.GetClusterResponse, error) {
	var resp apistructs.GetClusterResponse
	var b bytes.Buffer

	if cluster == "" {
		return apistructs.GetClusterResponse{}, fmt.Errorf(
			format.FormatErrMsg("get clusters", "missing required arg cluster", false))
	}

	response, err := ctx.Get().Path("/api/clusters/" + cluster).Do().Body(&b)
	if err != nil {
		return apistructs.GetClusterResponse{}, fmt.Errorf(format.FormatErrMsg(
			"get cluster detail", "failed to request ("+err.Error()+")", false))
	}

	if !response.IsOK() {
		return apistructs.GetClusterResponse{}, fmt.Errorf(format.FormatErrMsg("get cluster detail",
			fmt.Sprintf("failed to request, status-code: %d, content-type: %s, raw bod: %s",
				response.StatusCode(), response.ResponseHeader("Content-Type"), b.String()), false))
	}

	if err := json.Unmarshal(b.Bytes(), &resp); err != nil {
		return apistructs.GetClusterResponse{}, fmt.Errorf(format.FormatErrMsg("get cluster detail",
			fmt.Sprintf("failed to unmarshal build detail response ("+err.Error()+")"), false))
	}

	if !resp.Success {
		return apistructs.GetClusterResponse{}, fmt.Errorf(format.FormatErrMsg("get cluster detail",
			fmt.Sprintf("failed to request, error code: %s, error message: %s",
				resp.Error.Code, resp.Error.Msg), false))
	}

	return resp, nil
}
