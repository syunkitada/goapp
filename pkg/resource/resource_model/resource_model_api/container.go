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

func (modelApi *ResourceModelApi) GetContainer(req *resource_api_grpc_pb.GetContainerRequest) *resource_api_grpc_pb.GetContainerReply {
	rep := &resource_api_grpc_pb.GetContainerReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var containers []resource_model.Container
	if err = db.Where("name like ?", req.Target).Find(&containers).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Containers = modelApi.convertContainers(req.TraceId, containers)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) CreateContainer(req *resource_api_grpc_pb.CreateContainerRequest) *resource_api_grpc_pb.CreateContainerReply {
	rep := &resource_api_grpc_pb.CreateContainerReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateContainerSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	var container resource_model.Container
	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&container).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
			return rep
		}

		container = resource_model.Container{
			Cluster:      spec.Cluster,
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_model.StatusActive,
			StatusReason: fmt.Sprintf("CreateContainer: user=%v, project=%v", req.UserName, req.ProjectName),
		}
		if err = tx.Create(&container).Error; err != nil {
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

	containerPb, err := modelApi.convertContainer(&container)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Container = containerPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) UpdateContainer(req *resource_api_grpc_pb.UpdateContainerRequest) *resource_api_grpc_pb.UpdateContainerReply {
	rep := &resource_api_grpc_pb.UpdateContainerReply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateContainerSpec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	tx := db.Begin()
	defer tx.Rollback()
	var container resource_model.Container
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&container).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	container.Spec = req.Spec
	container.Status = resource_model.StatusActive
	container.StatusReason = fmt.Sprintf("UpdateContainer: user=%v, project=%v", req.UserName, req.ProjectName)
	tx.Save(container)
	tx.Commit()

	containerPb, err := modelApi.convertContainer(&container)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Container = containerPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) DeleteContainer(req *resource_api_grpc_pb.DeleteContainerRequest) *resource_api_grpc_pb.DeleteContainerReply {
	rep := &resource_api_grpc_pb.DeleteContainerReply{}

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
	var container resource_model.Container
	if err = tx.Where("name = ?", req.Target).Delete(&container).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	tx.Commit()

	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) convertContainers(traceId string, containers []resource_model.Container) []*resource_api_grpc_pb.Container {
	pbContainers := make([]*resource_api_grpc_pb.Container, len(containers))
	for i, container := range containers {
		updatedAt, err := ptypes.TimestampProto(container.Model.UpdatedAt)
		if err != nil {
			logger.TraceError(traceId, modelApi.host, modelApi.name, map[string]string{
				"Msg":    fmt.Sprintf("Failed ptypes.TimestampProto: %v", container.Model.UpdatedAt),
				"Err":    err.Error(),
				"Method": "CreateContainer",
			})
			continue
		}
		createdAt, err := ptypes.TimestampProto(container.Model.CreatedAt)
		if err != nil {
			logger.TraceError(traceId, modelApi.host, modelApi.name, map[string]string{
				"Msg":    fmt.Sprintf("Failed ptypes.TimestampProto: %v", container.Model.CreatedAt),
				"Err":    err.Error(),
				"Method": "CreateContainer",
			})
			continue
		}

		pbContainers[i] = &resource_api_grpc_pb.Container{
			Cluster:      container.Cluster,
			Name:         container.Name,
			Kind:         container.Kind,
			Labels:       container.Labels,
			Status:       container.Status,
			StatusReason: container.StatusReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		}
	}

	return pbContainers
}

func (modelApi *ResourceModelApi) convertContainer(container *resource_model.Container) (*resource_api_grpc_pb.Container, error) {
	updatedAt, err := ptypes.TimestampProto(container.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(container.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	containerPb := &resource_api_grpc_pb.Container{
		Cluster:      container.Cluster,
		Name:         container.Name,
		Kind:         container.Kind,
		Labels:       container.Labels,
		Status:       container.Status,
		StatusReason: container.StatusReason,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}

	return containerPb, nil
}

func (modelApi *ResourceModelApi) validateContainerSpec(db *gorm.DB, specStr string) (resource_model.ContainerSpec, int64, error) {
	var spec resource_model.ContainerSpec
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
	case resource_model.SpecKindContainerDocker:
		// TODO Implement Validate SpecKindContainerDocker
		logger.Warning(modelApi.host, modelApi.name, "Validate SpecKindContainerDocker is not implemented")

	default:
		errors = append(errors, fmt.Sprintf("Invalid kind: %v", spec.Spec.Kind))
	}

	if len(errors) > 0 {
		return spec, codes.ClientInvalidRequest, fmt.Errorf(strings.Join(errors, "\n"))
	}

	return spec, codes.Ok, nil
}
