package db_api

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec/genpkg"
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
		Api:          base_db_api.New(baseConf, &mainConf.Resource.App, genpkg.ApiQueryMap),
		databaseConf: mainConf.Resource.App.Database,
		baseConf:     baseConf,
		mainConf:     mainConf,
		appConf:      mainConf.Resource.App,
	}

	return &api
}
