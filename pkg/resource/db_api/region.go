package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (data *spec.Region, err error) {
	data = &spec.Region{}
	err = db.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data []spec.Region, err error) {
	err = db.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreateRegions(tctx *logger.TraceContext, db *gorm.DB, input []spec.Region) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.Region
			if err = tx.Where("name = ? AND deleted_at IS NULL", val.Name).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.Region{
					Name: val.Name,
					Kind: val.Kind,
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

func (api *Api) UpdateRegions(tctx *logger.TraceContext, db *gorm.DB, input []spec.Region) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.Region{}).
				Where("name = ?", val.Name).
				Updates(&db_model.Region{
					Kind: val.Kind,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRegion) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", input.Name).Delete(&db_model.Region{}).Error
		return
	})
	return
}

func (api *Api) DeleteRegions(tctx *logger.TraceContext, db *gorm.DB, input []spec.Region) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			if err = tx.Where("name = ?", data.Name).
				Delete(&db_model.Region{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
