package monitor_agent

import (
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/metric_plugins"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/metric_plugins/system"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/monitor_agent_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_client"
)

type MonitorAgentServer struct {
	base.BaseApp
	conf                  *config.Config
	monitorApiClient      *monitor_api_client.MonitorApiClient
	role                  string
	reportIndex           string
	reportProject         string
	reportSpan            int
	reportCount           int
	metricReaders         []metric_plugins.MetricReader
	logReaderMap          map[string]*LogReader
	logReaderRefreshSpan  int
	logReaderRefreshCount int
}

func NewMonitorAgentServer(conf *config.Config) *MonitorAgentServer {
	metricReaders := []metric_plugins.MetricReader{}
	if conf.Monitor.AgentApp.Metrics.System.Enable {
		metricReaders = append(metricReaders, system.NewSystemMetricReader(&conf.Monitor.AgentApp.Metrics.System))
	}

	conf.Monitor.AgentApp.Name = "monitor.agent"
	server := MonitorAgentServer{
		BaseApp:               base.NewBaseApp(conf, &conf.Monitor.AgentApp.AppConfig),
		conf:                  conf,
		monitorApiClient:      monitor_api_client.NewMonitorApiClient(conf),
		reportIndex:           conf.Monitor.AgentApp.ReportIndex,
		reportProject:         conf.Monitor.AgentApp.ReportProject,
		reportSpan:            conf.Monitor.AgentApp.ReportSpan,
		reportCount:           0,
		metricReaders:         metricReaders,
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
