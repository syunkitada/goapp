package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetPhysicalResource(tctx *logger.TraceContext, input *spec.GetPhysicalResource) (data *spec.PhysicalResource, err error) {
	data = &spec.PhysicalResource{}
	err = api.DB.Where("name = ? AND datacenter = ? AND deleted_at IS NULL", input.Name, input.Datacenter).
		First(data).Error
	return
}

func (api *Api) GetPhysicalResources(tctx *logger.TraceContext, input *spec.GetPhysicalResources) (data []spec.PhysicalResource, err error) {
	err = api.DB.Where("datacenter = ? AND deleted_at IS NULL", input.Datacenter).Find(&data).Error
	return
}

func (api *Api) CreatePhysicalResources(tctx *logger.TraceContext, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.PhysicalResource
			if err = tx.Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.PhysicalResource{
					Name:          val.Name,
					Kind:          val.Kind,
					Datacenter:    val.Datacenter,
					Cluster:       val.Cluster,
					Rack:          val.Rack,
					PhysicalModel: val.PhysicalModel,
					RackPosition:  val.RackPosition,
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

func (api *Api) UpdatePhysicalResources(tctx *logger.TraceContext, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.PhysicalResource{}).
				Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				Updates(&db_model.PhysicalResource{
					Kind:          val.Kind,
					Cluster:       val.Cluster,
					Rack:          val.Rack,
					PhysicalModel: val.PhysicalModel,
					RackPosition:  val.RackPosition,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeletePhysicalResource(tctx *logger.TraceContext, input *spec.DeletePhysicalResource) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).
			Delete(&db_model.PhysicalResource{}).Error
		return
	})
	return
}

func (api *Api) DeletePhysicalResources(tctx *logger.TraceContext, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			if err = tx.Where("name = ? AND datacenter = ?", data.Name, data.Datacenter).
				Delete(&db_model.PhysicalResource{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
