package base_db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type IApi interface {
	Open(tctx *logger.TraceContext) (*gorm.DB, error)
	Close(tctx *logger.TraceContext, db *gorm.DB)
	GetUserWithValidatePassword(tctx *logger.TraceContext, db *gorm.DB, name string, password string) (user *base_db_model.User, code uint8, err error)
	GetUserAuthority(tctx *logger.TraceContext, db *gorm.DB, username string) (*base_spec.UserAuthority, error)
}

type Api struct {
	baseConf     *base_config.Config
	appConf      *base_config.AppConfig
	databaseConf base_config.DatabaseConfig
	secrets      []string
}

func New(baseConf *base_config.Config, appConf *base_config.AppConfig) *Api {
	api := Api{
		baseConf:     baseConf,
		appConf:      appConf,
		databaseConf: appConf.Database,
		secrets:      appConf.Auth.Secrets,
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
