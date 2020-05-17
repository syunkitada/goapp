package db_api

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/home_api/spec/genpkg"
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
		Api:          base_db_api.New(baseConf, &mainConf.Home.Api, genpkg.ApiQueryMap),
		databaseConf: mainConf.Home.Api.Database,
		baseConf:     baseConf,
		mainConf:     mainConf,
		appConf:      mainConf.Home.Api,
	}

	return &api
}
