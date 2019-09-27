package base_db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateRole(tctx *logger.TraceContext, name string, projectName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var project base_db_model.Project
		if err = tx.First(&project, "name = ?", projectName).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
		}

		var role base_db_model.Role
		if err = tx.Where("name = ?", name).First(&role).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			role = base_db_model.Role{
				Name:      name,
				ProjectID: project.ID,
			}
			err = tx.Create(&role).Error
		}
		return
	})
	return
}

func (api *Api) AssignRoleToUser(tctx *logger.TraceContext, roleName string, userName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var role base_db_model.Role
		if err = tx.Where("name = ?", roleName).First(&role).Error; err != nil {
			return
		}

		var user base_db_model.User
		if err = tx.Preload("Roles").First(&user, "name = ?", userName).Error; err != nil {
			return
		}

		err = tx.Model(&user).Association("Roles").Append(&role).Error
		return
	})
	return
}
