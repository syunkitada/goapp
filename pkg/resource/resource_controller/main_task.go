package resource_controller

import (
	"fmt"

	"github.com/golang/glog"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (srv *ResourceControllerServer) MainTask() error {
	glog.Info("Run MainTask")
	srv.UpdateNode()
	srv.MonitorTask()

	return nil
}

func (srv *ResourceControllerServer) MonitorTask() error {
	var err error
	rep, err := srv.resourceApiClient.GetNode(&resource_api_grpc_pb.GetNodeRequest{
		Target: "%",
	})
	if err != nil {
		return err
	}
	existsSelfNode := false
	var selfRole string
	existsActiveLeader := false
	lenActiveMembers := 0
	for _, node := range rep.Nodes {
		if node.Kind != resource_model.KindResourceController {
			continue
		}
		if node.Name == srv.conf.Default.Name && node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
			existsSelfNode = true
			selfRole = node.Role
		}
		if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
			if node.Role == resource_model.RoleLeader {
				existsActiveLeader = true
			} else {
				lenActiveMembers += 1
			}
		}
	}

	if !existsSelfNode {
		return fmt.Errorf("This node is not activated")
	}

	if !existsActiveLeader {
		glog.Info("Active Leader is not exists, all node will be reassigned")
		rep, err := srv.resourceApiClient.ReassignRole(&resource_api_grpc_pb.ReassignRoleRequest{
			Kind: resource_model.KindResourceController,
		})
		if err != nil {
			return err
		}

		for _, node := range rep.Nodes {
			lenActiveMembers = 0
			if node.Kind != resource_model.KindResourceController {
				continue
			}
			if node.Name == srv.conf.Default.Name {
				existsSelfNode = true
				selfRole = node.Role
			}
			if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
				if node.Role == resource_model.RoleLeader {
					existsActiveLeader = true
				} else {
					lenActiveMembers += 1
				}
			}
		}
	}

	if !existsSelfNode {
		return fmt.Errorf("This node is not activated")
	}

	if !existsActiveLeader {
		return fmt.Errorf("Active Leader is not exists, after ReassignNode")
	}

	srv.role = selfRole

	glog.Info("Completed MonitorTask")
	return nil
}

func (server *ResourceControllerServer) UpdateNode() error {
	request := resource_api_grpc_pb.UpdateNodeRequest{
		Name:         server.conf.Default.Name,
		Kind:         resource_model.KindResourceController,
		Role:         resource_model.RoleMember,
		Status:       resource_model.StatusEnabled,
		StatusReason: "Always Enabled",
		State:        resource_model.StateUp,
		StateReason:  "UpdateNode",
	}
	server.resourceApiClient.UpdateNode(&request)

	glog.Info("UpdatedNode")
	return nil
}
