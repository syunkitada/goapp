package qemu_driver

import (
	"path/filepath"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type QemuDriver struct {
	name            string
	conf            *config.Config
	vmsDir          string
	imagesDir       string
	systemdDir      string
	userdataTmpl    string
	vmServiceTmpl   string
	vmServiceShTmpl string
}

func New(conf *config.Config) *QemuDriver {
	userdataTmpl := filepath.Join(conf.Resource.Node.Compute.ConfigDir, "user-data.tmpl")
	vmServiceTmpl := filepath.Join(conf.Resource.Node.Compute.ConfigDir, "vm-service.tmpl")
	vmServiceShTmpl := filepath.Join(conf.Resource.Node.Compute.ConfigDir, "vm-service.sh.tmpl")

	driver := QemuDriver{
		name:            "qemu",
		conf:            conf,
		vmsDir:          conf.Resource.Node.Compute.VmsDir,
		imagesDir:       conf.Resource.Node.Compute.ImagesDir,
		systemdDir:      "/etc/systemd/system",
		userdataTmpl:    userdataTmpl,
		vmServiceTmpl:   vmServiceTmpl,
		vmServiceShTmpl: vmServiceShTmpl,
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
	assignmentMap map[uint]resource_model.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
	return driver.syncActivatingAssignmentMap(tctx, assignmentMap, computeNetnsPortsMap)
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
