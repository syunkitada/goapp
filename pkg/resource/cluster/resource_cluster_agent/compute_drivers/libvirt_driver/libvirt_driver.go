package libvirt_driver

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type LibvirtDriver struct {
	name string
	conf *config.Config
}

func New(conf *config.Config) *LibvirtDriver {
	driver := LibvirtDriver{
		name: "libvirt",
		conf: conf,
	}
	return &driver
}

func (driver *LibvirtDriver) GetName() string {
	return ""
}

func (driver *LibvirtDriver) Deploy(tctx *logger.TraceContext) error {
	return driver.deploy(tctx)
}

func (driver *LibvirtDriver) ConfirmDeploy(tctx *logger.TraceContext) (bool, error) {
	return false, nil
}

func (driver *LibvirtDriver) SyncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	driver.syncActivatingAssignmentMap(tctx, assignmentMap)
	return nil
}

func (driver *LibvirtDriver) ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *LibvirtDriver) SyncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	return nil
}

func (driver *LibvirtDriver) ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error) {
	return true, nil
}
