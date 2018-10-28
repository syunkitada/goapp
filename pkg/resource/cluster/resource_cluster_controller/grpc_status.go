package resource_controller

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/resource/resource_controller/resource_controller_grpc_pb"
)

func (server *ResourceControllerServer) Status(ctx context.Context, statusRequest *resource_controller_grpc_pb.StatusRequest) (*resource_controller_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_controller_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
