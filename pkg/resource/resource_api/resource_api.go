package resource_api

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceApiServer struct {
	base.BaseApp
	conf             *config.Config
	resourceModelApi *resource_model_api.ResourceModelApi
}

func NewResourceApiServer(conf *config.Config) *ResourceApiServer {
	conf.Resource.ApiApp.Name = "resource.api"
	server := ResourceApiServer{
		BaseApp:          base.NewBaseApp(conf, &conf.Resource.ApiApp),
		conf:             conf,
		resourceModelApi: resource_model_api.NewResourceModelApi(conf, nil),
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *ResourceApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_api_grpc_pb.RegisterResourceApiServer(grpcServer, srv)
	return nil
}

func (srv *ResourceApiServer) Status(ctx context.Context, statusRequest *resource_api_grpc_pb.StatusRequest) (*resource_api_grpc_pb.StatusReply, error) {
	return &resource_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

//
// Action
//
func (srv *ResourceApiServer) Action(ctx context.Context, req *resource_api_grpc_pb.ActionRequest) (*resource_api_grpc_pb.ActionReply, error) {
	var err error
	rep := &resource_api_grpc_pb.ActionReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	switch req.Tctx.ActionName {
	case "CreatePhysicalResource", "CreateVirtualResource":
		srv.resourceModelApi.Create(tctx, req, rep)
	case "GetPhysicalIndex":
		srv.resourceModelApi.GetPhysicalIndex(tctx, req, rep)
	case "GetCluster":
		srv.resourceModelApi.GetCluster(tctx, req, rep)
	case "GetNode":
		srv.resourceModelApi.GetNode(tctx, req, rep)
	case "GetCompute":
		srv.resourceModelApi.GetCompute(tctx, req, rep)
	case "CreateCompute":
		srv.resourceModelApi.CreateCompute(tctx, req, rep)
	case "UpdateCompute":
		srv.resourceModelApi.UpdateCompute(tctx, req, rep)
	case "DeleteCompute":
		srv.resourceModelApi.DeleteCompute(tctx, req, rep)
	case "GetNetwork":
		srv.resourceModelApi.GetNetworkV4(tctx, req, rep)
	case "CreateNetwork":
		srv.resourceModelApi.CreateNetworkV4(tctx, req, rep)
	case "UpdateNetwork":
		srv.resourceModelApi.UpdateNetworkV4(tctx, req, rep)
	case "DeleteNetwork":
		srv.resourceModelApi.DeleteNetworkV4(tctx, req, rep)
	case "GetImage":
		srv.resourceModelApi.GetImage(tctx, req, rep)
	case "CreateImage":
		srv.resourceModelApi.CreateImage(tctx, req, rep)
	case "UpdateImage":
		srv.resourceModelApi.UpdateImage(tctx, req, rep)
	case "DeleteImage":
		srv.resourceModelApi.DeleteImage(tctx, req, rep)
	default:
		rep.Tctx.Err = fmt.Sprintf("InvalidAction: %v", req.Tctx.ActionName)
		rep.Tctx.StatusCode = codes.ClientNotFound
	}

	return rep, nil
}

//
// Node
//
func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	rep := &resource_api_grpc_pb.UpdateNodeReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	srv.resourceModelApi.UpdateNode(tctx, req, rep)
	return rep, nil
}
