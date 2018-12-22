package monitor_proxy

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model/monitor_model_api"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_proxy/monitor_proxy_grpc_pb"
)

type MonitorProxyServer struct {
	base.BaseApp
	conf            *config.Config
	monitorModelApi *monitor_model_api.MonitorModelApi
}

func NewMonitorProxyServer(conf *config.Config) *MonitorProxyServer {
	conf.Monitor.ProxyApp.Name = "monitor.proxy"
	server := MonitorProxyServer{
		BaseApp:         base.NewBaseApp(conf, &conf.Monitor.ProxyApp),
		conf:            conf,
		monitorModelApi: monitor_model_api.NewMonitorModelApi(conf),
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *MonitorProxyServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_proxy_grpc_pb.RegisterMonitorProxyServer(grpcServer, srv)
	return nil
}

func (srv *MonitorProxyServer) Status(ctx context.Context, statusRequest *monitor_proxy_grpc_pb.StatusRequest) (*monitor_proxy_grpc_pb.StatusReply, error) {
	return &monitor_proxy_grpc_pb.StatusReply{Msg: "Status"}, nil
}

//
// Node
//
func (srv *MonitorProxyServer) GetNode(ctx context.Context, req *monitor_proxy_grpc_pb.GetNodeRequest) (*monitor_proxy_grpc_pb.GetNodeReply, error) {
	startTime, clientIp := logger.StartGrpcTrace(req.TraceId, srv.Host, srv.Name, ctx)
	rep := srv.monitorModelApi.GetNode(req)
	logger.EndGrpcTrace(req.TraceId, srv.Host, srv.Name, startTime, clientIp, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *MonitorProxyServer) UpdateNode(ctx context.Context, req *monitor_proxy_grpc_pb.UpdateNodeRequest) (*monitor_proxy_grpc_pb.UpdateNodeReply, error) {
	startTime, clientIp := logger.StartGrpcTrace(req.TraceId, srv.Host, srv.Name, ctx)
	rep := srv.monitorModelApi.UpdateNode(req)
	logger.EndGrpcTrace(req.TraceId, srv.Host, srv.Name, startTime, clientIp, rep.StatusCode, rep.Err)
	return rep, nil
}
