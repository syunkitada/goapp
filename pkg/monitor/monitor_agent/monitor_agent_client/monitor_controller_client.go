package monitor_controller_client

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_controller/monitor_controller_grpc_pb"
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
		CaFilePath:         conf.Path(conf.Monitor.Grpc.CaFile),
		ServerHostOverride: conf.Monitor.Grpc.ServerHostOverride,
		Targets:            conf.Monitor.Grpc.Targets,
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

func (monitorClient *MonitorClient) Status() (*monitor_controller_grpc_pb.StatusReply, error) {
	conn, connErr := monitorClient.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		return nil, connErr
	}

	client := monitor_controller_grpc_pb.NewMonitorClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	statusResponse, err := client.Status(ctx, &monitor_controller_grpc_pb.StatusRequest{})
	if err != nil {
		return nil, err
	}

	return statusResponse, nil
}
