package resource_cluster_controller

import (
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb"
)

type ResourceClusterControllerServer struct {
	base.BaseApp
	Conf                     *config.Config
	cluster                  *config.ResourceClusterConfig
	resourceClusterApiClient *resource_cluster_api_client.ResourceClusterApiClient
}

func NewResourceClusterControllerServer(conf *config.Config) *ResourceClusterControllerServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Cluster.Name]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Cluster.Name))
	}

	server := ResourceClusterControllerServer{
		BaseApp: base.NewBaseApp(conf, &cluster.ApiApp),
		Conf:    conf,
		resourceClusterApiClient: resource_cluster_api_client.NewResourceClusterApiClient(conf),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterControllerServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_cluster_controller_grpc_pb.RegisterResourceClusterControllerServer(grpcServer, srv)
	return nil
}
