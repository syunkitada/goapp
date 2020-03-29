package db_api

import (
	"fmt"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	resource_cluster_api "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (api *Api) GetCompute(tctx *logger.TraceContext, input *spec.GetCompute, user *base_spec.UserAuthority) (data *spec.Compute, err error) {
	var dbData db_model.Compute
	if err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(&dbData).Error; err != nil {
		data = &spec.Compute{}
		return
	}

	data = &spec.Compute{
		Region:        dbData.Region,
		Cluster:       dbData.Cluster,
		RegionService: dbData.RegionService,
		Name:          dbData.Name,
		Kind:          dbData.Kind,
		Project:       dbData.Project,
		Status:        dbData.Status,
		StatusReason:  dbData.StatusReason,
	}
	var specData spec.RegionServiceComputeSpec
	if err = json_utils.Unmarshal(dbData.Spec, &specData); err != nil {
		return
	}
	data.Spec = &specData

	return
}

func (api *Api) GetComputes(tctx *logger.TraceContext, input *spec.GetComputes, user *base_spec.UserAuthority) (data []spec.Compute, err error) {
	err = api.DB.Table("computes").
		Select("region, cluster, region_service, project, name, kind, status, status_reason, updated_at, created_at").
		Where("region = ? AND deleted_at IS NULL", input.Region).Scan(&data).Error
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

		rspec.Name = name
		var ports []spec.PortSpec
		switch rspec.NetworkPolicy.Version {
		case 4:
			if ports, err = api.AssignNetworkV4Port(tctx, tx, &rspec.NetworkPolicy,
				clusterNetworkV4s, consts.KindCompute, name); err != nil {
				return
			}
		}
		rspec.Ports = ports

		var specBytes []byte
		specBytes, err = json_utils.Marshal(rspec)
		if err != nil {
			err = error_utils.NewInvalidDataError("spec", rspec, "Failed Marshal")
			return
		}

		var data []db_model.Compute
		if err = tx.Where("name = ?", name).Find(&data).Error; err != nil {
			return
		}

		if len(data) == 0 {
			compute := db_model.Compute{
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
			if err = tx.Create(&compute).Error; err != nil {
				return
			}

		} else { // Update
			err = tx.Table("computes").Where("id = ?", data[0].ID).Updates(map[string]interface{}{
				"image":         rspec.Image,
				"vcpus":         rspec.Vcpus,
				"memory":        rspec.Memory,
				"disk":          rspec.Disk,
				"spec":          string(specBytes),
				"status":        resource_model.StatusUpdating,
				"status_reason": "UpdateCompute",
			}).Error
		}
	}

	return
}

func (api *Api) ProxyComputeConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, user *base_spec.UserAuthority, wsConn *websocket.Conn) (err error) {
	var compute db_model.Compute
	var tmpErr error
	err = api.DB.Where("name = ? AND deleted_at IS NULL", input.Name).First(&compute).Error

	var clusters []db_model.Cluster
	if err = api.DB.Where("name = ?", compute.Cluster).Find(&clusters).Error; err != nil {
		return
	}
	if len(clusters) != 1 {
		err = fmt.Errorf("Invalid Cluster: %s", compute.Cluster)
		return
	}
	cluster := clusters[0]
	endpoints := strings.Split(cluster.Endpoints, ",")

	client := resource_cluster_api.NewClient(&base_config.ClientConfig{
		Endpoints:             endpoints,
		Token:                 cluster.Token,
		Project:               cluster.Project,
		TlsInsecureSkipVerify: true,
	})

	queries := []base_client.Query{
		base_client.Query{
			Name: "GetComputeConsole",
			Data: *input,
		},
	}
	_, proxyWsConn, tmpErr := client.ResourceVirtualAdminGetComputeConsole(tctx, queries)
	if tmpErr != nil {
		logger.Warningf(tctx, "Failed GetNodeServices: %s", tmpErr.Error())
		return
	}

	var isDone bool
	doneCh := make(chan bool, 2)
	conMutex := sync.Mutex{}
	defer func() {
		conMutex.Lock()
		if tmpErr = proxyWsConn.Close(); tmpErr != nil {
			logger.Warningf(tctx, "Failed proxyWsConn.Close: err=%s", tmpErr.Error())
		} else {
			logger.Info(tctx, "Success proxyWsConn.Close")
		}
		isDone = true
		close(doneCh)
		conMutex.Unlock()
	}()

	go func() {
		var messageType int
		var message []byte
		for {
			fmt.Println("Waiting Messages on client WebSocket")
			if messageType, message, tmpErr = wsConn.ReadMessage(); tmpErr != nil {
				conMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild ReadMessage: %s", tmpErr.Error())
					doneCh <- true
				}
				conMutex.Unlock()
				return
			}
			if tmpErr = proxyWsConn.WriteMessage(messageType, message); tmpErr != nil {
				conMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild WriteMessage: %s", tmpErr.Error())
					doneCh <- true
				}
				conMutex.Unlock()
				return
			}
		}
	}()

	go func() {
		var messageType int
		var message []byte
		for {
			fmt.Println("Waiting Messages on proxy WebSocket")
			if messageType, message, tmpErr = proxyWsConn.ReadMessage(); tmpErr != nil {
				conMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild ReadMessage: %s", tmpErr.Error())
					doneCh <- true
				}
				conMutex.Unlock()
				return
			}
			if tmpErr = wsConn.WriteMessage(messageType, message); tmpErr != nil {
				conMutex.Lock()
				if !isDone {
					logger.Warningf(tctx, "Faild WriteMessage: %s", tmpErr.Error())
					doneCh <- true
				}
				conMutex.Unlock()
				return
			}
		}
	}()

	<-doneCh
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
}

func (api *Api) DeleteClusterCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	client, ok := api.clusterClientMap[compute.Cluster]
	if !ok {
		err = error_utils.NewNotFoundError("clusterClient")
		return
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: "DeleteCompute",
			Data: spec.DeleteCompute{
				Name: compute.Name,
			},
		},
	}

	_, tmpErr := client.ResourceVirtualAdminDeleteCompute(tctx, queries)
	if tmpErr != nil {
		err = fmt.Errorf("Failed DeleteCompute: %s", tmpErr.Error())
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        resource_model.StatusDeletingScheduled,
			"status_reason": "DeleteClusterCompute",
		}).Error
		return
	})
}

func (api *Api) ConfirmDeletingScheduledCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, clusterComputeMap map[string]map[string]spec.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	computeMap, ok := clusterComputeMap[compute.Cluster]
	if !ok {
		err = error_utils.NewConflictNotFoundError(compute.Cluster)
		return
	}

	if _, ok := computeMap[compute.Name]; ok {
		logger.Info(tctx, "Waiting compute is deleted")
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		if err = tx.Where("resource_kind = ? AND resource_name = ?", consts.KindCompute, compute.Name).
			Unscoped().Delete(&db_model.NetworkV4Port{}).Error; err != nil {
			return
		}
		err = tx.Where("id = ?", compute.ID).Unscoped().Delete(&db_model.Compute{}).Error
		return
	})
}
