package resource_cluster_controller

import (
	"fmt"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model/resource_cluster_model_api"
)

type ResourceClusterControllerServer struct {
	base.BaseApp
	conf                     *config.Config
	cluster                  *config.ResourceClusterConfig
	resourceClusterModelApi  *resource_cluster_model_api.ResourceClusterModelApi
	resourceClusterApiClient *resource_cluster_api_client.ResourceClusterApiClient
	role                     string
}

func NewResourceClusterControllerServer(conf *config.Config) *ResourceClusterControllerServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
	}

	server := ResourceClusterControllerServer{
		BaseApp: base.NewBaseApp(conf, &cluster.ControllerApp),
		conf:    conf,
		resourceClusterModelApi:  resource_cluster_model_api.NewResourceClusterModelApi(conf),
		resourceClusterApiClient: resource_cluster_api_client.NewResourceClusterApiClient(conf, nil),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterControllerServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_cluster_controller_grpc_pb.RegisterResourceClusterControllerServer(grpcServer, srv)
	return nil
}
