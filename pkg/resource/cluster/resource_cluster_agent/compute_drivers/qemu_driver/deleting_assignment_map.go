package qemu_driver

import (
	"fmt"
	"path/filepath"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (driver *QemuDriver) syncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) error {
	var err error
	for _, assignment := range assignmentMap {
		if err = driver.syncDeletingAssignment(tctx, assignment); err != nil {
			return err
		}
	}
	return nil
}

func (driver *QemuDriver) syncDeletingAssignment(tctx *logger.TraceContext,
	assignment spec.ComputeAssignmentEx) error {
	var err error
	compute := assignment.Spec
	fmt.Println("DEBUG QEMU syncDeletingAssignment")

	if _, err = exec_utils.Cmdf(tctx, "systemctl stop %s", compute.Name); err != nil {
		return err
	}

	vmDir := filepath.Join(driver.conf.VmsDir, compute.Name)
	if err = os_utils.Rmdir(vmDir); err != nil {
		return err
	}

	return nil
}
