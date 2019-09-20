package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetPhysicalResource(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.PhysicalResource, err error) {
	data = &spec.PhysicalResource{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetPhysicalResources(tctx *logger.TraceContext, db *gorm.DB) (data []spec.PhysicalResource, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreatePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, physicalResources []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, physicalResource := range physicalResources {
			var tmpPhysicalResource db_model.PhysicalResource
			if err = tx.Where("name = ?", physicalResource.Name).First(&tmpPhysicalResource).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpPhysicalResource = db_model.PhysicalResource{
					Name:          physicalResource.Name,
					Kind:          physicalResource.Kind,
					Datacenter:    physicalResource.Datacenter,
					Cluster:       physicalResource.Cluster,
					Rack:          physicalResource.Rack,
					PhysicalModel: physicalResource.PhysicalModel,
					RackPosition:  physicalResource.RackPosition,
				}
				if err = tx.Create(&tmpPhysicalResource).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdatePhysicalResources(tctx *logger.TraceContext, db *gorm.DB, physicalResources []spec.PhysicalResource) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, physicalResource := range physicalResources {
			if err = tx.Model(&db_model.PhysicalResource{}).Where("name = ?", physicalResource.Name).Updates(&db_model.PhysicalResource{
				Kind:          physicalResource.Kind,
				Datacenter:    physicalResource.Datacenter,
				Cluster:       physicalResource.Cluster,
				Rack:          physicalResource.Rack,
				PhysicalModel: physicalResource.PhysicalModel,
				RackPosition:  physicalResource.RackPosition,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeletePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.PhysicalResource{}).Error
		return
	})
	return
}
