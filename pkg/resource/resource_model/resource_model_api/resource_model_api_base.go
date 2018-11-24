package resource_model_api

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNode(req *resource_api_grpc_pb.GetNodeRequest) (*resource_api_grpc_pb.GetNodeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	if req.Cluster != "" {
		clusterClient, ok := modelApi.clusterClientMap[req.Cluster]
		if !ok {
			return nil, fmt.Errorf("NotFound cluster: ", req.Cluster)
		}
		// TODO get cluster nodes
		glog.Info(clusterClient)
		getNodeReq := &resource_cluster_api_grpc_pb.GetNodeRequest{
			Target: req.Target,
		}
		rep, err := clusterClient.GetNode(getNodeReq)
		if err != nil {
			return nil, err
		}

		return &resource_api_grpc_pb.GetNodeReply{
			Nodes: modelApi.convertClusterNodes(rep.Nodes),
		}, nil
	}

	var nodes []resource_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetNodeReply{
		Nodes: modelApi.convertNodes(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) GetCluster(req *resource_api_grpc_pb.GetClusterRequest) (*resource_api_grpc_pb.GetClusterReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Cluster
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetClusterReply{
		Clusters: modelApi.convertClusters(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) UpdateNode(req *resource_api_grpc_pb.UpdateNodeRequest) (*resource_api_grpc_pb.UpdateNodeReply, error) {
	rep := &resource_api_grpc_pb.UpdateNodeReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var node resource_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		node = resource_model.Node{
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

func (modelApi *ResourceModelApi) SyncRole(kind string) ([]resource_model.Node, error) {
	var nodes []resource_model.Node
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
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
		if node.Role == resource_model.RoleLeader {
			if node.Status == resource_model.StatusEnabled && node.State == resource_model.StateUp && node.UpdatedAt.After(downTime) {
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
	newNodes := []resource_model.Node{}
	for _, node := range nodes {
		if isReassignLeader {
			node.Role = resource_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		} else if node.Status == resource_model.StatusEnabled &&
			node.State == resource_model.StateUp &&
			node.UpdatedAt.After(downTime) {

			node.Role = resource_model.RoleLeader
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
			isReassignLeader = true
			glog.Infof("Leader is assigned: %v", node.Name)
		} else {
			node.Role = resource_model.RoleMember
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

func (modelApi *ResourceModelApi) convertNodes(nodes []resource_model.Node) []*resource_api_grpc_pb.Node {
	pbNodes := make([]*resource_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
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
			CreatedAt:    createdAt,
		}
	}

	return pbNodes
}

func (modelApi *ResourceModelApi) convertClusterNodes(nodes []*resource_cluster_api_grpc_pb.Node) []*resource_api_grpc_pb.Node {
	pbNodes := make([]*resource_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &resource_api_grpc_pb.Node{
			Name:         node.Name,
			Kind:         node.Kind,
			Role:         node.Role,
			Status:       node.Status,
			StatusReason: node.StatusReason,
			State:        node.State,
			StateReason:  node.StateReason,
			UpdatedAt:    node.UpdatedAt,
			CreatedAt:    node.CreatedAt,
		}
	}

	return pbNodes
}

func (modelApi *ResourceModelApi) convertClusters(clusters []resource_model.Cluster) []*resource_api_grpc_pb.Cluster {
	pbClusters := make([]*resource_api_grpc_pb.Cluster, len(clusters))
	for i, cluster := range clusters {
		updatedAt, err := ptypes.TimestampProto(cluster.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(cluster.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbClusters[i] = &resource_api_grpc_pb.Cluster{
			Name:      cluster.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbClusters
}

func (modelApi *ResourceModelApi) CheckNodes() error {
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	tx := db.Begin()
	defer tx.Rollback()
	var nodes []resource_model.Node
	if err = tx.Find(&nodes).Error; err != nil {
		return err
	}

	downTimeDuration := -1 * time.Duration(modelApi.conf.Resource.AppDownTime) * time.Second
	downTime := time.Now().Add(downTimeDuration)

	for _, node := range nodes {
		if node.UpdatedAt.Before(downTime) {
			node.State = resource_model.StateDown
			if err = tx.Save(&node).Error; err != nil {
				return err
			}
		}
	}
	tx.Commit()

	glog.Info("Completed CheckNodes")
	return nil
}
