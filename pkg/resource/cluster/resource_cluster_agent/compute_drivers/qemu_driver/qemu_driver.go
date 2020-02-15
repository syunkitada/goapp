package qemu_driver

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type QemuDriver struct {
	conf *config.ResourceComputeExConfig
	name string
}

func New(conf *config.ResourceComputeExConfig) *QemuDriver {
	driver := QemuDriver{
		conf: conf,
		name: "qemu",
	}
	return &driver
}

func (driver *QemuDriver) GetName() string {
	return ""
}

func (driver *QemuDriver) Deploy(tctx *logger.TraceContext) error {
	return nil
}

func (driver *QemuDriver) ConfirmDeploy(tctx *logger.TraceContext) (bool, error) {
	return false, nil
}

func (driver *QemuDriver) SyncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
	return driver.syncActivatingAssignmentMap(tctx, assignmentMap, computeNetnsPortsMap)
}

func (driver *QemuDriver) ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *QemuDriver) SyncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) error {
	return driver.syncDeletingAssignmentMap(tctx, assignmentMap)
}

func (driver *QemuDriver) ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}
