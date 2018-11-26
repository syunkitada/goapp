package resource_cluster_model_api

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceClusterModelApi) GetCompute(req *resource_cluster_api_grpc_pb.GetComputeRequest) (*resource_cluster_api_grpc_pb.GetComputeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var computes []resource_cluster_model.Compute
	if err = db.Where("name like ?", req.Target).Find(&computes).Error; err != nil {
		return nil, err
	}

	return &resource_cluster_api_grpc_pb.GetComputeReply{
		Computes: modelApi.convertComputes(computes),
	}, nil
}

func (modelApi *ResourceClusterModelApi) CreateCompute(req *resource_cluster_api_grpc_pb.CreateComputeRequest) (*resource_cluster_api_grpc_pb.CreateComputeReply, error) {
	rep := &resource_cluster_api_grpc_pb.CreateComputeReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return rep, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.ComputeSpec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		return rep, err
	}
	glog.Info(spec.Name)

	// TODO Validate projectRole
	// TODO Validate cluster
	// TODO Validate spec

	var compute resource_cluster_model.Compute
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&compute).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		compute = resource_cluster_model.Compute{
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_model.StatusCreating,
			StatusReason: req.StatusReason,
		}
		if err = db.Create(&compute).Error; err != nil {
			return rep, err
		}
	} else {
		return rep, fmt.Errorf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Kind, spec.Name)
	}

	computePb, err := modelApi.convertCompute(&compute)
	if err != nil {
		return rep, err
	}
	rep.Compute = computePb
	glog.Info("Completed CreateCompute")
	return rep, err
}

func (modelApi *ResourceClusterModelApi) UpdateCompute(req *resource_cluster_api_grpc_pb.UpdateComputeRequest) (*resource_cluster_api_grpc_pb.UpdateComputeReply, error) {
	rep := &resource_cluster_api_grpc_pb.UpdateComputeReply{}
	var err error

	glog.Info("Completed UpdateCompute")
	return rep, err
}

func (modelApi *ResourceClusterModelApi) DeleteCompute(req *resource_cluster_api_grpc_pb.DeleteComputeRequest) (*resource_cluster_api_grpc_pb.DeleteComputeReply, error) {
	return nil, nil
}

func (modelApi *ResourceClusterModelApi) convertComputes(computes []resource_cluster_model.Compute) []*resource_cluster_api_grpc_pb.Compute {
	pbComputes := make([]*resource_cluster_api_grpc_pb.Compute, len(computes))
	for i, compute := range computes {
		updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbComputes[i] = &resource_cluster_api_grpc_pb.Compute{
			Name:         compute.Name,
			FullName:     compute.FullName,
			Kind:         compute.Kind,
			Labels:       compute.Labels,
			Status:       compute.Status,
			StatusReason: compute.StatusReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		}
	}

	return pbComputes
}

func (modelApi *ResourceClusterModelApi) convertCompute(compute *resource_cluster_model.Compute) (*resource_cluster_api_grpc_pb.Compute, error) {
	updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	computePb := &resource_cluster_api_grpc_pb.Compute{
		Name:         compute.Name,
		FullName:     compute.FullName,
		Kind:         compute.Kind,
		Labels:       compute.Labels,
		Status:       compute.Status,
		StatusReason: compute.StatusReason,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}

	return computePb, nil
}

func (modelApi *ResourceClusterModelApi) SyncCompute() error {
	glog.Info("Starting SyncCompute")

	var err error
	db, err := gorm.Open("mysql", modelApi.cluster.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var computes []resource_cluster_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	glog.Info(computes)

	glog.Info("Complete SyncCompute")
	return nil
}
