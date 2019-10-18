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

func (api *Api) CreateCompute(tctx *logger.TraceContext, tx *gorm.DB,
	service *db_model.RegionService, rspec *spec.RegionServiceComputeSpec,
	cluster *db_model.Cluster, clusterNetworkV4s []db_model.NetworkV4) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	for i := 0; i < rspec.SchedulePolicy.Replicas; i++ {
		name := fmt.Sprintf("%s.r%d.%s.%s", service.Name, i, service.Project, cluster.DomainSuffix)
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
					return err
				}
			}
			rspec.Ports = ports

			specBytes, err := json_utils.Marshal(rspec)
			if err != nil {
				return error_utils.NewInvalidDataError("spec", rspec, "Failed Marshal")
			}

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
				StatusReason:  "CreatingRegionService",
			}
			if err = tx.Create(&data).Error; err != nil {
				return err
			}
		}
	}

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

func (api *Api) CreateClusterCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var rspec spec.RegionServiceComputeSpec
	if err = json_utils.Unmarshal(compute.Spec, &rspec); err != nil {
		return
	}

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

		res, tmpErr := client.ResourceVirtualAdminCreateCompute(tctx, queries)
		if tmpErr != nil {
			err = fmt.Errorf("Failed CreateCompute: %s", tmpErr.Error())
			return
		}

		fmt.Println("DEBUG CreateCompute res", res)
	}

	// clusterApiClient, ok := modelApi.clusterClientMap[compute.Cluster]
	// if !ok {
	// 	err = error_utils.NewNotFoundError("cluster")
	// 	return
	// }

	// if _, ok := computeMap[compute.Name]; !ok {
	// 	specs := "[" + compute.Spec + "]"
	// 	queries := []authproxy_model.Query{
	// 		authproxy_model.Query{
	// 			Kind: "create_compute",
	// 			StrParams: map[string]string{
	// 				"Specs": specs,
	// 			},
	// 		},
	// 	}
	// 	var rep *authproxy_grpc_pb.ActionReply
	// 	if rep, err = clusterApiClient.Action(
	// 		logger.NewActionTraceContext(tctx, compute.Project, "", queries)); err != nil {
	// 		return
	// 	}

	// 	var resp resource_model.ActionResponse
	// 	if err = json_utils.Unmarshal(rep.Response, &resp); err != nil {
	// 		return
	// 	}
	// 	if resp.Tctx.StatusCode != codes.OkCreated {
	// 		logger.Warningf(tctx, "Failed create: %s", resp.Tctx.Err)
	// 		return
	// 	}
	// }

	// compute.Status = resource_model.StatusCreatingScheduled
	// compute.StatusReason = "CreateClusterCompute"

	// err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
	// 	err = tx.Save(compute).Error
	// })
	return
}

func (api *Api) ConfirmCreatingScheduledCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	fmt.Println("DEBUG Api.ConfirmClusterCompute")
}

// TODO
// func (modelApi *ResourceModelApi) CreateClusterCompute(tctx *logger.TraceContext, db *gorm.DB,
// 	compute *resource_model.Compute, clusterComputeMap map[string]map[string]resource_model.Compute) {
// 	var err error
// 	startTime := logger.StartTrace(tctx)
// 	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
//
// 	var spec resource_model.RegionServiceSpec
// 	if err = json_utils.Unmarshal(compute.Spec, &spec); err != nil {
// 		return
// 	}
//
// 	computeMap, ok := clusterComputeMap[compute.Cluster]
// 	if !ok {
// 		err = error_utils.NewNotFoundError("cluster")
// 		return
// 	}
//
// 	clusterApiClient, ok := modelApi.clusterClientMap[compute.Cluster]
// 	if !ok {
// 		err = error_utils.NewNotFoundError("cluster")
// 		return
// 	}
//
// 	if _, ok := computeMap[compute.Name]; !ok {
// 		specs := "[" + compute.Spec + "]"
// 		queries := []authproxy_model.Query{
// 			authproxy_model.Query{
// 				Kind: "create_compute",
// 				StrParams: map[string]string{
// 					"Specs": specs,
// 				},
// 			},
// 		}
// 		var rep *authproxy_grpc_pb.ActionReply
// 		if rep, err = clusterApiClient.Action(
// 			logger.NewActionTraceContext(tctx, compute.Project, "", queries)); err != nil {
// 			return
// 		}
//
// 		var resp resource_model.ActionResponse
// 		if err = json_utils.Unmarshal(rep.Response, &resp); err != nil {
// 			return
// 		}
// 		if resp.Tctx.StatusCode != codes.OkCreated {
// 			logger.Warningf(tctx, "Failed create: %s", resp.Tctx.Err)
// 			return
// 		}
// 	}
//
// 	compute.Status = resource_model.StatusCreatingScheduled
// 	compute.StatusReason = "CreateClusterCompute"
//
// 	tx := db.Begin()
// 	defer tx.Rollback()
// 	if err = tx.Save(compute).Error; err != nil {
// 		return
// 	}
//
// 	tx.Commit()
// 	return
// }
//
// func (modelApi *ResourceModelApi) ConfirmCreatingScheduledCompute(tctx *logger.TraceContext, db *gorm.DB,
// 	compute *resource_model.Compute, clusterComputeMap map[string]map[string]resource_model.Compute) {
// 	var err error
// 	startTime := logger.StartTrace(tctx)
// 	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
//
// 	computeMap, ok := clusterComputeMap[compute.Cluster]
// 	if !ok {
// 		err = error_utils.NewConflictNotFoundError(compute.Cluster)
// 		return
// 	}
//
// 	clusterCompute, ok := computeMap[compute.Name]
// 	if !ok {
// 		err = error_utils.NewConflictNotFoundError(compute.Name)
// 		return
// 	}
//
// 	if clusterCompute.Status != resource_model.StatusActive {
// 		logger.Info(tctx, "Waiting: status is not Active")
// 		return
// 	}
//
// 	tx := db.Begin()
// 	defer tx.Rollback()
//
// 	tmpCompute := resource_model.Compute{
// 		Status:       resource_model.StatusActive,
// 		StatusReason: "ConfirmedCreagingScheduled",
// 	}
// 	if err = tx.Model(&tmpCompute).Where("id = ?", compute.ID).Updates(&tmpCompute).Error; err != nil {
// 		return
// 	}
// 	tx.Commit()
//
// 	return
// }
