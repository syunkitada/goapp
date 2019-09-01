package base_db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateProject(tctx *logger.TraceContext, db *gorm.DB, name string, projectRoleName string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var projectRole base_db_model.ProjectRole
	if err = tx.First(&projectRole, "name = ?", projectRoleName).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}
	}

	var project base_db_model.Project
	if err = tx.Where("name = ?", name).First(&project).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		project = base_db_model.Project{
			Name:          name,
			ProjectRoleID: projectRole.ID,
		}
		tx.Create(&project)
		err = tx.Commit().Error
	}
	return err
}

func (api *Api) CreateProjectRole(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var projectRole base_db_model.ProjectRole
	if err = db.Where("name = ?", name).First(&projectRole).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return err
		}

		projectRole = base_db_model.ProjectRole{
			Name: name,
		}
		tx.Create(&projectRole)
		err = tx.Commit().Error
	}
	return err
}
