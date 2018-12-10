package resource_cluster_api

import (
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model/resource_cluster_model_api"
)

type ResourceClusterApiServer struct {
	base.BaseApp
	conf                    *config.Config
	cluster                 *config.ResourceClusterConfig
	resourceClusterModelApi *resource_cluster_model_api.ResourceClusterModelApi
}

func NewResourceClusterApiServer(conf *config.Config) *ResourceClusterApiServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
	}

	cluster.ApiApp.Name = "resource.cluster.api"
	server := ResourceClusterApiServer{
		BaseApp:                 base.NewBaseApp(conf, &cluster.ApiApp),
		conf:                    conf,
		cluster:                 &cluster,
		resourceClusterModelApi: resource_cluster_model_api.NewResourceClusterModelApi(conf),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_cluster_api_grpc_pb.RegisterResourceClusterApiServer(grpcServer, srv)
	return nil
}

func (srv *ResourceClusterApiServer) Status(ctx context.Context, statusRequest *resource_cluster_api_grpc_pb.StatusRequest) (*resource_cluster_api_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_cluster_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

func (srv *ResourceClusterApiServer) GetNode(ctx context.Context, req *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	var err error
	var rep *resource_cluster_api_grpc_pb.GetNodeReply
	glog.Info("DEBUGlalala")
	if rep, err = srv.resourceClusterModelApi.GetNode(req); err != nil {
		glog.Error(err)
	}
	return rep, err
}

func (srv *ResourceClusterApiServer) UpdateNode(ctx context.Context, req *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_cluster_api_grpc_pb.UpdateNodeReply
	var err error
	glog.Infof("UpdateNode: %v, %v", req.Name, req.Kind)
	rep, err = srv.resourceClusterModelApi.UpdateNode(req)
	return rep, err
}

//
// Compute
//
func (srv *ResourceClusterApiServer) GetCompute(ctx context.Context, req *resource_cluster_api_grpc_pb.GetComputeRequest) (*resource_cluster_api_grpc_pb.GetComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.GetComputeReply
	var err error
	rep, err = srv.resourceClusterModelApi.GetCompute(req)
	glog.Infof("Completed GetCompute: %v", err)
	return rep, err
}

func (srv *ResourceClusterApiServer) CreateCompute(ctx context.Context, req *resource_cluster_api_grpc_pb.CreateComputeRequest) (*resource_cluster_api_grpc_pb.CreateComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.CreateComputeReply
	var err error
	rep, err = srv.resourceClusterModelApi.CreateCompute(req)
	glog.Infof("Completed CreateCompute: %v", err)
	return rep, err
}

func (srv *ResourceClusterApiServer) UpdateCompute(ctx context.Context, req *resource_cluster_api_grpc_pb.UpdateComputeRequest) (*resource_cluster_api_grpc_pb.UpdateComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.UpdateComputeReply
	var err error
	rep, err = srv.resourceClusterModelApi.UpdateCompute(req)
	glog.Infof("Completed UpdateCompute: %v", err)
	return rep, err
}

func (srv *ResourceClusterApiServer) DeleteCompute(ctx context.Context, req *resource_cluster_api_grpc_pb.DeleteComputeRequest) (*resource_cluster_api_grpc_pb.DeleteComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.DeleteComputeReply
	var err error
	rep, err = srv.resourceClusterModelApi.DeleteCompute(req)
	glog.Infof("Completed DeleteCompute: %v", err)
	return rep, err
}
