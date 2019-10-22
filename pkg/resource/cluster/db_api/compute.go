package db_api

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (api *Api) GetCompute(tctx *logger.TraceContext, input *spec.GetCompute) (data *spec.Compute, err error) {
	data = &spec.Compute{}
	err = api.DB.Where("name = ?", input.Name).First(data).Error
	return
}

func (api *Api) GetComputes(tctx *logger.TraceContext, input *spec.GetComputes) (data []spec.Compute, err error) {
	err = api.DB.Find(&data).Error
	return
}

func (api *Api) CreateComputes(tctx *logger.TraceContext, specs []spec.RegionServiceComputeSpec) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, spec := range specs {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(spec); err != nil {
				return
			}
			data := db_model.Compute{
				Name:         spec.Name,
				Spec:         string(specBytes),
				Status:       base_const.StatusCreating,
				StatusReason: "CreateComputes",
			}
			if err = tx.Create(&data).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) UpdateComputes(tctx *logger.TraceContext, specs []spec.RegionServiceComputeSpec) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, spec := range specs {
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(spec); err != nil {
				return
			}
			query := tx.Table("computes").Where("name = ?", spec.Name).Updates(map[string]interface{}{
				"spec":          string(specBytes),
				"status":        base_const.StatusUpdating,
				"status_reason": "UpdateComputes",
			})

			var rows int64
			if rows, err = query.RowsAffected, query.Error; err != nil {
				return
			} else if rows != 1 {
				err = fmt.Errorf("updated rows is nothing: count=%d", rows)
				return
			}
			return
		}
		return
	})
	return
}

func (api *Api) DeleteCompute(tctx *logger.TraceContext, input *spec.DeleteCompute) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		query := tx.Table("computes").
			Where("name = ?", input.Name).Updates(map[string]interface{}{
			"status":        base_const.StatusDeleting,
			"status_reason": "DeleteCompute",
		})

		var rows int64
		if rows, err = query.RowsAffected, query.Error; err != nil {
			return
		} else if rows != 1 {
			err = fmt.Errorf("deleted rows is nothing: count=%d", rows)
			return
		}
		return
	})
	return
}

func (api *Api) DeleteComputes(tctx *logger.TraceContext, specs []spec.RegionServiceComputeSpec) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, spec := range specs {
			query := tx.Table("computes").
				Where("name = ?", spec.Name).Updates(map[string]interface{}{
				"status":        base_const.StatusDeleting,
				"status_reason": "DeleteComputes"})

			var rows int64
			if rows, err = query.RowsAffected, query.Error; err != nil {
				return
			} else if rows != 1 {
				err = fmt.Errorf("deleted rows is nothing: count=%d", rows)
				return
			}
		}
		return
	})
	return
}

func (api *Api) SyncCompute(tctx *logger.TraceContext) (err error) {
	fmt.Println("SyncCompute")

	var computes []db_model.Compute
	var nodes []db_model.NodeWithMeta
	var computeAssignments []db_model.ComputeAssignmentWithComputeAndNode
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		if err = tx.Find(&computes).Error; err != nil {
			return
		}

		// TODO filter by resource driver
		if err = tx.Table("nodes as n").Select("*").
			Joins("INNER JOIN node_meta as nm ON n.id = nm.node_id").
			Where("n.kind = ?", consts.KindResourceClusterAgent).Scan(&nodes).Error; err != nil {
			return
		}

		if computeAssignments, err = api.GetComputeAssignments(tctx, tx, ""); err != nil {
			return
		}
		return
	})

	nodeMap := map[uint]*db_model.NodeWithMeta{}
	nodeAssignmentsMap := map[uint][]db_model.ComputeAssignmentWithComputeAndNode{}
	for _, node := range nodes {
		nodeAssignmentsMap[node.ID] = []db_model.ComputeAssignmentWithComputeAndNode{}
		nodeMap[node.ID] = &node
	}

	computeAssignmentsMap := map[string][]db_model.ComputeAssignmentWithComputeAndNode{}
	for _, assignment := range computeAssignments {
		assignments, ok := computeAssignmentsMap[assignment.ComputeName]
		if !ok {
			assignments = []db_model.ComputeAssignmentWithComputeAndNode{}
		}
		assignments = append(assignments, assignment)
		computeAssignmentsMap[assignment.ComputeName] = assignments

		nodeAssignments := nodeAssignmentsMap[assignment.NodeID]
		nodeAssignments = append(nodeAssignments, assignment)
		nodeAssignmentsMap[assignment.NodeID] = nodeAssignments
	}

	for _, compute := range computes {
		switch compute.Status {
		case base_const.StatusCreating:
			api.AssignCompute(tctx, &compute, nodeMap, nodeAssignmentsMap, computeAssignmentsMap, false)
		case base_const.StatusCreatingScheduled:
			api.ConfirmCreatingOrUpdatingScheduledCompute(tctx, &compute, computeAssignmentsMap)
		case base_const.StatusUpdating:
			api.AssignCompute(tctx, &compute, nodeMap, nodeAssignmentsMap, computeAssignmentsMap, false)
		case base_const.StatusUpdatingScheduled:
			api.ConfirmCreatingOrUpdatingScheduledCompute(tctx, &compute, computeAssignmentsMap)
		case base_const.StatusDeleting:
			api.DeleteComputeAssignments(tctx, &compute)
		case base_const.StatusDeletingScheduled:
			api.ConfirmDeletingScheduledCompute(tctx, &compute, computeAssignmentsMap)
		}
	}

	return
}

func (api *Api) GetComputeAssignments(tctx *logger.TraceContext, db *gorm.DB,
	nodeName string) (assignments []db_model.ComputeAssignmentWithComputeAndNode, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	query := db.Table("compute_assignments as ca").
		Select("ca.id, ca.status, ca.updated_at, ca.compute_id, c.name as compute_name, c.spec as compute_spec, ca.node_id, n.name as node_name").
		Joins("INNER JOIN computes AS c ON c.id = ca.compute_id").
		Joins("INNER JOIN nodes AS n ON n.id = ca.node_id")
	if nodeName != "" {
		query = query.Where("n.name = ?", nodeName)
	}

	err = query.Find(&assignments).Error
	return
}

func (api *Api) AssignCompute(tctx *logger.TraceContext,
	compute *db_model.Compute, nodeMap map[uint]*db_model.NodeWithMeta,
	nodeAssignmentsMap map[uint][]db_model.ComputeAssignmentWithComputeAndNode,
	assignmentsMap map[string][]db_model.ComputeAssignmentWithComputeAndNode,
	isReschedule bool) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var rspec spec.RegionServiceComputeSpec
	if err = json_utils.Unmarshal(compute.Spec, &rspec); err != nil {
		return
	}

	policy := rspec.SchedulePolicy
	assignNodes := []uint{}
	updateNodes := []uint{}
	updateAssignments := []uint{}
	unassignNodes := []uint{}

	currentAssignments, ok := assignmentsMap[compute.Name]
	if ok {
		infoMsg := []string{}
		for _, currentAssignment := range currentAssignments {
			infoMsg = append(infoMsg, currentAssignment.NodeName)
		}
		logger.Infof(tctx, "currentAssignments: %v", infoMsg)
	}

	// filtering node
	enableNodeFilters := false
	if len(policy.NodeFilters) > 0 {
		enableNodeFilters = true
	}
	enableLabelFilters := false
	if len(policy.NodeLabelFilters) > 0 {
		enableLabelFilters = true
	}
	enableHardAffinites := false
	if len(policy.NodeLabelHardAffinities) > 0 {
		enableHardAffinites = true
	}
	enableHardAntiAffinites := false
	if len(policy.NodeLabelHardAntiAffinities) > 0 {
		enableHardAntiAffinites = true
	}
	enableSoftAffinites := false
	if len(policy.NodeLabelSoftAffinities) > 0 {
		enableSoftAffinites = true
	}
	enableSoftAntiAffinites := false
	if len(policy.NodeLabelSoftAntiAffinities) > 0 {
		enableSoftAntiAffinites = true
	}

	labelFilterNodeMap := map[uint]*db_model.NodeWithMeta{}
	filteredNodes := []*db_model.NodeWithMeta{}
	labelNodesMap := map[string][]*db_model.NodeWithMeta{} // LabelごとのNode候補
	for _, node := range nodeMap {
		labels := []string{}
		ok := true
		if enableNodeFilters {
			ok = false
			for _, nodeName := range policy.NodeFilters {
				if node.Name == nodeName {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableLabelFilters {
			ok = false
			for _, label := range policy.NodeLabelFilters {
				if strings.Index(node.Labels, label) >= 0 {
					ok = true
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableHardAffinites {
			ok = false
			for _, label := range policy.NodeLabelHardAffinities {
				if strings.Index(node.Labels, label) >= 0 {
					ok = true
					labels = append(labels, label)
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableHardAntiAffinites {
			ok = false
			for _, label := range policy.NodeLabelHardAntiAffinities {
				if strings.Index(node.Labels, label) >= 0 {
					ok = true
					labels = append(labels, label)
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableSoftAffinites {
			ok = false
			for _, label := range policy.NodeLabelSoftAffinities {
				if strings.Index(node.Labels, label) >= 0 {
					ok = true
					labels = append(labels, label)
					break
				}
			}
			if !ok {
				continue
			}
		}

		if enableSoftAntiAffinites {
			ok = false
			for _, label := range policy.NodeLabelSoftAntiAffinities {
				if strings.Index(node.Labels, label) >= 0 {
					ok = true
					labels = append(labels, label)
					break
				}
			}
			if !ok {
				continue
			}
		}

		// labelFilterNodeMapには、LabelのみによるNodeのフィルタリング結果を格納する
		labelFilterNodeMap[node.ID] = node

		// Filter node by status, state
		if node.Status != base_const.StatusEnabled {
			continue
		}

		if node.State != base_const.StateUp {
			continue
		}

		// TODO
		// Filter node by cpu, memory, disk

		filteredNodes = append(filteredNodes, node)

		for _, label := range labels {
			nodes, lok := labelNodesMap[label]
			if !lok {
				nodes = []*db_model.NodeWithMeta{}
			}
			nodes = append(nodes, node)
			labelNodesMap[label] = nodes
		}
	}

	replicas := policy.Replicas
	if !isReschedule {
		for _, assignment := range currentAssignments {
			// labelFilterNodeMapには、LabelのみによるNodeのフィルタリング結果が格納されている
			// label変更されてNodeが候補から外された場合は、unassignNodesに追加する
			_, ok := labelFilterNodeMap[assignment.NodeID]
			if ok {
				updateNodes = append(updateNodes, assignment.NodeID)
				updateAssignments = append(updateAssignments, assignment.ID)
			} else {
				unassignNodes = append(unassignNodes, assignment.NodeID)
			}
		}
		replicas = replicas - len(currentAssignments) + len(unassignNodes)
	}

	if replicas != 0 {
		for i := 0; i < replicas; i++ {
			candidates := []*db_model.NodeWithMeta{}
			for _, label := range policy.NodeLabelHardAntiAffinities {
				tmpCandidates := []*db_model.NodeWithMeta{}
				nodes := labelNodesMap[label]
				for _, node := range nodes {
					existsNode := false
					for _, n := range assignNodes {
						if node.ID == n {
							existsNode = true
							break
						}
					}
					if existsNode {
						continue
					}
					for _, n := range updateNodes {
						if node.ID == n {
							existsNode = true
							break
						}
					}
					if existsNode {
						continue
					}
					tmpCandidates = append(candidates, node)
				}
				if len(candidates) == 0 {
					candidates = tmpCandidates
				} else {
					newCandidates := []*db_model.NodeWithMeta{}
					for _, c := range candidates {
						for _, tc := range tmpCandidates {
							if c == tc {
								newCandidates = append(newCandidates, c)
								break
							}
						}
					}
					candidates = newCandidates
				}
			}

			for _, label := range policy.NodeLabelHardAffinities {
				tmpCandidates := []*db_model.NodeWithMeta{}
				nodes := labelNodesMap[label]
				if len(candidates) == 0 && len(assignNodes) == 0 && len(updateNodes) == 0 {
					for _, node := range nodes {
						tmpCandidates = append(tmpCandidates, node)
					}
					candidates = tmpCandidates
					break
				} else if len(assignNodes) > 0 {
					for _, node := range nodes {
						for _, assignNodeID := range assignNodes {
							if node.ID == assignNodeID {
								candidates = append(candidates, node)
								break
							}
						}
					}
					break
				} else if len(updateNodes) > 0 {
					for _, node := range nodes {
						for _, updateNodeID := range updateNodes {
							if node.ID == updateNodeID {
								candidates = append(candidates, node)
								break
							}
						}
					}
					break
				}
			}

			if !enableNodeFilters && !enableLabelFilters && !enableHardAffinites && !enableHardAntiAffinites {
				if len(candidates) == 0 {
					for _, node := range filteredNodes {
						candidates = append(candidates, node)
					}
				}
			}

			// candidatesのweightを調整する
			for _, label := range policy.NodeLabelSoftAffinities {
				nodes := labelNodesMap[label]
				for _, node := range nodes {
					for _, assignNodeId := range assignNodes {
						if node.ID == assignNodeId {
							node.Weight += 1000
							break
						}
					}
					for _, updateNodeId := range updateNodes {
						if node.ID == updateNodeId {
							node.Weight += 1000
							break
						}
					}
				}
			}

			for _, label := range policy.NodeLabelSoftAntiAffinities {
				nodes := labelNodesMap[label]
				for _, node := range nodes {
					for _, assignNodeId := range assignNodes {
						if node.ID == assignNodeId {
							node.Weight -= 1000
							break
						}
					}
					for _, updateNodeId := range updateNodes {
						if node.ID == updateNodeId {
							node.Weight -= 1000
							break
						}
					}
				}
			}

			if len(candidates) == 0 {
				break
			}

			for _, candidate := range candidates {
				assignments, ok := nodeAssignmentsMap[candidate.ID]
				if ok {
					for _, assignment := range assignments {
						// TODO calucurate assignment.Cost before scheduling
						candidate.Weight -= (10 + assignment.Cost)
					}
				}
			}

			// TODO Sort candidates by weight

			assignNodes = append(assignNodes, candidates[0].ID)
		}
	}

	if policy.Replicas != len(assignNodes)+len(updateNodes)-len(unassignNodes) {
		logger.Warningf(tctx, "Failed assign: compute=%v", compute.Name)
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, id := range updateAssignments {
			switch compute.Status {
			case base_const.StatusUpdating:
				if err = tx.Table("compute_assignments").Where("id = ?", id).Updates(map[string]interface{}{
					"status":        base_const.StatusUpdating,
					"status_reason": "Updating",
				}).Error; err != nil {
					return
				}
			}
		}

		for _, nodeID := range assignNodes {
			switch compute.Status {
			case base_const.StatusCreating:
				if err = tx.Create(&db_model.ComputeAssignment{
					ComputeID:    compute.ID,
					NodeID:       nodeID,
					Status:       base_const.StatusCreating,
					StatusReason: "Creating",
				}).Error; err != nil {
					return
				}
			}
		}

		switch compute.Status {
		case base_const.StatusCreating:
			err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
				"status":        base_const.StatusCreatingScheduled,
				"status_reason": "CreatingScheduled",
			}).Error
		case base_const.StatusUpdating:
			err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
				"status":        base_const.StatusUpdatingScheduled,
				"status_reason": "UpdatingScheduled",
			}).Error
		}
		return
	})
}

func (api *Api) ConfirmCreatingOrUpdatingScheduledCompute(tctx *logger.TraceContext,
	compute *db_model.Compute,
	assignmentsMap map[string][]db_model.ComputeAssignmentWithComputeAndNode) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	assignments, ok := assignmentsMap[compute.Name]
	if !ok {
		err = error_utils.NewConflictNotFoundError(compute.Name)
		return
	}

	existsNonActiveAssignments := false
	for _, assignment := range assignments {
		if assignment.Status != base_const.StatusActive {
			existsNonActiveAssignments = true
			break
		}
	}

	if existsNonActiveAssignments {
		logger.Info(tctx, "Waiting: exists non active assignments")
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        resource_model.StatusActive,
			"status_reason": "ConfirmedActive",
		}).Error
		return
	})
}

func (api *Api) DeleteComputeAssignments(tctx *logger.TraceContext, compute *db_model.Compute) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Table("computes").Where("id = ?", compute.ID).Updates(map[string]interface{}{
			"status":        base_const.StatusDeletingScheduled,
			"status_reason": "DeleteComputeAssignments",
		}).Error

		err = tx.Table("compute_assignments").Where("compute_id = ?", compute.ID).
			Updates(map[string]interface{}{
				"status":        base_const.StatusDeleting,
				"status_reason": "Deleting",
			}).Error
		return
	})
	return
}

func (api *Api) ConfirmDeletingScheduledCompute(tctx *logger.TraceContext,
	compute *db_model.Compute,
	assignmentsMap map[string][]db_model.ComputeAssignmentWithComputeAndNode) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	_, ok := assignmentsMap[compute.Name]
	if ok {
		logger.Infof(tctx, "waiting compute_assignments is deleted")
		return
	}

	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("id = ?", compute.ID).Unscoped().Delete(&db_model.Compute{}).Error
		return
	})
}
