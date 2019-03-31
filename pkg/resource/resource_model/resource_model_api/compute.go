package resource_model_api

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

// import (
// 	"encoding/json"
// 	"fmt"
// 	"strconv"
// 	"strings"
// 	"time"
//
// 	"github.com/golang/protobuf/ptypes"
// 	"github.com/jinzhu/gorm"
//
// 	"github.com/syunkitada/goapp/pkg/lib/codes"
// 	"github.com/syunkitada/goapp/pkg/lib/logger"
// 	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
// 	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
// 	"github.com/syunkitada/goapp/pkg/resource/resource_model"
// )
//
// func (modelApi *ResourceModelApi) CreateCompute(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
// 	var err error
// 	startTime := logger.StartTrace(tctx)
// 	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
//
// 	var db *gorm.DB
// 	if db, err = modelApi.open(tctx); err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.RemoteDbError
// 		return
// 	}
// 	defer func() { err = db.Close() }()
//
// 	spec, statusCode, err := modelApi.validateComputeSpec(db, req.Spec)
// 	if err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = statusCode
// 		return
// 	}
//
// 	var compute resource_model.Compute
// 	tx := db.Begin()
// 	defer tx.Rollback()
// 	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&compute).Error; err != nil {
// 		if !gorm.IsRecordNotFoundError(err) {
// 			rep.Tctx.Err = err.Error()
// 			rep.Tctx.StatusCode = codes.RemoteDbError
// 			return
// 		}
//
// 		compute = resource_model.Compute{
// 			Cluster:      spec.Cluster,
// 			Kind:         spec.Kind,
// 			Name:         spec.Name,
// 			Domain:       spec.Spec.Domain,
// 			Spec:         req.Spec,
// 			Status:       resource_model.StatusCreating,
// 			StatusReason: fmt.Sprintf("CreateCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
// 		}
// 		if err = tx.Create(&compute).Error; err != nil {
// 			rep.Tctx.Err = err.Error()
// 			rep.Tctx.StatusCode = codes.RemoteDbError
// 			return
// 		}
// 	} else {
// 		rep.Tctx.Err = fmt.Sprintf("Already Exists: cluster=%v, name=%v", spec.Cluster, spec.Name)
// 		rep.Tctx.StatusCode = codes.ClientAlreadyExists
// 		return
// 	}
// 	tx.Commit()
//
// 	computePb, err := modelApi.convertCompute(&compute)
// 	if err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.ServerInternalError
// 		return
// 	}
//
// 	rep.Computes = []*resource_api_grpc_pb.Compute{computePb}
// 	rep.Tctx.StatusCode = codes.Ok
// 	return
// }
//
// func (modelApi *ResourceModelApi) UpdateCompute(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
// 	var err error
// 	startTime := logger.StartTrace(tctx)
// 	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
//
// 	var db *gorm.DB
// 	if db, err = modelApi.open(tctx); err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.RemoteDbError
// 		return
// 	}
// 	defer func() { err = db.Close() }()
//
// 	spec, statusCode, err := modelApi.validateComputeSpec(db, req.Spec)
// 	if err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = statusCode
// 		return
// 	}
//
// 	tx := db.Begin()
// 	defer tx.Rollback()
// 	var compute resource_model.Compute
// 	if err = tx.Model(&compute).Where(resource_model.Compute{
// 		Name:    spec.Name,
// 		Cluster: spec.Cluster,
// 	}).Updates(resource_model.Compute{
// 		Spec:         req.Spec,
// 		Status:       resource_model.StatusActive,
// 		StatusReason: fmt.Sprintf("UpdateCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
// 	}).Error; err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.RemoteDbError
// 		return
// 	}
//
// 	computePb, err := modelApi.convertCompute(&compute)
// 	if err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.ServerInternalError
// 		return
// 	}
// 	computePb.Name = spec.Name
// 	computePb.Cluster = spec.Cluster
//
// 	rep.Computes = []*resource_api_grpc_pb.Compute{computePb}
// 	rep.Tctx.StatusCode = codes.Ok
// }
//
// func (modelApi *ResourceModelApi) DeleteCompute(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
// 	var err error
// 	startTime := logger.StartTrace(tctx)
// 	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()
//
// 	var db *gorm.DB
// 	if db, err = modelApi.open(tctx); err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.RemoteDbError
// 		return
// 	}
// 	defer func() { err = db.Close() }()
//
// 	tx := db.Begin()
// 	defer tx.Rollback()
// 	var compute resource_model.Compute
// 	now := time.Now()
// 	if err = tx.Model(&compute).Where(resource_model.Compute{
// 		Name:    req.Target,
// 		Cluster: req.Cluster,
// 	}).Updates(resource_model.Compute{
// 		Model: gorm.Model{
// 			DeletedAt: &now,
// 		},
// 		Status:       resource_model.StatusDeleted,
// 		StatusReason: fmt.Sprintf("DeleteCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
// 	}).Error; err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.RemoteDbError
// 		return
// 	}
// 	tx.Commit()
//
// 	computePb, err := modelApi.convertCompute(&compute)
// 	if err != nil {
// 		rep.Tctx.Err = err.Error()
// 		rep.Tctx.StatusCode = codes.ServerInternalError
// 		return
// 	}
//
// 	computePb.Name = req.Target
// 	computePb.Cluster = req.Cluster
// 	rep.Computes = []*resource_api_grpc_pb.Compute{computePb}
// 	rep.Tctx.StatusCode = codes.Ok
// }
//
// func (modelApi *ResourceModelApi) convertComputes(tctx *logger.TraceContext, computes []resource_model.Compute) []*resource_api_grpc_pb.Compute {
// 	pbComputes := make([]*resource_api_grpc_pb.Compute, len(computes))
// 	for i, compute := range computes {
// 		updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
// 		if err != nil {
// 			logger.Warningf(tctx, err,
// 				"Failed ptypes.TimestampProto: %v", compute.Model.UpdatedAt)
// 			continue
// 		}
// 		createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
// 		if err != nil {
// 			logger.Warningf(tctx, err,
// 				"Failed ptypes.TimestampProto: %v", compute.Model.CreatedAt)
// 			continue
// 		}
//
// 		pbComputes[i] = &resource_api_grpc_pb.Compute{
// 			Cluster:      compute.Cluster,
// 			Name:         compute.Name,
// 			Kind:         compute.Kind,
// 			Labels:       compute.Labels,
// 			Status:       compute.Status,
// 			StatusReason: compute.StatusReason,
// 			UpdatedAt:    updatedAt,
// 			CreatedAt:    createdAt,
// 		}
// 	}
//
// 	return pbComputes
// }
//
// func (modelApi *ResourceModelApi) convertCompute(compute *resource_model.Compute) (*resource_api_grpc_pb.Compute, error) {
// 	updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
// 	createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	computePb := &resource_api_grpc_pb.Compute{
// 		Cluster:      compute.Cluster,
// 		Name:         compute.Name,
// 		Kind:         compute.Kind,
// 		Labels:       compute.Labels,
// 		Status:       compute.Status,
// 		StatusReason: compute.StatusReason,
// 		UpdatedAt:    updatedAt,
// 		CreatedAt:    createdAt,
// 	}
//
// 	return computePb, nil
// }
//
// func (modelApi *ResourceModelApi) validateComputeSpec(db *gorm.DB, specStr string) (resource_model.ComputeSpec, int64, error) {
// 	var spec resource_model.ComputeSpec
// 	var err error
// 	if err = json.Unmarshal([]byte(specStr), &spec); err != nil {
// 		return spec, codes.ClientBadRequest, err
// 	}
// 	if err = modelApi.validate.Struct(spec); err != nil {
// 		return spec, codes.ClientInvalidRequest, err
// 	}
//
// 	ok, err := modelApi.ValidateClusterName(db, spec.Cluster)
// 	if err != nil {
// 		return spec, codes.RemoteDbError, err
// 	}
// 	if !ok {
// 		return spec, codes.ClientInvalidRequest, fmt.Errorf("Invalid cluster: %v", spec.Cluster)
// 	}
//
// 	errors := []string{}
// 	switch spec.Spec.Kind {
// 	case resource_model.SpecKindComputeLibvirt:
// 		// TODO Implement Validate SpecKindComputeLibvirt
// 		fmt.Printf("Validate SpecKindComputeLibvirt is not implemented")
//
// 	default:
// 		errors = append(errors, fmt.Sprintf("Invalid kind: %v", spec.Spec.Kind))
// 	}
//
// 	if len(errors) > 0 {
// 		return spec, codes.ClientInvalidRequest, fmt.Errorf(strings.Join(errors, "\n"))
// 	}
//
// 	return spec, codes.Ok, nil
// }

func (modelApi *ResourceModelApi) SyncCompute(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	computeMap := map[uint64]resource_cluster_api_grpc_pb.Compute{}
	for clusterName, clusterClient := range modelApi.clusterClientMap {
		result, err := clusterClient.GetCompute(tctx, "")
		if err != nil {
			logger.Errorf(tctx, err, "Failed GetCompute from %v", clusterName)
		}
		for _, compute := range result.Computes {
			computeMap[compute.Compute.Id] = *compute
		}
	}

	for _, compute := range computes {
		tctx.Metadata["ComputeId"] = strconv.FormatUint(uint64(compute.ID), 10)
		switch compute.Status {
		case resource_model.StatusCreating:
			modelApi.InitializeCompute(tctx, db, &compute)
		case resource_model.StatusCreatingInitialized:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusCreatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdating:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdatingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeleting:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeletingScheduled:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
		}
		tctx.Metadata = map[string]string{}
	}

	return nil
}

func (modelApi *ResourceModelApi) InitializeCompute(tctx *logger.TraceContext, db *gorm.DB, compute *resource_model.Compute) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	if err = modelApi.AssignPort(tctx, db, compute); err != nil {
		return
	}

	if err = modelApi.RegisterRecord(tctx, db, compute); err != nil {
		return
	}

	// Update Creating Initialized

	return
}
