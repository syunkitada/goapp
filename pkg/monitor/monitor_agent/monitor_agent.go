package monitor_agent

import (
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/monitor_agent_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_proxy/monitor_proxy_client"
)

type MonitorAgentServer struct {
	base.BaseApp
	conf               *config.Config
	monitorProxyClient *monitor_proxy_client.MonitorProxyClient
	role               string
}

func NewMonitorAgentServer(conf *config.Config) *MonitorAgentServer {
	conf.Monitor.AgentApp.Name = "monitor.agent"
	server := MonitorAgentServer{
		BaseApp:            base.NewBaseApp(conf, &conf.Monitor.AgentApp),
		conf:               conf,
		monitorProxyClient: monitor_proxy_client.NewMonitorProxyClient(conf),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *MonitorAgentServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_agent_grpc_pb.RegisterMonitorAgentServer(grpcServer, srv)
	return nil
}
