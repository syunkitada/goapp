package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/db_model"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateOrUpdateService(tctx *logger.TraceContext, db *gorm.DB, input *base_config.AuthService) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			err = error_utils.NewRecoveredError(tmpErr)
		}
		api.Rollback(tctx, tx, err)
	}()

	var service db_model.Service
	if err = tx.Where("name = ?", input.Name).First(&service).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return
		}

		service = db_model.Service{
			Name:      input.Name,
			Scope:     input.Scope,
			Endpoints: "",
		}
		if err = tx.Create(&service).Error; err != nil {
			return
		}
	} else {
		service.Scope = input.Scope
		service.Endpoints = ""
		if err = tx.Save(&service).Error; err != nil {
			return
		}
	}

	for _, projectRoleName := range input.ProjectRoles {
		var projectRole db_model.ProjectRole
		if err = db.Where("name = ?", projectRoleName).First(&projectRole).Error; err != nil {
			return err
		}

		if err = tx.Model(&projectRole).Association("Services").Append(&service).Error; err != nil {
			return err
		}
	}
	err = tx.Commit().Error
	return
}
