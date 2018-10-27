package resource_api_client

import (
	"errors"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResourceApiClient struct {
	Conf               *config.Config
	CaFilePath         string
	ServerHostOverride string
	Targets            []string
}

func NewResourceApiClient(conf *config.Config) *ResourceApiClient {
	resourceClient := ResourceApiClient{
		Conf:               conf,
		CaFilePath:         conf.Path(conf.Resource.ApiGrpc.CaFile),
		ServerHostOverride: conf.Resource.ApiGrpc.ServerHostOverride,
		Targets:            conf.Resource.ApiGrpc.Targets,
	}
	return &resourceClient
}

func (client *ResourceApiClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range client.Targets {
		creds, credsErr := credentials.NewClientTLSFromFile(client.CaFilePath, client.ServerHostOverride)
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

func (client *ResourceApiClient) Status() (*resource_api_grpc_pb.StatusReply, error) {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	statusResponse, err := grpcClient.Status(ctx, &resource_api_grpc_pb.StatusRequest{})
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", grpcClient, err)
		return nil, err
	}

	return statusResponse, nil
}

func (client *ResourceApiClient) UpdateNode(request *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	conn, connErr := client.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	grpcClient := resource_api_grpc_pb.NewResourceApiClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	reply, err := grpcClient.UpdateNode(ctx, request)
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", grpcClient, err)
		return nil, err
	}

	return reply, nil
}
