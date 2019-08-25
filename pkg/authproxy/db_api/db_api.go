package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Api struct {
	databaseConf base_config.DatabaseConfig
	baseConf     *base_config.Config
	mainConf     *config.Config
	appConf      base_config.AppConfig
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Api {
	api := Api{
		databaseConf: mainConf.Authproxy.App.Database,
		baseConf:     baseConf,
		mainConf:     mainConf,
		appConf:      mainConf.Authproxy.App,
	}

	return &api
}

func (api *Api) Open(tctx *logger.TraceContext) (*gorm.DB, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = gorm.Open("mysql", api.databaseConf.Connection)
	if err != nil {
		return nil, err
	}
	db.LogMode(api.baseConf.EnableDatabaseLog)

	return db, nil
}

func (api *Api) Close(tctx *logger.TraceContext, db *gorm.DB) {
	if err := db.Close(); err != nil {
		logger.Error(tctx, err)
	}
}

func (api *Api) Rollback(tctx *logger.TraceContext, tx *gorm.DB, err error) {
	if err != nil {
		logger.Error(tctx, err)
		if err = tx.Rollback().Error; err != nil {
			logger.Error(tctx, err)
		}
	}
}
