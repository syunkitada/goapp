package resource_cluster_compute_agent

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"google.golang.org/grpc"
)

type ResourceClusterComputeAgentServer struct {
	base.BaseApp
	conf *config.Config
}

func New(conf *config.Config) *ResourceClusterComputeAgentServer {
	conf.Resource.Node.ComputeAgent.Name = "resource.cluster.agent"
	fmt.Println(conf.Resource.Node.ComputeAgent.HttpListen)
	server := ResourceClusterComputeAgentServer{
		BaseApp: base.NewBaseApp(conf, &conf.Resource.Node.ComputeAgent),
		conf:    conf,
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterComputeAgentServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	return nil
}
