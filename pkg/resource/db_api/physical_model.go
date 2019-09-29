package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetPhysicalModel(tctx *logger.TraceContext, input *spec.GetPhysicalModel, user *base_spec.UserAuthority) (data *spec.PhysicalModel, err error) {
	data = &spec.PhysicalModel{}
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetPhysicalModels(tctx *logger.TraceContext, input *spec.GetPhysicalModels, user *base_spec.UserAuthority) (data []spec.PhysicalModel, err error) {
	err = api.DB.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreatePhysicalModels(tctx *logger.TraceContext, input []spec.PhysicalModel, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.PhysicalModel
			if err = tx.Where("name = ?", val.Name).First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.PhysicalModel{
					Name:        val.Name,
					Kind:        val.Kind,
					Unit:        val.Unit,
					Description: val.Description,
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

func (api *Api) UpdatePhysicalModels(tctx *logger.TraceContext, input []spec.PhysicalModel, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.PhysicalModel{}).
				Where("name = ?", val.Name).
				Updates(&db_model.PhysicalModel{
					Kind:        val.Kind,
					Unit:        val.Unit,
					Description: val.Description,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeletePhysicalModel(tctx *logger.TraceContext, input *spec.DeletePhysicalModel, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", input.Name).Delete(&db_model.PhysicalModel{}).Error
		return
	})
	return
}

func (api *Api) DeletePhysicalModels(tctx *logger.TraceContext, input []spec.PhysicalModel, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ?", val.Name).
				Delete(&db_model.PhysicalModel{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
