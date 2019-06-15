package resource_api_client

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResourceApiClient struct {
	*base.BaseClient
	conf        *config.Config
	localServer *resource_api.ResourceApiServer
}

func NewResourceApiClient(conf *config.Config) *ResourceApiClient {
	resourceClient := ResourceApiClient{
		BaseClient:  base.NewBaseClient(conf, &conf.Resource.ApiApp),
		conf:        conf,
		localServer: resource_api.NewResourceApiServer(conf),
	}
	return &resourceClient
}

func (cli *ResourceApiClient) Status(tctx *logger.TraceContext) (*resource_api_grpc_pb.StatusReply, error) {
	var rep *resource_api_grpc_pb.StatusReply
	var err error

	conn, err := cli.NewClientConnection()
	if err != nil {
		return rep, err
	}
	defer func() { err = conn.Close() }()

	req := &resource_api_grpc_pb.StatusRequest{}
	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Status(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.Status(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) PhysicalAction(tctx *logger.ActionTraceContext) (*authproxy_grpc_pb.ActionReply, error) {
	var err error
	startTime := logger.StartTrace(&tctx.TraceContext)
	defer func() { logger.EndTrace(&tctx.TraceContext, startTime, err, 1) }()

	queries := []*authproxy_grpc_pb.Query{}
	for _, query := range tctx.Queries {
		queries = append(queries, &authproxy_grpc_pb.Query{
			Kind:      query.Kind,
			StrParams: query.StrParams,
			NumParams: query.NumParams,
		})
	}

	req := authproxy_grpc_pb.ActionRequest{
		Tctx:    logger.NewAuthproxyTraceContext(nil, tctx),
		Queries: queries,
	}

	conn, err := cli.NewClientConnection()
	if err != nil {
		return nil, err
	}
	defer func() { err = conn.Close() }()

	ctx, cancel := cli.GetContext()
	defer cancel()

	var rep *authproxy_grpc_pb.ActionReply
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.PhysicalAction(ctx, &req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.PhysicalAction(ctx, &req)
	}

	return rep, err
}

func (cli *ResourceApiClient) VirtualAction(tctx *logger.ActionTraceContext) (*authproxy_grpc_pb.ActionReply, error) {
	var err error
	startTime := logger.StartTrace(&tctx.TraceContext)
	defer func() { logger.EndTrace(&tctx.TraceContext, startTime, err, 1) }()

	queries := []*authproxy_grpc_pb.Query{}
	for _, query := range tctx.Queries {
		queries = append(queries, &authproxy_grpc_pb.Query{
			Kind:      query.Kind,
			StrParams: query.StrParams,
			NumParams: query.NumParams,
		})
	}

	req := authproxy_grpc_pb.ActionRequest{
		Tctx:    logger.NewAuthproxyTraceContext(nil, tctx),
		Queries: queries,
	}

	conn, err := cli.NewClientConnection()
	if err != nil {
		return nil, err
	}
	defer func() { err = conn.Close() }()

	ctx, cancel := cli.GetContext()
	defer cancel()

	var rep *authproxy_grpc_pb.ActionReply
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.VirtualAction(ctx, &req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.VirtualAction(ctx, &req)
	}

	return rep, err
}

func (cli *ResourceApiClient) UpdateNode(tctx *logger.TraceContext, node *resource_api_grpc_pb.Node) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	atctx := logger.NewAuthproxyTraceContext(tctx, nil)
	atctx.ActionName = "UpdateNode"

	req := &resource_api_grpc_pb.UpdateNodeRequest{
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

	var rep *resource_api_grpc_pb.UpdateNodeReply
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.UpdateNode(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.UpdateNode(ctx, req)
	}

	return rep, err
}
