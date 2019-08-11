package resource_cluster_model_api

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceClusterModelApi) GetCompute(tctx *logger.TraceContext, db *gorm.DB,
	req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	name, ok := query.StrParams["name"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("name is None")
	}

	var compute resource_model.Compute
	if err = db.Where(&resource_model.Compute{
		Name: name,
	}).First(&compute).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Compute"] = compute
	return codes.OkRead, nil
}

func (modelApi *ResourceClusterModelApi) GetComputes(tctx *logger.TraceContext, db *gorm.DB,
	req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["Computes"] = computes
	return codes.OkRead, nil
}

func (modelApi *ResourceClusterModelApi) CreateCompute(tctx *logger.TraceContext, db *gorm.DB,
	req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.RegionServiceSpec
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

		var specBytes []byte
		if specBytes, err = json_utils.Marshal(spec); err != nil {
			return codes.ClientBadRequest, err
		}

		computeSpec := spec.Compute

		var data resource_model.Compute
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Compute{
				Project:       req.Tctx.ProjectName,
				Kind:          computeSpec.Kind,
				Name:          computeSpec.Name,
				RegionService: spec.Name,
				Region:        spec.Region,
				Cluster:       computeSpec.Cluster,
				Image:         computeSpec.Image,
				Vcpus:         computeSpec.Vcpus,
				Memory:        computeSpec.Memory,
				Disk:          computeSpec.Disk,
				Spec:          string(specBytes),
				Status:        resource_model.StatusInitializing,
				StatusReason:  resource_model.StatusMsgInitializing,
			}
			if err = tx.Create(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return codes.ClientAlreadyExists, err
		}
	}

	tx.Commit()
	return codes.OkCreated, nil
}

func (modelApi *ResourceClusterModelApi) UpdateCompute(tctx *logger.TraceContext, db *gorm.DB,
	req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceClusterModelApi) DeleteCompute(tctx *logger.TraceContext, db *gorm.DB,
	req *authproxy_grpc_pb.ActionRequest, query *authproxy_grpc_pb.Query) (int64, error) {
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

		if err = tx.Delete(&resource_model.Compute{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceClusterModelApi) SyncCompute(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer modelApi.close(tctx, db)

	tx := db.Begin()
	defer tx.Rollback()
	var computes []resource_model.Compute
	if err = tx.Find(&computes).Error; err != nil {
		return err
	}

	var nodes []resource_model.Node
	// TODO filter by resource driver
	if err = tx.Where(&resource_model.Node{
		Kind: resource_model.KindResourceClusterAgent,
	}).Find(&nodes).Error; err != nil {
		return err
	}

	var computeAssignments []resource_model.ComputeAssignmentWithComputeAndNode
	if computeAssignments, err = modelApi.GetComputeAssignments(tctx, tx, ""); err != nil {
		return err
	}
	tx.Commit()

	nodeMap := map[uint]*resource_model.Node{}
	nodeAssignmentsMap := map[uint][]resource_model.ComputeAssignmentWithComputeAndNode{}
	for _, node := range nodes {
		nodeAssignmentsMap[node.ID] = []resource_model.ComputeAssignmentWithComputeAndNode{}
		nodeMap[node.ID] = &node
	}

	computeAssignmentsMap := map[string][]resource_model.ComputeAssignmentWithComputeAndNode{}
	for _, assignment := range computeAssignments {
		assignments, ok := computeAssignmentsMap[assignment.ComputeName]
		if !ok {
			assignments = []resource_model.ComputeAssignmentWithComputeAndNode{}
		}
		assignments = append(assignments, assignment)
		computeAssignmentsMap[assignment.ComputeName] = assignments

		nodeAssignments := nodeAssignmentsMap[assignment.NodeID]
		nodeAssignments = append(nodeAssignments, assignment)
		nodeAssignmentsMap[assignment.NodeID] = nodeAssignments
	}

	for _, compute := range computes {
		switch compute.Status {
		case resource_model.StatusInitializing:
			modelApi.AssignCompute(tctx, db, &compute, nodeMap, nodeAssignmentsMap, computeAssignmentsMap, false)
		}
	}

	fmt.Println("TODO SyncCompute")
	return nil
}

func (modelApi *ResourceClusterModelApi) AssignCompute(tctx *logger.TraceContext, db *gorm.DB,
	compute *resource_model.Compute, nodeMap map[uint]*resource_model.Node,
	nodeAssignmentsMap map[uint][]resource_model.ComputeAssignmentWithComputeAndNode,
	assignmentsMap map[string][]resource_model.ComputeAssignmentWithComputeAndNode,
	isReschedule bool) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var spec resource_model.RegionServiceSpec
	if err = json_utils.Unmarshal(compute.Spec, &spec); err != nil {
		return
	}

	policy := spec.Compute.SchedulePolicy
	assignNodes := []uint{}
	updateNodes := []uint{}
	unassignNodes := []uint{}

	currentAssignments, ok := assignmentsMap[compute.Name]
	if ok {
		infoMsg := []string{}
		for _, currentAssignment := range currentAssignments {
			infoMsg = append(infoMsg, currentAssignment.NodeName)
		}
		logger.Infof(tctx, "currentAssignments: %v", infoMsg)
	}

	fmt.Println("DEBUG nodeAssignments: ", nodeAssignmentsMap)

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

	labelFilterNodeMap := map[uint]*resource_model.Node{}
	filteredNodes := []*resource_model.Node{}
	labelNodesMap := map[string][]*resource_model.Node{} // LabelごとのNode候補
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
		if node.Status != resource_model.StatusEnabled {
			continue
		}

		if node.State != resource_model.StateUp {
			continue
		}

		// TODO
		// Filter node by cpu, memory, disk

		filteredNodes = append(filteredNodes, node)

		for _, label := range labels {
			nodes, lok := labelNodesMap[label]
			if !lok {
				nodes = []*resource_model.Node{}
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
			} else {
				unassignNodes = append(unassignNodes, assignment.NodeID)
			}
		}
		replicas = replicas - len(currentAssignments) + len(unassignNodes)
	}

	if replicas != 0 {
		for i := 0; i < replicas; i++ {
			candidates := []*resource_model.Node{}
			for _, label := range policy.NodeLabelHardAntiAffinities {
				tmpCandidates := []*resource_model.Node{}
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
					newCandidates := []*resource_model.Node{}
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
				tmpCandidates := []*resource_model.Node{}
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

	tx := db.Begin()
	defer tx.Rollback()
	for _, nodeID := range updateNodes {
		switch compute.Status {
		case resource_model.StatusUpdating:
			tx.Create(&resource_model.ComputeAssignment{
				ComputeID:    compute.ID,
				NodeID:       nodeID,
				Status:       resource_model.StatusUpdating,
				StatusReason: resource_model.StatusMsgUpdating,
			})
		}
	}

	for _, nodeID := range assignNodes {
		switch compute.Status {
		case resource_model.StatusInitializing:
			tx.Create(&resource_model.ComputeAssignment{
				ComputeID:    compute.ID,
				NodeID:       nodeID,
				Status:       resource_model.StatusCreating,
				StatusReason: resource_model.StatusMsgCreating,
			})
		}
	}

	switch compute.Status {
	case resource_model.StatusInitializing:
		compute.Status = resource_model.StatusCreatingScheduled
		compute.StatusReason = resource_model.StatusMsgInitializeSuccess
	}

	tx.Save(compute)
	tx.Commit()
}

func (modelApi *ResourceClusterModelApi) GetComputeAssignments(tctx *logger.TraceContext, db *gorm.DB,
	nodeName string) ([]resource_model.ComputeAssignmentWithComputeAndNode, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	query := db.Table("compute_assignments as ca").
		Select("ca.id, ca.status, ca.updated_at, ca.compute_id, c.name as compute_name, c.spec as compute_spec, ca.node_id, n.name as node_name").
		Joins("INNER JOIN computes AS c ON c.id = ca.compute_id").
		Joins("INNER JOIN nodes AS n ON n.id = ca.node_id")
	if nodeName != "" {
		query = query.Where("n.name = ?", nodeName)
	}

	var computeAssignments []resource_model.ComputeAssignmentWithComputeAndNode
	if err = query.Find(&computeAssignments).Error; err != nil {
		return nil, err
	}

	return computeAssignments, nil
}
