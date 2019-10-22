package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_resolver"

	"github.com/syunkitada/goapp/pkg/resource/config"
)

type Resolver struct {
	*base_resolver.Resolver
	appConf base_config.AppConfig
}

func New(baseConf *base_config.Config, clusterConf *config.ResourceClusterConfig) *Resolver {
	return &Resolver{
		Resolver: base_resolver.New(baseConf, &clusterConf.Api, nil),
		appConf:  clusterConf.Api,
	}
}
