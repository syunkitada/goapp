package resource_cluster_model_api

import (
	"encoding/json"
	"fmt"

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
	if err = tx.Find(&nodes).Error; err != nil {
		return err
	}

	query := tx.Table("compute_assignments as ca").
		Select("ca.status, c.name as compute_name, c.spec as compute_spec, ca.node_id, n.name as node_name").
		Joins("INNER JOIN computes AS c ON c.id = ca.compute_id").
		Joins("INNER JOIN nodes AS n ON n.id = ca.node_id")
	var computeAssignments []resource_model.ComputeAssignmentWithComputeAndNode
	if err = query.Find(&computeAssignments).Error; err != nil {
		return nil
	}
	tx.Commit()

	nodeAssignmentsMap := map[uint][]resource_model.ComputeAssignmentWithComputeAndNode{}
	for _, node := range nodes {
		nodeAssignmentsMap[node.ID] = []resource_model.ComputeAssignmentWithComputeAndNode{}
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
			modelApi.AssignCompute(tctx, db, &compute, nodeAssignmentsMap, computeAssignmentsMap, false)
		}
	}

	fmt.Println("TODO SyncCompute")
	return nil
}

func (modelApi *ResourceClusterModelApi) AssignCompute(tctx *logger.TraceContext, db *gorm.DB,
	compute *resource_model.Compute,
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

	replicas := policy.Replicas
	if !isReschedule {
		for _, assignment := range currentAssignments {
			updateNodes = append(updateNodes, assignment.NodeID)
		}
		replicas -= len(currentAssignments)
	}

	if replicas != 0 {
		// assignNodes := []uint{}
		// updateNodes := []uint{}
		// unassignNodes := []uint{}
		for i := 0; i < replicas; i++ {
			fmt.Println("DEBUG")
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
		case resource_model.StatusInitializing:
			tx.Create(&resource_model.ComputeAssignment{
				ComputeID:    compute.ID,
				NodeID:       nodeID,
				Status:       resource_model.StatusUpdating,
				StatusReason: resource_model.StatusMsgUpdating,
			})
		}
	}

	switch compute.Status {
	case resource_model.StatusInitializing:
		compute.Status = resource_model.StatusCreatingScheduled
		compute.Status = resource_model.StatusMsgInitializeSuccess
	}

	tx.Commit()
}
