package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"

	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_api"
)

type Resolver struct {
	dbApi   *db_api.Api
	appConf base_config.AppConfig
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Resolver {
	return &Resolver{
		appConf: mainConf.Resource.App,
		dbApi:   db_api.New(baseConf, mainConf),
	}
}
