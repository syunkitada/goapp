package compute_drivers

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/libvirt_driver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/qemu_driver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

type ComputeDriver interface {
	GetName() string
	Deploy(tctx *logger.TraceContext) error
	ConfirmDeploy(tctx *logger.TraceContext) (bool, error)
	SyncActivatingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]resource_model.ComputeAssignmentEx,
		computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error
	ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error)
	SyncDeletingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]resource_model.ComputeAssignmentEx) error
	ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]resource_model.ComputeAssignmentEx) (bool, error)
}

func Load(conf *config.Config) ComputeDriver {
	switch conf.Resource.Node.Compute.Driver {
	case "libvirt":
		driver := libvirt_driver.New(conf)
		return driver
	case "qemu":
		driver := qemu_driver.New(conf)
		return driver
	}

	return nil
}
