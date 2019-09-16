package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetDatacenter(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.Datacenter, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	data = &spec.Datacenter{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetDatacenters(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Datacenter, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateDatacenters(tctx *logger.TraceContext, db *gorm.DB, datacenters []spec.Datacenter) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, datacenter := range datacenters {
			var tmpDatacenter db_model.Datacenter
			if err = tx.Where("name = ?", datacenter.Name).First(&tmpDatacenter).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpDatacenter = db_model.Datacenter{
					Name:         datacenter.Name,
					Kind:         datacenter.Kind,
					Description:  datacenter.Description,
					Region:       datacenter.Region,
					DomainSuffix: datacenter.DomainSuffix,
				}
				if err = tx.Create(&tmpDatacenter).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateDatacenters(tctx *logger.TraceContext, db *gorm.DB, datacenters []spec.Datacenter) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, datacenter := range datacenters {
			if err = tx.Model(&db_model.Datacenter{}).Where("name = ?", datacenter.Name).Updates(&db_model.Datacenter{
				Kind:         datacenter.Kind,
				Description:  datacenter.Description,
				Region:       datacenter.Region,
				DomainSuffix: datacenter.DomainSuffix,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteDatacenter(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Datacenter{}).Error
		return
	})
	return
}
