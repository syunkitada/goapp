package db_api

import (
	"fmt"
	"strconv"
	"strings"

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

func (api *Api) GetRegionService(tctx *logger.TraceContext, input *spec.GetRegionService, user *base_spec.UserAuthority) (data *spec.RegionService, err error) {
	data = &spec.RegionService{}
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(data).Error
	return
}

func (api *Api) GetRegionServices(tctx *logger.TraceContext, input *spec.GetRegionServices, user *base_spec.UserAuthority) (data []spec.RegionService, err error) {
	err = api.DB.Where("region = ? AND deleted_at IS NULL", input.Region).Find(&data).Error
	return
}

func (api *Api) CreateRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			var tmp db_model.RegionService
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.RegionService{
					Project:      user.ProjectName,
					Name:         val.Name,
					Region:       val.Region,
					Kind:         val.Kind,
					Status:       db_model.StatusInitializing,
					StatusReason: "CreateRegionService",
					Spec:         string(specBytes),
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

func (api *Api) UpdateRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
			if err = tx.Model(&db_model.RegionService{}).
				Where("name = ? AND region = ?", val.Name, val.Region).
				Updates(&db_model.RegionService{
					Spec: string(specBytes),
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteRegionService(tctx *logger.TraceContext, input *spec.DeleteRegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).
			Delete(&db_model.RegionService{}).Error
		return
	})
	return
}

func (api *Api) DeleteRegionServices(tctx *logger.TraceContext, input []spec.RegionService, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND region = ?", val.Name, val.Region).
				Delete(&db_model.RegionService{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) SyncRegionService(tctx *logger.TraceContext) (err error) {
	fmt.Println("DEBUG SyncRegionService")

	clusterNetworkV4sMap := map[string][]db_model.NetworkV4{}
	var networks []db_model.NetworkV4
	if err = api.DB.Find(&networks).Error; err != nil {
		return
	}
	for _, network := range networks {
		if networks, ok := clusterNetworkV4sMap[network.Cluster]; ok {
			networks = append(networks, network)
		} else {
			clusterNetworkV4sMap[network.Cluster] = []db_model.NetworkV4{network}
		}
	}

	regionClustersMap := map[string][]db_model.Cluster{}
	var clusters []db_model.Cluster
	if err = api.DB.Find(&clusters).Error; err != nil {
		return
	}
	for _, cluster := range clusters {
		if rclusters, ok := regionClustersMap[cluster.Region]; ok {
			rclusters = append(rclusters, cluster)
		} else {
			regionClustersMap[cluster.Region] = []db_model.Cluster{cluster}
		}
	}

	regionImageMap := map[string]map[string]db_model.Image{}
	var images []db_model.Image
	if err = api.DB.Where(&db_model.Image{Status: db_model.StatusActive}).
		Find(&images).Error; err != nil {
		return err
	}
	for _, image := range images {
		imageMap, ok := regionImageMap[image.Region]
		if !ok {
			imageMap = map[string]db_model.Image{}
		}
		imageMap[image.Name] = image
		regionImageMap[image.Region] = imageMap
	}

	var regionServices []db_model.RegionService
	if err = api.DB.Find(&regionServices).Error; err != nil {
		return err
	}

	fmt.Println("DEBUG regionServices", regionServices)

	for _, service := range regionServices {
		tctx.Metadata["RegionServiceId"] = strconv.FormatUint(uint64(service.ID), 10)
		switch service.Status {
		case db_model.StatusInitializing:
			switch service.Kind {
			case "Compute":
				api.InitializeRegionServiceCompute(
					tctx, &service, regionClustersMap, clusterNetworkV4sMap, regionImageMap)
			}
		case db_model.StatusCreatingInitialized:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case db_model.StatusCreatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case db_model.StatusUpdating:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case db_model.StatusUpdatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case db_model.StatusDeleting:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		case db_model.StatusDeletingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", service.Status, service.Name)
		}
		tctx.Metadata = map[string]string{}
	}

	clusterComputeMap := map[string]map[string]spec.Compute{}
	for clusterName, clusterApiClient := range api.clusterClientMap {
		computeMap, ok := clusterComputeMap[clusterName]
		if !ok {
			computeMap = map[string]spec.Compute{}
		}

		queries := []base_client.Query{
			base_client.Query{
				Name: "GetComputes",
				Data: spec.GetCompute{},
			},
		}

		res, tmpErr := clusterApiClient.ResourceVirtualAdminGetComputes(tctx, queries)
		if tmpErr != nil {
			err = fmt.Errorf("Failed GetComputes: %s", tmpErr.Error())
			continue
		}

		fmt.Println("DEBUG GetComputes res", res, computeMap)
		// TODO update computeMap

		clusterComputeMap[clusterName] = computeMap
	}

	// SyncComputes
	var computes []db_model.Compute
	if err = api.DB.Find(&computes).Error; err != nil {
		return err
	}

	for _, compute := range computes {
		tctx.Metadata["ComputeId"] = strconv.FormatUint(uint64(compute.ID), 10)
		switch compute.Status {
		case db_model.StatusCreating:
			api.CreateClusterCompute(tctx, &compute, clusterComputeMap)
		case db_model.StatusCreatingScheduled:
			api.ConfirmCreatingScheduledCompute(tctx, &compute, clusterComputeMap)
		case db_model.StatusUpdating:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case db_model.StatusUpdatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case db_model.StatusDeleting:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case db_model.StatusDeletingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		}
		tctx.Metadata = map[string]string{}
	}

	return
}

func (api *Api) InitializeRegionServiceCompute(tctx *logger.TraceContext,
	service *db_model.RegionService, regionClustersMap map[string][]db_model.Cluster,
	clusterNetworkV4sMap map[string][]db_model.NetworkV4,
	regionImageMap map[string]map[string]db_model.Image) {

	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var rspec spec.RegionServiceComputeSpec
	if err = json_utils.Unmarshal(service.Spec, &rspec); err != nil {
		return
	}

	tmpClusters, ok := regionClustersMap[service.Region]
	if !ok {
		logger.Warningf(tctx, "cluster not found: region=%v", service.Region)
		return
	}

	imageMap, ok := regionImageMap[service.Region]
	if !ok {
		logger.Warningf(tctx, "image not found: region=%v", service.Region)
		return
	}
	image, ok := imageMap[rspec.Image]
	if !ok {
		logger.Warningf(tctx, "image not found: region=%s, image=%v", service.Region, rspec.Image)
		return
	}

	switch image.Kind {
	case "Url":
		var imageSpec spec.ImageUrlSpec
		if err = json_utils.Unmarshal(image.Spec, &imageSpec); err != nil {
			return
		}
		rspec.ImageSpec = spec.Image{
			Region: image.Region,
			Name:   image.Name,
			Kind:   image.Kind,
			Spec:   imageSpec,
		}
	default:
		logger.Warningf(tctx, "invalid image kind: kind=%s", image.Kind)
		return
	}

	policy := rspec.SchedulePolicy
	enableClusterFilters := false
	if len(policy.ClusterFilters) > 0 {
		enableClusterFilters = true
	}
	enableLabelFilters := false
	if len(policy.ClusterLabelFilters) > 0 {
		enableLabelFilters = true
	}

	clusters := []db_model.Cluster{}
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

	if len(clusters) == 0 {
		err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
			err = error_utils.NewNotFoundError("NoValidCluster")
			service.Status = base_const.StatusError
			service.StatusReason = err.Error()
			tx.Save(service)
			return
		})
		if err != nil {
			return
		}
	}

	// TODO Sort clusters by weight
	// TODO Sort clusters by resource
	cluster := clusters[0]
	clusterNetworkV4s := clusterNetworkV4sMap[cluster.Name]

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		if err = api.CreateCompute(tctx, tx, service, &rspec,
			&cluster, clusterNetworkV4s); err != nil {
			return
		}

		service.Status = base_const.StatusCreating
		service.StatusReason = "CreatingCompute"
		tx.Save(service)
		return
	})
	return
}
