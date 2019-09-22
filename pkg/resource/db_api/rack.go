package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRack) (data *spec.Rack, err error) {
	data = &spec.Rack{}
	err = db.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).First(data).Error
	return
}

func (api *Api) GetRacks(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Rack, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateRacks(tctx *logger.TraceContext, db *gorm.DB, input []spec.Rack) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			var tmpRack db_model.Rack
			if err = tx.Where("name = ? AND datacenter = ?", data.Name, data.Datacenter).
				First(&tmpRack).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpRack = db_model.Rack{
					Name:       data.Name,
					Datacenter: data.Datacenter,
					Kind:       data.Kind,
					Floor:      data.Floor,
					Unit:       data.Unit,
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

func (api *Api) UpdateRacks(tctx *logger.TraceContext, db *gorm.DB, input []spec.Rack) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			if err = tx.Model(&db_model.Rack{}).
				Where("name = ? AND datacenter = ?", data.Name, data.Datacenter).
				Updates(&db_model.Rack{
					Kind:  data.Kind,
					Floor: data.Floor,
					Unit:  data.Unit,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRack(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRack) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).
			Delete(&db_model.Rack{}).Error
		return
	})
	return
}
