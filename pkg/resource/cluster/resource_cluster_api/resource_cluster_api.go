package resource_cluster_api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model_api"
)

type ResourceClusterApiServer struct {
	base.BaseApp
	conf                    *config.Config
	cluster                 *config.ResourceClusterConfig
	resourceClusterModelApi *resource_cluster_model_api.ResourceClusterModelApi
}

func NewResourceClusterApiServer(conf *config.Config) *ResourceClusterApiServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		logger.StdoutFatalf("Cluster(%s) is not found in ClusterMap", conf.Resource.Node.ClusterName)
	}

	cluster.ApiApp.Name = "resource.cluster.api"
	server := ResourceClusterApiServer{
		BaseApp:                 base.NewBaseApp(conf, &cluster.ApiApp),
		conf:                    conf,
		cluster:                 &cluster,
		resourceClusterModelApi: resource_cluster_model_api.NewResourceClusterModelApi(conf),
	}

	server.RegisterDriver(&server)
	return &server
}

func (srv *ResourceClusterApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_cluster_api_grpc_pb.RegisterResourceClusterApiServer(grpcServer, srv)
	return nil
}

func (srv *ResourceClusterApiServer) Action(ctx context.Context,
	req *authproxy_grpc_pb.ActionRequest) (*authproxy_grpc_pb.ActionReply, error) {
	var err error
	rep := &authproxy_grpc_pb.ActionReply{Tctx: req.Tctx}
	tctx := logger.NewGrpcAuthproxyTraceContext(srv.Host, srv.Name, ctx, req.Tctx)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	srv.resourceClusterModelApi.Action(tctx, req, rep)
	return rep, nil
}

func (srv *ResourceClusterApiServer) localAction(tctx *logger.TraceContext, atctx *logger.ActionTraceContext) (*authproxy_grpc_pb.ActionReply, error) {
	queries := []*authproxy_grpc_pb.Query{}
	for _, query := range atctx.Queries {
		queries = append(queries, &authproxy_grpc_pb.Query{
			Kind:      query.Kind,
			StrParams: query.StrParams,
			NumParams: query.NumParams,
		})
	}

	req := authproxy_grpc_pb.ActionRequest{
		Tctx:    logger.NewAuthproxyTraceContext(nil, atctx),
		Queries: queries,
	}
	rep := &authproxy_grpc_pb.ActionReply{Tctx: req.Tctx}

	srv.resourceClusterModelApi.Action(tctx, &req, rep)
	return rep, nil
}
