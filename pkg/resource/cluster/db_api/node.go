package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

func (api *Api) SyncNode(tctx *logger.TraceContext, input *api_spec.SyncNode) (nodeTask *api_spec.NodeTask, err error) {
	node := input.Node
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		var tmpNode base_db_model.Node
		if err = tx.Table("nodes").Where(
			"name = ? and kind = ?", node.Name, node.Kind).First(&tmpNode).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			tmpNode = base_db_model.Node{
				Name:         node.Name,
				Kind:         node.Kind,
				Role:         node.Role,
				Status:       node.Status,
				StatusReason: node.StatusReason,
				State:        node.State,
				StateReason:  node.StateReason,
			}
			if err = tx.Create(&tmpNode).Error; err != nil {
				return
			}
		} else {
			tmpNode.State = node.State
			tmpNode.StateReason = node.StateReason
			if err = tx.Save(&tmpNode).Error; err != nil {
				return
			}
		}

		var tmpNodeMeta db_model.NodeMeta
		if err = tx.Table("node_meta").Where(
			"node_id = ?", tmpNode.ID).First(&tmpNodeMeta).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			tmpNodeMeta = db_model.NodeMeta{
				NodeID: tmpNode.ID,
				Weight: 0,
			}
			if err = tx.Create(&tmpNodeMeta).Error; err != nil {
				return
			}
		} else {
			tmpNodeMeta.Weight = 0
			if err = tx.Save(&tmpNodeMeta).Error; err != nil {
				return
			}
		}
		return
	})

	// generate node tasks
	var computeAssignments []db_model.ComputeAssignmentWithComputeAndNode
	if computeAssignments, err = api.GetComputeAssignments(tctx, api.DB, ""); err != nil {
		return
	}
	computeAssignmentExs := []api_spec.ComputeAssignmentEx{}
	for _, assignment := range computeAssignments {
		var rspec api_spec.RegionServiceComputeSpec
		if err = json_utils.Unmarshal(assignment.ComputeSpec, &rspec); err != nil {
			return
		}
		computeAssignmentExs = append(computeAssignmentExs, api_spec.ComputeAssignmentEx{
			ID:        assignment.ID,
			UpdatedAt: assignment.UpdatedAt,
			Status:    assignment.Status,
			Spec:      rspec,
		})
	}
	nodeTask = &api_spec.NodeTask{
		ComputeAssignments: computeAssignmentExs,
	}
	return
}
