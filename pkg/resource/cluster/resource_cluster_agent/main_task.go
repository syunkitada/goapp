package resource_cluster_agent

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (srv *ResourceClusterAgentServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterAgentServer) UpdateNode(tctx *logger.TraceContext) error {
	node := &resource_cluster_api_grpc_pb.Node{
		Node: &resource_api_grpc_pb.Node{
			Name:         srv.conf.Default.Host,
			Kind:         resource_cluster_model.KindResourceClusterApi,
			Role:         resource_cluster_model.RoleMember,
			Status:       resource_cluster_model.StatusEnabled,
			StatusReason: "Default",
			State:        resource_cluster_model.StateUp,
			StateReason:  "UpdateNode",
		},
	}

	if _, err := srv.resourceClusterApiClient.UpdateNode(tctx, node); err != nil {
		return err
	}

	return nil
}
