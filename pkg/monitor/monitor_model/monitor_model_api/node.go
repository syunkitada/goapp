package monitor_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

func (modelApi *MonitorModelApi) GetNode(tctx *logger.TraceContext, req *monitor_api_grpc_pb.GetNodeRequest) *monitor_api_grpc_pb.GetNodeReply {
	rep := &monitor_api_grpc_pb.GetNodeReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	var nodes []monitor_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Nodes = modelApi.convertNodes(tctx, nodes)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) UpdateNode(tctx *logger.TraceContext, req *monitor_api_grpc_pb.UpdateNodeRequest) *monitor_api_grpc_pb.UpdateNodeReply {
	rep := &monitor_api_grpc_pb.UpdateNodeReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	var node monitor_model.Node
	if err = db.Where("name = ? and kind = ?", req.Name, req.Kind).First(&node).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}

		node = monitor_model.Node{
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

func (modelApi *MonitorModelApi) SyncRole(tctx *logger.TraceContext, kind string) ([]monitor_model.Node, error) {
	var nodes []monitor_model.Node
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("kind = ?", kind).Find(&nodes).Error; err != nil {
		return nil, err
	}

	downTime := time.Now().Add(modelApi.downTimeDuration)
	existsActiveLeader := false
	for _, node := range nodes {
		if node.Role == monitor_model.RoleLeader {
			if node.Status == monitor_model.StatusEnabled && node.State == monitor_model.StateUp && node.UpdatedAt.After(downTime) {
				logger.Infof(tctx, "Found Active Leader: %v", node.Name)
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
	newNodes := []monitor_model.Node{}
	for _, node := range nodes {
		if isReassignLeader {
			node.Role = monitor_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		} else if node.Status == monitor_model.StatusEnabled &&
			node.State == monitor_model.StateUp &&
			node.UpdatedAt.After(downTime) {

			node.Role = monitor_model.RoleLeader
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
			isReassignLeader = true
			logger.Infof(tctx, "Leader is assigned: %v", node.Name)
		} else {
			node.Role = monitor_model.RoleMember
			if err = tx.Save(&node).Error; err != nil {
				return nil, err
			}
		}
		newNodes = append(newNodes, node)
	}
	tx.Commit()

	return newNodes, nil
}

func (modelApi *MonitorModelApi) convertNodes(tctx *logger.TraceContext, nodes []monitor_model.Node) []*monitor_api_grpc_pb.Node {
	pbNodes := make([]*monitor_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
		if err != nil {
			logger.Warning(tctx, err, "Invalid timestamp")
			continue
		}

		pbNodes[i] = &monitor_api_grpc_pb.Node{
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

func (modelApi *MonitorModelApi) CheckNodes(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.Begin()
	defer tx.Rollback()
	var nodes []monitor_model.Node
	if err = tx.Find(&nodes).Error; err != nil {
		return err
	}

	downTimeDuration := -1 * time.Duration(modelApi.conf.Monitor.AppDownTime) * time.Second
	downTime := time.Now().Add(downTimeDuration)

	for _, node := range nodes {
		if node.UpdatedAt.Before(downTime) {
			node.State = monitor_model.StateDown
			if err = tx.Save(&node).Error; err != nil {
				return err
			}
		}
	}
	tx.Commit()

	return nil
}
