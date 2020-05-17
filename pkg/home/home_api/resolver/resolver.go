package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_resolver"

	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/db_api"
)

type Resolver struct {
	*base_resolver.Resolver
	dbApi   *db_api.Api
	appConf base_config.AppConfig
}

func New(baseConf *base_config.Config, mainConf *config.Config, dbApi *db_api.Api) *Resolver {
	return &Resolver{
		Resolver: base_resolver.New(baseConf, &mainConf.Home.Api, dbApi),
		appConf:  mainConf.Home.Api,
		dbApi:    dbApi,
	}
}
