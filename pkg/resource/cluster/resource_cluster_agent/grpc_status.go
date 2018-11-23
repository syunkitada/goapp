package resource_cluster_agent

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/resource_cluster_agent_grpc_pb"
)

func (server *ResourceClusterAgentServer) Status(ctx context.Context, statusRequest *resource_cluster_agent_grpc_pb.StatusRequest) (*resource_cluster_agent_grpc_pb.StatusReply, error) {
	glog.Info("Status")
	return &resource_cluster_agent_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
