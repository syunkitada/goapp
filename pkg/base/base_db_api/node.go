package base_db_api

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (api *Api) GetNodeService(tctx *logger.TraceContext, input *base_spec.GetNodeService, user *base_spec.UserAuthority) (data *base_spec.NodeService, err error) {
	data = &base_spec.NodeService{}
	err = api.DB.Where("name = ? AND kind = ? AND deleted_at IS NULL", input.Name, input.Kind).
		First(data).Error
	return
}

func (api *Api) GetNodeServices(tctx *logger.TraceContext, input *base_spec.GetNodeServices, user *base_spec.UserAuthority) (data []base_spec.NodeService, err error) {
	err = api.DB.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreateOrUpdateNodeService(tctx *logger.TraceContext, input *base_spec.UpdateNodeService) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var tmp base_db_model.NodeService
		node := input.NodeService
		if err = tx.Where("name = ? and kind = ?", node.Name, node.Kind).First(&tmp).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			tmp = base_db_model.NodeService{
				Name:         node.Name,
				Kind:         node.Kind,
				Role:         node.Role,
				Status:       node.Status,
				StatusReason: node.StatusReason,
				State:        node.State,
				StateReason:  node.StateReason,
			}
			if err = tx.Create(&tmp).Error; err != nil {
				return
			}
		} else {
			tmp.State = node.State
			tmp.StateReason = node.StateReason
			if err = tx.Save(&tmp).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteNodeService(tctx *logger.TraceContext, input *base_spec.DeleteNodeService) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND kind = ?", input.Name, input.Kind).Delete(&base_db_model.NodeService{}).Error
		return
	})
	return
}

func (api *Api) DeleteNodeServices(tctx *logger.TraceContext, input []base_spec.NodeService) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND kind = ?", val.Name, val.Kind).
				Delete(&db_model.Image{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) SyncNodeServiceRole(tctx *logger.TraceContext, kind string) (nodes []base_db_model.NodeService, err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		if err = tx.Where("kind = ? AND deleted_at IS NULL", kind).Find(&nodes).Error; err != nil {
			return
		}

		downTime := time.Now().Add(api.nodeDownTimeDuration)
		existsActiveLeader := false
		for _, node := range nodes {
			if node.Role == base_const.RoleLeader {
				if node.Status == base_const.StatusEnabled && node.State == base_const.StateUp && node.UpdatedAt.After(downTime) {
					existsActiveLeader = true
				}
				break
			}
		}
		if existsActiveLeader {
			return
		}
		logger.Info(tctx, "Active Leader is not exists, Leader will be assigned.")

		isReassignLeader := false
		newNodeServices := []base_db_model.NodeService{}
		for _, node := range nodes {
			if isReassignLeader {
				node.Role = base_const.RoleMember
				if err = tx.Save(&node).Error; err != nil {
					return
				}
			} else if node.Status == base_const.StatusEnabled &&
				node.State == base_const.StateUp &&
				node.UpdatedAt.Before(downTime) {

				node.Role = base_const.RoleLeader
				if err = tx.Save(&node).Error; err != nil {
					return
				}
				isReassignLeader = true
				logger.Infof(tctx, "Leader is assigned: %v", node.Name)
			} else {
				node.Role = base_const.RoleMember
				if err = tx.Save(&node).Error; err != nil {
					return
				}
			}
			newNodeServices = append(newNodeServices, node)
		}
		nodes = newNodeServices
		return
	})
	return
}

func (api *Api) SyncNodeServiceState(tctx *logger.TraceContext) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var nodes []base_db_model.NodeService
		if err = tx.Find(&nodes).Error; err != nil {
			return
		}

		downTime := time.Now().Add(api.nodeDownTimeDuration)
		for _, node := range nodes {
			if node.UpdatedAt.After(downTime) {
				node.State = resource_model.StateDown
				if err = tx.Save(&node).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}
