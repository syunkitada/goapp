package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRack(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.Rack, err error) {
	data = &spec.Rack{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetRacks(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Rack, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateRacks(tctx *logger.TraceContext, db *gorm.DB, racks []spec.Rack) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, rack := range racks {
			var tmpRack db_model.Rack
			if err = tx.Where("name = ? AND datacenter = ?", rack.Name, rack.Datacenter).First(&tmpRack).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpRack = db_model.Rack{
					Name:       rack.Name,
					Datacenter: rack.Datacenter,
					Kind:       rack.Kind,
					Floor:      rack.Floor,
					Unit:       rack.Unit,
				}
				if err = tx.Create(&tmpRack).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateRacks(tctx *logger.TraceContext, db *gorm.DB, racks []spec.Rack) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, rack := range racks {
			if err = tx.Model(&db_model.Rack{}).Where("name = ?", rack.Name).Updates(&db_model.Rack{
				Kind:  rack.Kind,
				Floor: rack.Floor,
				Unit:  rack.Unit,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRack(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Rack{}).Error
		return
	})
	return
}
