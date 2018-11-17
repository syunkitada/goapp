package resource_cluster_api_client

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
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
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Cluster.Name]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Cluster.Name))
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
	defer conn.Close()

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

func (cli *ResourceClusterApiClient) GetNode(req *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	var rep *resource_cluster_api_grpc_pb.GetNodeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetNode(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.GetNode(ctx, req)
		glog.Info(err)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) UpdateNode(req *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_cluster_api_grpc_pb.UpdateNodeReply
	var err error

	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()

	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.UpdateNode(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.UpdateNode(ctx, req)
	}

	return rep, err
}
