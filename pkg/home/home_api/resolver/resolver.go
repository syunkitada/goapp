package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_resolver"

	authproxy_db_api "github.com/syunkitada/goapp/pkg/authproxy/db_api"

	"github.com/syunkitada/goapp/pkg/home/config"
	"github.com/syunkitada/goapp/pkg/home/db_api"
)

type Resolver struct {
	*base_resolver.Resolver
	appConf        base_config.AppConfig
	dbApi          *db_api.Api
	authproxyDbApi *authproxy_db_api.Api
}

func New(baseConf *base_config.Config, mainConf *config.Config, dbApi *db_api.Api, authproxyDbApi *authproxy_db_api.Api) *Resolver {
	return &Resolver{
		Resolver:       base_resolver.New(baseConf, &mainConf.Home.Api, dbApi),
		appConf:        mainConf.Home.Api,
		dbApi:          dbApi,
		authproxyDbApi: authproxyDbApi,
	}
}
