package mock_driver

import (
	"github.com/gorilla/websocket"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type MockDriver struct {
	conf *config.ResourceComputeExConfig
	name string
}

func New(conf *config.ResourceComputeExConfig) *MockDriver {
	driver := MockDriver{
		conf: conf,
		name: "mock",
	}
	return &driver
}

func (driver *MockDriver) GetName() string {
	return ""
}

func (driver *MockDriver) Deploy(tctx *logger.TraceContext) error {
	return nil
}

func (driver *MockDriver) ConfirmDeploy(tctx *logger.TraceContext) (bool, error) {
	return false, nil
}

func (driver *MockDriver) SyncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx,
	computeNetnsPortsMap map[uint][]compute_models.NetnsPort) error {
	return nil
}

func (driver *MockDriver) ConfirmActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *MockDriver) SyncDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) error {
	return nil
}

func (driver *MockDriver) ConfirmDeletingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]spec.ComputeAssignmentEx) (bool, error) {
	return true, nil
}

func (driver *MockDriver) ProxyConsole(tctx *logger.TraceContext, input *spec.GetComputeConsole, conn *websocket.Conn) (err error) {
	return
}
