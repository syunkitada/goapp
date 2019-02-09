package monitor_controller_client

import (
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/monitor_agent_grpc_pb"
)

type MonitorClient struct {
	Conf               *config.Config
	CaFilePath         string
	ServerHostOverride string
	Targets            []string
}

func NewMonitorClient(conf *config.Config) *MonitorClient {
	monitorClient := MonitorClient{
		Conf:               conf,
		CaFilePath:         conf.Path(conf.Monitor.ApiApp.CaFile),
		ServerHostOverride: conf.Monitor.ApiApp.ServerHostOverride,
		Targets:            conf.Monitor.ApiApp.Targets,
	}
	return &monitorClient
}

func (monitorClient *MonitorClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range monitorClient.Targets {
		creds, credsErr := credentials.NewClientTLSFromFile(monitorClient.CaFilePath, monitorClient.ServerHostOverride)
		if credsErr != nil {
			continue
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(target, opts...)
		if err != nil {
			continue
		}

		return conn, nil
	}

	return nil, errors.New("Failed NewGrpcConnection")
}

func (monitorClient *MonitorClient) Status() (*monitor_agent_grpc_pb.StatusReply, error) {
	return nil, nil
}
