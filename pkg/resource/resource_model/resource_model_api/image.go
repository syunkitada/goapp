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

func (modelApi *ResourceModelApi) GetImage(req *resource_api_grpc_pb.GetImageRequest) (*resource_api_grpc_pb.GetImageReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Image
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetImageReply{
		Images: modelApi.convertImages(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) CreateImage(req *resource_api_grpc_pb.CreateImageRequest) (*resource_api_grpc_pb.CreateImageReply, error) {
	rep := &resource_api_grpc_pb.CreateImageReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.ImageSpec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		return rep, err
	}
	glog.Info(spec.Name)

	// TODO Validate spec

	var image resource_model.Image
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&image).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		image = resource_model.Image{
			Cluster: spec.Cluster,
			Kind:    spec.Kind,
			Name:    spec.Name,
			Spec:    req.Spec,
		}
		if err = db.Create(&image).Error; err != nil {
			return rep, err
		}
	} else {
		return rep, fmt.Errorf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Kind, spec.Name)
	}

	imagePb, err := modelApi.convertImage(&image)
	if err != nil {
		return rep, err
	}
	rep.Image = imagePb
	glog.Info("Completed CreateImage")
	return rep, err
}

func (modelApi *ResourceModelApi) UpdateImage(req *resource_api_grpc_pb.UpdateImageRequest) (*resource_api_grpc_pb.UpdateImageReply, error) {
	rep := &resource_api_grpc_pb.UpdateImageReply{}
	var err error

	glog.Info("Completed UpdateImage")
	return rep, err
}

func (modelApi *ResourceModelApi) DeleteImage(req *resource_api_grpc_pb.DeleteImageRequest) (*resource_api_grpc_pb.DeleteImageReply, error) {
	return nil, nil
}

func (modelApi *ResourceModelApi) convertImages(images []resource_model.Image) []*resource_api_grpc_pb.Image {
	pbImages := make([]*resource_api_grpc_pb.Image, len(images))
	for i, image := range images {
		updatedAt, err := ptypes.TimestampProto(image.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(image.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbImages[i] = &resource_api_grpc_pb.Image{
			Name:      image.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbImages
}

func (modelApi *ResourceModelApi) convertImage(image *resource_model.Image) (*resource_api_grpc_pb.Image, error) {
	updatedAt, err := ptypes.TimestampProto(image.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(image.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	imagePb := &resource_api_grpc_pb.Image{
		Name:      image.Name,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return imagePb, nil
}
