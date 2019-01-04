package monitor_api

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

func (srv *MonitorApiServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNodeTask(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *MonitorApiServer) UpdateNodeTask(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() {
		logger.EndTrace(tctx, startTime, err)
	}()

	req := &monitor_api_grpc_pb.UpdateNodeRequest{
		Name:         srv.Host,
		Kind:         monitor_model.KindMonitorApi,
		Role:         monitor_model.RoleMember,
		Status:       monitor_model.StatusEnabled,
		StatusReason: "Default",
		State:        monitor_model.StateUp,
		StateReason:  "UpdateNode",
	}

	rep := srv.monitorModelApi.UpdateNode(req)
	if rep.Err != "" {
		err = fmt.Errorf(rep.Err)
		return err
	}
	return nil
}
