package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRegionService(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegionService) (data *spec.RegionService, err error) {
	data = &spec.RegionService{}
	err = db.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetRegionServices(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegionServices) (data []spec.RegionService, err error) {
	err = db.Where("region = ? AND deleted_at IS NULL", input.Region).Find(&data).Error
	return
}

func (api *Api) CreateRegionServices(tctx *logger.TraceContext, db *gorm.DB, input []spec.RegionService) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.RegionService
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.RegionService{
					Name:   val.Name,
					Region: val.Region,
					Kind:   val.Kind,
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

func (api *Api) UpdateRegionServices(tctx *logger.TraceContext, db *gorm.DB, input []spec.RegionService) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.RegionService{}).
				Where("name = ? AND region = ?", val.Name, val.Region).
				Updates(&db_model.RegionService{
					Kind: val.Kind,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRegionService(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteRegionService) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).
			Delete(&db_model.RegionService{}).Error
		return
	})
	return
}

func (api *Api) DeleteRegionServices(tctx *logger.TraceContext, db *gorm.DB, input []spec.RegionService) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				Delete(&db_model.RegionService{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
