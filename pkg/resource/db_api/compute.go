package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetCompute(tctx *logger.TraceContext, name string, user *base_spec.UserAuthority) (data *spec.Compute, err error) {
	data = &spec.Compute{}
	err = api.DB.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetComputes(tctx *logger.TraceContext, db *gorm.DB, user *base_spec.UserAuthority) (data []spec.Compute, err error) {
	err = api.DB.Find(&data).Error
	return
}

func (api *Api) CreateComputes(tctx *logger.TraceContext, regions []spec.Compute) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
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

func (api *Api) UpdateComputes(tctx *logger.TraceContext, regions []spec.Compute) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
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

func (api *Api) DeleteCompute(tctx *logger.TraceContext, name string) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Compute{}).Error
		return
	})
	return
}
