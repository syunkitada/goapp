package resource_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ResourceModelApi struct {
	conf             *config.Config
	downTimeDuration time.Duration
	clusterClientMap map[string]*resource_cluster_api_client.ResourceClusterApiClient
}

func NewResourceModelApi(conf *config.Config, clusterApiMap map[string]resource_cluster_api.ResourceClusterApiServer) *ResourceModelApi {
	clusterClientMap := map[string]*resource_cluster_api_client.ResourceClusterApiClient{}
	if clusterApiMap == nil {
		for clusterName, _ := range conf.Resource.ClusterMap {
			newConf := *conf
			newConf.Resource.Cluster.Name = clusterName
			clusterClientMap[clusterName] = resource_cluster_api_client.NewResourceClusterApiClient(conf, nil)
		}
	} else {
		for clusterName, _ := range conf.Resource.ClusterMap {
			api, ok := clusterApiMap[clusterName]
			if !ok {
				glog.Fatalf("NotFound cluster: %v", clusterName)
			}
			newConf := *conf
			newConf.Resource.Cluster.Name = clusterName
			clusterClientMap[clusterName] = resource_cluster_api_client.NewResourceClusterApiClient(conf, &api)
		}
	}

	modelApi := ResourceModelApi{
		conf:             conf,
		downTimeDuration: -1 * time.Duration(conf.Resource.AppDownTime) * time.Second,
		clusterClientMap: clusterClientMap,
	}

	return &modelApi
}

func (modelApi *ResourceModelApi) Bootstrap() error {
	db, dbErr := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	db.AutoMigrate(&resource_model.Node{})
	db.AutoMigrate(&resource_model.Cluster{})
	db.AutoMigrate(&resource_model.Compute{})
	db.AutoMigrate(&resource_model.Volume{})
	db.AutoMigrate(&resource_model.Image{})
	db.AutoMigrate(&resource_model.Loadbalancer{})
	db.AutoMigrate(&resource_model.NetworkV4{})
	db.AutoMigrate(&resource_model.NetworkV4Port{})

	if err := modelApi.bootstrapClusters(); err != nil {
		return err
	}

	return nil
}

func (modelApi *ResourceModelApi) bootstrapClusters() error {
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	for clusterName, clusterConf := range modelApi.conf.Resource.ClusterMap {
		glog.Info(clusterConf)
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
		} else {
			continue
			// node.State = req.State
			// node.StateReason = req.StateReason
			// if err = db.Save(&node).Error; err != nil {
			// 	return err
			// }
		}
	}

	return nil
}
