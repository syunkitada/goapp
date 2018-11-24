package resource_model_api

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetContainer(req *resource_api_grpc_pb.GetContainerRequest) (*resource_api_grpc_pb.GetContainerReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Container
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetContainerReply{
		Containers: modelApi.convertContainers(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) CreateContainer(req *resource_api_grpc_pb.CreateContainerRequest) (*resource_api_grpc_pb.CreateContainerReply, error) {
	rep := &resource_api_grpc_pb.CreateContainerReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.ContainerSpec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		return rep, err
	}
	glog.Info(spec.Name)

	// TODO Validate spec

	var container resource_model.Container
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&container).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		container = resource_model.Container{
			Cluster: spec.Cluster,
			Kind:    spec.Kind,
			Name:    spec.Name,
			Spec:    req.Spec,
		}
		if err = db.Create(&container).Error; err != nil {
			return rep, err
		}
	} else {
		return rep, fmt.Errorf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Kind, spec.Name)
	}

	containerPb, err := modelApi.convertContainer(&container)
	if err != nil {
		return rep, err
	}
	rep.Container = containerPb
	glog.Info("Completed CreateContainer")
	return rep, err
}

func (modelApi *ResourceModelApi) UpdateContainer(req *resource_api_grpc_pb.UpdateContainerRequest) (*resource_api_grpc_pb.UpdateContainerReply, error) {
	rep := &resource_api_grpc_pb.UpdateContainerReply{}
	var err error

	glog.Info("Completed UpdateContainer")
	return rep, err
}

func (modelApi *ResourceModelApi) DeleteContainer(req *resource_api_grpc_pb.DeleteContainerRequest) (*resource_api_grpc_pb.DeleteContainerReply, error) {
	return nil, nil
}

func (modelApi *ResourceModelApi) convertContainers(containers []resource_model.Container) []*resource_api_grpc_pb.Container {
	pbContainers := make([]*resource_api_grpc_pb.Container, len(containers))
	for i, container := range containers {
		updatedAt, err := ptypes.TimestampProto(container.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(container.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbContainers[i] = &resource_api_grpc_pb.Container{
			Name:      container.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
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
		Name:      container.Name,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return containerPb, nil
}
