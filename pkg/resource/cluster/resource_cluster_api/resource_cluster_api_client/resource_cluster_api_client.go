package resource_cluster_api_client

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
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
		logger.StdoutFatalf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName)
	}

	resourceClient := ResourceClusterApiClient{
		BaseClient:  base.NewBaseClient(conf, &cluster.ApiApp),
		conf:        conf,
		cluster:     &cluster,
		localServer: localServer,
	}
	return &resourceClient
}

func (cli *ResourceClusterApiClient) Action(tctx *logger.ActionTraceContext) (*authproxy_grpc_pb.ActionReply, error) {
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
		rep, err = cli.localServer.Action(ctx, &req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.Action(ctx, &req)
	}

	return rep, err
}
