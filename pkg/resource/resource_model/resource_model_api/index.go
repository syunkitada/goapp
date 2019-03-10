package resource_model_api

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetPhysicalIndex(tctx *logger.TraceContext, req *resource_api_grpc_pb.ActionRequest, rep *resource_api_grpc_pb.ActionReply) {
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

	var datacenters []resource_model.Datacenter
	if err = db.Find(&datacenters).Error; err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}

	rep.Datacenters = modelApi.convertDatacenters(tctx, datacenters)

	rep.Tctx.StatusCode = codes.Ok
}

func (modelApi *ResourceModelApi) convertDatacenters(tctx *logger.TraceContext, datacenters []resource_model.Datacenter) []*resource_api_grpc_pb.Datacenter {
	pbDatacenters := make([]*resource_api_grpc_pb.Datacenter, len(datacenters))
	for i, datacenter := range datacenters {
		updatedAt, err := ptypes.TimestampProto(datacenter.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", datacenter.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(datacenter.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", datacenter.Model.CreatedAt)
			continue
		}

		pbDatacenters[i] = &resource_api_grpc_pb.Datacenter{
			Region:    datacenter.Region,
			Name:      datacenter.Name,
			Kind:      datacenter.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbDatacenters
}
