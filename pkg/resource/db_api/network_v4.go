package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetNetworkV4(tctx *logger.TraceContext, input *spec.GetNetworkV4) (data *spec.NetworkV4, err error) {
	data = &spec.NetworkV4{}
	err = api.DB.Where("name = ? AND cluster = ? AND deleted_at IS NULL", input.Name, input.Cluster).
		First(data).Error
	return
}

func (api *Api) GetNetworkV4s(tctx *logger.TraceContext, input *spec.GetNetworkV4s) (data []spec.NetworkV4, err error) {
	err = api.DB.Where("cluster = ? AND deleted_at IS NULL", input.Cluster).Find(&data).Error
	return
}

// TODO FIXME
func (api *Api) CreateNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.NetworkV4
			if err = tx.Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.NetworkV4{
					Name:        val.Name,
					Cluster:     val.Cluster,
					Kind:        val.Kind,
					Description: val.Description,
					Status:      base_const.StatusActive,
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

func (api *Api) UpdateNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.NetworkV4{}).
				Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				Updates(&db_model.NetworkV4{
					Kind:        val.Kind,
					Description: val.Description,
					Status:      base_const.StatusActive,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteNetworkV4(tctx *logger.TraceContext, input *spec.DeleteNetworkV4) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).Delete(&db_model.NetworkV4{}).Error
		return
	})
	return
}

func (api *Api) DeleteNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				Delete(&db_model.NetworkV4{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
