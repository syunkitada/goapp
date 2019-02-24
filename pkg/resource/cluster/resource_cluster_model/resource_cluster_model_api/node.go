package resource_cluster_model_api

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (modelApi *ResourceClusterModelApi) GetNode(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.ActionRequest, rep *resource_cluster_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	pbNodes := []*resource_cluster_api_grpc_pb.Node{}

	var nodes []resource_cluster_model.Node
	if err = db.Find(&nodes).Error; err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	rootNodes := modelApi.convertNodes(tctx, nodes)
	pbNodes = append(pbNodes, rootNodes...)
	rep.Tctx.StatusCode = codes.Ok

	rep.Nodes = pbNodes
}

func (modelApi *ResourceClusterModelApi) UpdateNode(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.UpdateNodeRequest, rep *resource_cluster_api_grpc_pb.UpdateNodeReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	var node resource_cluster_model.Node
	reqNode := req.Node.Node
	if err = db.Where("name = ? and kind = ?", reqNode.Name, reqNode.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}

		node = resource_cluster_model.Node{
			Name:         reqNode.Name,
			Kind:         reqNode.Kind,
			Role:         reqNode.Role,
			Status:       reqNode.Status,
			StatusReason: reqNode.StatusReason,
			State:        reqNode.State,
			StateReason:  reqNode.StateReason,
		}
		if err = db.Create(&node).Error; err != nil {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}
	} else {
		node.State = reqNode.State
		node.StateReason = reqNode.StateReason
		if err = db.Save(&node).Error; err != nil {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}
	}

	rep.Tctx.StatusCode = codes.Ok
	return
}

func (modelApi *ResourceClusterModelApi) SyncRole(tctx *logger.TraceContext, kind string) ([]resource_cluster_model.Node, error) {
	var nodes []resource_cluster_model.Node
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
		if node.Role == resource_cluster_model.RoleLeader {
			if node.Status == resource_cluster_model.StatusEnabled && node.State == resource_cluster_model.StateUp && node.UpdatedAt.After(downTime) {
				existsActiveLeader = true
			}
			break
		}
	}
	if existsActiveLeader {
		return nodes, nil
	}
	logger.Info(tctx, "Active Leader is not exists, Leader will be assigned.")

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
			logger.Infof(tctx, "Leader is assigned: %v", node.Name)
		} else {
			node.Role = resource_cluster_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		}
		newNodes = append(newNodes, node)
	}
	tx.Commit()

	return newNodes, nil
}

func (modelApi *ResourceClusterModelApi) convertNodes(tctx *logger.TraceContext, nodes []resource_cluster_model.Node) []*resource_cluster_api_grpc_pb.Node {
	pbNodes := make([]*resource_cluster_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
		if err != nil {
			logger.Warning(tctx, err, "Invalid timestamp")
			continue
		}

		pbNodes[i] = &resource_cluster_api_grpc_pb.Node{
			Node: &resource_api_grpc_pb.Node{
				Name:         node.Name,
				Kind:         node.Kind,
				Role:         node.Role,
				Status:       node.Status,
				StatusReason: node.StatusReason,
				State:        node.State,
				StateReason:  node.StateReason,
				UpdatedAt:    updatedAt,
				CreatedAt:    createdAt,
			},
		}
	}

	return pbNodes
}

func (modelApi *ResourceClusterModelApi) CheckNodes(tctx *logger.TraceContext) error {
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

	return nil
}
