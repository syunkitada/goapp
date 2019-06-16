package resource_api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
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

//
// Action
//
func (srv *ResourceApiServer) PhysicalAction(ctx context.Context,
	req *authproxy_grpc_pb.ActionRequest) (*authproxy_grpc_pb.ActionReply, error) {
	var err error
	rep := &authproxy_grpc_pb.ActionReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	srv.resourceModelApi.PhysicalAction(tctx, req, rep)
	return rep, nil
}

func (srv *ResourceApiServer) VirtualAction(ctx context.Context,
	req *authproxy_grpc_pb.ActionRequest) (*authproxy_grpc_pb.ActionReply, error) {
	var err error
	rep := &authproxy_grpc_pb.ActionReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	srv.resourceModelApi.VirtualAction(tctx, req, rep)
	return rep, nil
}
