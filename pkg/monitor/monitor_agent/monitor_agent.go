package monitor_agent

import (
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/monitor_agent_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_client"
)

type MonitorAgentServer struct {
	base.BaseApp
	conf                  *config.Config
	monitorApiClient      *monitor_api_client.MonitorApiClient
	role                  string
	reportIndex           string
	flushSpan             int
	flushCount            int
	logReaderMap          map[string]*LogReader
	logReaderRefreshSpan  int
	logReaderRefreshCount int
}

func NewMonitorAgentServer(conf *config.Config) *MonitorAgentServer {
	conf.Monitor.AgentApp.Name = "monitor.agent"
	server := MonitorAgentServer{
		BaseApp:               base.NewBaseApp(conf, &conf.Monitor.AgentApp.AppConfig),
		conf:                  conf,
		monitorApiClient:      monitor_api_client.NewMonitorApiClient(conf),
		reportIndex:           conf.Monitor.AgentApp.Index,
		flushSpan:             conf.Monitor.AgentApp.FlushSpan,
		flushCount:            0,
		logReaderMap:          map[string]*LogReader{},
		logReaderRefreshSpan:  conf.Monitor.AgentApp.LogReaderRefreshSpan,
		logReaderRefreshCount: 0,
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *MonitorAgentServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_agent_grpc_pb.RegisterMonitorAgentServer(grpcServer, srv)
	return nil
}
