package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetCompute(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.Compute, err error) {
	data = &spec.Compute{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetComputes(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Compute, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateComputes(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Compute) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			var tmpCompute db_model.Compute
			if err = tx.Where("name = ?", region.Name).First(&tmpCompute).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpCompute = db_model.Compute{
					Name: region.Name,
					Kind: region.Kind,
				}
				if err = tx.Create(&tmpCompute).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateComputes(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Compute) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			if err = tx.Model(&db_model.Compute{}).Where("name = ?", region.Name).Updates(&db_model.Compute{
				Kind: region.Kind,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteCompute(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Compute{}).Error
		return
	})
	return
}
