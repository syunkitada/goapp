package resource_api

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type ResourceApiServer struct {
	base.BaseApp
	conf             *config.Config
	resourceModelApi *resource_model_api.ResourceModelApi
}

func NewResourceApiServer(conf *config.Config) *ResourceApiServer {
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

func (srv *ResourceApiServer) GetNode(ctx context.Context, req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	var rep *resource_api_grpc_pb.GetNodeReply
	var err error
	rep, err = srv.resourceModelApi.GetNode(req)
	glog.Infof("Completed GetNode: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) GetCluster(ctx context.Context, req *resource_api_grpc_pb.GetClusterRequest) (*resource_api_grpc_pb.GetClusterReply, error) {
	var rep *resource_api_grpc_pb.GetClusterReply
	var err error
	rep, err = srv.resourceModelApi.GetCluster(req)
	glog.Infof("Completed GetCluster: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) GetCompute(ctx context.Context, req *resource_api_grpc_pb.GetComputeRequest) (*resource_api_grpc_pb.GetComputeReply, error) {
	var rep *resource_api_grpc_pb.GetComputeReply
	var err error
	rep, err = srv.resourceModelApi.GetCompute(req)
	glog.Infof("Completed GetCompute: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) GetImage(ctx context.Context, req *resource_api_grpc_pb.GetImageRequest) (*resource_api_grpc_pb.GetImageReply, error) {
	var rep *resource_api_grpc_pb.GetImageReply
	var err error
	rep, err = srv.resourceModelApi.GetImage(req)
	glog.Infof("Completed GetImage: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) GetVolume(ctx context.Context, req *resource_api_grpc_pb.GetVolumeRequest) (*resource_api_grpc_pb.GetVolumeReply, error) {
	var rep *resource_api_grpc_pb.GetVolumeReply
	var err error
	rep, err = srv.resourceModelApi.GetVolume(req)
	glog.Infof("Completed GetVolume: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_api_grpc_pb.UpdateNodeReply
	var err error
	rep, err = srv.resourceModelApi.UpdateNode(req)
	glog.Infof("Completed UpdateNode: %v", err)
	return rep, err
}
