package db_api

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
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

func (api *Api) CreateOrUpdateCompute(tctx *logger.TraceContext, tx *gorm.DB,
	service *db_model.RegionService, rspec *spec.RegionServiceComputeSpec,
	clusters []db_model.Cluster, clusterNetworkV4sMap map[string][]db_model.NetworkV4) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	cluster := clusters[0]
	clusterNetworkV4s := clusterNetworkV4sMap[cluster.Name]

	for i := 0; i < rspec.SchedulePolicy.Replicas; i++ {
		name := fmt.Sprintf("%s.r%d.%s.%s", service.Name, i, service.Project, cluster.DomainSuffix)

		var specBytes []byte
		specBytes, err = json_utils.Marshal(rspec)
		if err != nil {
			err = error_utils.NewInvalidDataError("spec", rspec, "Failed Marshal")
			return
		}

		var data db_model.Compute
		if err = tx.Where("name = ?", name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			rspec.Name = name

			var ports []spec.PortSpec
			switch rspec.NetworkPolicy.Version {
			case 4:
				// TODO FIX Kind Compute
				if ports, err = api.AssignNetworkV4Port(tctx, tx, &rspec.NetworkPolicy,
					clusterNetworkV4s, "Compute", name); err != nil {
					return
				}
			}
			rspec.Ports = ports

			data = db_model.Compute{
				Project:       service.Project,
				Kind:          "Compute", // TODO Fix Kind Compute
				Name:          name,
				RegionService: service.Name,
				Region:        cluster.Region,
				Cluster:       cluster.Name,
				Image:         rspec.Image,
				Vcpus:         rspec.Vcpus,
				Memory:        rspec.Memory,
				Disk:          rspec.Disk,
				Spec:          string(specBytes),
				Status:        base_const.StatusCreating,
				StatusReason:  "CreateRegionService",
			}
			if err = tx.Create(&data).Error; err != nil {
				return
			}

		} else { // Update
			if err = tx.Save(&data).Error; err != nil {
				return
			}
			err = tx.Table("computes").Where("id = ?", data.ID).Updates(map[string]interface{}{
				"image":         rspec.Image,
				"vcpus":         rspec.Vcpus,
				"memory":        rspec.Memory,
				"disk":          rspec.Disk,
				"spec":          string(specBytes),
				"status":        base_const.StatusUpdating,
				"status_reason": "UpdateRegionService",
			}).Error
		}
	}

	return
}

func (api *Api) DeleteCompute(tctx *logger.TraceContext, name string) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("name = ?", name).Updates(map[string]interface{}{
			"status":        resource_model.StatusDeleting,
			"status_reason": "DeleteCompute",
		}).Error
		return
	})
	return
}

func (api *Api) CreateClusterCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	computeMap, ok := clusterComputeMap[compute.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("cluster")
		return
	}

	// すでに作成済みかを確認し、作成済みの場合はスキップされる
	if _, ok := computeMap[compute.Name]; !ok {
		// そうでない場合は、作成する
		client, ok := api.clusterClientMap[compute.Cluster]
		if !ok {
			err = error_utils.NewNotFoundError("clusterClient")
			return
		}

		specStr := "[" + compute.Spec + "]"
		queries := []base_client.Query{
			base_client.Query{
				Name: "CreateCompute",
				Data: spec.CreateCompute{
					Spec: specStr,
				},
			},
		}

		_, tmpErr := client.ResourceVirtualAdminCreateCompute(tctx, queries)
		if tmpErr != nil {
			err = fmt.Errorf("Failed CreateCompute: %s", tmpErr.Error())
			return
		}
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        resource_model.StatusCreatingScheduled,
			"status_reason": "CreateClusterCompute",
		}).Error
		return
	})
	return
}

func (api *Api) ConfirmCreatingOrUpdatingScheduledCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	computeMap, ok := clusterComputeMap[compute.Cluster]
	if !ok {
		err = error_utils.NewConflictNotFoundError(compute.Cluster)
		return
	}

	clusterCompute, ok := computeMap[compute.Name]
	if !ok {
		err = error_utils.NewConflictNotFoundError(compute.Name)
		return
	}

	if clusterCompute.Status != resource_model.StatusActive {
		logger.Info(tctx, "Waiting: status is not Active")
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        resource_model.StatusActive,
			"status_reason": "ConfirmedActive",
		}).Error
		return
	})
	return
}

func (api *Api) UpdateClusterCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	client, ok := api.clusterClientMap[compute.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	specStr := "[" + compute.Spec + "]"
	queries := []base_client.Query{
		base_client.Query{
			Name: "UpdateCompute",
			Data: spec.UpdateCompute{
				Spec: specStr,
			},
		},
	}

	_, tmpErr := client.ResourceVirtualAdminUpdateCompute(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed UpdateCompute: %s", tmpErr.Error())
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        resource_model.StatusUpdatingScheduled,
			"status_reason": "UpdateClusterCompute",
		}).Error
		return
	})
	return
}
