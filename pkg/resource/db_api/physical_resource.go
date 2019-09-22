package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetPhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetPhysicalResource) (data *spec.PhysicalResource, err error) {
	data = &spec.PhysicalResource{}
	err = db.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).First(data).Error
	return
}

func (api *Api) GetPhysicalResources(tctx *logger.TraceContext, db *gorm.DB) (data []spec.PhysicalResource, err error) {
	err = db.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreatePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			var tmp db_model.PhysicalResource
			if err = tx.Where("name = ? AND datacenter = ?", data.Name, data.Datacenter).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.PhysicalResource{
					Name:          data.Name,
					Kind:          data.Kind,
					Datacenter:    data.Datacenter,
					Cluster:       data.Cluster,
					Rack:          data.Rack,
					PhysicalModel: data.PhysicalModel,
					RackPosition:  data.RackPosition,
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

func (api *Api) UpdatePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			if err = tx.Debug().Model(&db_model.PhysicalResource{}).
				Where("name = ? AND datacenter = ?", data.Name, data.Datacenter).
				Updates(&db_model.PhysicalResource{
					Kind:          data.Kind,
					Cluster:       data.Cluster,
					Rack:          data.Rack,
					PhysicalModel: data.PhysicalModel,
					RackPosition:  data.RackPosition,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeletePhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).
			Delete(&db_model.PhysicalResource{}).Error
		return
	})
	return
}

func (api *Api) DeletePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, input []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
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
