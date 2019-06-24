package resource_model_api

import (
	"encoding/json"
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetRegionService(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var region resource_model.RegionService
	if err = db.Where(&resource_model.RegionService{
		Name: resource,
	}).First(&region).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["RegionService"] = region
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetRegionServices(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error

	var regions []resource_model.RegionService
	if err = db.Find(&regions).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["RegionServices"] = regions
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateRegionService(tctx *logger.TraceContext,
	db *gorm.DB, req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.RegionServiceSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.RegionService
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}
			specBytes, err := json_utils.Marshal(spec)
			if err != nil {
				return codes.ServerInternalError, error_utils.NewInvalidDataError("spec", spec, "Failed Marshal")
			}

			data = resource_model.RegionService{
				Project:      req.Tctx.ProjectName,
				Kind:         spec.Kind,
				Name:         spec.Name,
				Region:       spec.Region,
				Status:       resource_model.StatusInitializing,
				StatusReason: "CreateRegionService",
				Spec:         string(specBytes),
			}
			if err = tx.Create(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return codes.ClientAlreadyExists, err
		}
	}

	tx.Commit()
	return codes.OkCreated, nil
}

func (modelApi *ResourceModelApi) UpdateRegionService(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.RegionServiceSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}
		specBytes, err := json_utils.Marshal(spec)
		if err != nil {
			return codes.ServerInternalError, error_utils.NewInvalidDataError("spec", spec, "Failed Marshal")
		}
		region := &resource_model.RegionService{
			Kind:         spec.Kind,
			Region:       spec.Region,
			Status:       resource_model.StatusUpdating,
			StatusReason: "UpdateRegionService",
			Spec:         string(specBytes),
		}
		if err = tx.Model(region).Where("name = ?", spec.Name).Updates(region).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteRegionService(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NameSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		if err = tx.Delete(&resource_model.RegionService{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) SyncRegionService(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	clusterNetworksMap := map[string][]resource_model.NetworkV4{}
	var networks []resource_model.NetworkV4
	if err = db.Find(&networks).Error; err != nil {
		return err
	}
	for _, network := range networks {
		if networks, ok := clusterNetworksMap[network.Cluster]; ok {
			networks = append(networks, network)
		} else {
			clusterNetworksMap[network.Cluster] = []resource_model.NetworkV4{network}
		}
	}

	regionClustersMap := map[string][]resource_model.Cluster{}
	var clusters []resource_model.Cluster
	if err = db.Find(&clusters).Error; err != nil {
		return err
	}
	for _, cluster := range clusters {
		if rclusters, ok := regionClustersMap[cluster.Region]; ok {
			rclusters = append(rclusters, cluster)
		} else {
			regionClustersMap[cluster.Region] = []resource_model.Cluster{cluster}
		}
	}

	var regionServices []resource_model.RegionService
	if err = db.Find(&regionServices).Error; err != nil {
		return err
	}

	for _, service := range regionServices {
		tctx.Metadata["RegionServiceId"] = strconv.FormatUint(uint64(service.ID), 10)
		switch service.Status {
		case resource_model.StatusInitializing:
			modelApi.InitializeRegionService(tctx, db, &service, regionClustersMap)
		case resource_model.StatusCreatingInitialized:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case resource_model.StatusCreatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case resource_model.StatusUpdating:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case resource_model.StatusUpdatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case resource_model.StatusDeleting:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case resource_model.StatusDeletingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		}
		tctx.Metadata = map[string]string{}
	}

	// SyncComputes
	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	for _, compute := range computes {
		tctx.Metadata["ComputeId"] = strconv.FormatUint(uint64(compute.ID), 10)
		switch compute.Status {
		case resource_model.StatusInitializing:
			modelApi.InitializeCompute(tctx, db, &compute, clusterNetworksMap)
		case resource_model.StatusCreatingInitialized:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusCreatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdating:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeleting:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeletingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		}
		tctx.Metadata = map[string]string{}
	}

	return nil
}

func (modelApi *ResourceModelApi) InitializeRegionService(tctx *logger.TraceContext, db *gorm.DB,
	regionService *resource_model.RegionService, regionClustersMap map[string][]resource_model.Cluster) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var spec resource_model.RegionServiceSpec
	if err = json_utils.Unmarshal(regionService.Spec, &spec); err != nil {
		return
	}

	clusters, ok := regionClustersMap[spec.Region]
	if !ok {
		err = error_utils.NewNotFoundError("cluster")
		logger.Warningf(tctx, err, "cluster not found: region=%v", spec.Region)
		return
	}

	var cluster resource_model.Cluster
	switch spec.SchedulePolicy.Cluster {
	case resource_model.SchedulePolicyAffinity:
		cluster = clusters[0]
	}

	tx := db.Begin()
	defer tx.Rollback()
	if err = modelApi.CreateCompute(tctx, tx, regionService, &spec, &cluster); err != nil {
		return
	}

	regionService.Status = resource_model.StatusCreating
	regionService.StatusReason = "InitializeRegionService"
	tx.Save(regionService)

	tx.Commit()
	return
}