package resource_cluster_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
)

func (modelApi *ResourceClusterModelApi) GetNode(req *resource_cluster_api_grpc_pb.GetNodeRequest) (*resource_cluster_api_grpc_pb.GetNodeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_cluster_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	reply := &resource_cluster_api_grpc_pb.GetNodeReply{
		Nodes: modelApi.convertNodes(nodes),
	}

	return reply, nil
}

func (modelApi *ResourceClusterModelApi) UpdateNode(req *resource_cluster_api_grpc_pb.UpdateNodeRequest) (*resource_cluster_api_grpc_pb.UpdateNodeReply, error) {
	var rep *resource_cluster_api_grpc_pb.UpdateNodeReply
	var err error

	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var node resource_cluster_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		node = resource_cluster_model.Node{
			Name:         req.Name,
			Kind:         req.Kind,
			Role:         req.Role,
			Status:       req.Status,
			StatusReason: req.StatusReason,
			State:        req.State,
			StateReason:  req.StateReason,
		}
		if err = db.Create(&node).Error; err != nil {
			return rep, err
		}
	} else {
		node.State = req.State
		node.StateReason = req.StateReason
		if err = db.Save(&node).Error; err != nil {
			return rep, err
		}
	}

	glog.Info("Completed UpdateNode")
	return rep, err
}

func (modelApi *ResourceClusterModelApi) SyncRole(kind string) ([]resource_cluster_model.Node, error) {
	var nodes []resource_cluster_model.Node
	var err error

	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return nodes, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("kind = ?", kind).Find(&nodes).Error; err != nil {
		return nil, err
	}

	downTime := time.Now().Add(modelApi.downTimeDuration)
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Role == resource_cluster_model.RoleLeader {
			if node.Status == resource_cluster_model.StatusEnabled && node.State == resource_cluster_model.StateUp && node.UpdatedAt.After(downTime) {
				glog.Infof("Found Active Leader: %v", node.Name)
				existsActiveLeader = true
			}
			break
		}
	}
	if existsActiveLeader {
		return nodes, nil
	}
	glog.Info("Active Leader is not exists, Leader will be assigned.")

	isReassignLeader := false
	newNodes := []resource_cluster_model.Node{}
	for _, node := range nodes {
		if isReassignLeader {
			node.Role = resource_cluster_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		} else if node.Status == resource_cluster_model.StatusEnabled &&
			node.State == resource_cluster_model.StateUp &&
			node.UpdatedAt.After(downTime) {

			node.Role = resource_cluster_model.RoleLeader
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
			isReassignLeader = true
			glog.Infof("Leader is assigned: %v", node.Name)
		} else {
			node.Role = resource_cluster_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		}
		newNodes = append(newNodes, node)
	}
	tx.Commit()

	glog.Info("Completed SyncNode")
	return newNodes, nil
}

func (modelApi *ResourceClusterModelApi) convertNodes(nodes []resource_cluster_model.Node) []*resource_cluster_api_grpc_pb.Node {
	pbNodes := make([]*resource_cluster_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbNodes[i] = &resource_cluster_api_grpc_pb.Node{
			Name:         node.Name,
			Kind:         node.Kind,
			Role:         node.Role,
			Status:       node.Status,
			StatusReason: node.StatusReason,
			State:        node.State,
			StateReason:  node.StateReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		}
	}

	return pbNodes
}

func (modelApi *ResourceClusterModelApi) CheckNodes() error {
	var err error

	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	tx := db.Begin()
	defer tx.Rollback()
	var nodes []resource_cluster_model.Node
	if err = tx.Find(&nodes).Error; err != nil {
		return err
	}

	downTimeDuration := -1 * time.Duration(modelApi.conf.Resource.AppDownTime) * time.Second
	downTime := time.Now().Add(downTimeDuration)

	for _, node := range nodes {
		if node.UpdatedAt.Before(downTime) {
			node.State = resource_cluster_model.StateDown
			if err = tx.Save(&node).Error; err != nil {
				return err
			}
		}
	}
	tx.Commit()

	glog.Info("Completed CheckNodes")
	return nil
}
