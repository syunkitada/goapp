package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetFloor(tctx *logger.TraceContext, input *spec.GetFloor) (data *spec.Floor, err error) {
	data = &spec.Floor{}
	err = api.DB.Where("name = ? AND datacenter = ? AND deleted_at IS NULL", input.Name, input.Datacenter).
		First(data).Error
	return
}

func (api *Api) GetFloors(tctx *logger.TraceContext, input *spec.GetFloors) (data []spec.Floor, err error) {
	err = api.DB.Where("datacenter = ? AND deleted_at IS NULL", input.Datacenter).Find(&data).Error
	return
}

func (api *Api) CreateFloors(tctx *logger.TraceContext, input []spec.Floor) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.Floor
			if err = tx.Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.Floor{
					Name:       val.Name,
					Kind:       val.Kind,
					Datacenter: val.Datacenter,
					Zone:       val.Zone,
					Floor:      val.Floor,
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

func (api *Api) UpdateFloors(tctx *logger.TraceContext, input []spec.Floor) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.Floor{}).
				Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				Updates(&db_model.Floor{
					Kind:  val.Kind,
					Zone:  val.Zone,
					Floor: val.Floor,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteFloor(tctx *logger.TraceContext, input *spec.DeleteFloor) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).Delete(&db_model.Floor{}).Error
		return
	})
	return
}

func (api *Api) DeleteFloors(tctx *logger.TraceContext, input []spec.Floor) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				Delete(&db_model.Floor{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
