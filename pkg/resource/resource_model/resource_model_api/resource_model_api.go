package resource_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/golang/glog"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ResourceModelApi struct {
	conf             *config.Config
	downTimeDuration time.Duration
}

func NewResourceModelApi(conf *config.Config) *ResourceModelApi {
	modelApi := ResourceModelApi{
		conf:             conf,
		downTimeDuration: -1 * time.Duration(conf.Resource.AppDownTime) * time.Second,
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
	db.AutoMigrate(&resource_model.Region{})
	db.AutoMigrate(&resource_model.ComputeResource{})
	db.AutoMigrate(&resource_model.VolumeResource{})
	db.AutoMigrate(&resource_model.ImageResource{})
	db.AutoMigrate(&resource_model.LoadbalancerResource{})
	db.AutoMigrate(&resource_model.NetworkV4{})
	db.AutoMigrate(&resource_model.NetworkV4Port{})

	return nil
}
