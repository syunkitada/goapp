package resource_api_client

import (
	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/base"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResourceApiClient struct {
	*base.BaseClient
	conf        *config.Config
	localServer *resource_api.ResourceApiServer
}

func NewResourceApiClient(conf *config.Config) *ResourceApiClient {
	resourceClient := ResourceApiClient{
		BaseClient:  base.NewBaseClient(conf, &conf.Resource.ApiApp),
		conf:        conf,
		localServer: resource_api.NewResourceApiServer(conf),
	}
	return &resourceClient
}

func (cli *ResourceApiClient) Status() (*resource_api_grpc_pb.StatusReply, error) {
	var rep *resource_api_grpc_pb.StatusReply
	var err error

	conn, err := cli.NewClientConnection()
	if err != nil {
		return rep, err
	}
	defer conn.Close()

	req := &resource_api_grpc_pb.StatusRequest{}
	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.Status(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.Status(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) GetCluster(req *resource_api_grpc_pb.GetClusterRequest) (*resource_api_grpc_pb.GetClusterReply, error) {
	glog.V(2).Info("Called GetCluster")
	var rep *resource_api_grpc_pb.GetClusterReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetCluster(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.GetCluster(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) GetNode(req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	var rep *resource_api_grpc_pb.GetNodeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.GetNode(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) UpdateNode(req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_api_grpc_pb.UpdateNodeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.UpdateNode(ctx, req)
	}

	return rep, err
}

//
// Compute
//
func (cli *ResourceApiClient) GetCompute(req *resource_api_grpc_pb.GetComputeRequest) (*resource_api_grpc_pb.GetComputeReply, error) {
	var rep *resource_api_grpc_pb.GetComputeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.GetCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) CreateCompute(req *resource_api_grpc_pb.CreateComputeRequest) (*resource_api_grpc_pb.CreateComputeReply, error) {
	var rep *resource_api_grpc_pb.CreateComputeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.CreateCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) UpdateCompute(req *resource_api_grpc_pb.UpdateComputeRequest) (*resource_api_grpc_pb.UpdateComputeReply, error) {
	var rep *resource_api_grpc_pb.UpdateComputeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.UpdateCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) DeleteCompute(req *resource_api_grpc_pb.DeleteComputeRequest) (*resource_api_grpc_pb.DeleteComputeReply, error) {
	var rep *resource_api_grpc_pb.DeleteComputeReply
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
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.DeleteCompute(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) GetImage(req *resource_api_grpc_pb.GetImageRequest) (*resource_api_grpc_pb.GetImageReply, error) {
	var rep *resource_api_grpc_pb.GetImageReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetImage(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.GetImage(ctx, req)
	}

	return rep, err
}

func (cli *ResourceApiClient) GetVolume(req *resource_api_grpc_pb.GetVolumeRequest) (*resource_api_grpc_pb.GetVolumeReply, error) {
	var rep *resource_api_grpc_pb.GetVolumeReply
	var err error
	conn, err := cli.NewClientConnection()
	defer conn.Close()
	if err != nil {
		return rep, err
	}

	ctx, cancel := cli.GetContext()
	defer cancel()
	if cli.conf.Default.EnableTest {
		rep, err = cli.localServer.GetVolume(ctx, req)
	} else {
		grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)
		rep, err = grpcClient.GetVolume(ctx, req)
	}

	return rep, err
}
