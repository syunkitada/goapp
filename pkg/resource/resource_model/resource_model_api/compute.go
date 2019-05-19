package resource_model_api

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.VirtualActionReply) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var compute resource_model.Compute
	if err = db.Where(&resource_model.Compute{
		Name: resource,
	}).First(&compute).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Compute = modelApi.convertCompute(tctx, &compute)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetComputes(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.VirtualActionReply) (int64, error) {
	var err error
	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Computes = modelApi.convertComputes(tctx, computes)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.ComputeSpecData
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

		var data resource_model.Compute
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Compute{
				Kind:        spec.Kind,
				Name:        spec.Name,
				Description: spec.Description,
				Cluster:     spec.Cluster,
				Domain:      spec.Domain,
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

func (modelApi *ResourceModelApi) UpdateCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query) (int64, error) {
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

	var specs []resource_model.ComputeSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}
		compute := &resource_model.Compute{
			Kind:        spec.Kind,
			Description: spec.Description,
			Cluster:     spec.Cluster,
			Domain:      spec.Domain,
		}
		if err = tx.Model(compute).Where("name = ?", spec.Name).Updates(compute).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteCompute(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query) (int64, error) {
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

func (modelApi *ResourceModelApi) convertCompute(tctx *logger.TraceContext,
	compute *resource_model.Compute) *resource_api_grpc_pb.Compute {
	updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", compute.Model.UpdatedAt)
	}
	createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", compute.Model.CreatedAt)
	}

	return &resource_api_grpc_pb.Compute{
		Name:        compute.Name,
		Kind:        compute.Kind,
		Description: compute.Description,
		Cluster:     compute.Cluster,
		UpdatedAt:   updatedAt,
		CreatedAt:   createdAt,
	}
}

func (modelApi *ResourceModelApi) convertComputes(tctx *logger.TraceContext,
	computes []resource_model.Compute) []*resource_api_grpc_pb.Compute {
	pbComputes := make([]*resource_api_grpc_pb.Compute, len(computes))
	for i, compute := range computes {
		pbComputes[i] = modelApi.convertCompute(tctx, &compute)
	}

	return pbComputes
}

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
