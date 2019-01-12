package monitor_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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

func (modelApi *MonitorModelApi) Bootstrap(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer db.Close()

	db.AutoMigrate(&monitor_model.Node{})

	return nil
}

func (modelApi *MonitorModelApi) open(tctx *logger.TraceContext) (*gorm.DB, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	db, err := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	return db, nil
}
