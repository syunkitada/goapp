package resource_model_api

import (
	"encoding/json"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNode(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, error_utils.NewInvalidRequestError("resource is None")
	}

	var node resource_model.Node
	if err = db.Where(&resource_model.Node{
		Name: resource,
	}).First(&node).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Node"] = node
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetNodes(tctx *logger.TraceContext,
	db *gorm.DB, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error

	var nodes []resource_model.Node
	if err = db.Find(&nodes).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Nodes"] = nodes
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) UpdateNode(tctx *logger.TraceContext, req *resource_api_grpc_pb.UpdateNodeRequest, rep *resource_api_grpc_pb.UpdateNodeReply) {
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

	var node resource_model.Node
	reqNode := req.Node
	if err = db.Where("name = ? and kind = ?", reqNode.Name, reqNode.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}

		node = resource_model.Node{
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
}

func (modelApi *ResourceModelApi) DeleteNode(tctx *logger.TraceContext, db *gorm.DB, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NameSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		if err = tx.Delete(&resource_model.Node{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
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
			logger.Infof(tctx, "Leader is assigned: %v", node.Name)
		} else {
			node.Role = resource_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		}
		newNodes = append(newNodes, node)
	}
	tx.Commit()

	return newNodes, nil
}

func (modelApi *ResourceModelApi) convertNodes(tctx *logger.TraceContext, clusterName string, nodes []resource_model.Node) []*resource_api_grpc_pb.Node {
	pbNodes := make([]*resource_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
		if err != nil {
			logger.Warning(tctx, err, "Invalid timestamp")
			continue
		}

		pbNodes[i] = &resource_api_grpc_pb.Node{
			Cluster:      clusterName,
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

func (modelApi *ResourceModelApi) convertClusterNodes(clusterName string, nodes []*resource_cluster_api_grpc_pb.Node) []*resource_api_grpc_pb.Node {
	pbNodes := make([]*resource_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		pbNodes[i] = &resource_api_grpc_pb.Node{
			Cluster:      clusterName,
			Name:         node.Node.Name,
			Kind:         node.Node.Kind,
			Role:         node.Node.Role,
			Status:       node.Node.Status,
			StatusReason: node.Node.StatusReason,
			State:        node.Node.State,
			StateReason:  node.Node.StateReason,
			UpdatedAt:    node.Node.UpdatedAt,
			CreatedAt:    node.Node.CreatedAt,
		}
	}

	return pbNodes
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
