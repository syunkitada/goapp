package resource_model_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

func (modelApi *ResourceModelApi) PhysicalAction(tctx *logger.TraceContext,
	req *resource_api_grpc_pb.PhysicalActionRequest, rep *resource_api_grpc_pb.PhysicalActionReply) {
	var err error
	var statusCode int64
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "GetDatacenter":
			statusCode, err = modelApi.GetDatacenter(tctx, db, query, rep)
		case "GetDatacenters", "GetIndex":
			statusCode, err = modelApi.GetDatacenters(tctx, db, query, rep)
		case "CreateDatacenter":
			statusCode, err = modelApi.CreateDatacenter(tctx, db, query)
		case "UpdateDatacenter":
			statusCode, err = modelApi.UpdateDatacenter(tctx, db, query)
		case "DeleteDatacenter":
			statusCode, err = modelApi.DeleteDatacenter(tctx, db, query)

		case "GetFloor":
			statusCode, err = modelApi.GetFloor(tctx, db, query, rep)
		case "GetFloors":
			statusCode, err = modelApi.GetFloors(tctx, db, query, rep)
		case "CreateFloor":
			statusCode, err = modelApi.CreateFloor(tctx, db, query)
		case "UpdateFloor":
			statusCode, err = modelApi.UpdateFloor(tctx, db, query)
		case "DeleteFloor":
			statusCode, err = modelApi.DeleteFloor(tctx, db, query)

		case "GetRack":
			statusCode, err = modelApi.GetRack(tctx, db, query, rep)
		case "GetRacks":
			statusCode, err = modelApi.GetRacks(tctx, db, query, rep)
		case "CreateRack":
			statusCode, err = modelApi.CreateRack(tctx, db, query)
		case "UpdateRack":
			statusCode, err = modelApi.UpdateRack(tctx, db, query)
		case "DeleteRack":
			statusCode, err = modelApi.DeleteRack(tctx, db, query)

		case "GetPhysicalResource":
			statusCode, err = modelApi.GetPhysicalResource(tctx, db, query, rep)
		case "GetPhysicalResources":
			statusCode, err = modelApi.GetPhysicalResources(tctx, db, query, rep)
		case "CreatePhysicalResource":
			statusCode, err = modelApi.CreatePhysicalResource(tctx, db, query)
		case "UpdatePhysicalResource":
			statusCode, err = modelApi.UpdatePhysicalResource(tctx, db, query)
		case "DeletePhysicalResource":
			statusCode, err = modelApi.DeletePhysicalResource(tctx, db, query)

		case "GetPhysicalModel":
			statusCode, err = modelApi.GetPhysicalModel(tctx, db, query, rep)
		case "GetPhysicalModels":
			statusCode, err = modelApi.GetPhysicalModels(tctx, db, query, rep)
		case "CreatePhysicalModel":
			statusCode, err = modelApi.CreatePhysicalModel(tctx, db, query)
		case "UpdatePhysicalModel":
			statusCode, err = modelApi.UpdatePhysicalModel(tctx, db, query)
		case "DeletePhysicalModel":
			statusCode, err = modelApi.DeletePhysicalModel(tctx, db, query)
		}

		if err != nil {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = statusCode
			return
		}
	}

	rep.Tctx.StatusCode = statusCode
}

func (modelApi *ResourceModelApi) VirtualAction(tctx *logger.TraceContext,
	req *resource_api_grpc_pb.VirtualActionRequest, rep *resource_api_grpc_pb.VirtualActionReply) {
	var err error
	var statusCode int64
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		rep.Tctx.Err = err.Error()
		rep.Tctx.StatusCode = codes.RemoteDbError
		return
	}
	defer func() { err = db.Close() }()

	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "GetCluster":
			statusCode, err = modelApi.GetCluster(tctx, db, query, rep)
		case "GetClusters", "GetIndex":
			statusCode, err = modelApi.GetClusters(tctx, db, query, rep)
		case "CreateCluster":
			statusCode, err = modelApi.CreateCluster(tctx, db, query)
		case "UpdateCluster":
			statusCode, err = modelApi.UpdateCluster(tctx, db, query)
		case "DeleteCluster":
			statusCode, err = modelApi.DeleteCluster(tctx, db, query)

		case "GetCompute":
			statusCode, err = modelApi.GetCompute(tctx, db, query, rep)
		case "GetComputes":
			statusCode, err = modelApi.GetComputes(tctx, db, query, rep)
		case "CreateCompute":
			statusCode, err = modelApi.CreateCompute(tctx, db, query)
		case "UpdateCompute":
			statusCode, err = modelApi.UpdateCompute(tctx, db, query)
		case "DeleteCompute":
			statusCode, err = modelApi.DeleteCompute(tctx, db, query)
		}

		if err != nil {
			rep.Tctx.Err = err.Error()
			rep.Tctx.StatusCode = statusCode
			return
		}
	}

	rep.Tctx.StatusCode = statusCode
}
