package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService, user *base_spec.UserAuthority) (data *spec.RegionService, err error) {
	data = &spec.RegionService{}
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices, user *base_spec.UserAuthority) (data []spec.RegionService, err error) {
	err = api.DB.Where("region = ? AND deleted_at IS NULL", input.Region).Find(&data).Error
	return
}

func (api *Api) CreateRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			var tmp db_model.RegionService
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.RegionService{
					Project:      user.ProjectName,
					Name:         val.Name,
					Region:       val.Region,
					Kind:         val.Kind,
					Status:       resource_model.StatusInitializing,
					StatusReason: "CreateRegionService",
					Spec:         string(specBytes),
				}
				if err = tx.Create(&tmp).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			if err = tx.Model(&db_model.RegionService{}).
				Where("name = ? AND region = ?", val.Name, val.Region).
				Updates(&db_model.RegionService{
					Spec: string(specBytes),
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).
			Delete(&db_model.RegionService{}).Error
		return
	})
	return
}

func (api *Api) DeleteRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				Delete(&db_model.RegionService{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
