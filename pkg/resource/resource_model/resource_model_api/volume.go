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

func (modelApi *ResourceModelApi) GetVolume(req *resource_api_grpc_pb.GetVolumeRequest) (*resource_api_grpc_pb.GetVolumeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Volume
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetVolumeReply{
		Volumes: modelApi.convertVolumes(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) CreateVolume(req *resource_api_grpc_pb.CreateVolumeRequest) (*resource_api_grpc_pb.CreateVolumeReply, error) {
	rep := &resource_api_grpc_pb.CreateVolumeReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.VolumeSpec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		return rep, err
	}
	glog.Info(spec.Name)

	// TODO Validate spec

	var volume resource_model.Volume
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&volume).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		volume = resource_model.Volume{
			Cluster: spec.Cluster,
			Kind:    spec.Kind,
			Name:    spec.Name,
			Spec:    req.Spec,
		}
		if err = db.Create(&volume).Error; err != nil {
			return rep, err
		}
	} else {
		return rep, fmt.Errorf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Kind, spec.Name)
	}

	volumePb, err := modelApi.convertVolume(&volume)
	if err != nil {
		return rep, err
	}
	rep.Volume = volumePb
	glog.Info("Completed CreateVolume")
	return rep, err
}

func (modelApi *ResourceModelApi) UpdateVolume(req *resource_api_grpc_pb.UpdateVolumeRequest) (*resource_api_grpc_pb.UpdateVolumeReply, error) {
	rep := &resource_api_grpc_pb.UpdateVolumeReply{}
	var err error

	glog.Info("Completed UpdateVolume")
	return rep, err
}

func (modelApi *ResourceModelApi) DeleteVolume(req *resource_api_grpc_pb.DeleteVolumeRequest) (*resource_api_grpc_pb.DeleteVolumeReply, error) {
	return nil, nil
}

func (modelApi *ResourceModelApi) convertVolumes(volumes []resource_model.Volume) []*resource_api_grpc_pb.Volume {
	pbVolumes := make([]*resource_api_grpc_pb.Volume, len(volumes))
	for i, volume := range volumes {
		updatedAt, err := ptypes.TimestampProto(volume.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(volume.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbVolumes[i] = &resource_api_grpc_pb.Volume{
			Name:      volume.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
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
		Name:      volume.Name,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return volumePb, nil
}
