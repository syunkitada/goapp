package resource_model_api

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNode(tctx *logger.TraceContext, req *resource_api_grpc_pb.GetNodeRequest) *resource_api_grpc_pb.GetNodeReply {
	rep := &resource_api_grpc_pb.GetNodeReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer func() { err = db.Close() }()

	clusterNodesMap := map[string]*resource_api_grpc_pb.Nodes{}
	if req.Cluster == "" {
		var nodes []resource_model.Node
		if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
		clusterNodesMap["root"] = modelApi.convertNodes(nodes)
		rep.StatusCode = codes.Ok

		for clusterName, clusterClient := range modelApi.clusterClientMap {
			getNodeReq := &resource_cluster_api_grpc_pb.GetNodeRequest{
				Target: req.Target,
			}
			remoteRep, err := clusterClient.GetNode(getNodeReq)
			if err != nil {
				rep.Err += fmt.Sprintf("Failed get node from cluster(%v) by error(%v), ",
					clusterName, err.Error())
				rep.StatusCode = codes.RemoteClusterError
			}

			clusterNodesMap[clusterName] = modelApi.convertClusterNodes(remoteRep.Nodes)
		}

	} else {
		clusterClient, ok := modelApi.clusterClientMap[req.Cluster]
		if !ok {
			rep.Err = fmt.Sprintf("NotFound cluster: %v", req.Cluster)
			rep.StatusCode = codes.ClientNotFound
			return rep
		}

		getNodeReq := &resource_cluster_api_grpc_pb.GetNodeRequest{
			Target: req.Target,
		}
		remoteRep, err := clusterClient.GetNode(getNodeReq)
		if err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteClusterError
			return rep
		}

		clusterNodesMap[req.Cluster] = modelApi.convertClusterNodes(remoteRep.Nodes)
	}

	rep.ClusterNodesMap = clusterNodesMap
	return rep
}

func (modelApi *ResourceModelApi) UpdateNode(tctx *logger.TraceContext, req *resource_api_grpc_pb.UpdateNodeRequest) *resource_api_grpc_pb.UpdateNodeReply {
	rep := &resource_api_grpc_pb.UpdateNodeReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer func() { err = db.Close() }()

	var node resource_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
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
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
	} else {
		node.State = req.State
		node.StateReason = req.StateReason
		if err = db.Save(&node).Error; err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
	}

	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) SyncRole(tctx *logger.TraceContext, kind string) ([]resource_model.Node, error) {
	var nodes []resource_model.Node
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return nodes, err
	}
	defer func() { err = db.Close() }()

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

func (modelApi *ResourceModelApi) convertNodes(nodes []resource_model.Node) *resource_api_grpc_pb.Nodes {
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

	return &resource_api_grpc_pb.Nodes{
		Nodes: pbNodes,
	}
}

func (modelApi *ResourceModelApi) convertClusterNodes(nodes []*resource_cluster_api_grpc_pb.Node) *resource_api_grpc_pb.Nodes {
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

	return &resource_api_grpc_pb.Nodes{
		Nodes: pbNodes,
	}
}

func (modelApi *ResourceModelApi) CheckNodes(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

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

	return nil
}
