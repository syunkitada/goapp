package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetCluster(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetCluster) (data *spec.Cluster, err error) {
	data = &spec.Cluster{}
	err = db.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetClusters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetClusters) (data []spec.Cluster, err error) {
	err = db.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreateClusters(tctx *logger.TraceContext, db *gorm.DB, input []spec.Cluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.Cluster
			if err = tx.Where("name = ? AND deleted_at IS NULL", val.Name).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.Cluster{
					Name:         val.Name,
					Kind:         val.Kind,
					Region:       val.Region,
					Datacenter:   val.Datacenter,
					DomainSuffix: val.DomainSuffix,
					Description:  val.Description,
					Weight:       val.Weight,
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

func (api *Api) UpdateClusters(tctx *logger.TraceContext, db *gorm.DB, input []spec.Cluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.Cluster{}).
				Where("name = ?", val.Name).
				Updates(&db_model.Cluster{
					Kind:        val.Kind,
					Description: val.Description,
					Weight:      val.Weight,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteCluster(tctx *logger.TraceContext, db *gorm.DB, input *spec.DeleteCluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", input.Name).Delete(&db_model.Cluster{}).Error
		return
	})
	return
}

func (api *Api) DeleteClusters(tctx *logger.TraceContext, db *gorm.DB, input []spec.Cluster) (err error) {
	err = api.Transact(tctx, db, func(tx *gorm.DB) (err error) {
		for _, data := range input {
			if err = tx.Where("name = ?", data.Name).
				Delete(&db_model.Cluster{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}
