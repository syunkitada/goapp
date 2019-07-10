package resource_cluster_agent

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceClusterAgentServer) SyncComputeAssignments(tctx *logger.TraceContext,
	assignments []resource_model.ComputeAssignmentWithComputeAndNode) error {

	fmt.Println("SyncComputeAssignments: ", assignments)

	return nil
}
