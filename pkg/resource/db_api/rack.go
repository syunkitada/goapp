package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRack(tctx *logger.TraceContext, input *spec.GetRack, user *base_spec.UserAuthority) (data *spec.Rack, err error) {
	data = &spec.Rack{}
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetRacks(tctx *logger.TraceContext, input *spec.GetRacks, user *base_spec.UserAuthority) (data []spec.Rack, err error) {
	err = api.DB.Where("datacenter = ? AND deleted_at IS NULL", input.Datacenter).Find(&data).Error
	return
}

func (api *Api) CreateRacks(tctx *logger.TraceContext, input []spec.Rack, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.Rack
			if err = tx.Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.Rack{
					Name:       val.Name,
					Datacenter: val.Datacenter,
					Kind:       val.Kind,
					Floor:      val.Floor,
					Unit:       val.Unit,
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

func (api *Api) UpdateRacks(tctx *logger.TraceContext, input []spec.Rack, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.Rack{}).
				Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				Updates(&db_model.Rack{
					Kind:  val.Kind,
					Floor: val.Floor,
					Unit:  val.Unit,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRack(tctx *logger.TraceContext, input *spec.DeleteRack, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND datacenter = ?", input.Name, input.Datacenter).Delete(&db_model.Rack{}).Error
		return
	})
	return
}

func (api *Api) DeleteRacks(tctx *logger.TraceContext, input []spec.Rack, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND datacenter = ?", val.Name, val.Datacenter).
				Delete(&db_model.Rack{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
