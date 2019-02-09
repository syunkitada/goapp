package resource_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ResourceModelApi struct {
	host             string
	name             string
	conf             *config.Config
	downTimeDuration time.Duration
	clusterClientMap map[string]*resource_cluster_api_client.ResourceClusterApiClient
	validate         *validator.Validate
}

func NewResourceModelApi(conf *config.Config, clusterApiMap map[string]resource_cluster_api.ResourceClusterApiServer) *ResourceModelApi {
	clusterClientMap := map[string]*resource_cluster_api_client.ResourceClusterApiClient{}
	if clusterApiMap == nil {
		for clusterName := range conf.Resource.ClusterMap {
			newConf := *conf
			newConf.Resource.Node.ClusterName = clusterName
			clusterClientMap[clusterName] = resource_cluster_api_client.NewResourceClusterApiClient(conf, nil)
		}
	} else {
		for clusterName := range conf.Resource.ClusterMap {
			api, ok := clusterApiMap[clusterName]
			if !ok {
				glog.Fatalf("NotFound cluster: %v", clusterName)
			}
			newConf := *conf
			newConf.Resource.Node.ClusterName = clusterName
			clusterClientMap[clusterName] = resource_cluster_api_client.NewResourceClusterApiClient(conf, &api)
		}
	}

	modelApi := ResourceModelApi{
		host:             conf.Default.Host,
		name:             "resource.model_api",
		conf:             conf,
		downTimeDuration: -1 * time.Duration(conf.Resource.AppDownTime) * time.Second,
		clusterClientMap: clusterClientMap,
		validate:         validator.New(),
	}

	return &modelApi
}

func (modelApi *ResourceModelApi) Bootstrap(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	db.AutoMigrate(&resource_model.Node{})
	db.AutoMigrate(&resource_model.Datacenter{})
	db.AutoMigrate(&resource_model.Cluster{})
	db.AutoMigrate(&resource_model.Rack{})
	db.AutoMigrate(&resource_model.PhysicalResource{})
	db.AutoMigrate(&resource_model.Compute{})
	db.AutoMigrate(&resource_model.Volume{})
	db.AutoMigrate(&resource_model.Image{})
	db.AutoMigrate(&resource_model.Loadbalancer{})
	db.AutoMigrate(&resource_model.NetworkV4{})
	db.AutoMigrate(&resource_model.NetworkV4Port{})
	glog.V(2).Info("ResourceModelApi: Complete AutoMigrate")

	if err = modelApi.bootstrapClusters(tctx); err != nil {
		return err
	}

	return nil
}

func (modelApi *ResourceModelApi) bootstrapClusters(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	for clusterName, clusterConf := range modelApi.conf.Resource.ClusterMap {
		var cluster resource_model.Cluster
		if err = db.Where("name = ?", clusterName).First(&cluster).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return err
			}

			cluster = resource_model.Cluster{
				Name: clusterName,
			}
			if err = db.Create(&cluster).Error; err != nil {
				return err
			}
			glog.V(2).Infof("Resource Cluster: Created: cluster=%v, conf=%v", clusterName, clusterConf)
		} else {
			glog.V(2).Infof("Resource Cluster: Already Exists: cluster=%v", clusterName, clusterConf)
		}
	}

	return nil
}

func (modelApi *ResourceModelApi) open(tctx *logger.TraceContext) (*gorm.DB, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = gorm.Open("mysql", modelApi.conf.Resource.Database.Connection); err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	return db, err
}
