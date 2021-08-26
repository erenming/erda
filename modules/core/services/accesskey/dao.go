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

package accesskey

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/erda-project/erda-proto-go/core/services/accesskey/pb"
	"github.com/erda-project/erda/pkg/common/errors"
	"github.com/erda-project/erda/pkg/secret"
)

type Dao interface {
	QueryAccessKey(ctx context.Context, req *pb.QueryAccessKeysRequest) ([]AccessKey, error)
	CreateAccessKey(ctx context.Context, req *pb.CreateAccessKeysRequest) (*AccessKey, error)
	GetAccessKey(ctx context.Context, req *pb.GetAccessKeysRequest) (*AccessKey, error)
	UpdateAccessKey(ctx context.Context, req *pb.UpdateAccessKeysRequest) error
	DeleteAccessKey(ctx context.Context, req *pb.DeleteAccessKeysRequest) error
}

type dao struct {
	db *gorm.DB
}

func (d *dao) QueryAccessKey(ctx context.Context, req *pb.QueryAccessKeysRequest) ([]AccessKey, error) {
	var objs []AccessKey
	q := d.db
	where := make(map[string]interface{})
	if req.Status != pb.StatusEnum_NOT_SPECIFIED {
		where["status"] = req.Status
	} else {
		q = q.Where("status != ?", pb.StatusEnum_DELETED)
	}
	if req.Subject != "" {
		where["subject"] = req.Subject
	}
	if req.SubjectType != pb.SubjectTypeEnum_NOT_SPECIFIED {
		where["subject_type"] = req.SubjectType
	}
	if req.AccessKey != "" {
		where["access_key"] = req.AccessKey
	}

	if res := q.Where(where).Find(&objs); res.Error != nil {
		return nil, res.Error
	}
	return objs, nil
}

func (d *dao) CreateAccessKey(ctx context.Context, req *pb.CreateAccessKeysRequest) (*AccessKey, error) {
	obj := toModel(req)
	q := d.db.Create(&obj)
	if q.Error != nil {
		return nil, q.Error
	}
	return &obj, nil
}

func (d *dao) GetAccessKey(ctx context.Context, req *pb.GetAccessKeysRequest) (*AccessKey, error) {
	var obj AccessKey
	q := d.db.Where(&AccessKey{ID: req.Id}).Find(&obj)
	if q.RecordNotFound() {
		return nil, errors.NewNotFoundError("access-key")
	}
	if q.Error != nil {
		return nil, q.Error
	}
	return &obj, nil
}

func (d *dao) UpdateAccessKey(ctx context.Context, req *pb.UpdateAccessKeysRequest) error {
	q := d.db.Model(&AccessKey{}).Where(&AccessKey{ID: req.Id})
	updated := AccessKey{}
	if req.Status.String() != "" {
		updated.Status = req.Status
	}
	if req.Description != "" {
		updated.Description = req.Description
	}
	q = q.Update(updated)
	return q.Error
}

func (d *dao) DeleteAccessKey(ctx context.Context, req *pb.DeleteAccessKeysRequest) error {
	q := d.db.Model(&AccessKey{}).Where(&AccessKey{ID: req.Id}).Update(&AccessKey{Status: pb.StatusEnum_DELETED})
	return q.Error
}

func toModel(req *pb.CreateAccessKeysRequest) AccessKey {
	pair := secret.CreateAkSkPair()
	return AccessKey{
		AccessKey:   pair.AccessKeyID,
		SecretKey:   pair.SecretKey,
		Status:      pb.StatusEnum_ACTIVATE,
		SubjectType: req.SubjectType,
		Subject:     req.Subject,
		Description: req.Description,
	}
}
