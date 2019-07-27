package qemu_driver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type QemuDriver struct {
	name      string
	conf      *config.Config
	vmsDir    string
	imagesDir string
}

func New(conf *config.Config) *QemuDriver {
	driver := QemuDriver{
		name:      "qemu",
		conf:      conf,
		vmsDir:    conf.Resource.Node.Compute.VmsDir,
		imagesDir: conf.Resource.Node.Compute.ImagesDir,
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
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	fmt.Println("DEBUG Activating", assignmentMap)
	return nil
}

func (driver *QemuDriver) ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *QemuDriver) SyncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	return nil
}

func (driver *QemuDriver) ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error) {
	return true, nil
}
