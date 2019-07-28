package resource_cluster_agent

import (
	"fmt"
	"strings"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_utils"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceClusterAgentServer) SyncComputeAssignments(tctx *logger.TraceContext,
	assignments []resource_model.ComputeAssignmentEx) error {
	var err error
	var ok bool
	var retryCount int
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tctx.SetTimeout(1)

	fmt.Println("SyncComputeAssignments: ", assignments)

	netnsSet, err := os_utils.GetNetnsSet(tctx)
	if err != nil {
		return err
	}

	assignedNetnsPortIds := make([]bool, 4096)
	computeNetnsPortsMap := map[uint][]compute_models.NetnsPort{}
	activatingAssignmentMap := map[uint]resource_model.ComputeAssignmentEx{}
	deletingAssignmentMap := map[uint]resource_model.ComputeAssignmentEx{}
	for _, assignment := range assignments {
		switch assignment.Status {
		case resource_model.StatusActive, resource_model.StatusCreating, resource_model.StatusUpdating,
			resource_model.StatusUnknownActivating, resource_model.StatusRebalancingUnassigned:
			activatingAssignmentMap[assignment.ID] = assignment
		case resource_model.StatusDeleting:
			deletingAssignmentMap[assignment.ID] = assignment
		}

		netnsPorts := []compute_models.NetnsPort{}
		for _, port := range assignment.Spec.Compute.Ports {
			netnsPortId := compute_utils.AssignNetnsPortId(assignedNetnsPortIds)
			netnsName := fmt.Sprintf("com%d", netnsPortId)
			gatewayIp := ip_utils.AddIntToIp(srv.vmNetnsStartIp, netnsPortId*4)
			netnsGateway := gatewayIp.String()
			ip_utils.IncrementIp(gatewayIp)
			netnsAddr := fmt.Sprintf("%s/30", gatewayIp.String())

			shareNetnsIp := ip_utils.AddIntToIp(srv.shareNetnsVmStartIp, netnsPortId)

			splitedSubnet := strings.Split(port.Subnet, "/")
			netnsPort := compute_models.NetnsPort{
				Id:             netnsPortId,
				Name:           netnsName,
				ShareNetnsAddr: fmt.Sprintf("%s/%s", shareNetnsIp, srv.shareNetnsAddrSuffix),
				NetnsGateway:   netnsGateway,
				NetnsAddr:      netnsAddr,
				VmGateway:      port.Gateway,
				VmIp:           port.Ip,
				VmAddr:         fmt.Sprintf("%s/%s", port.Ip, splitedSubnet[1]),
				VmMac:          port.Mac,
			}

			if _, ok := netnsSet[netnsName]; !ok {
				if err = os_utils.AddNetns(tctx, netnsName); err != nil {
					return err
				}
				if err = compute_utils.InitNetns(tctx, &srv.conf.Resource.Node.Compute, netnsName, netnsPort); err != nil {
					return err
				}
			}

			netnsPorts = append(netnsPorts, netnsPort)
		}
		computeNetnsPortsMap[assignment.ID] = netnsPorts
	}

	if err = srv.computeDriver.SyncActivatingAssignmentMap(tctx, activatingAssignmentMap, computeNetnsPortsMap); err != nil {
		return err
	}

	ok = false
	retryCount = srv.computeConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmActivatingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
			return err
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				return error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
			}
			time.Sleep(srv.computeConfirmRetryInterval)
			retryCount -= 1
		}
	}

	if err = srv.computeDriver.SyncDeletingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
		return err
	}

	ok = false
	retryCount = srv.computeConfirmRetryCount
	for {
		if ok, err = srv.computeDriver.ConfirmDeletingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
			return err
		}
		if ok {
			break
		} else {
			if retryCount == 0 {
				return error_utils.NewTimeoutExceededError("ConfirmActivatingAssignmentMap")
			}
			time.Sleep(srv.computeConfirmRetryInterval)
			retryCount -= 1
		}
	}

	return nil
}
