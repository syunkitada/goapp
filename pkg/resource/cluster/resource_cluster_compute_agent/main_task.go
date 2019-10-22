package resource_cluster_compute_agent

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (srv *ResourceClusterComputeAgentServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterComputeAgentServer) UpdateNode(tctx *logger.TraceContext) error {
	fmt.Println("HOGE")
	return nil
}
