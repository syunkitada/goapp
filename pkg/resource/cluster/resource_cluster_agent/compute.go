package resource_cluster_agent

import (
	"fmt"
	"time"

	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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
	}

	if err = srv.computeDriver.SyncActivatingAssignmentMap(tctx, activatingAssignmentMap); err != nil {
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
