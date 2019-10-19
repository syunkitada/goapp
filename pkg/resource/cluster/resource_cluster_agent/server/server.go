package server

import (
	"github.com/syunkitada/goapp/pkg/base/base_app"
	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/resolver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/spec/genpkg"
	resource_cluster_api "github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/spec/genpkg"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/consts"
)

type Server struct {
	base_app.BaseApp
	baseConf     *base_config.Config
	clusterConf  *config.ResourceClusterConfig
	queryHandler *genpkg.QueryHandler
	apiClient    *resource_cluster_api.Client
}

func New(baseConf *base_config.Config, mainConf *config.Config) *Server {
	clusterConf, ok := mainConf.Resource.ClusterMap[mainConf.Resource.ClusterName]
	if !ok {
		logger.StdoutFatalf("cluster config is not found: cluster=%s", mainConf.Resource.ClusterName)
	}

	clusterConf.Agent.Name = consts.KindResourceClusterAgent
	resolver := resolver.New(baseConf, &clusterConf)
	queryHandler := genpkg.NewQueryHandler(baseConf, &clusterConf.Agent, resolver)
	baseApp := base_app.New(baseConf, &clusterConf.Agent, nil, queryHandler)
	apiClient := resource_cluster_api.NewClient(&clusterConf.Agent.RootClient)

	srv := &Server{
		BaseApp:      baseApp,
		baseConf:     baseConf,
		clusterConf:  &clusterConf,
		queryHandler: queryHandler,
		apiClient:    apiClient,
	}
	srv.SetDriver(srv)
	return srv
}
