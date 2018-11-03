package resource_cluster_api_client

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
)

type ResourceClusterApiClient struct {
	conf               *config.Config
	cluster            *config.ResourceClusterConfig
	caFilePath         string
	serverHostOverride string
	targets            []string
}

func NewResourceClusterApiClient(conf *config.Config) *ResourceClusterApiClient {
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Cluster.Name]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Cluster.Name))
	}

	resourceClient := ResourceClusterApiClient{
		conf:               conf,
		cluster:            cluster,
		caFilePath:         conf.Path(cluster.ApiApp.CaFile),
		serverHostOverride: cluster.ApiApp.ServerHostOverride,
		targets:            cluster.ApiApp.Targets,
	}
	return &resourceClient
}

func (client *ResourceClusterApiClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range client.targets {
		creds, credsErr := credentials.NewClientTLSFromFile(client.caFilePath, client.serverHostOverride)
		if credsErr != nil {
			glog.Warning("Failed to create TLS credentials %v", credsErr)
			continue
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(target, opts...)
		if err != nil {
			glog.Warning("fail to dial: %v", err)
			continue
		}

		return conn, nil
	}

	return nil, errors.New("Failed NewGrpcConnection")
}

func (client *ResourceClusterApiClient) Status() (*resource_cluster_api_grpc_pb.StatusReply, error) {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	statusResponse, err := grpcClient.Status(ctx, &resource_cluster_api_grpc_pb.StatusRequest{})
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", grpcClient, err)
		return nil, err
	}

	return statusResponse, nil
}

func (cli *ResourceClusterApiClient) GetNode(request *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	conn, connErr := cli.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	reply, err := grpcClient.GetNode(ctx, request)
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", grpcClient, err)
		return nil, err
	}

	return reply, nil
}

func (client *ResourceClusterApiClient) UpdateNode(request *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	reply, err := grpcClient.UpdateNode(ctx, request)
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", grpcClient, err)
		return nil, err
	}

	return reply, nil
}
