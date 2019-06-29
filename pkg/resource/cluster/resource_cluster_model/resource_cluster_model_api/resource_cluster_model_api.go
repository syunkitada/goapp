package resource_cluster_model_api

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ResourceClusterModelApi struct {
	conf             *config.Config
	cluster          *config.ResourceClusterConfig
	downTimeDuration time.Duration
	validate         *validator.Validate
}

func NewResourceClusterModelApi(conf *config.Config) *ResourceClusterModelApi {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
	}

	modelApi := ResourceClusterModelApi{
		conf:             conf,
		cluster:          &cluster,
		downTimeDuration: -1 * time.Duration(conf.Resource.AppDownTime) * time.Second,
		validate:         validator.New(),
	}

	return &modelApi
}

func (modelApi *ResourceClusterModelApi) Bootstrap(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer modelApi.close(tctx, db)

	db.AutoMigrate(&resource_model.Node{})
	db.AutoMigrate(&resource_model.Compute{})
	db.AutoMigrate(&resource_model.ComputeAssignment{})
	return nil
}

func (modelApi *ResourceClusterModelApi) open(tctx *logger.TraceContext) (*gorm.DB, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = gorm.Open("mysql", modelApi.cluster.Database.Connection); err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	return db, err
}

func (modelApi *ResourceClusterModelApi) close(tctx *logger.TraceContext, db *gorm.DB) {
	if err := db.Close(); err != nil {
		logger.Error(tctx, err)
	}
}
