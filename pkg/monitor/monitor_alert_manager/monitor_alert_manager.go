package monitor_alert_manager

import (
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_alert_manager/monitor_alert_manager_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_client"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model/monitor_model_api"
)

type MonitorAlertManagerServer struct {
	base.BaseApp
	conf             *config.Config
	monitorApiClient *monitor_api_client.MonitorApiClient
	monitorModelApi  *monitor_model_api.MonitorModelApi
	role             string
}

func NewMonitorAlertManagerServer(conf *config.Config) *MonitorAlertManagerServer {
	conf.Monitor.AlertManagerApp.Name = "monitor.alert_manager"
	server := MonitorAlertManagerServer{
		BaseApp:          base.NewBaseApp(conf, &conf.Monitor.AlertManagerApp.AppConfig),
		conf:             conf,
		monitorApiClient: monitor_api_client.NewMonitorApiClient(conf),
		monitorModelApi:  monitor_model_api.NewMonitorModelApi(conf),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *MonitorAlertManagerServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	monitor_alert_manager_grpc_pb.RegisterMonitorAlertManagerServer(grpcServer, srv)
	return nil
}
