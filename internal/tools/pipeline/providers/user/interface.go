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

package user

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"

	"github.com/erda-project/erda-proto-go/core/pipeline/base/pb"
)

type Interface interface {
	TryGetUser(ctx context.Context, userID string) *pb.PipelineUser
}

func (p *provider) TryGetUser(ctx context.Context, userID string) *pb.PipelineUser {
	user, err := p.bdl.GetCurrentUser(userID)
	if err != nil {
		p.Log.Warnf("failed to get user info, userID: %s, err: %v", userID, err)
		// return basic user just with id
		return &pb.PipelineUser{ID: structpb.NewStringValue(userID)}
	}
	if user == nil {
		p.Log.Warnf("failed to get user info, userID: %s, err: %v", userID, fmt.Errorf("get empty user info"))
		// return basic user just with id
		return &pb.PipelineUser{ID: structpb.NewStringValue(userID)}
	}
	// return queried user
	return user.ConvertToPipelineUser()
}
