package monitor_api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model/monitor_model_api"
)

type MonitorApiServer struct {
	base.BaseApp
	conf            *config.Config
	monitorModelApi *monitor_model_api.MonitorModelApi
}

func NewMonitorApiServer(conf *config.Config) *MonitorApiServer {
	conf.Monitor.ApiApp.Name = "monitor.api"
	server := MonitorApiServer{
		BaseApp:         base.NewBaseApp(conf, &conf.Monitor.ApiApp),
		conf:            conf,
		monitorModelApi: monitor_model_api.NewMonitorModelApi(conf),
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *MonitorApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_api_grpc_pb.RegisterMonitorApiServer(grpcServer, srv)
	return nil
}

func (srv *MonitorApiServer) Status(ctx context.Context, statusRequest *monitor_api_grpc_pb.StatusRequest) (*monitor_api_grpc_pb.StatusReply, error) {
	return &monitor_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

//
// Node
//
func (srv *MonitorApiServer) GetNode(ctx context.Context, req *monitor_api_grpc_pb.GetNodeRequest) (*monitor_api_grpc_pb.GetNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.GetNode(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorApiServer) UpdateNode(ctx context.Context, req *monitor_api_grpc_pb.UpdateNodeRequest) (*monitor_api_grpc_pb.UpdateNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.monitorModelApi.UpdateNode(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}
