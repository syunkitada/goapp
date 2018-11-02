package resource_cluster_api

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_client"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model/resource_cluster_model_api"
)

type ResourceClusterApiServer struct {
	Conf                     *config.Config
	GrpcServer               *grpc.Server
	ShutdownTimeout          time.Duration
	cluster                  *config.ResourceClusterConfig
	loopInterval             time.Duration
	isGracefulShutdown       bool
	resourceClusterModelApi  *resource_cluster_model_api.ResourceClusterModelApi
	resourceClusterApiClient *resource_cluster_api_client.ResourceClusterApiClient
}

func NewResourceClusterApiServer(conf *config.Config) *ResourceClusterApiServer {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Cluster.Name]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Cluster.Name))
	}

	server := ResourceClusterApiServer{
		Conf:                     conf,
		cluster:                  cluster,
		ShutdownTimeout:          time.Duration(10) * time.Second,
		loopInterval:             time.Duration(5) * time.Second,
		isGracefulShutdown:       false,
		resourceClusterModelApi:  resource_cluster_model_api.NewResourceClusterModelApi(conf),
		resourceClusterApiClient: resource_cluster_api_client.NewResourceClusterApiClient(conf),
	}
	return &server
}

func (server *ResourceClusterApiServer) StartMainLoop() error {
	go server.MainLoop()
	return nil
}

func (server *ResourceClusterApiServer) MainLoop() error {
	glog.Info("Starting MainLoop")
	for {
		if err := server.MainTask(); err != nil {
			glog.Warning(err)
		}

		if server.isGracefulShutdown {
			glog.Info("End MainLoop on GracefulShutdown")
			glog.Info("Start GrpcServer.GracefulStop")
			server.GrpcServer.GracefulStop()
			glog.Info("Success GrpcServer.GracefulStop")
			glog.Info("Success GracefulShutdown")
			os.Exit(0)
		}
		glog.Infof("Success MainTask, and sleep %v", server.loopInterval)
		time.Sleep(server.loopInterval)
	}
	return nil
}

func (server *ResourceClusterApiServer) Serv() error {
	grpcConfig := server.cluster.ApiGrpc

	lis, err := net.Listen("tcp", grpcConfig.Listen)
	if err != nil {
		return err
	}

	var opts []grpc.ServerOption
	creds, err := credentials.NewServerTLSFromFile(
		server.Conf.Path(grpcConfig.CertFile),
		server.Conf.Path(grpcConfig.KeyFile),
	)
	if err != nil {
		return err
	}
	opts = []grpc.ServerOption{grpc.Creds(creds)}

	server.GrpcServer = grpc.NewServer(opts...)

	resource_cluster_api_grpc_pb.RegisterResourceClusterApiServer(server.GrpcServer, server)
	glog.Infof("Serve: %v", grpcConfig.Listen)

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, syscall.SIGTERM)
		signal.Notify(shutdown, syscall.SIGINT)
		<-shutdown
		if err := server.GracefulShutdown(context.Background()); err != nil {
			glog.Errorf("App Shutdown: %v\n", err)
		}
	}()

	if err := server.GrpcServer.Serve(lis); err != nil {
		glog.Infof("App Serv Failed: %v\n", err)
	}

	glog.Infof("End Serv")
	return nil
}

func (server *ResourceClusterApiServer) GracefulShutdown(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, server.ShutdownTimeout)
	defer cancel()

	server.isGracefulShutdown = true

	select {
	case <-ctx.Done():
		glog.Warning(ctx.Err())
		os.Exit(1)
	}

	return nil
}

func (srv *ResourceClusterApiServer) Status(ctx context.Context, statusRequest *resource_cluster_api_grpc_pb.StatusRequest) (*resource_cluster_api_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_cluster_api_grpc_pb.StatusReply{Msg: "Status"}, nil
}

func (srv *ResourceClusterApiServer) GetNode(ctx context.Context, req *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	glog.Info("GetNode")
	var err error
	var rep *resource_cluster_api_grpc_pb.GetNodeReply
	if rep, err = srv.resourceClusterModelApi.GetNode(req); err != nil {
		glog.Error(err)
	}
	return rep, err
}

func (srv *ResourceClusterApiServer) UpdateNode(ctx context.Context, req *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var err error
	glog.Info("UpdateNode")
	if err = srv.resourceClusterModelApi.UpdateNode(req); err != nil {
		glog.Error(err)
	}
	return &resource_cluster_api_grpc_pb.UpdateNodeReply{}, err
}
