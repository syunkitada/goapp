package grpc_client

import (
	"github.com/golang/glog"

	"golang.org/x/net/context"
	"time"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/health/grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/lib_grpc"
)

var (
	Conf = &config.Conf
)

type HealthClient struct{}

func NewHealthClient() *HealthClient {
	healthClient := HealthClient{}
	return &healthClient
}

func (healthClient *HealthClient) Status() (*grpc_pb.StatusReply, error) {
	conn, connErr := lib_grpc.NewClientConnection(&Conf.HealthGrpc)
	defer conn.Close()
	if connErr != nil {
		glog.Warning("Failed NewClientConnection")
		return nil, connErr
	}

	client := grpc_pb.NewHealthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	statusResponse, err := client.Status(ctx, &grpc_pb.StatusRequest{})
	if err != nil {
		glog.Error("%v.GetFeatures(_) = _, %v: ", client, err)
		return nil, err
	}

	return statusResponse, nil
}
