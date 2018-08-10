package grpc_server

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/health/grpc_pb"
)

type HealthServer struct{}

func (s *HealthServer) Status(ctx context.Context, statusRequest *grpc_pb.StatusRequest) (*grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}

func NewHealthServer() *HealthServer {
	healthServer := &HealthServer{}
	return healthServer
}
