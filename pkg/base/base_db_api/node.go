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

func (api *Api) GetNode(tctx *logger.TraceContext, input *base_spec.GetNode, user *base_spec.UserAuthority) (data *base_spec.Node, err error) {
	data = &base_spec.Node{}
	err = api.DB.Where("name = ? AND kind = ? AND deleted_at IS NULL", input.Name, input.Kind).
		First(data).Error
	return
}

func (api *Api) GetNodes(tctx *logger.TraceContext, input *base_spec.GetNodes, user *base_spec.UserAuthority) (data []base_spec.Node, err error) {
	err = api.DB.Where("deleted_at IS NULL").Find(&data).Error
	return
}

func (api *Api) CreateOrUpdateNode(tctx *logger.TraceContext, input *base_spec.UpdateNode) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var tmp base_db_model.Node
		node := input.Node
		if err = tx.Where("name = ? and kind = ?", node.Name, node.Kind).First(&tmp).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}

			tmp = base_db_model.Node{
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

func (api *Api) DeleteNode(tctx *logger.TraceContext, input *base_spec.DeleteNode) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND kind = ?", input.Name, input.Kind).Delete(&base_db_model.Node{}).Error
		return
	})
	return
}

func (api *Api) DeleteNodes(tctx *logger.TraceContext, input []base_spec.Node) (err error) {
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

func (api *Api) SyncNodeRole(tctx *logger.TraceContext, kind string) (nodes []base_db_model.Node, err error) {
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
		newNodes := []base_db_model.Node{}
		for _, node := range nodes {
			if isReassignLeader {
				node.Role = base_const.RoleMember
				if err = tx.Save(&node).Error; err != nil {
					return
				}
			} else if node.Status == base_const.StatusEnabled &&
				node.State == base_const.StateUp &&
				node.UpdatedAt.After(downTime) {

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
			newNodes = append(newNodes, node)
		}
		nodes = newNodes
		return
	})
	return
}

func (api *Api) CheckNodes(tctx *logger.TraceContext) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var nodes []base_db_model.Node
		if err = tx.Find(&nodes).Error; err != nil {
			return
		}

		downTime := time.Now().Add(-1 * api.nodeDownTimeDuration)
		for _, node := range nodes {
			if node.UpdatedAt.Before(downTime) {
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
