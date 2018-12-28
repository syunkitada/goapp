package monitor_api

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	// "github.com/syunkitada/goapp/pkg/monitor/monitor_model"
	// "github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
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

	// req := &monitor_api_grpc_pb.UpdateNodeRequest{
	// }

	// rep := srv.monitorModelApi.UpdateNode(req)
	// if rep.Err != "" {
	// 	err = fmt.Errorf(rep.Err)
	// 	return err
	// }
	return nil
}
