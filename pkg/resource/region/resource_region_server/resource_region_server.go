package resource_region_server

import (
	"net"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/region/resource_region_grpc_pb"
)

type ResourceRegionServer struct {
	Conf           *config.Config
	ResourceRegion *config.ResourceRegionConfig
}

func NewResourceRegionServer(conf *config.Config) *ResourceRegionServer {
	resourceRegion, resourceRegionOk := conf.Resource.RegionMap[conf.Resource.Region.Name]
	if !resourceRegionOk {
		glog.Fatalf("NotFound %v in conf.ResourceRegionMap", conf.Resource.Region.Name)
	}

	server := ResourceRegionServer{
		Conf:           conf,
		ResourceRegion: resourceRegion,
	}
	return &server
}

func (resourceRegionServer *ResourceRegionServer) Health(ctx context.Context, statusRequest *resource_region_grpc_pb.HealthRequest) (*resource_region_grpc_pb.HealthReply, error) {
	glog.Info("Health")
	return &resource_region_grpc_pb.HealthReply{Msg: "Health", Err: ""}, nil
}

func (resourceRegionServer *ResourceRegionServer) Serv() error {
	grpcConfig := resourceRegionServer.ResourceRegion.ApiGrpc

	lis, err := net.Listen("tcp", grpcConfig.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		resourceRegionServer.Conf.Path(grpcConfig.CertFile),
		resourceRegionServer.Conf.Path(grpcConfig.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	resource_region_grpc_pb.RegisterResourceRegionServer(grpcServer, resourceRegionServer)
	glog.Infof("Serve: %v", grpcConfig.Listen)
	grpcServer.Serve(lis)

	return nil
}
