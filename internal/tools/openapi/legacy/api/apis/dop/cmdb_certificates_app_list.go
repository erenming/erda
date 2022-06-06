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
	"github.com/erda-project/erda/internal/tools/openapi/legacy/api/apis"
)

var CMDB_CERTIFICATES_APP_LIST = apis.ApiSpec{
	Path:         "/api/certificates/actions/list-application-quotes",
	BackendPath:  "/api/certificates/actions/list-application-quotes",
	Host:         "dop.marathon.l4lb.thisdcos.directory:9527",
	Scheme:       "http",
	Method:       "GET",
	CheckLogin:   true,
	CheckToken:   true,
	RequestType:  apistructs.AppCertificateListRequest{},
	ResponseType: apistructs.PagingAppCertificateDTO{},
	IsOpenAPI:    true,
	Doc:          "summary: 获取指定应用的所有引用证书列表",
}
