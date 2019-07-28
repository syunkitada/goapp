package libvirt_driver

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type LibvirtDriver struct {
	name      string
	conf      *config.Config
	vmsDir    string
	imagesDir string
}

func New(conf *config.Config) *LibvirtDriver {
	driver := LibvirtDriver{
		name:      "libvirt",
		conf:      conf,
		vmsDir:    conf.Resource.Node.Compute.VmsDir,
		imagesDir: conf.Resource.Node.Compute.ImagesDir,
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
	assignmentMap map[uint]resource_model.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
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
