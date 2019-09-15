package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetPhysicalModel(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.PhysicalModel, err error) {
	data = &spec.PhysicalModel{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetPhysicalModels(tctx *logger.TraceContext, db *gorm.DB) (data []spec.PhysicalModel, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreatePhysicalModels(tctx *logger.TraceContext, db *gorm.DB, regions []spec.PhysicalModel) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			var tmpPhysicalModel db_model.PhysicalModel
			if err = tx.Where("name = ?", region.Name).First(&tmpPhysicalModel).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpPhysicalModel = db_model.PhysicalModel{
					Name: region.Name,
					Kind: region.Kind,
				}
				if err = tx.Create(&tmpPhysicalModel).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdatePhysicalModels(tctx *logger.TraceContext, db *gorm.DB, regions []spec.PhysicalModel) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			if err = tx.Model(&db_model.PhysicalModel{}).Where("name = ?", region.Name).Updates(&db_model.PhysicalModel{
				Kind: region.Kind,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeletePhysicalModel(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.PhysicalModel{}).Error
		return
	})
	return
}
