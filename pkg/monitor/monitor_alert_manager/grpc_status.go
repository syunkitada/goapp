package monitor_alert_manager

import (
	"golang.org/x/net/context"

	"github.com/syunkitada/goapp/pkg/monitor/monitor_alert_manager/monitor_alert_manager_grpc_pb"
)

func (server *MonitorAlertManagerServer) Status(ctx context.Context, statusRequest *monitor_alert_manager_grpc_pb.StatusRequest) (*monitor_alert_manager_grpc_pb.StatusReply, error) {
	return &monitor_alert_manager_grpc_pb.StatusReply{Msg: "Health", Err: ""}, nil
}
