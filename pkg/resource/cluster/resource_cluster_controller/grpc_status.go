package resource_cluster_controller

import (
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller/resource_cluster_controller_grpc_pb"
)

func (server *ResourceClusterControllerServer) Status(ctx context.Context, statusRequest *resource_cluster_controller_grpc_pb.StatusRequest) (*resource_cluster_controller_grpc_pb.StatusReply, error) {
	return &resource_cluster_controller_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
