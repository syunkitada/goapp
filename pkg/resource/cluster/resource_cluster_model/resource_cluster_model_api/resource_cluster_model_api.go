package resource_cluster_model_api

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

type ResourceClusterModelApi struct {
	conf    *config.Config
	cluster *config.ResourceClusterConfig
}

func NewResourceClusterModelApi(conf *config.Config) *ResourceClusterModelApi {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Cluster.Name]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Cluster.Name))
	}

	modelApi := ResourceClusterModelApi{
		conf:    conf,
		cluster: cluster,
	}

	return &modelApi
}

func (modelApi *ResourceClusterModelApi) Bootstrap() error {
	db, dbErr := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	db.AutoMigrate(&resource_cluster_model.Node{})
	db.AutoMigrate(&resource_cluster_model.Region{})
	db.AutoMigrate(&resource_cluster_model.ComputeResource{})
	db.AutoMigrate(&resource_cluster_model.VolumeResource{})
	db.AutoMigrate(&resource_cluster_model.ImageResource{})
	db.AutoMigrate(&resource_cluster_model.LoadbalancerResource{})

	return nil
}
