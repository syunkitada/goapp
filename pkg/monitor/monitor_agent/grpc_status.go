package monitor_agent

import (
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/monitor/monitor_agent/monitor_agent_grpc_pb"
)

func (server *MonitorAgentServer) Status(ctx context.Context, statusRequest *monitor_agent_grpc_pb.StatusRequest) (*monitor_agent_grpc_pb.StatusReply, error) {
	return &monitor_agent_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
