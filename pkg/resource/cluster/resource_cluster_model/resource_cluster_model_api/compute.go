package resource_cluster_model_api

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceClusterModelApi) GetCompute(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.ActionRequest, rep *resource_cluster_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	var computes []resource_cluster_model.Compute
	if req.Target == "" {
		err = db.Find(&computes).Error
	} else {
		err = db.Where("name like ?", req.Target).Find(&computes).Error
	}
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}

	rep.Computes = modelApi.convertComputes(tctx, computes)
	rep.Tctx.StatusCode = codes.Ok
	return
}

func (modelApi *ResourceClusterModelApi) CreateCompute(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.ActionRequest, rep *resource_cluster_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	spec, statusCode, err := modelApi.validateComputeSpec(db, req.Spec)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = statusCode
		return
	}

	var compute resource_cluster_model.Compute
	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&compute).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}

		compute = resource_cluster_model.Compute{
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_cluster_model.StatusCreating,
			StatusReason: fmt.Sprintf("CreateCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
		}
		if err = tx.Create(&compute).Error; err != nil {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = codes.RemoteDbError
			return
		}
	} else {
		rep.Tctx.Err = fmt.Sprintf("Already Exists: cluster=%v, name=%v", spec.Cluster, spec.Name)
		rep.Tctx.StatusCode = codes.ClientAlreadyExists
		return
	}
	tx.Commit()

	computePb, err := modelApi.convertCompute(&compute)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.ServerInternalError
		return
	}

	rep.Computes = []*resource_cluster_api_grpc_pb.Compute{computePb}
	rep.Tctx.StatusCode = codes.Ok
	return
}

func (modelApi *ResourceClusterModelApi) UpdateCompute(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.ActionRequest, rep *resource_cluster_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	spec, statusCode, err := modelApi.validateComputeSpec(db, req.Spec)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = statusCode
		return
	}

	tx := db.Begin()
	defer tx.Rollback()
	var compute resource_cluster_model.Compute
	if err = tx.Model(&compute).Where(resource_cluster_model.Compute{
		Name: spec.Name,
	}).Updates(resource_cluster_model.Compute{
		Spec:         req.Spec,
		Status:       resource_cluster_model.StatusUpdating,
		StatusReason: fmt.Sprintf("UpdateCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
	}).Error; err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}

	computePb, err := modelApi.convertCompute(&compute)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.ServerInternalError
		return
	}
	computePb.Compute.Name = spec.Name

	rep.Computes = []*resource_cluster_api_grpc_pb.Compute{computePb}
	rep.Tctx.StatusCode = codes.Ok
}

func (modelApi *ResourceClusterModelApi) DeleteCompute(tctx *logger.TraceContext, req *resource_cluster_api_grpc_pb.ActionRequest, rep *resource_cluster_api_grpc_pb.ActionReply) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	tx := db.Begin()
	defer tx.Rollback()
	var compute resource_cluster_model.Compute
	now := time.Now()
	if err = tx.Model(&compute).Where(resource_cluster_model.Compute{
		Name: req.Target,
	}).Updates(resource_cluster_model.Compute{
		Model: gorm.Model{
			DeletedAt: &now,
		},
		Status:       resource_cluster_model.StatusDeleting,
		StatusReason: fmt.Sprintf("DeleteCompute: user=%v, project=%v", req.Tctx.UserName, req.Tctx.ProjectName),
	}).Error; err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	tx.Commit()

	computePb, err := modelApi.convertCompute(&compute)
	if err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.ServerInternalError
		return
	}

	computePb.Compute.Name = req.Target
	rep.Computes = []*resource_cluster_api_grpc_pb.Compute{computePb}
	rep.Tctx.StatusCode = codes.Ok
}

func (modelApi *ResourceClusterModelApi) convertComputes(tctx *logger.TraceContext, computes []resource_cluster_model.Compute) []*resource_cluster_api_grpc_pb.Compute {
	pbComputes := make([]*resource_cluster_api_grpc_pb.Compute, len(computes))
	for i, compute := range computes {
		updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", compute.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", compute.Model.CreatedAt)
			continue
		}

		pbComputes[i] = &resource_cluster_api_grpc_pb.Compute{
			Compute: &resource_api_grpc_pb.Compute{
				Name:         compute.Name,
				Kind:         compute.Kind,
				Labels:       compute.Labels,
				Status:       compute.Status,
				StatusReason: compute.StatusReason,
				UpdatedAt:    updatedAt,
				CreatedAt:    createdAt,
			},
		}
	}

	return pbComputes
}

func (modelApi *ResourceClusterModelApi) convertCompute(compute *resource_cluster_model.Compute) (*resource_cluster_api_grpc_pb.Compute, error) {
	updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	computePb := &resource_cluster_api_grpc_pb.Compute{
		Compute: &resource_api_grpc_pb.Compute{
			Name:         compute.Name,
			Kind:         compute.Kind,
			Labels:       compute.Labels,
			Status:       compute.Status,
			StatusReason: compute.StatusReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		},
	}

	return computePb, nil
}

func (modelApi *ResourceClusterModelApi) validateComputeSpec(db *gorm.DB, specStr string) (resource_model.ComputeSpec, int64, error) {
	var spec resource_model.ComputeSpec
	var err error
	if err = json.Unmarshal([]byte(specStr), &spec); err != nil {
		return spec, codes.ClientBadRequest, err
	}
	if err = modelApi.validate.Struct(spec); err != nil {
		return spec, codes.ClientInvalidRequest, err
	}

	errors := []string{}
	switch spec.Kind {
	case resource_model.SpecKindComputeLibvirt:
		// TODO Implement Validate SpecKindComputeLibvirt
		fmt.Printf("Validate SpecKindComputeLibvirt is not implemented")

	default:
		errors = append(errors, fmt.Sprintf("Invalid kind: %v", spec.Kind))
	}

	if len(errors) > 0 {
		return spec, codes.ClientInvalidRequest, fmt.Errorf(strings.Join(errors, "\n"))
	}

	return spec, codes.Ok, nil
}

func (modelApi *ResourceClusterModelApi) SyncCompute(tctx *logger.TraceContext) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		return err
	}
	defer func() { err = db.Close() }()

	var computes []resource_cluster_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	for _, compute := range computes {
		switch compute.Status {
		case resource_cluster_model.StatusCreating:
			logger.Infof(tctx, "Found %v resource: %v", compute.Status, compute.Name)
			modelApi.InitializeCompute(db, &compute)
		}
	}

	return nil
}

func (modelApi *ResourceClusterModelApi) InitializeCompute(db *gorm.DB, compute *resource_cluster_model.Compute) {
	// TODO
	// Assgin IP address
	return
}
