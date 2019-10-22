package base_db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateProject(tctx *logger.TraceContext, name string, projectRoleName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var projectRole base_db_model.ProjectRole
		if err = tx.First(&projectRole, "name = ?", projectRoleName).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
		}

		var project base_db_model.Project
		if err = tx.Where("name = ?", name).First(&project).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			project = base_db_model.Project{
				Name:          name,
				ProjectRoleID: projectRole.ID,
			}
			err = tx.Create(&project).Error
		}
		return
	})
	return
}

func (api *Api) CreateProjectRole(tctx *logger.TraceContext, name string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var projectRole base_db_model.ProjectRole
		if err = tx.Where("name = ?", name).First(&projectRole).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			projectRole = base_db_model.ProjectRole{
				Name: name,
			}
			err = tx.Create(&projectRole).Error
		}
		return
	})
	return
}
