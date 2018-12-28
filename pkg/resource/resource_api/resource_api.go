package resource_api

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceApiServer struct {
	base.BaseApp
	conf             *config.Config
	resourceModelApi *resource_model_api.ResourceModelApi
}

func NewResourceApiServer(conf *config.Config) *ResourceApiServer {
	conf.Resource.ApiApp.Name = "resource.api"
	server := ResourceApiServer{
		BaseApp:          base.NewBaseApp(conf, &conf.Resource.ApiApp),
		conf:             conf,
		resourceModelApi: resource_model_api.NewResourceModelApi(conf, nil),
	}

	server.RegisterDriver(&server)

	return &server
}

func (srv *ResourceApiServer) RegisterGrpcServer(grpcServer *grpc.Server) error {
	resource_api_grpc_pb.RegisterResourceApiServer(grpcServer, srv)
	return nil
}

func (srv *ResourceApiServer) Status(ctx context.Context, statusRequest *resource_api_grpc_pb.StatusRequest) (*resource_api_grpc_pb.StatusReply, error) {
	return &resource_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

//
// Cluster
//
func (srv *ResourceApiServer) GetCluster(ctx context.Context, req *resource_api_grpc_pb.GetClusterRequest) (*resource_api_grpc_pb.GetClusterReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetCluster(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Node
//
func (srv *ResourceApiServer) GetNode(ctx context.Context, req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetNode(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateNode(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// NetworkV4
//
func (srv *ResourceApiServer) GetNetworkV4(ctx context.Context, req *resource_api_grpc_pb.GetNetworkV4Request) (*resource_api_grpc_pb.GetNetworkV4Reply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetNetworkV4(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateNetworkV4(ctx context.Context, req *resource_api_grpc_pb.CreateNetworkV4Request) (*resource_api_grpc_pb.CreateNetworkV4Reply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateNetworkV4(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateNetworkV4(ctx context.Context, req *resource_api_grpc_pb.UpdateNetworkV4Request) (*resource_api_grpc_pb.UpdateNetworkV4Reply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateNetworkV4(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteNetworkV4(ctx context.Context, req *resource_api_grpc_pb.DeleteNetworkV4Request) (*resource_api_grpc_pb.DeleteNetworkV4Reply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteNetworkV4(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Compute
//
func (srv *ResourceApiServer) GetCompute(ctx context.Context, req *resource_api_grpc_pb.GetComputeRequest) (*resource_api_grpc_pb.GetComputeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetCompute(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateCompute(ctx context.Context, req *resource_api_grpc_pb.CreateComputeRequest) (*resource_api_grpc_pb.CreateComputeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateCompute(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateCompute(ctx context.Context, req *resource_api_grpc_pb.UpdateComputeRequest) (*resource_api_grpc_pb.UpdateComputeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateCompute(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteCompute(ctx context.Context, req *resource_api_grpc_pb.DeleteComputeRequest) (*resource_api_grpc_pb.DeleteComputeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteCompute(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Container
//
func (srv *ResourceApiServer) GetContainer(ctx context.Context, req *resource_api_grpc_pb.GetContainerRequest) (*resource_api_grpc_pb.GetContainerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetContainer(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateContainer(ctx context.Context, req *resource_api_grpc_pb.CreateContainerRequest) (*resource_api_grpc_pb.CreateContainerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateContainer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateContainer(ctx context.Context, req *resource_api_grpc_pb.UpdateContainerRequest) (*resource_api_grpc_pb.UpdateContainerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateContainer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteContainer(ctx context.Context, req *resource_api_grpc_pb.DeleteContainerRequest) (*resource_api_grpc_pb.DeleteContainerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteContainer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Image
//
func (srv *ResourceApiServer) GetImage(ctx context.Context, req *resource_api_grpc_pb.GetImageRequest) (*resource_api_grpc_pb.GetImageReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetImage(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateImage(ctx context.Context, req *resource_api_grpc_pb.CreateImageRequest) (*resource_api_grpc_pb.CreateImageReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateImage(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateImage(ctx context.Context, req *resource_api_grpc_pb.UpdateImageRequest) (*resource_api_grpc_pb.UpdateImageReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateImage(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteImage(ctx context.Context, req *resource_api_grpc_pb.DeleteImageRequest) (*resource_api_grpc_pb.DeleteImageReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteImage(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Volume
//
func (srv *ResourceApiServer) GetVolume(ctx context.Context, req *resource_api_grpc_pb.GetVolumeRequest) (*resource_api_grpc_pb.GetVolumeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetVolume(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateVolume(ctx context.Context, req *resource_api_grpc_pb.CreateVolumeRequest) (*resource_api_grpc_pb.CreateVolumeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateVolume(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateVolume(ctx context.Context, req *resource_api_grpc_pb.UpdateVolumeRequest) (*resource_api_grpc_pb.UpdateVolumeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateVolume(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteVolume(ctx context.Context, req *resource_api_grpc_pb.DeleteVolumeRequest) (*resource_api_grpc_pb.DeleteVolumeReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteVolume(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

//
// Loadbalancer
//
func (srv *ResourceApiServer) GetLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.GetLoadbalancerRequest) (*resource_api_grpc_pb.GetLoadbalancerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.GetLoadbalancer(tctx, req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) CreateLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.CreateLoadbalancerRequest) (*resource_api_grpc_pb.CreateLoadbalancerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.CreateLoadbalancer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) UpdateLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.UpdateLoadbalancerRequest) (*resource_api_grpc_pb.UpdateLoadbalancerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.UpdateLoadbalancer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}

func (srv *ResourceApiServer) DeleteLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.DeleteLoadbalancerRequest) (*resource_api_grpc_pb.DeleteLoadbalancerReply, error) {
	tctx := logger.NewGrpcTraceContext(srv.Host, srv.Name, ctx)
	startTime := logger.StartTrace(tctx)
	rep := srv.resourceModelApi.DeleteLoadbalancer(req)
	logger.EndGrpcTrace(tctx, startTime, rep.StatusCode, rep.Err)
	return rep, nil
}
