package resolver

import (
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_resolver"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers"
	"github.com/syunkitada/goapp/pkg/resource/config"
)

type Resolver struct {
	*base_resolver.Resolver
	appConf       base_config.AppConfig
	computeDriver compute_drivers.ComputeDriver
}

func New(baseConf *base_config.Config, clusterConf *config.ResourceClusterConfig) *Resolver {
	return &Resolver{
		Resolver: base_resolver.New(baseConf, &clusterConf.Api, nil),
		appConf:  clusterConf.Api,
	}
}

func (resolver *Resolver) SetComputeDriver(computeDriver compute_drivers.ComputeDriver) {
	resolver.computeDriver = computeDriver
}
