package resource_model_api

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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

func (modelApi *ResourceModelApi) UpdateNode(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NodeSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestError("Specs is empty")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		var data resource_model.Node
		if err = db.Where("name = ? and kind = ?", spec.Name, spec.Kind).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Node{
				Name:         spec.Name,
				Kind:         spec.Kind,
				Role:         spec.Role,
				Status:       spec.Status,
				StatusReason: spec.StatusReason,
				State:        spec.State,
				StateReason:  spec.StateReason,
			}
			if err = tx.Create(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		} else {
			data.State = spec.State
			data.StateReason = spec.StateReason
			if err = db.Save(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
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
