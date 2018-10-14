package resource_client

import (
	"errors"
	"time"

	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/grpc_pb"
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

func (resourceClient *ResourceClient) Status() (*grpc_pb.StatusReply, error) {
	conn, connErr := resourceClient.NewClientConnection()
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	client := grpc_pb.NewResourceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	statusResponse, err := client.Status(ctx, &grpc_pb.StatusRequest{})
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", client, err)
		return nil, err
	}

	return statusResponse, nil
}
