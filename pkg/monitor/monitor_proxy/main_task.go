package monitor_proxy

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	// "github.com/syunkitada/goapp/pkg/monitor/monitor_model"
	// "github.com/syunkitada/goapp/pkg/monitor/monitor_proxy/monitor_proxy_grpc_pb"
)

func (srv *MonitorProxyServer) MainTask(traceId string) error {
	if err := srv.UpdateNodeTask(traceId); err != nil {
		return err
	}

	return nil
}

func (srv *MonitorProxyServer) UpdateNodeTask(traceId string) error {
	var err error
	startTime := logger.StartTaskTrace(traceId, srv.Host, srv.Name)
	defer func() {
		logger.EndTaskTrace(traceId, srv.Host, srv.Name, startTime, err)
	}()

	// req := &monitor_proxy_grpc_pb.UpdateNodeRequest{
	// }

	// rep := srv.monitorModelProxy.UpdateNode(req)
	// if rep.Err != "" {
	// 	err = fmt.Errorf(rep.Err)
	// 	return err
	// }
	return nil
}
