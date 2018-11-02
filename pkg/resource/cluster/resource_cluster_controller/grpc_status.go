package resource_cluster_controller

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb"
)

func (server *ResourceClusterControllerServer) Status(ctx context.Context, statusRequest *resource_cluster_controller_grpc_pb.StatusRequest) (*resource_cluster_controller_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_cluster_controller_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
