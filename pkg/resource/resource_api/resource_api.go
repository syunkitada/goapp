package resource_api

import (
	"fmt"
	"time"

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

func (srv *ResourceApiServer) GetCluster(ctx context.Context, req *resource_api_grpc_pb.GetClusterRequest) (*resource_api_grpc_pb.GetClusterReply, error) {
	var rep *resource_api_grpc_pb.GetClusterReply
	var err error
	rep, err = srv.resourceModelApi.GetCluster(req)
	glog.Info(rep)
	glog.Infof("Completed GetCluster: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) GetNode(ctx context.Context, req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	var rep *resource_api_grpc_pb.GetNodeReply
	var err error
	rep, err = srv.resourceModelApi.GetNode(req)
	glog.Infof("Completed GetNode: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateNode(ctx context.Context, req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_api_grpc_pb.UpdateNodeReply
	var err error
	rep, err = srv.resourceModelApi.UpdateNode(req)
	glog.Infof("Completed UpdateNode: %v", err)
	return rep, err
}

//
// Compute
//
func (srv *ResourceApiServer) GetCompute(ctx context.Context, req *resource_api_grpc_pb.GetComputeRequest) (*resource_api_grpc_pb.GetComputeReply, error) {
	var rep *resource_api_grpc_pb.GetComputeReply
	var err error
	rep, err = srv.resourceModelApi.GetCompute(req)

	glog.Infof("Completed GetCompute: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) CreateCompute(ctx context.Context, req *resource_api_grpc_pb.CreateComputeRequest) (*resource_api_grpc_pb.CreateComputeReply, error) {
	startTime := time.Now()
	var rep *resource_api_grpc_pb.CreateComputeReply
	var err error
	rep, err = srv.resourceModelApi.CreateCompute(req)
	if err != nil {
		return rep, fmt.Errorf("@@ApiCreateCompute: time=%v, error=%v", time.Now().Sub(startTime), err)
	}
	glog.Infof("Completed CreateCompute: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateCompute(ctx context.Context, req *resource_api_grpc_pb.UpdateComputeRequest) (*resource_api_grpc_pb.UpdateComputeReply, error) {
	var rep *resource_api_grpc_pb.UpdateComputeReply
	var err error
	rep, err = srv.resourceModelApi.UpdateCompute(req)
	glog.Infof("Completed UpdateCompute: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) DeleteCompute(ctx context.Context, req *resource_api_grpc_pb.DeleteComputeRequest) (*resource_api_grpc_pb.DeleteComputeReply, error) {
	var rep *resource_api_grpc_pb.DeleteComputeReply
	var err error
	rep, err = srv.resourceModelApi.DeleteCompute(req)
	glog.Infof("Completed DeleteCompute: %v", err)
	return rep, err
}

//
// Container
//
func (srv *ResourceApiServer) GetContainer(ctx context.Context, req *resource_api_grpc_pb.GetContainerRequest) (*resource_api_grpc_pb.GetContainerReply, error) {
	var rep *resource_api_grpc_pb.GetContainerReply
	var err error
	rep, err = srv.resourceModelApi.GetContainer(req)
	glog.Infof("Completed GetContainer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) CreateContainer(ctx context.Context, req *resource_api_grpc_pb.CreateContainerRequest) (*resource_api_grpc_pb.CreateContainerReply, error) {
	var rep *resource_api_grpc_pb.CreateContainerReply
	var err error
	rep, err = srv.resourceModelApi.CreateContainer(req)
	glog.Infof("Completed CreateContainer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateContainer(ctx context.Context, req *resource_api_grpc_pb.UpdateContainerRequest) (*resource_api_grpc_pb.UpdateContainerReply, error) {
	var rep *resource_api_grpc_pb.UpdateContainerReply
	var err error
	rep, err = srv.resourceModelApi.UpdateContainer(req)
	glog.Infof("Completed UpdateContainer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) DeleteContainer(ctx context.Context, req *resource_api_grpc_pb.DeleteContainerRequest) (*resource_api_grpc_pb.DeleteContainerReply, error) {
	var rep *resource_api_grpc_pb.DeleteContainerReply
	var err error
	rep, err = srv.resourceModelApi.DeleteContainer(req)
	glog.Infof("Completed DeleteContainer: %v", err)
	return rep, err
}

//
// Image
//
func (srv *ResourceApiServer) GetImage(ctx context.Context, req *resource_api_grpc_pb.GetImageRequest) (*resource_api_grpc_pb.GetImageReply, error) {
	var rep *resource_api_grpc_pb.GetImageReply
	var err error
	rep, err = srv.resourceModelApi.GetImage(req)
	glog.Infof("Completed GetImage: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) CreateImage(ctx context.Context, req *resource_api_grpc_pb.CreateImageRequest) (*resource_api_grpc_pb.CreateImageReply, error) {
	var rep *resource_api_grpc_pb.CreateImageReply
	var err error
	rep, err = srv.resourceModelApi.CreateImage(req)
	glog.Infof("Completed CreateImage: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateImage(ctx context.Context, req *resource_api_grpc_pb.UpdateImageRequest) (*resource_api_grpc_pb.UpdateImageReply, error) {
	var rep *resource_api_grpc_pb.UpdateImageReply
	var err error
	rep, err = srv.resourceModelApi.UpdateImage(req)
	glog.Infof("Completed UpdateImage: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) DeleteImage(ctx context.Context, req *resource_api_grpc_pb.DeleteImageRequest) (*resource_api_grpc_pb.DeleteImageReply, error) {
	var rep *resource_api_grpc_pb.DeleteImageReply
	var err error
	rep, err = srv.resourceModelApi.DeleteImage(req)
	glog.Infof("Completed DeleteImage: %v", err)
	return rep, err
}

//
// Volume
//
func (srv *ResourceApiServer) GetVolume(ctx context.Context, req *resource_api_grpc_pb.GetVolumeRequest) (*resource_api_grpc_pb.GetVolumeReply, error) {
	var rep *resource_api_grpc_pb.GetVolumeReply
	var err error
	rep, err = srv.resourceModelApi.GetVolume(req)
	glog.Infof("Completed GetVolume: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) CreateVolume(ctx context.Context, req *resource_api_grpc_pb.CreateVolumeRequest) (*resource_api_grpc_pb.CreateVolumeReply, error) {
	var rep *resource_api_grpc_pb.CreateVolumeReply
	var err error
	rep, err = srv.resourceModelApi.CreateVolume(req)
	glog.Infof("Completed CreateVolume: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateVolume(ctx context.Context, req *resource_api_grpc_pb.UpdateVolumeRequest) (*resource_api_grpc_pb.UpdateVolumeReply, error) {
	var rep *resource_api_grpc_pb.UpdateVolumeReply
	var err error
	rep, err = srv.resourceModelApi.UpdateVolume(req)
	glog.Infof("Completed UpdateVolume: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) DeleteVolume(ctx context.Context, req *resource_api_grpc_pb.DeleteVolumeRequest) (*resource_api_grpc_pb.DeleteVolumeReply, error) {
	var rep *resource_api_grpc_pb.DeleteVolumeReply
	var err error
	rep, err = srv.resourceModelApi.DeleteVolume(req)
	glog.Infof("Completed DeleteVolume: %v", err)
	return rep, err
}

//
// Loadbalancer
//
func (srv *ResourceApiServer) GetLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.GetLoadbalancerRequest) (*resource_api_grpc_pb.GetLoadbalancerReply, error) {
	var rep *resource_api_grpc_pb.GetLoadbalancerReply
	var err error
	rep, err = srv.resourceModelApi.GetLoadbalancer(req)
	glog.Infof("Completed GetLoadbalancer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) CreateLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.CreateLoadbalancerRequest) (*resource_api_grpc_pb.CreateLoadbalancerReply, error) {
	var rep *resource_api_grpc_pb.CreateLoadbalancerReply
	var err error
	rep, err = srv.resourceModelApi.CreateLoadbalancer(req)
	glog.Infof("Completed CreateLoadbalancer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) UpdateLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.UpdateLoadbalancerRequest) (*resource_api_grpc_pb.UpdateLoadbalancerReply, error) {
	var rep *resource_api_grpc_pb.UpdateLoadbalancerReply
	var err error
	rep, err = srv.resourceModelApi.UpdateLoadbalancer(req)
	glog.Infof("Completed UpdateLoadbalancer: %v", err)
	return rep, err
}

func (srv *ResourceApiServer) DeleteLoadbalancer(ctx context.Context, req *resource_api_grpc_pb.DeleteLoadbalancerRequest) (*resource_api_grpc_pb.DeleteLoadbalancerReply, error) {
	var rep *resource_api_grpc_pb.DeleteLoadbalancerReply
	var err error
	rep, err = srv.resourceModelApi.DeleteLoadbalancer(req)
	glog.Infof("Completed DeleteLoadbalancer: %v", err)
	return rep, err
}
