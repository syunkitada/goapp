package resource_cluster_controller

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (srv *ResourceClusterControllerServer) MainTask(tctx *logger.TraceContext) error {
	if err := srv.UpdateNode(tctx); err != nil {
		return err
	}
	if err := srv.SyncRole(tctx); err != nil {
		return err
	}
	if srv.role == resource_cluster_model.RoleMember {
		return nil
	}

	if err := srv.resourceClusterModelApi.CheckNodes(tctx); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterControllerServer) UpdateNode(tctx *logger.TraceContext) error {
	node := &resource_cluster_api_grpc_pb.Node{
		Node: &resource_api_grpc_pb.Node{
			Name:         srv.conf.Default.Host,
			Kind:         resource_cluster_model.KindResourceClusterController,
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

func (srv *ResourceClusterControllerServer) SyncRole(tctx *logger.TraceContext) error {
	var err error
	nodes, err := srv.resourceClusterModelApi.SyncRole(tctx, resource_cluster_model.KindResourceClusterController)
	if err != nil {
		return err
	}

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != resource_cluster_model.KindResourceClusterController {
			continue
		}
		if node.Name == srv.conf.Default.Host && node.Status == resource_cluster_model.StatusEnabled && node.State == resource_cluster_model.StateUp {
			existsSelfNode = true
			srv.role = node.Role
		}
		if node.Status == resource_cluster_model.StatusEnabled && node.State == resource_cluster_model.StateUp {
			if node.Role == resource_cluster_model.RoleLeader {
				existsActiveLeader = true
			}
		}
	}

	if !existsSelfNode {
		return fmt.Errorf("This node is not activated")
	}

	if !existsActiveLeader {
		return fmt.Errorf("Active Leader is not exists, after ReassignNode")
	}

	return nil
}
