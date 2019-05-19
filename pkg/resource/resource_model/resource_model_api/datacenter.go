package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetDatacenter(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.PhysicalActionReply) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var datacenter resource_model.Datacenter
	if err = db.Where(&resource_model.Datacenter{
		Name: resource,
	}).First(&datacenter).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Datacenter = modelApi.convertDatacenter(tctx, &datacenter)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetDatacenters(tctx *logger.TraceContext, db *gorm.DB,
	query *resource_api_grpc_pb.Query, rep *resource_api_grpc_pb.PhysicalActionReply) (int64, error) {
	var err error
	var datacenters []resource_model.Datacenter
	if err = db.Find(&datacenters).Error; err != nil {
		return codes.RemoteDbError, err
	}
	rep.Datacenters = modelApi.convertDatacenters(tctx, datacenters)
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateDatacenter(tctx *logger.TraceContext, db *gorm.DB,
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

	var specs []resource_model.DatacenterSpecData
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

		var data resource_model.Datacenter
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.Datacenter{
				Kind:         spec.Kind,
				Name:         spec.Name,
				Description:  spec.Description,
				Region:       spec.Region,
				DomainSuffix: spec.DomainSuffix,
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

func (modelApi *ResourceModelApi) UpdateDatacenter(tctx *logger.TraceContext, db *gorm.DB,
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

	var specs []resource_model.DatacenterSpecData
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
		datacenter := &resource_model.Datacenter{
			Kind:         spec.Kind,
			Description:  spec.Description,
			Region:       spec.Region,
			DomainSuffix: spec.DomainSuffix,
		}
		if err = tx.Model(datacenter).Where("name = ?", spec.Name).Updates(datacenter).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteDatacenter(tctx *logger.TraceContext, db *gorm.DB,
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

		if err = tx.Delete(&resource_model.Datacenter{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) convertDatacenter(tctx *logger.TraceContext,
	datacenter *resource_model.Datacenter) *resource_api_grpc_pb.Datacenter {
	updatedAt, err := ptypes.TimestampProto(datacenter.Model.UpdatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", datacenter.Model.UpdatedAt)
	}
	createdAt, err := ptypes.TimestampProto(datacenter.Model.CreatedAt)
	if err != nil {
		logger.Warningf(tctx, err,
			"Failed ptypes.TimestampProto: %v", datacenter.Model.CreatedAt)
	}

	return &resource_api_grpc_pb.Datacenter{
		Name:         datacenter.Name,
		Kind:         datacenter.Kind,
		Description:  datacenter.Description,
		Region:       datacenter.Region,
		DomainSuffix: datacenter.DomainSuffix,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}
}

func (modelApi *ResourceModelApi) convertDatacenters(tctx *logger.TraceContext,
	datacenters []resource_model.Datacenter) []*resource_api_grpc_pb.Datacenter {
	pbDatacenters := make([]*resource_api_grpc_pb.Datacenter, len(datacenters))
	for i, datacenter := range datacenters {
		pbDatacenters[i] = modelApi.convertDatacenter(tctx, &datacenter)
	}

	return pbDatacenters
}
