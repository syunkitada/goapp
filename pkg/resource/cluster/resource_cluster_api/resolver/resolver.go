package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_resolver"

	"github.com/syunkitada/goapp/pkg/resource/cluster/db_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/tsdb_api"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

type Resolver struct {
	*base_resolver.Resolver
	dbApi   *db_api.Api
	tsdbApi *tsdb_api.Api
	appConf base_config.AppConfig
}

func New(baseConf *base_config.Config, clusterConf *config.ResourceClusterConfig, dbApi *db_api.Api) *Resolver {
	return &Resolver{
		Resolver: base_resolver.New(baseConf, &clusterConf.Api, dbApi),
		appConf:  clusterConf.Api,
		dbApi:    dbApi,
		tsdbApi:  tsdb_api.New(baseConf, clusterConf),
	}
}
