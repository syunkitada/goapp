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

func (cli *MonitorApiClient) GetIndex(req *monitor_api_grpc_pb.GetIndexRequest) (*monitor_api_grpc_pb.GetIndexReply, error) {
	var rep *monitor_api_grpc_pb.GetIndexReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetIndex(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetIndex(ctx, req)
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

func (cli *MonitorApiClient) GetUserState(req *monitor_api_grpc_pb.GetUserStateRequest) (*monitor_api_grpc_pb.GetUserStateReply, error) {
	var rep *monitor_api_grpc_pb.GetUserStateReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetUserState(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetUserState(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) GetIndexState(req *monitor_api_grpc_pb.GetIndexStateRequest) (*monitor_api_grpc_pb.GetIndexStateReply, error) {
	var rep *monitor_api_grpc_pb.GetIndexStateReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetIndexState(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetIndexState(ctx, req)
	}

	return rep, err
}

//
// IgnoreAlert
//
func (cli *MonitorApiClient) GetIgnoreAlert(req *monitor_api_grpc_pb.GetIgnoreAlertRequest) (*monitor_api_grpc_pb.GetIgnoreAlertReply, error) {
	var rep *monitor_api_grpc_pb.GetIgnoreAlertReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetIgnoreAlert(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.GetIgnoreAlert(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) CreateIgnoreAlert(req *monitor_api_grpc_pb.CreateIgnoreAlertRequest) (*monitor_api_grpc_pb.CreateIgnoreAlertReply, error) {
	var rep *monitor_api_grpc_pb.CreateIgnoreAlertReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.CreateIgnoreAlert(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.CreateIgnoreAlert(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) UpdateIgnoreAlert(req *monitor_api_grpc_pb.UpdateIgnoreAlertRequest) (*monitor_api_grpc_pb.UpdateIgnoreAlertReply, error) {
	var rep *monitor_api_grpc_pb.UpdateIgnoreAlertReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.UpdateIgnoreAlert(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.UpdateIgnoreAlert(ctx, req)
	}

	return rep, err
}

func (cli *MonitorApiClient) DeleteIgnoreAlert(req *monitor_api_grpc_pb.DeleteIgnoreAlertRequest) (*monitor_api_grpc_pb.DeleteIgnoreAlertReply, error) {
	var rep *monitor_api_grpc_pb.DeleteIgnoreAlertReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.DeleteIgnoreAlert(ctx, req)
	} else {
		grpcClient := monitor_api_grpc_pb.NewMonitorApiClient(conn)
		rep, err = grpcClient.DeleteIgnoreAlert(ctx, req)
	}

	return rep, err
}
