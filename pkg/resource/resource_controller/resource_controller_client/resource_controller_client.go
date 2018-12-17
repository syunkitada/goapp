package resource_controller_client

import (
	"errors"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller/resource_controller_grpc_pb"
)

type ResourceClient struct {
	Conf               *config.Config
	CaFilePath         string
	ServerHostOverride string
	Targets            []string
}

func NewResourceClient(conf *config.Config) *ResourceClient {
	resourceClient := ResourceClient{
		Conf:               conf,
		CaFilePath:         conf.Path(conf.Resource.Grpc.CaFile),
		ServerHostOverride: conf.Resource.Grpc.ServerHostOverride,
		Targets:            conf.Resource.Grpc.Targets,
	}
	return &resourceClient
}

func (resourceClient *ResourceClient) NewClientConnection() (*grpc.ClientConn, error) {
	var opts []grpc.DialOption

	for _, target := range resourceClient.Targets {
		creds, credsErr := credentials.NewClientTLSFromFile(resourceClient.CaFilePath, resourceClient.ServerHostOverride)
		if credsErr != nil {
			continue
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

		conn, err := grpc.Dial(target, opts...)
		if err != nil {
			continue
		}

		return conn, nil
	}

	return nil, errors.New("Failed NewGrpcConnection")
}

func (resourceClient *ResourceClient) Status() (*resource_controller_grpc_pb.StatusReply, error) {
	conn, connErr := resourceClient.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		return nil, connErr
	}

	client := resource_controller_grpc_pb.NewResourceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Second)
	defer cancel()

	statusResponse, err := client.Status(ctx, &resource_controller_grpc_pb.StatusRequest{})
	if err != nil {
		return nil, err
	}

	return statusResponse, nil
}
