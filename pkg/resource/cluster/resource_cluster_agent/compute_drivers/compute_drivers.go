package compute_drivers

import (
	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/mock_driver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_drivers/qemu_driver"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type ComputeDriver interface {
	GetName() string
	Deploy(tctx *logger.TraceContext) error
	ConfirmDeploy(tctx *logger.TraceContext) (bool, error)
	SyncActivatingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]spec.ComputeAssignmentEx,
		computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error
	ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error)
	SyncDeletingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]spec.ComputeAssignmentEx) error
	ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
		assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error)
	ProxyConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, conn *websocket.Conn) error
}

func Load(conf *config.ResourceComputeExConfig) ComputeDriver {
	switch conf.Driver {
	case "mock":
		driver := mock_driver.New(conf)
		return driver
	case "qemu":
		driver := qemu_driver.New(conf)
		return driver
	}

	return nil
}
