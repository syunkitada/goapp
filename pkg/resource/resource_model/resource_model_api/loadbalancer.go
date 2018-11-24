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

func (modelApi *ResourceModelApi) GetLoadbalancer(req *resource_api_grpc_pb.GetLoadbalancerRequest) (*resource_api_grpc_pb.GetLoadbalancerReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var nodes []resource_model.Loadbalancer
	if err = db.Where("name like ?", req.Target).Find(&nodes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetLoadbalancerReply{
		Loadbalancers: modelApi.convertLoadbalancers(nodes),
	}, nil
}

func (modelApi *ResourceModelApi) CreateLoadbalancer(req *resource_api_grpc_pb.CreateLoadbalancerRequest) (*resource_api_grpc_pb.CreateLoadbalancerReply, error) {
	rep := &resource_api_grpc_pb.CreateLoadbalancerReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.LoadbalancerSpec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		return rep, err
	}
	glog.Info(spec.Name)

	// TODO Validate spec

	var loadbalancer resource_model.Loadbalancer
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&loadbalancer).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		loadbalancer = resource_model.Loadbalancer{
			Cluster: spec.Cluster,
			Kind:    spec.Kind,
			Name:    spec.Name,
			Spec:    req.Spec,
		}
		if err = db.Create(&loadbalancer).Error; err != nil {
			return rep, err
		}
	} else {
		return rep, fmt.Errorf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Kind, spec.Name)
	}

	loadbalancerPb, err := modelApi.convertLoadbalancer(&loadbalancer)
	if err != nil {
		return rep, err
	}
	rep.Loadbalancer = loadbalancerPb
	glog.Info("Completed CreateLoadbalancer")
	return rep, err
}

func (modelApi *ResourceModelApi) UpdateLoadbalancer(req *resource_api_grpc_pb.UpdateLoadbalancerRequest) (*resource_api_grpc_pb.UpdateLoadbalancerReply, error) {
	rep := &resource_api_grpc_pb.UpdateLoadbalancerReply{}
	var err error

	glog.Info("Completed UpdateLoadbalancer")
	return rep, err
}

func (modelApi *ResourceModelApi) DeleteLoadbalancer(req *resource_api_grpc_pb.DeleteLoadbalancerRequest) (*resource_api_grpc_pb.DeleteLoadbalancerReply, error) {
	return nil, nil
}

func (modelApi *ResourceModelApi) convertLoadbalancers(loadbalancers []resource_model.Loadbalancer) []*resource_api_grpc_pb.Loadbalancer {
	pbLoadbalancers := make([]*resource_api_grpc_pb.Loadbalancer, len(loadbalancers))
	for i, loadbalancer := range loadbalancers {
		updatedAt, err := ptypes.TimestampProto(loadbalancer.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(loadbalancer.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbLoadbalancers[i] = &resource_api_grpc_pb.Loadbalancer{
			Name:      loadbalancer.Name,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbLoadbalancers
}

func (modelApi *ResourceModelApi) convertLoadbalancer(loadbalancer *resource_model.Loadbalancer) (*resource_api_grpc_pb.Loadbalancer, error) {
	updatedAt, err := ptypes.TimestampProto(loadbalancer.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(loadbalancer.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	loadbalancerPb := &resource_api_grpc_pb.Loadbalancer{
		Name:      loadbalancer.Name,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}

	return loadbalancerPb, nil
}
