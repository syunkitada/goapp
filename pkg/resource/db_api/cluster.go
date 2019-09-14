package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetCluster(tctx *logger.TraceContext, db *gorm.DB, name string) (data *spec.Cluster, err error) {
	data = &spec.Cluster{}
	err = db.Where("name = ?", name).First(data).Error
	return
}

func (api *Api) GetClusters(tctx *logger.TraceContext, db *gorm.DB) (data []spec.Cluster, err error) {
	err = db.Find(&data).Error
	return
}

func (api *Api) CreateClusters(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Cluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			var tmpCluster db_model.Cluster
			if err = tx.Where("name = ?", region.Name).First(&tmpCluster).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmpCluster = db_model.Cluster{
					Name: region.Name,
					Kind: region.Kind,
				}
				if err = tx.Create(&tmpCluster).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateClusters(tctx *logger.TraceContext, db *gorm.DB, regions []spec.Cluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, region := range regions {
			if err = tx.Model(&db_model.Cluster{}).Where("name = ?", region.Name).Updates(&db_model.Cluster{
				Kind: region.Kind,
			}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteCluster(tctx *logger.TraceContext, db *gorm.DB, name string) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", name).Unscoped().Delete(&db_model.Cluster{}).Error
		return
	})
	return
}
