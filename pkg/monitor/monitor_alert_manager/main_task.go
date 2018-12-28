package monitor_alert_manager

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

func (srv *MonitorAlertManagerServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}
	if err := srv.SyncRole(tctx); err != nil {
		return err
	}
	if srv.role == monitor_model.RoleMember {
		return nil
	}

	if err := srv.monitorModelApi.CheckNodes(); err != nil {
		return err
	}

	// TODO

	return nil
}

func (srv *MonitorAlertManagerServer) UpdateNode(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err) }()

	req := &monitor_api_grpc_pb.UpdateNodeRequest{
		TraceId:      tctx.TraceId,
		Name:         srv.Host,
		Kind:         monitor_model.KindMonitorAlertManager,
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

func (srv *MonitorAlertManagerServer) SyncRole(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err) }()

	nodes, err := srv.monitorModelApi.SyncRole(monitor_model.KindMonitorAlertManager)
	if err != nil {
		return err
	}

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != monitor_model.KindMonitorAlertManager {
			continue
		}
		if node.Name == srv.conf.Default.Host && node.Status == monitor_model.StatusEnabled && node.State == monitor_model.StateUp {
			existsSelfNode = true
			srv.role = node.Role
		}
		if node.Status == monitor_model.StatusEnabled && node.State == monitor_model.StateUp {
			if node.Role == monitor_model.RoleLeader {
				existsActiveLeader = true
			}
		}
	}

	if !existsSelfNode {
		err = fmt.Errorf("This node is not activated")
		return err
	}

	if !existsActiveLeader {
		err = fmt.Errorf("Active Leader is not exists, after ReassignNode")
		return err
	}

	return nil
}
