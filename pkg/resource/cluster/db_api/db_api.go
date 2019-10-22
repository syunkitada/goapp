package db_api

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

type Api struct {
	*base_db_api.Api
	databaseConf base_config.DatabaseConfig
	baseConf     *base_config.Config
	clusterConf  *config.ResourceClusterConfig
	appConf      base_config.AppConfig
}

func New(baseConf *base_config.Config, clusterConf *config.ResourceClusterConfig) *Api {
	api := Api{
		Api:          base_db_api.New(baseConf, &clusterConf.Api, genpkg.ApiQueryMap),
		databaseConf: clusterConf.Api.Database,
		baseConf:     baseConf,
		clusterConf:  clusterConf,
		appConf:      clusterConf.Api,
	}

	return &api
}
