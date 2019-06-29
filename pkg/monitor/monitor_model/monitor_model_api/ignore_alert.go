package monitor_model_api

import (
	// "time"
	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model"
)

func (modelApi *MonitorModelApi) GetIgnoreAlert(tctx *logger.TraceContext, req *monitor_api_grpc_pb.GetIgnoreAlertRequest) *monitor_api_grpc_pb.GetIgnoreAlertReply {
	rep := &monitor_api_grpc_pb.GetIgnoreAlertReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	var ignoreAlerts []monitor_model.IgnoreAlert
	if req.Index == "%" || req.Index == "" {
		if err = db.Find(&ignoreAlerts).Error; err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
	} else {
		if err = db.Where("index like ?", req.Index).Find(&ignoreAlerts).Error; err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
	}

	rep.IgnoreAlerts = modelApi.convertIgnoreAlerts(tctx, ignoreAlerts)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) CreateIgnoreAlert(tctx *logger.TraceContext, req *monitor_api_grpc_pb.CreateIgnoreAlertRequest) *monitor_api_grpc_pb.CreateIgnoreAlertReply {
	rep := &monitor_api_grpc_pb.CreateIgnoreAlertReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	spec, statusCode, err := modelApi.validateIgnoreAlertSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	ignoreAlert := monitor_model.IgnoreAlert{
		Index:  spec.Index,
		Host:   spec.Host,
		Name:   spec.Name,
		Level:  spec.Level,
		Reason: spec.Reason,
		Until:  spec.Until,
	}
	if err = db.Create(&ignoreAlert).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	ignoreAlertPb, err := modelApi.convertIgnoreAlert(&ignoreAlert)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.IgnoreAlert = ignoreAlertPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) UpdateIgnoreAlert(tctx *logger.TraceContext, req *monitor_api_grpc_pb.UpdateIgnoreAlertRequest) *monitor_api_grpc_pb.UpdateIgnoreAlertReply {
	rep := &monitor_api_grpc_pb.UpdateIgnoreAlertReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	spec, statusCode, err := modelApi.validateIgnoreAlertSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	tx := db.Begin()
	defer tx.Rollback()
	var ignoreAlert monitor_model.IgnoreAlert
	if err = tx.Where("id = ?", req.Id).First(&ignoreAlert).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	ignoreAlert.Index = spec.Index
	ignoreAlert.Host = spec.Host
	ignoreAlert.Name = spec.Name
	ignoreAlert.Level = spec.Level
	ignoreAlert.Reason = spec.Reason
	ignoreAlert.Until = spec.Until
	tx.Save(ignoreAlert)
	tx.Commit()

	ignoreAlertPb, err := modelApi.convertIgnoreAlert(&ignoreAlert)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.IgnoreAlert = ignoreAlertPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) DeleteIgnoreAlert(tctx *logger.TraceContext, req *monitor_api_grpc_pb.DeleteIgnoreAlertRequest) *monitor_api_grpc_pb.DeleteIgnoreAlertReply {
	rep := &monitor_api_grpc_pb.DeleteIgnoreAlertReply{}
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var db *gorm.DB
	db, err = modelApi.open(tctx)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	defer db.Close()

	tx := db.Begin()
	defer tx.Rollback()
	var ignoreAlert monitor_model.IgnoreAlert
	if err = tx.Where("id = ?", req.Id).Delete(&ignoreAlert).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	tx.Commit()

	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *MonitorModelApi) convertIgnoreAlerts(tctx *logger.TraceContext, ignoreAlerts []monitor_model.IgnoreAlert) []*monitor_api_grpc_pb.IgnoreAlert {
	pbIgnoreAlerts := make([]*monitor_api_grpc_pb.IgnoreAlert, len(ignoreAlerts))
	for i, ignoreAlert := range ignoreAlerts {
		updatedAt, err := ptypes.TimestampProto(ignoreAlert.Model.UpdatedAt)
		if err != nil {
			continue
		}
		createdAt, err := ptypes.TimestampProto(ignoreAlert.Model.CreatedAt)
		if err != nil {
			continue
		}

		pbIgnoreAlerts[i] = &monitor_api_grpc_pb.IgnoreAlert{
			Index:     ignoreAlert.Index,
			Host:      ignoreAlert.Host,
			Name:      ignoreAlert.Name,
			Level:     ignoreAlert.Level,
			User:      ignoreAlert.User,
			Reason:    ignoreAlert.Reason,
			Until:     ignoreAlert.Until,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbIgnoreAlerts
}

func (modelApi *MonitorModelApi) convertIgnoreAlert(ignoreAlert *monitor_model.IgnoreAlert) (*monitor_api_grpc_pb.IgnoreAlert, error) {
	updatedAt, err := ptypes.TimestampProto(ignoreAlert.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(ignoreAlert.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	ignoreAlertPb := &monitor_api_grpc_pb.IgnoreAlert{
		Index:     ignoreAlert.Index,
		Host:      ignoreAlert.Host,
		Name:      ignoreAlert.Name,
		Level:     ignoreAlert.Level,
		User:      ignoreAlert.User,
		Reason:    ignoreAlert.Reason,
		Until:     ignoreAlert.Until,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return ignoreAlertPb, nil
}

func (modelApi *MonitorModelApi) validateIgnoreAlertSpec(db *gorm.DB, specStr string) (monitor_model.IgnoreAlertSpec, int64, error) {
	var spec monitor_model.IgnoreAlertSpec
	var err error
	if err = json.Unmarshal([]byte(specStr), &spec); err != nil {
		return spec, codes.ClientBadRequest, err
	}
	if err = modelApi.validate.Struct(spec); err != nil {
		return spec, codes.ClientInvalidRequest, err
	}

	return spec, codes.Ok, nil
}
