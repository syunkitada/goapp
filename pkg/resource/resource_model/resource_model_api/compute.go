package resource_model_api

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api/resource_cluster_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetCompute(req *resource_api_grpc_pb.GetComputeRequest) (*resource_api_grpc_pb.GetComputeReply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var computes []resource_model.Compute
	if err = db.Where("name like ?", req.Target).Find(&computes).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetComputeReply{
		Computes: modelApi.convertComputes(computes),
	}, nil
}

func (modelApi *ResourceModelApi) CreateCompute(req *resource_api_grpc_pb.CreateComputeRequest) (*resource_api_grpc_pb.CreateComputeReply, error) {
	rep := &resource_api_grpc_pb.CreateComputeReply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
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
	// TODO Validate image

	var compute resource_model.Compute
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&compute).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			return rep, err
		}

		compute = resource_model.Compute{
			Cluster:      spec.Cluster,
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_model.StatusCreating,
			StatusReason: fmt.Sprintf("CreateCompute: user=%v, project=%v", req.UserName, req.ProjectName),
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

func (modelApi *ResourceModelApi) UpdateCompute(req *resource_api_grpc_pb.UpdateComputeRequest) (*resource_api_grpc_pb.UpdateComputeReply, error) {
	rep := &resource_api_grpc_pb.UpdateComputeReply{}
	var err error

	glog.Info("Completed UpdateCompute")
	return rep, err
}

func (modelApi *ResourceModelApi) DeleteCompute(req *resource_api_grpc_pb.DeleteComputeRequest) (*resource_api_grpc_pb.DeleteComputeReply, error) {
	return nil, nil
}

func (modelApi *ResourceModelApi) convertComputes(computes []resource_model.Compute) []*resource_api_grpc_pb.Compute {
	pbComputes := make([]*resource_api_grpc_pb.Compute, len(computes))
	for i, compute := range computes {
		updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbComputes[i] = &resource_api_grpc_pb.Compute{
			Cluster:      compute.Cluster,
			Name:         compute.Name,
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

func (modelApi *ResourceModelApi) convertCompute(compute *resource_model.Compute) (*resource_api_grpc_pb.Compute, error) {
	updatedAt, err := ptypes.TimestampProto(compute.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(compute.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	computePb := &resource_api_grpc_pb.Compute{
		Cluster:      compute.Cluster,
		Name:         compute.Name,
		Kind:         compute.Kind,
		Labels:       compute.Labels,
		Status:       compute.Status,
		StatusReason: compute.StatusReason,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}

	return computePb, nil
}

func (modelApi *ResourceModelApi) SyncCompute() error {
	glog.Info("Starting SyncCompute")

	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var computes []resource_model.Compute
	if err = db.Find(&computes).Error; err != nil {
		return err
	}

	glog.Info(computes)

	computeMap := map[string]resource_cluster_api_grpc_pb.Compute{}
	req := resource_cluster_api_grpc_pb.GetComputeRequest{Target: "%"}
	for clusterName, clusterClient := range modelApi.clusterClientMap {
		result, err := clusterClient.GetCompute(&req)
		if err != nil {
			glog.Errorf("Failed GetCompute from %v: %v", clusterName, err)
		}
		for _, compute := range result.Computes {
			computeMap[compute.FullName] = *compute
		}
	}

	glog.Info(computeMap)

	for _, compute := range computes {
		switch compute.Status {
		case resource_model.StatusCreating:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
			modelApi.InitializeCompute(db, &compute, computeMap)
		case resource_model.StatusCreatingInitialized:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusCreatingScheduled:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdating:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusUpdatingScheduled:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeleting:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		case resource_model.StatusDeletingScheduled:
			glog.Infof("Found %v resource: %v", compute.Status, compute.Name)
		}
	}

	glog.Info("Complete SyncCompute")
	return nil
}

func (modelApi *ResourceModelApi) InitializeCompute(db *gorm.DB, compute *resource_model.Compute, computeMap map[string]resource_cluster_api_grpc_pb.Compute) error {
	// TODO
	// Assgin IP address
	return nil
}
