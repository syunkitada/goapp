package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRegion(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegion) (data *spec.Region, err error) {
	data = &spec.Region{}
	err = db.Where("name = ?", input.Name).First(data).Error
	return
}

func (api *Api) GetRegions(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Region, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateRegions(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Region) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			var tmpRegion db_model.Region
			if err = tx.Where("name = ?", region.Name).First(&tmpRegion).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpRegion = db_model.Region{
					Name: region.Name,
					Kind: region.Kind,
				}
				if err = tx.Create(&tmpRegion).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateRegions(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Region) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			if err = tx.Model(&db_model.Region{}).Where("name = ?", region.Name).Updates(&db_model.Region{
				Kind: region.Kind,
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
		err = tx.Where("name = ?", input.Name).Unscoped().Delete(&db_model.Region{}).Error
		return
	})
	return
}
