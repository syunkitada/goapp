package resolver

import (
	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Resolver struct {
	dbApi   *db_api.Api
	appConf base_config.AppConfig
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Resolver {
	return &Resolver{
		appConf: mainConf.Authproxy.App,
		dbApi:   db_api.New(baseConf, mainConf),
	}
}
