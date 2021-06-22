// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package model

import (
	"time"
)

// store secret key pair
type AkSk struct {
	ID          int64     `json:"id" gorm:"primary_key;comment:'Primary Key'"`
	Ak          string    `json:"ak" gorm:"size:24;unique;comment:'Access Key ID'"`
	Sk          string    `json:"sk" gorm:"size:32;unique;comment:'Secret Key'"`
	IsInternal  bool      `json:"internal" gorm:"comment:'identify weather used for internal component communication'"`
	Scope       string    `json:"scope" gorm:"comment:'affect scope. eg: organization, micro_service'"`
	Owner       string    `json:"owner" gorm:"comment:'owner identifier. eg: <orgID>'"`
	Description string    `json:"description" gorm:"comment:'description'"`
	CreatedAt   time.Time `json:"createdAt" gorm:"comment:'created time'"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (ak AkSk) TableName() string {
	return "erda_aksk"
}
