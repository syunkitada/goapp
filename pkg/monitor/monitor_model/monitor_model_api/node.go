package monitor_model_api

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

func (modelApi *MonitorModelApi) GetNode(req *monitor_api_grpc_pb.GetNodeRequest) *monitor_api_grpc_pb.GetNodeReply {
	rep := &monitor_api_grpc_pb.GetNodeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []monitor_model.Node
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Nodes = modelApi.convertNodes(nodes)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) UpdateNode(req *monitor_api_grpc_pb.UpdateNodeRequest) *monitor_api_grpc_pb.UpdateNodeReply {
	rep := &monitor_api_grpc_pb.UpdateNodeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

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

func (modelApi *MonitorModelApi) SyncRole(kind string) ([]monitor_model.Node, error) {
	var nodes []monitor_model.Node
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
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
		if node.Role == monitor_model.RoleLeader {
			if node.Status == monitor_model.StatusEnabled && node.State == monitor_model.StateUp && node.UpdatedAt.After(downTime) {
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
			glog.Infof("Leader is assigned: %v", node.Name)
		} else {
			node.Role = monitor_model.RoleMember
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

func (modelApi *MonitorModelApi) convertNodes(nodes []monitor_model.Node) []*monitor_api_grpc_pb.Node {
	pbNodes := make([]*monitor_api_grpc_pb.Node, len(nodes))
	for i, node := range nodes {
		updatedAt, err := ptypes.TimestampProto(node.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(node.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
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

func (modelApi *MonitorModelApi) CheckNodes() error {
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Monitor.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

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

	glog.Info("Completed CheckNodes")
	return nil
}
