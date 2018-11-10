package resource_cluster_controller

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (srv *ResourceClusterControllerServer) MainTask() error {
	glog.Info("Run MainTask")
	if err := srv.UpdateNode(); err != nil {
		return err
	}
	if err := srv.SyncRole(); err != nil {
		return err
	}
	if srv.role == resource_cluster_model.RoleMember {
		return nil
	}

	if err := srv.resourceClusterModelApi.CheckNodes(); err != nil {
		return err
	}

	return nil
}

func (srv *ResourceClusterControllerServer) UpdateNode() error {
	req := resource_cluster_api_grpc_pb.UpdateNodeRequest{
		Name:         srv.conf.Default.Name,
		Kind:         resource_cluster_model.KindResourceClusterController,
		Role:         resource_cluster_model.RoleMember,
		Status:       resource_cluster_model.StatusEnabled,
		StatusReason: "Always Enabled",
		State:        resource_cluster_model.StateUp,
		StateReason:  "UpdateNode",
	}
	if _, err := srv.resourceClusterApiClient.UpdateNode(&req); err != nil {
		return err
	}

	glog.Info("UpdatedNode")
	return nil
}

func (srv *ResourceClusterControllerServer) SyncRole() error {
	var err error
	nodes, err := srv.resourceClusterModelApi.SyncRole(resource_cluster_model.KindResourceClusterController)
	if err != nil {
		return err
	}

	existsSelfNode := false
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Kind != resource_cluster_model.KindResourceClusterController {
			continue
		}
		if node.Name == srv.conf.Default.Name && node.Status == resource_cluster_model.StatusEnabled && node.State == resource_cluster_model.StateUp {
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

	glog.Infof("Completed SyncRole: role=%v", srv.role)
	return nil
}
