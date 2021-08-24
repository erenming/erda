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

package configcenter

import "github.com/erda-project/erda/modules/openapi/api/apis"

var GET_GROUP_PROPERTIES = apis.ApiSpec{
	Path:        "/api/config/tenants/<tenantID>/groups/<groupID>",
	BackendPath: "/api/msp/config/tenants/<tenantID>/groups/<groupID>",
	Host:        "msp.marathon.l4lb.thisdcos.directory:8080",
	Scheme:      "http",
	Method:      "GET",
	CheckLogin:  true,
	CheckToken:  true,
	Doc:         "summary: 查询配置中心租户下组的具体配置",
}
