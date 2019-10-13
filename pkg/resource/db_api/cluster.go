package db_api

import (
	"strings"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) GetCluster(tctx *logger.TraceContext, input *spec.GetCluster, user *base_spec.UserAuthority) (data *spec.Cluster, err error) {
	data = &spec.Cluster{}
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetClusters(tctx *logger.TraceContext, input *spec.GetClusters, user *base_spec.UserAuthority) (data []spec.Cluster, err error) {
	err = api.DB.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreateClusters(tctx *logger.TraceContext, input []spec.Cluster, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
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

func (api *Api) CreateOrUpdateCluster(tctx *logger.TraceContext, input *spec.UpdateCluster) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		cluster := db_model.Cluster{
			Region:       input.Region,
			Datacenter:   input.Datacenter,
			Name:         input.Name,
			Kind:         input.Kind,
			DomainSuffix: input.DomainSuffix,
			Token:        input.Token,
			Project:      input.Project,
			Endpoints:    strings.Join(input.Endpoints, ","),
			Weight:       input.Weight,
		}
		var tmpCluster db_model.Cluster
		if err = tx.Where("name = ?", input.Name).First(&tmpCluster).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			err = tx.Create(&cluster).Error
			return
		}
		err = tx.Model(&tmpCluster).Updates(&cluster).Error

		return
	})
	return
}

func (api *Api) DeleteCluster(tctx *logger.TraceContext, input *spec.DeleteCluster, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ?", input.Name).Delete(&db_model.Cluster{}).Error
		return
	})
	return
}

func (api *Api) DeleteClusters(tctx *logger.TraceContext, input []spec.Cluster, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
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
