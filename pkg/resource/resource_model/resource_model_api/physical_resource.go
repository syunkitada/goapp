package resource_model_api

import (
	"encoding/json"
	"strings"

	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) CreatePhysicalResource(tctx *logger.TraceContext, db *gorm.DB, query *resource_api_grpc_pb.Query) (error, int64) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return error_utils.NewInvalidRequestError("NotFound Specs"), codes.ClientBadRequest
	}

	var specs []resource_model.PhysicalResourceSpecData
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return err, codes.ClientBadRequest
	}
	// TODO validate

	tx := db.Begin()
	defer tx.Rollback()

	for _, spec := range specs {
		var data resource_model.PhysicalResource
		if err = tx.Where("name = ? and datacenter = ?", spec.Name, spec.Datacenter).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return err, codes.RemoteDbError
			}

			data = resource_model.PhysicalResource{
				Kind:          spec.Kind,
				Name:          spec.Name,
				Datacenter:    spec.Datacenter,
				Cluster:       spec.Cluster,
				Rack:          spec.Rack,
				PhysicalModel: spec.Model,
				RackPosition:  spec.RackPosition,
				PowerLinks:    strings.Join(spec.PowerLinks, ","),
				NetLinks:      strings.Join(spec.NetLinks, ","),
				Spec:          spec.Spec,
			}
			if err = tx.Create(&data).Error; err != nil {
				return err, codes.RemoteDbError
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return err, codes.ClientAlreadyExists
		}
	}

	tx.Commit()

	return nil, codes.Ok
}

func (modelApi *ResourceModelApi) convertPhysicalResource(tctx *logger.TraceContext, physicalResource resource_model.PhysicalResource) (*resource_api_grpc_pb.PhysicalResource, error) {
	updatedAt, err := ptypes.TimestampProto(physicalResource.Model.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(physicalResource.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.PhysicalResource{
		Name:      physicalResource.Name,
		Kind:      physicalResource.Kind,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func (modelApi *ResourceModelApi) convertPhysicalResources(tctx *logger.TraceContext, physicalResourcess []resource_model.PhysicalResource) []*resource_api_grpc_pb.PhysicalResource {
	pbPhysicalResources := make([]*resource_api_grpc_pb.PhysicalResource, len(physicalResourcess))
	for i, physicalResources := range physicalResourcess {
		updatedAt, err := ptypes.TimestampProto(physicalResources.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", physicalResources.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(physicalResources.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", physicalResources.Model.CreatedAt)
			continue
		}

		pbPhysicalResources[i] = &resource_api_grpc_pb.PhysicalResource{
			Name:      physicalResources.Name,
			Kind:      physicalResources.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbPhysicalResources
}
