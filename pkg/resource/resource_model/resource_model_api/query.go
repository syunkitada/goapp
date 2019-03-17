package resource_model_api

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) UserQuery(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
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

	for _, query := range req.Queries {
		switch query.Kind {
		case "GetDatacenters":
			var datacenters []resource_model.Datacenter
			if err = db.Find(&datacenters).Error; err != nil {
				rep.Tctx.Err = err.Error()
				rep.Tctx.StatusCode = codes.RemoteDbError
				return
			}
			rep.Datacenters = modelApi.convertDatacenters(tctx, datacenters)
		case "GetFloors":
			var floors []resource_model.Floor
			if err = db.Where("datacenter = ?", query.DatacenterName).Find(&floors).Error; err != nil {
				rep.Tctx.Err = err.Error()
				rep.Tctx.StatusCode = codes.RemoteDbError
				return
			}
			rep.Floors = modelApi.convertFloors(tctx, floors)
		case "GetRacks":
			var racks []resource_model.Rack
			if err = db.Where("datacenter = ?", query.DatacenterName).Find(&racks).Error; err != nil {
				rep.Tctx.Err = err.Error()
				rep.Tctx.StatusCode = codes.RemoteDbError
				return
			}
			rep.Racks = modelApi.convertRacks(tctx, racks)
		case "GetPhysicalResources":
			var physicalResources []resource_model.PhysicalResource
			if err = db.Where("datacenter = ?", query.DatacenterName).Find(&physicalResources).Error; err != nil {
				rep.Tctx.Err = err.Error()
				rep.Tctx.StatusCode = codes.RemoteDbError
				return
			}
			rep.PhysicalResources = modelApi.convertPhysicalResources(tctx, physicalResources)
		}
	}

	rep.Tctx.StatusCode = codes.Ok
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

func (modelApi *ResourceModelApi) convertFloors(tctx *logger.TraceContext, floorss []resource_model.Floor) []*resource_api_grpc_pb.Floor {
	pbFloors := make([]*resource_api_grpc_pb.Floor, len(floorss))
	for i, floors := range floorss {
		updatedAt, err := ptypes.TimestampProto(floors.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", floors.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(floors.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", floors.Model.CreatedAt)
			continue
		}

		pbFloors[i] = &resource_api_grpc_pb.Floor{
			Name:      floors.Name,
			Kind:      floors.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbFloors
}

func (modelApi *ResourceModelApi) convertRacks(tctx *logger.TraceContext, rackss []resource_model.Rack) []*resource_api_grpc_pb.Rack {
	pbRacks := make([]*resource_api_grpc_pb.Rack, len(rackss))
	for i, racks := range rackss {
		updatedAt, err := ptypes.TimestampProto(racks.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", racks.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(racks.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", racks.Model.CreatedAt)
			continue
		}

		pbRacks[i] = &resource_api_grpc_pb.Rack{
			Name:      racks.Name,
			Kind:      racks.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbRacks
}
