package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetFloor(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.Floor, err error) {
	data = &spec.Floor{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetFloors(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Floor, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateFloors(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Floor) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			var tmpFloor db_model.Floor
			if err = tx.Where("name = ? AND datacenter = ?", region.Name, region.Datacenter).
				First(&tmpFloor).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpFloor = db_model.Floor{
					Name:       region.Name,
					Kind:       region.Kind,
					Datacenter: region.Datacenter,
					Zone:       region.Zone,
					Floor:      region.Floor,
				}
				if err = tx.Create(&tmpFloor).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateFloors(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Floor) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			if err = tx.Model(&db_model.Floor{}).
				Where("name = ? AND datacenter = ?", region.Name, region.Datacenter).
				Updates(&db_model.Floor{
					Kind:  region.Kind,
					Zone:  region.Zone,
					Floor: region.Floor,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteFloor(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		// FIXME
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Floor{}).Error
		return
	})
	return
}
