package resource_controller

import (
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/resource_controller_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceControllerServer struct {
	base.BaseApp
	conf              *config.Config
	resourceApiClient *resource_api_client.ResourceApiClient
	role              string
	resourceModelApi  *resource_model_api.ResourceModelApi
}

func NewResourceControllerServer(conf *config.Config) *ResourceControllerServer {
	server := ResourceControllerServer{
		BaseApp:           base.NewBaseApp(conf, &conf.Resource.ControllerApp),
		conf:              conf,
		resourceApiClient: resource_api_client.NewResourceApiClient(conf),
		resourceModelApi:  resource_model_api.NewResourceModelApi(conf, nil),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceControllerServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_controller_grpc_pb.RegisterResourceControllerServer(grpcServer, srv)
	return nil
}