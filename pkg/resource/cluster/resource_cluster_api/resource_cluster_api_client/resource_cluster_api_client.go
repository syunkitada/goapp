package resource_cluster_api_client

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
)

type ResourceClusterApiClient struct {
	*base.BaseClient
	conf        *config.Config
	cluster     *config.ResourceClusterConfig
	localServer *resource_cluster_api.ResourceClusterApiServer
}

func NewResourceClusterApiClient(conf *config.Config, localServer *resource_cluster_api.ResourceClusterApiServer) *ResourceClusterApiClient {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
	}

	resourceClient := ResourceClusterApiClient{
		BaseClient:  base.NewBaseClient(conf, &cluster.ApiApp),
		conf:        conf,
		cluster:     &cluster,
		localServer: localServer,
	}
	return &resourceClient
}

func (cli *ResourceClusterApiClient) Status() (*resource_cluster_api_grpc_pb.StatusReply, error) {
	var rep *resource_cluster_api_grpc_pb.StatusReply
	var err error

	conn, err := cli.NewClientConnection()
	if err != nil {
		return rep, err
	}
	defer func() { err = conn.Close() }()

	req := &resource_cluster_api_grpc_pb.StatusRequest{}
	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Status(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.Status(ctx, req)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) GetCompute(tctx *logger.TraceContext, target string) (*resource_cluster_api_grpc_pb.ActionReply, error) {
	return cli.GetAction(tctx, "GetCompute", target)
}

func (cli *ResourceClusterApiClient) GetNode(tctx *logger.TraceContext, target string) (*resource_cluster_api_grpc_pb.ActionReply, error) {
	return cli.GetAction(tctx, "GetNode", target)
}

func (cli *ResourceClusterApiClient) GetAction(tctx *logger.TraceContext, actionName string, target string) (*resource_cluster_api_grpc_pb.ActionReply, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	atctx := logger.NewAuthproxyTraceContext(tctx, nil)
	atctx.ActionName = actionName

	req := &resource_cluster_api_grpc_pb.ActionRequest{
		Tctx:   atctx,
		Target: target,
	}

	conn, err := cli.NewClientConnection()
	if err != nil {
		return nil, err
	}
	defer func() { err = conn.Close() }()

	ctx, cancel := cli.GetContext()
	defer cancel()

	var rep *resource_cluster_api_grpc_pb.ActionReply
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Action(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.Action(ctx, req)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) UpdateNode(tctx *logger.TraceContext, node *resource_cluster_api_grpc_pb.Node) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	atctx := logger.NewAuthproxyTraceContext(tctx, nil)
	atctx.ActionName = "UpdateNode"

	req := &resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Tctx: atctx,
		Node: node,
	}

	conn, err := cli.NewClientConnection()
	if err != nil {
		return nil, err
	}
	defer func() { err = conn.Close() }()

	ctx, cancel := cli.GetContext()
	defer cancel()

	var rep *resource_cluster_api_grpc_pb.UpdateNodeReply
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.UpdateNode(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.UpdateNode(ctx, req)
	}

	return rep, err
}
