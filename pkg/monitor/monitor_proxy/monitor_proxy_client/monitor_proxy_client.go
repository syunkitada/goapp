package monitor_proxy_client

import (
	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_proxy"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_proxy/monitor_proxy_grpc_pb"
)

type MonitorProxyClient struct {
	*base.BaseClient
	conf        *config.Config
	localServer *monitor_proxy.MonitorProxyServer
}

func NewMonitorProxyClient(conf *config.Config) *MonitorProxyClient {
	monitorClient := MonitorProxyClient{
		BaseClient:  base.NewBaseClient(conf, &conf.Monitor.ProxyApp),
		conf:        conf,
		localServer: monitor_proxy.NewMonitorProxyServer(conf),
	}
	return &monitorClient
}

func (cli *MonitorProxyClient) Status() (*monitor_proxy_grpc_pb.StatusReply, error) {
	var rep *monitor_proxy_grpc_pb.StatusReply
	var err error

	conn, err := cli.NewClientConnection()
	if err != nil {
		return rep, err
	}
	defer conn.Close()

	req := &monitor_proxy_grpc_pb.StatusRequest{}
	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Status(ctx, req)
	} else {
		grpcClient := monitor_proxy_grpc_pb.NewMonitorProxyClient(conn)
		rep, err = grpcClient.Status(ctx, req)
	}

	return rep, err
}
