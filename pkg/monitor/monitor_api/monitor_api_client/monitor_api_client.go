package monitor_api_client

import (
	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
)

type MonitorApiClient struct {
	*base.BaseClient
	conf        *config.Config
	localServer *monitor_api.MonitorApiServer
}

func NewMonitorApiClient(conf *config.Config) *MonitorApiClient {
	monitorClient := MonitorApiClient{
		BaseClient:  base.NewBaseClient(conf, &conf.Monitor.ApiApp.AppConfig),
		conf:        conf,
		localServer: monitor_api.NewMonitorApiServer(conf),
	}
	return &monitorClient
}

func (cli *MonitorApiClient) Status() (*monitor_api_grpc_pb.StatusReply, error) {
	var rep *monitor_api_grpc_pb.StatusReply
	var err error

	conn, err := cli.NewClientConnection()
	if err != nil {
		return rep, err
	}
	defer conn.Close()

	req := &monitor_api_grpc_pb.StatusRequest{}
	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Status(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.Status(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) GetNode(req *monitor_api_grpc_pb.GetNodeRequest) (*monitor_api_grpc_pb.GetNodeReply, error) {
	var rep *monitor_api_grpc_pb.GetNodeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetNode(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetNode(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) Report(req *monitor_api_grpc_pb.ReportRequest) (*monitor_api_grpc_pb.ReportReply, error) {
	var rep *monitor_api_grpc_pb.ReportReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Report(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.Report(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) GetHost(req *monitor_api_grpc_pb.GetHostRequest) (*monitor_api_grpc_pb.GetHostReply, error) {
	var rep *monitor_api_grpc_pb.GetHostReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetHost(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetHost(ctx, req)
	}

	return rep, err
}
