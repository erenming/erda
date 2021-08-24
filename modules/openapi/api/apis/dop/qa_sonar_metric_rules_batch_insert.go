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

package dop

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/openapi/api/apis"
)

var QA_SONAR_METRIC_RULES_BATCH_INSERT = apis.ApiSpec{
	Path:        "/api/sonar-metric-rules/actions/batch-insert",
	BackendPath: "/api/sonar-metric-rules/actions/batch-insert",
	Host:        "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:      "http",
	Method:      "POST",
	CheckLogin:  true,
	CheckToken:  true,
	Doc:         "summary: 批量插入 sonar 扫描规则",
	RequestType: apistructs.SonarMetricRulesBatchInsertRequest{},
	IsOpenAPI:   true,
}
