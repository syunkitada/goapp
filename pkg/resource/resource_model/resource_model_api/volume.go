package resource_model_api

import (
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetVolume(req *resource_api_grpc_pb.GetVolumeRequest) *resource_api_grpc_pb.GetVolumeReply {
	rep := &resource_api_grpc_pb.GetVolumeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var volumes []resource_model.Volume
	if err = db.Where("name like ?", req.Target).Find(&volumes).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Volumes = modelApi.convertVolumes(req.TraceId, volumes)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) CreateVolume(req *resource_api_grpc_pb.CreateVolumeRequest) *resource_api_grpc_pb.CreateVolumeReply {
	rep := &resource_api_grpc_pb.CreateVolumeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateVolumeSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	var volume resource_model.Volume
	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&volume).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}

		volume = resource_model.Volume{
			Cluster:      spec.Cluster,
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_model.StatusActive,
			StatusReason: fmt.Sprintf("CreateVolume: user=%v, project=%v", req.UserName, req.ProjectName),
		}
		if err = tx.Create(&volume).Error; err != nil {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}
	} else {
		rep.Err = fmt.Sprintf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Name)
		rep.StatusCode = codes.ClientAlreadyExists
		return rep
	}
	tx.Commit()

	volumePb, err := modelApi.convertVolume(&volume)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Volume = volumePb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) UpdateVolume(req *resource_api_grpc_pb.UpdateVolumeRequest) *resource_api_grpc_pb.UpdateVolumeReply {
	rep := &resource_api_grpc_pb.UpdateVolumeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateVolumeSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	tx := db.Begin()
	defer tx.Rollback()
	var volume resource_model.Volume
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&volume).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	volume.Spec = req.Spec
	volume.Status = resource_model.StatusActive
	volume.StatusReason = fmt.Sprintf("UpdateVolume: user=%v, project=%v", req.UserName, req.ProjectName)
	tx.Save(volume)
	tx.Commit()

	volumePb, err := modelApi.convertVolume(&volume)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Volume = volumePb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) DeleteVolume(req *resource_api_grpc_pb.DeleteVolumeRequest) *resource_api_grpc_pb.DeleteVolumeReply {
	rep := &resource_api_grpc_pb.DeleteVolumeReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	tx := db.Begin()
	defer tx.Rollback()
	var volume resource_model.Volume
	if err = tx.Where("name = ?", req.Target).Delete(&volume).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	tx.Commit()

	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) convertVolumes(traceId string, volumes []resource_model.Volume) []*resource_api_grpc_pb.Volume {
	pbVolumes := make([]*resource_api_grpc_pb.Volume, len(volumes))
	for i, volume := range volumes {
		updatedAt, err := ptypes.TimestampProto(volume.Model.UpdatedAt)
		if err != nil {
			logger.TraceError(traceId, modelApi.host, modelApi.name, map[string]string{
				"Msg":    fmt.Sprintf("Failed ptypes.TimestampProto: %v", volume.Model.UpdatedAt),
				"Err":    err.Error(),
				"Method": "CreateVolume",
			})
			continue
		}
		createdAt, err := ptypes.TimestampProto(volume.Model.CreatedAt)
		if err != nil {
			logger.TraceError(traceId, modelApi.host, modelApi.name, map[string]string{
				"Msg":    fmt.Sprintf("Failed ptypes.TimestampProto: %v", volume.Model.CreatedAt),
				"Err":    err.Error(),
				"Method": "CreateVolume",
			})
			continue
		}

		pbVolumes[i] = &resource_api_grpc_pb.Volume{
			Cluster:      volume.Cluster,
			Name:         volume.Name,
			Kind:         volume.Kind,
			Labels:       volume.Labels,
			Status:       volume.Status,
			StatusReason: volume.StatusReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		}
	}

	return pbVolumes
}

func (modelApi *ResourceModelApi) convertVolume(volume *resource_model.Volume) (*resource_api_grpc_pb.Volume, error) {
	updatedAt, err := ptypes.TimestampProto(volume.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(volume.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	volumePb := &resource_api_grpc_pb.Volume{
		Cluster:      volume.Cluster,
		Name:         volume.Name,
		Kind:         volume.Kind,
		Labels:       volume.Labels,
		Status:       volume.Status,
		StatusReason: volume.StatusReason,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}

	return volumePb, nil
}

func (modelApi *ResourceModelApi) validateVolumeSpec(db *gorm.DB, specStr string) (resource_model.VolumeSpec, int64, error) {
	var spec resource_model.VolumeSpec
	var err error
	if err = json.Unmarshal([]byte(specStr), &spec); err != nil {
		return spec, codes.ClientBadRequest, err
	}
	if err = modelApi.validate.Struct(spec); err != nil {
		return spec, codes.ClientInvalidRequest, err
	}

	ok, err := modelApi.ValidateClusterName(db, spec.Cluster)
	if err != nil {
		return spec, codes.RemoteDbError, err
	}
	if !ok {
		return spec, codes.ClientInvalidRequest, fmt.Errorf("Invalid cluster: %v", spec.Cluster)
	}

	errors := []string{}
	switch spec.Spec.Kind {
	case resource_model.SpecKindVolumeNfs:
		// TODO Implement Validate SpecKindVolumeNfs
		logger.Warning(modelApi.host, modelApi.name, "Validate SpecKindVolumeNfs is not implemented")

	default:
		errors = append(errors, fmt.Sprintf("Invalid kind: %v", spec.Spec.Kind))
	}

	if len(errors) > 0 {
		return spec, codes.ClientInvalidRequest, fmt.Errorf(strings.Join(errors, "\n"))
	}

	return spec, codes.Ok, nil
}
