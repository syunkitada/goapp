package server

import (
	"fmt"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *Server) SyncComputeAssignments(tctx *logger.TraceContext,
	assignments []spec.ComputeAssignmentEx) ([]spec.AssignmentReport, error) {
	var err error
	var ok bool
	var retryCount int
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tctx.SetTimeout(1)

	fmt.Println("SyncComputeAssignments: ", assignments)

	assignmentReports := []spec.AssignmentReport{}

	// ユニークなID管理用
	// ファイルに保存しておく
	assignedNetnsIds := make([]bool, 4096)

	computeNetnsPortsMap := map[uint][]compute_models.NetnsPort{}
	activatingAssignmentMap := map[uint]spec.ComputeAssignmentEx{}
	deletingAssignmentMap := map[uint]spec.ComputeAssignmentEx{}
	for _, assignment := range assignments {
		switch assignment.Status {
		case base_const.StatusActive, base_const.StatusCreating, base_const.StatusUpdating,
			base_const.StatusUnknownActivating, base_const.StatusRebalancingUnassigned:
			activatingAssignmentMap[assignment.ID] = assignment

			// ポートごとにveth, netns名を割り当てる(NodeServiceないでユニーク)
			netnsPorts := []compute_models.NetnsPort{}
			for j, port := range assignment.Spec.Ports {
				// インターフェイスの最大文字数が15なので、ベース文字数は12とする
				var netnsId int
				for id, assigned := range assignedNetnsIds {
					if !assigned {
						netnsId = id
						break
					}
				}
				netnsName := fmt.Sprintf("com-%d", netnsId)
				netnsGateway := ip_utils.AddIntToIp(srv.computeConf.VmNetnsGatewayStartIp, j)
				netnsIp := ip_utils.AddIntToIp(srv.computeConf.VmNetnsStartIp, netnsId)

				netnsPort := compute_models.NetnsPort{
					Id:           netnsId,
					Name:         netnsName,
					NetnsGateway: netnsGateway.String(),
					NetnsIp:      netnsIp.String(),
					VmIp:         port.Ip,
					VmMac:        port.Mac,
				}

				netnsPorts = append(netnsPorts, netnsPort)
				computeNetnsPortsMap[assignment.ID] = netnsPorts
			}

		case base_const.StatusDeleting:
			deletingAssignmentMap[assignment.ID] = assignment
		}
	}

	if err = srv.computeDriver.SyncActivatingAssignmentMap(tctx, activatingAssignmentMap, computeNetnsPortsMap); err != nil {
		return nil, err
	}

	ok = false
	retryCount = srv.computeConf.ConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmActivatingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
			return nil, err
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				return nil, error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
			}
			time.Sleep(srv.computeConf.ConfirmRetryInterval)
			retryCount -= 1
		}
	}

	if err = srv.computeDriver.SyncDeletingAssignmentMap(tctx, deletingAssignmentMap); err != nil {
		return nil, err
	}

	ok = false
	retryCount = srv.computeConf.ConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmDeletingAssignmentMap(tctx, deletingAssignmentMap); err != nil {
			return nil, err
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				return nil, error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
			}
			time.Sleep(srv.computeConf.ConfirmRetryInterval)
			retryCount -= 1
		}
	}

	for _, assignment := range activatingAssignmentMap {
		assignmentReports = append(assignmentReports, spec.AssignmentReport{
			ID:           assignment.ID,
			UpdatedAt:    assignment.UpdatedAt,
			Status:       resource_model.StatusActive,
			StatusReason: "Created",
		})
	}

	for _, assignment := range deletingAssignmentMap {
		assignmentReports = append(assignmentReports, spec.AssignmentReport{
			ID:           assignment.ID,
			UpdatedAt:    assignment.UpdatedAt,
			Status:       resource_model.StatusDeleted,
			StatusReason: "Deleted",
		})
	}

	return assignmentReports, nil
}
