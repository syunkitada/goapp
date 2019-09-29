package db_api

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
)

type Api struct {
	*base_db_api.Api
	databaseConf base_config.DatabaseConfig
	baseConf     *base_config.Config
	mainConf     *config.Config
	appConf      base_config.AppConfig
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Api {
	api := Api{
		Api:          base_db_api.New(baseConf, &mainConf.Authproxy.Api, genpkg.ApiQueryMap),
		databaseConf: mainConf.Authproxy.Api.Database,
		baseConf:     baseConf,
		mainConf:     mainConf,
		appConf:      mainConf.Authproxy.Api,
	}

	return &api
}
