package resource_model_api

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNode(req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.Conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.Conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetNodeReply{
		Nodes: modelApi.ConvertNodes(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) UpdateNode(req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_api_grpc_pb.UpdateNodeReply
	var err error

	db, err := gorm.Open("mysql", modelApi.Conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.Conf.Default.EnableDatabaseLog)

	var node resource_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		node = resource_model.Node{
			Name:         req.Name,
			Kind:         req.Kind,
			Role:         req.Role,
			Status:       resource_model.StatusDisabled,
			StatusReason: "Default status",
			State:        req.State,
			StateReason:  req.StateReason,
		}
		if err = db.Create(&node).Error; err != nil {
			return rep, err
		}
	} else {
		if req.Status != "" && req.StatusReason != "" {
			node.Status = req.Status
			node.StatusReason = req.StatusReason
		}
		node.State = req.State
		node.StateReason = req.StateReason
		if err = db.Save(&node).Error; err != nil {
			return rep, err
		}
	}

	glog.Info("Completed UpdateNode")
	return rep, err
}

func (modelApi *ResourceModelApi) ReassignRole(req *resource_api_grpc_pb.ReassignRoleRequest) (*resource_api_grpc_pb.ReassignRoleReply, error) {
	var rep *resource_api_grpc_pb.ReassignRoleReply
	var err error

	db, err := gorm.Open("mysql", modelApi.Conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.Conf.Default.EnableDatabaseLog)

	tx := db.Begin()
	defer tx.Rollback()
	var nodes []resource_model.Node
	if err = tx.Where("kind = ?", req.Kind).Find(&nodes).Error; err != nil {
		return nil, err
	}
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Role == resource_model.RoleLeader {
			if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
				existsActiveLeader = true
			}
			break
		}
	}
	if existsActiveLeader {
		return &resource_api_grpc_pb.ReassignRoleReply{
			Nodes: modelApi.ConvertNodes(nodes),
		}, nil
	}

	isReassignLeader := false
	for _, node := range nodes {
		if isReassignLeader {
			node.Role = resource_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		} else if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp {
			node.Role = resource_model.RoleLeader
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
			isReassignLeader = true
		}
	}

	glog.Info("Completed UpdateNode")
	tx.Commit()

	return &resource_api_grpc_pb.ReassignRoleReply{
		Nodes: modelApi.ConvertNodes(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) ConvertNodes(nodes []resource_model.Node) []*resource_api_grpc_pb.Node {
	pbNodes := make([]*resource_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		glog.Info(node.Model.UpdatedAt)
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbNodes[i] = &resource_api_grpc_pb.Node{
			Name:         node.Name,
			Kind:         node.Kind,
			Role:         node.Role,
			Status:       node.Status,
			StatusReason: node.StatusReason,
			State:        node.State,
			StateReason:  node.StateReason,
			UpdatedAt:    updatedAt,
		}
	}

	return pbNodes
}
