package base_resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
)

type Resolver struct {
	baseConf *base_config.Config
	appConf  *base_config.AppConfig
	dbApi    base_db_api.IApi
}

func New(baseConf *base_config.Config, appConf *base_config.AppConfig, dbApi base_db_api.IApi) *Resolver {
	return &Resolver{
		baseConf: baseConf,
		appConf:  appConf,
		dbApi:    dbApi,
	}
}
