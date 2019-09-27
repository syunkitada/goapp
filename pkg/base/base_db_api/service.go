package base_db_api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) CreateOrUpdateService(tctx *logger.TraceContext, input *base_spec.UpdateService) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var queryMapBytes []byte
	queryMapBytes, err = json.Marshal(&input.QueryMap)

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var service base_db_model.Service
		if err = tx.Where("name = ?", input.Name).First(&service).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			service = base_db_model.Service{
				Name:            input.Name,
				Token:           input.Token,
				Scope:           input.Scope,
				SyncRootCluster: input.SyncRootCluster,
				Endpoints:       strings.Join(input.Endpoints, ","),
				ProjectRoles:    strings.Join(input.ProjectRoles, ","),
				QueryMap:        string(queryMapBytes),
			}

			if err = tx.Create(&service).Error; err != nil {
				return
			}
		} else {
			service.Token = input.Token
			service.Scope = input.Scope
			service.SyncRootCluster = input.SyncRootCluster
			service.Endpoints = strings.Join(input.Endpoints, ",")
			service.ProjectRoles = strings.Join(input.ProjectRoles, ",")
			service.QueryMap = string(queryMapBytes)
			if err = tx.Save(&service).Error; err != nil {
				return
			}
		}

		for _, projectRoleName := range input.ProjectRoles {
			var projectRole base_db_model.ProjectRole
			if err = tx.Where("name = ?", projectRoleName).First(&projectRole).Error; err != nil {
				err = fmt.Errorf("Failed find projectRole: name=%s, err=%v", projectRoleName, err)
				return
			}

			if err = tx.Model(&projectRole).Association("Services").Append(&service).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) GetServices(tctx *logger.TraceContext, input *base_spec.GetServices) (data *base_spec.GetServicesData, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var services []base_spec.Service
	if err = api.DB.Find(&services).Error; err != nil {
		return
	}
	data = &base_spec.GetServicesData{
		Services: services,
	}
	return
}
