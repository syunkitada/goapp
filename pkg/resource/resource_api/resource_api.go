package resource_api

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceApiServer struct {
	base.BaseApp
	conf              *config.Config
	resourceModelApi  *resource_model_api.ResourceModelApi
	resourceApiClient *resource_api_client.ResourceApiClient
}

func NewResourceApiServer(conf *config.Config) *ResourceApiServer {
	server := ResourceApiServer{
		BaseApp:           base.NewBaseApp(conf, conf.Resource.ApiApp),
		conf:              conf,
		resourceModelApi:  resource_model_api.NewResourceModelApi(conf),
		resourceApiClient: resource_api_client.NewResourceApiClient(conf),
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *ResourceApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_api_grpc_pb.RegisterResourceApiServer(grpcServer, srv)
	return nil
}

func (srv *ResourceApiServer) Status(ctx context.Context, statusRequest *resource_api_grpc_pb.StatusRequest) (*resource_api_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

func (srv *ResourceApiServer) GetNode(ctx context.Context, req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	glog.Info("GetNode")
	var err error
	var rep *resource_api_grpc_pb.GetNodeReply
	if rep, err = srv.resourceModelApi.GetNode(req); err != nil {
		glog.Error(err)
	}
	return rep, err
}

func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	glog.Info("UpdateNode")
	if err = srv.resourceModelApi.UpdateNode(req); err != nil {
		glog.Error(err)
	}
	return &resource_api_grpc_pb.UpdateNodeReply{}, err
}
