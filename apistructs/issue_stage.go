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

package apistructs

type IssueStageRequest struct {
	OrgID     int64        `json:"orgID"`
	IssueType IssueType    `json:"issueType"`
	List      []IssueStage `json:"list"`
	IdentityInfo
}

type IssueStage struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type IssueStageResponse struct {
	Header
	Data []IssueStage `json:"data"`
}
