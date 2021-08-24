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

package testplan

import (
	"github.com/erda-project/erda/apistructs"
	dao2 "github.com/erda-project/erda/modules/dop/dao"
)

func (t *TestPlan) ConvertMember(dbMem dao2.TestPlanMember) apistructs.TestPlanMember {
	return apistructs.TestPlanMember{
		ID:         uint64(dbMem.ID),
		TestPlanID: dbMem.TestPlanID,
		Role:       dbMem.Role,
		UserID:     dbMem.UserID,
		CreatedAt:  dbMem.CreatedAt,
		UpdatedAt:  dbMem.UpdatedAt,
	}
}

func (t *TestPlan) BatchConvertMembers(dbMems []dao2.TestPlanMember) []apistructs.TestPlanMember {
	var results []apistructs.TestPlanMember
	for _, dbMem := range dbMems {
		results = append(results, t.ConvertMember(dbMem))
	}
	return results
}
