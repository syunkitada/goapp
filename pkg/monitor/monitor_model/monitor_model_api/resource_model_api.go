package monitor_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

type MonitorModelApi struct {
	host             string
	name             string
	conf             *config.Config
	downTimeDuration time.Duration
	validate         *validator.Validate
}

func NewMonitorModelApi(conf *config.Config) *MonitorModelApi {
	modelApi := MonitorModelApi{
		host:             conf.Default.Host,
		name:             "monitor.model_api",
		conf:             conf,
		downTimeDuration: -1 * time.Duration(conf.Monitor.AppDownTime) * time.Second,
		validate:         validator.New(),
	}

	return &modelApi
}

func (modelApi *MonitorModelApi) Bootstrap() error {
	glog.V(2).Info("MonitorModelApi: Complete AutoMigrate")
	db, dbErr := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
	defer db.Close()
	if dbErr != nil {
		return dbErr
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	db.AutoMigrate(&monitor_model.Node{})
	glog.V(2).Info("MonitorModelApi: Complete AutoMigrate")

	return nil
}
