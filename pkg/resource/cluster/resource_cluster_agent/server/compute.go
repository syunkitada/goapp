package server

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *Server) ComputeLoop() {
	var tctx *logger.TraceContext
	var err error
	var startTime time.Time
	loopInterval := 10 * time.Second
	logger.StdoutInfof("Start ComputeLoop")
	for {
		tctx = srv.NewTraceContext()
		startTime = logger.StartTrace(tctx)
		err = srv.SyncComputeAssignments(tctx)
		logger.EndTrace(tctx, startTime, err, 0)
		time.Sleep(loopInterval)
	}
}

func (srv *Server) SyncComputeAssignments(tctx *logger.TraceContext) (err error) {
	var ok bool
	var retryCount int
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tctx.SetTimeout(1)

	srv.computeAssignmentsMutex.Lock()
	assignments := srv.computeAssignments
	srv.computeAssignmentsMutex.Unlock()

	fmt.Println("SyncComputeAssignments: ", assignments)

	assignmentReports := []spec.AssignmentReport{}

	// ユニークなnetns管理用
	// TODO vm単位でnetnsを保存する
	// 全vmのnetnsを取得する
	assignedNetnsIds := make([]bool, 4096)
	var netnsSet map[string]bool
	if netnsSet, err = os_utils.GetNetnsSet(tctx); err != nil {
		return
	}
	for netns := range netnsSet {
		splitedNetns := strings.Split(netns, "com-")
		if len(splitedNetns) == 2 {
			if id, tmpErr := strconv.Atoi(splitedNetns[1]); tmpErr != nil {
				continue
			} else if id < 4096 {
				assignedNetnsIds[id] = true
			}
		}
	}

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
						assignedNetnsIds[id] = true
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
					VmSubnet:     port.Subnet,
					Kind:         port.Kind,
				}

				switch port.Kind {
				case "Local":
					fmt.Println("DEBUG Kind", port.Spec)
					var netSpec spec.NetworkV4LocalSpec
					if tmpErr := json_utils.Unmarshal(port.Spec, &netSpec); tmpErr != nil {
						logger.Warningf(tctx, "Invalid port.Spec: %s", tmpErr.Error())
						continue
					}
					netnsPort.NetworkV4LocalSpec = netSpec
				}

				netnsPorts = append(netnsPorts, netnsPort)
				computeNetnsPortsMap[assignment.ID] = netnsPorts
			}

		case base_const.StatusDeleting:
			deletingAssignmentMap[assignment.ID] = assignment
		}
	}

	fmt.Println("DEBUG Activating")
	if err = srv.computeDriver.SyncActivatingAssignmentMap(tctx, activatingAssignmentMap, computeNetnsPortsMap); err != nil {
		return
	}

	fmt.Println("DEBUG Activating2")
	retryCount = srv.computeConf.ConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmActivatingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
			return
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				err = error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
			}
			time.Sleep(srv.computeConf.ConfirmRetryInterval)
			retryCount -= 1
		}
	}

	fmt.Println("DEBUG Activating3", deletingAssignmentMap)

	if err = srv.computeDriver.SyncDeletingAssignmentMap(tctx, deletingAssignmentMap); err != nil {
		return
	}
	fmt.Println("DEBUG Activating4")

	retryCount = srv.computeConf.ConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmDeletingAssignmentMap(tctx, deletingAssignmentMap); err != nil {
			return
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				err = error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
				return
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

	srv.computeAssignmentReportsMutex.Lock()
	srv.computeAssignmentReports = assignmentReports
	srv.computeAssignmentReportsMutex.Unlock()
	return
}
