package db_api

import (
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_db_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (api *Api) SyncNodeService(tctx *logger.TraceContext, input *api_spec.SyncNodeService) (nodeTask *api_spec.NodeServiceTask, err error) {
	node := input.NodeService
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		endpoints := strings.Join(node.Endpoints, ",")
		var tmpNodeService base_db_model.NodeService
		if err = tx.Table("node_services").Where(
			"name = ? and kind = ?", node.Name, node.Kind).First(&tmpNodeService).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			tmpNodeService = base_db_model.NodeService{
				Name:         node.Name,
				Kind:         node.Kind,
				Role:         node.Role,
				Status:       node.Status,
				StatusReason: node.StatusReason,
				State:        node.State,
				StateReason:  node.StateReason,
				Token:        node.Token,
				Endpoints:    endpoints,
			}
			if err = tx.Debug().Create(&tmpNodeService).Error; err != nil {
				return
			}
		} else {
			if err = tx.Table("node_services").
				Where("name = ? AND kind = ?", node.Name, node.Kind).
				Updates(map[string]interface{}{
					"state":        node.State,
					"state_reason": node.StateReason,
					"token":        node.Token,
					"endpoints":    endpoints,
				}).Error; err != nil {
				return
			}
		}

		var tmpNodeServiceMeta db_model.NodeServiceMeta
		if err = tx.Table("node_service_meta").Where(
			"node_service_id = ?", tmpNodeService.ID).First(&tmpNodeServiceMeta).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return
			}
			tmpNodeServiceMeta = db_model.NodeServiceMeta{
				NodeServiceID: tmpNodeService.ID,
				Weight:        0,
			}
			if err = tx.Create(&tmpNodeServiceMeta).Error; err != nil {
				return
			}
		} else {
			tmpNodeServiceMeta.Weight = 0
			if err = tx.Save(&tmpNodeServiceMeta).Error; err != nil {
				return
			}
		}
		return
	})
	if err != nil {
		return
	}

	// generate node tasks
	var computeAssignments []db_model.ComputeAssignmentWithComputeAndNodeService
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
	nodeTask = &api_spec.NodeServiceTask{
		ComputeAssignments: computeAssignmentExs,
	}
	return
}

func (api *Api) ReportNodeServiceTask(tctx *logger.TraceContext, input *api_spec.ReportNodeServiceTask) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, report := range input.ComputeAssignmentReports {
			switch report.Status {
			case resource_model.StatusDeleted:
				if err = tx.Where("id = ? AND updated_at = ?", report.ID, report.UpdatedAt).
					Unscoped().Delete(&db_model.ComputeAssignment{}).Error; err != nil {
					return
				}
			default:
				if err = tx.Table("compute_assignments").
					Where("id = ? AND updated_at = ?", report.ID, report.UpdatedAt).
					Updates(map[string]interface{}{
						"status":        report.Status,
						"status_reason": report.StatusReason,
					}).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}
