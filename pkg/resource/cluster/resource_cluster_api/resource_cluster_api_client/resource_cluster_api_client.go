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
	cluster, ok := conf.Resource.ClusterMap[conf.Resource.Node.ClusterName]
	if !ok {
		glog.Fatal(fmt.Errorf("Cluster(%v) is not found in ClusterMap", conf.Resource.Node.ClusterName))
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

//
// Compute
//
func (cli *ResourceClusterApiClient) GetCompute(req *resource_cluster_api_grpc_pb.GetComputeRequest) (*resource_cluster_api_grpc_pb.GetComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.GetComputeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetCompute(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.GetCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) CreateCompute(req *resource_cluster_api_grpc_pb.CreateComputeRequest) (*resource_cluster_api_grpc_pb.CreateComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.CreateComputeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.CreateCompute(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.CreateCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) UpdateCompute(req *resource_cluster_api_grpc_pb.UpdateComputeRequest) (*resource_cluster_api_grpc_pb.UpdateComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.UpdateComputeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.UpdateCompute(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.UpdateCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceClusterApiClient) DeleteCompute(req *resource_cluster_api_grpc_pb.DeleteComputeRequest) (*resource_cluster_api_grpc_pb.DeleteComputeReply, error) {
	var rep *resource_cluster_api_grpc_pb.DeleteComputeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.DeleteCompute(ctx, req)
	} else {
		grpcClient := resource_cluster_api_grpc_pb.NewResourceClusterApiClient(conn)
		rep, err = grpcClient.DeleteCompute(ctx, req)
	}

	return rep, err
}
