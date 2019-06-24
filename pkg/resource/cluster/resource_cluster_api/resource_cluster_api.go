package resource_cluster_api

import (
	"fmt"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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

//
// Action
//
func (srv *ResourceClusterApiServer) Action(ctx context.Context, req *resource_cluster_api_grpc_pb.ActionRequest) (*resource_cluster_api_grpc_pb.ActionReply, error) {
	var err error
	rep := &resource_cluster_api_grpc_pb.ActionReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	switch req.Tctx.ActionName {
	case "GetNode":
		srv.resourceClusterModelApi.GetNode(tctx, req, rep)
	default:
		rep.Tctx.Err = fmt.Sprintf("InvalidAction: %v", req.Tctx.ActionName)
		rep.Tctx.StatusCode = codes.ClientNotFound
	}

	return rep, nil
}

func (srv *ResourceClusterApiServer) UpdateNode(ctx context.Context, req *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	rep := &resource_cluster_api_grpc_pb.UpdateNodeReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	srv.resourceClusterModelApi.UpdateNode(tctx, req, rep)
	return rep, nil
}
