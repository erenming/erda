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

package dbclient

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/erda-project/erda/apistructs"
)

// ListDeploymentsByOrgID list deployment by org id
func (c *DBClient) ListDeploymentsByOrgID(param *apistructs.OrgRunningTasksListRequest,
	orgID uint64) (int64, *[]Deployments, error) {
	var (
		total       int64
		deployments []Deployments
		endedAt     time.Time
		err         error
	)

	db := c.DB.Where("org_id = ?", orgID)
	if param.Cluster != "" {
		db = db.Where("cluster_name = ?", param.Cluster)
	}
	if param.ProjectName != "" {
		db = db.Where("project_name = ?", param.ProjectName)
	}
	if param.AppName != "" {
		db = db.Where("application_name = ?", param.AppName)
	}
	if param.PipelineID != 0 {
		db = db.Where("pipeline_id = ?", param.PipelineID)
	}
	if param.Status != "" {
		db = db.Where("status = ?", param.Status)
	}
	if param.UserID != "" {
		db = db.Where("user_id = ?", param.UserID)
	}
	if param.Env != "" {
		db = db.Where("env = ?", param.Env)
	}

	if param.EndTime == 0 {
		endedAt = time.Now()
	} else {
		endedAt = time.Unix(0, param.EndTime*1000000)
	}
	db = db.Where("created_at < ?", endedAt)
	if param.StartTime != 0 {
		startedAt := time.Unix(0, param.StartTime*1000000)
		db = db.Where("created_at > ?", startedAt)
	}

	if err = db.Order("updated_at DESC").
		Offset((param.PageNo - 1) * param.PageSize).
		Limit(param.PageSize).Find(&deployments).Error; err != nil {
		return 0, nil, err
	}
	// 符合条件的 deployment 总量
	if err = db.Model(&Deployments{}).Count(&total).Error; err != nil {
		return 0, nil, err
	}

	return total, &deployments, nil
}

// DeleteDeployment 删除 deployment 信息
func (c *DBClient) DeleteDeployment(orgID string, taskID uint64) error {
	return c.Where("org_id = ?", orgID).Where("task_id = ?", taskID).
		Delete(&Deployments{}).Error
}

// GetDeployment get deployment info
func (c *DBClient) GetDeployment(orgID string, taskID uint64) []Deployments {
	var deployments []Deployments
	if err := c.Where("org_id = ?", orgID).
		Where("task_id = ?", taskID).
		Find(&deployments).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nil
	}
	return deployments
}

// UpdateDeploymentStatus update deployment status
func (c *DBClient) UpdateDeploymentStatus(deployment *Deployments) error {
	return c.Save(deployment).Error
}

// CreateDeployment create running deployment
func (c *DBClient) CreateDeployment(deployment *Deployments) error {
	return c.Create(deployment).Error
}

// ListExpiredDeployments list expire deployment from specific time
func (c *DBClient) ListExpiredDeployments(startTime string) []Deployments {
	var deployments []Deployments
	if err := c.Where("created_at < ?", startTime).
		Find(&deployments).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nil
	}
	return deployments
}

// ListRunningDeployments list running deployments
func (c *DBClient) ListRunningDeployments() []Deployments {
	var deployments []Deployments
	if err := c.Where("status not in (?)",
		[]string{"Success", "AnalyzeFailed", "Timeout", "StopByUser", "Failed", "NoNeedBySystem"}).
		Find(&deployments).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nil
	}
	return deployments
}
