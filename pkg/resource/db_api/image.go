package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetImage(tctx *logger.TraceContext, input *spec.GetImage) (data *spec.Image, err error) {
	data = &spec.Image{}
	err = api.DB.Where("name = ? AND region = ? AND deleted_at IS NULL", input.Name, input.Region).
		First(data).Error
	return
}

func (api *Api) GetImages(tctx *logger.TraceContext, input *spec.GetImages) (data []spec.Image, err error) {
	err = api.DB.Where("region = ? AND deleted_at IS NULL", input.Region).Find(&data).Error
	return
}

func (api *Api) CreateImages(tctx *logger.TraceContext, input []spec.Image) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			var tmp db_model.Image
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.Image{
					Name:        val.Name,
					Region:      val.Region,
					Kind:        val.Kind,
					Description: val.Description,
					Status:      base_const.StatusActive,
					Spec:        string(specBytes),
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

func (api *Api) UpdateImages(tctx *logger.TraceContext, input []spec.Image) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.Image{}).
				Where("name = ? AND region = ?", val.Name, val.Region).
				Updates(&db_model.Image{
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

func (api *Api) DeleteImage(tctx *logger.TraceContext, input *spec.DeleteImage) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).Delete(&db_model.Image{}).Error
		return
	})
	return
}

func (api *Api) DeleteImages(tctx *logger.TraceContext, input []spec.Image) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				Delete(&db_model.Image{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
