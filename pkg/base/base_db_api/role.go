package base_db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateRole(tctx *logger.TraceContext, db *gorm.DB, name string, projectName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var project base_db_model.Project
	if err = tx.First(&project, "name = ?", projectName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	var role base_db_model.Role
	if err = tx.Where("name = ?", name).First(&role).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		role = base_db_model.Role{
			Name:      name,
			ProjectID: project.ID,
		}
		tx.Create(&role)
		err = tx.Commit().Error
	}
	return err
}

func (api *Api) AssignRoleToUser(tctx *logger.TraceContext, db *gorm.DB, roleName string, userName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var role base_db_model.Role
	if err = tx.Where("name = ?", roleName).First(&role).Error; err != nil {
		return err
	}

	var user base_db_model.User
	if err = tx.Preload("Roles").First(&user, "name = ?", userName).Error; err != nil {
		return err
	}
	if err = tx.Model(&user).Association("Roles").Append(&role).Error; err != nil {
		return err
	}
	err = tx.Commit().Error
	return err
}
