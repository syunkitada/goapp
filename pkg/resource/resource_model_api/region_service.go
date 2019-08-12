package resource_model_api

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
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

	strArgs, ok := query.StrParams["Args"]
	if !ok || len(strArgs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Args")
		return codes.ClientBadRequest, err
	}

	var args []string
	if err = json.Unmarshal([]byte(strArgs), &args); err != nil {
		return codes.ClientBadRequest, err
	}

	for _, arg := range args {
		if err = tx.Delete(&resource_model.RegionService{}, "name = ?", arg).Error; err != nil {
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
	defer modelApi.close(tctx, db)

	clusterNetworkV4sMap := map[string][]resource_model.NetworkV4{}
	var networks []resource_model.NetworkV4
	if err = db.Find(&networks).Error; err != nil {
		return err
	}
	for _, network := range networks {
		if networks, ok := clusterNetworkV4sMap[network.Cluster]; ok {
			networks = append(networks, network)
		} else {
			clusterNetworkV4sMap[network.Cluster] = []resource_model.NetworkV4{network}
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

	regionImageMap := map[string]map[string]resource_model.Image{}
	var images []resource_model.Image
	if err = db.Where(&resource_model.Image{Status: resource_model.StatusActive}).
		Find(&images).Error; err != nil {
		return err
	}
	for _, image := range images {
		imageMap, ok := regionImageMap[image.Region]
		if !ok {
			imageMap = map[string]resource_model.Image{}
		}
		imageMap[image.Name] = image
		regionImageMap[image.Region] = imageMap
	}

	var regionServices []resource_model.RegionService
	if err = db.Find(&regionServices).Error; err != nil {
		return err
	}

	for _, service := range regionServices {
		tctx.Metadata["RegionServiceId"] = strconv.FormatUint(uint64(service.ID), 10)
		switch service.Status {
		case resource_model.StatusInitializing:
			modelApi.InitializeRegionService(
				tctx, db, &service, regionClustersMap, clusterNetworkV4sMap, regionImageMap)
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

	clusterComputeMap := map[string]map[string]resource_model.Compute{}
	for clusterName, clusterApiClient := range modelApi.clusterClientMap {
		computeMap, ok := clusterComputeMap[clusterName]
		if !ok {
			computeMap = map[string]resource_model.Compute{}
		}

		queries := []authproxy_model.Query{
			authproxy_model.Query{
				Kind: "get_computes",
			},
		}
		var rep *authproxy_grpc_pb.ActionReply
		if rep, err = clusterApiClient.Action(
			logger.NewActionTraceContext(tctx, "system", "system", queries)); err != nil {
			return err
		}

		var resp resource_model.ActionResponse
		if err = json_utils.Unmarshal(rep.Response, &resp); err != nil {
			return err
		}
		if resp.Tctx.StatusCode == codes.OkRead {
			for _, compute := range resp.Data.Computes {
				fmt.Println("DEBUG NAME", compute.Name)
				computeMap[compute.Name] = compute
			}
		}

		clusterComputeMap[clusterName] = computeMap
	}
	fmt.Println("DEBUG clusterComputeMap", clusterComputeMap)

	// SyncComputes
	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	for _, compute := range computes {
		tctx.Metadata["ComputeId"] = strconv.FormatUint(uint64(compute.ID), 10)
		switch compute.Status {
		case resource_model.StatusCreating:
			modelApi.CreateClusterCompute(tctx, db, &compute, clusterComputeMap)
		case resource_model.StatusCreatingScheduled:
			modelApi.ConfirmCreatingScheduledCompute(tctx, db, &compute, clusterComputeMap)
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
	regionService *resource_model.RegionService, regionClustersMap map[string][]resource_model.Cluster,
	clusterNetworkV4sMap map[string][]resource_model.NetworkV4,
	regionImageMap map[string]map[string]resource_model.Image) {

	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var spec resource_model.RegionServiceSpec
	if err = json_utils.Unmarshal(regionService.Spec, &spec); err != nil {
		return
	}

	tmpClusters, ok := regionClustersMap[spec.Region]
	if !ok {
		logger.Warningf(tctx, "cluster not found: region=%v", spec.Region)
		return
	}

	imageMap, ok := regionImageMap[spec.Region]
	if !ok {
		logger.Warningf(tctx, "image not found: region=%v", spec.Region)
		return
	}
	image, ok := imageMap[spec.Compute.Image]
	if !ok {
		logger.Warningf(tctx, "image not found: region=%v, image=%v", spec.Region, spec.Compute.Image)
		return
	}
	var imageSpec resource_model.ImageSpec
	if err = json_utils.Unmarshal(image.Spec, &imageSpec); err != nil {
		return
	}
	spec.Compute.ImageSpec = imageSpec

	policy := spec.Compute.SchedulePolicy
	enableClusterFilters := false
	if len(policy.ClusterFilters) > 0 {
		enableClusterFilters = true
	}
	enableLabelFilters := false
	if len(policy.ClusterLabelFilters) > 0 {
		enableLabelFilters = true
	}

	clusters := []resource_model.Cluster{}
	for _, cluster := range tmpClusters {
		_, ok := clusterNetworkV4sMap[cluster.Name]
		if !ok {
			continue
		}

		if enableClusterFilters {
			ok = false
			for _, filter := range policy.ClusterFilters {
				if filter == cluster.Name {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableLabelFilters {
			ok = false
			for _, labelFilter := range policy.ClusterLabelFilters {
				if strings.Index(cluster.Labels, labelFilter) >= 0 {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		clusters = append(clusters, cluster)
	}

	tx := db.Begin()
	defer tx.Rollback()
	if len(clusters) == 0 {
		err = error_utils.NewNotFoundError(resource_model.StatusMsgInitializeErrorNoValidCluster)
		regionService.Status = resource_model.StatusError
		regionService.StatusReason = resource_model.StatusMsgInitializeErrorNoValidCluster
		tx.Save(regionService)
		tx.Commit()
		return
	}

	// TODO Sort clusters by weight
	// TODO Sort clusters by resource
	cluster := clusters[0]
	clusterNetworkV4s := clusterNetworkV4sMap[cluster.Name]

	if err = modelApi.CreateCompute(tctx, tx, regionService, &spec,
		&cluster, clusterNetworkV4s); err != nil {
		return
	}

	regionService.Status = resource_model.StatusCreating
	regionService.StatusReason = resource_model.StatusMsgInitializeSuccess
	tx.Save(regionService)

	tx.Commit()
	return
}
