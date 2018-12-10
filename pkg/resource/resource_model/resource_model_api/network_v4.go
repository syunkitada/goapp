package resource_model_api

import (
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNetworkV4(req *resource_api_grpc_pb.GetNetworkV4Request) (*resource_api_grpc_pb.GetNetworkV4Reply, error) {
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var networks []resource_model.NetworkV4
	if err = db.Where("name like ?", req.Target).Find(&networks).Error; err != nil {
		return nil, err
	}

	return &resource_api_grpc_pb.GetNetworkV4Reply{
		Networks: modelApi.convertNetworkV4s(networks),
	}, nil
}

func (modelApi *ResourceModelApi) CreateNetworkV4(req *resource_api_grpc_pb.CreateNetworkV4Request) *resource_api_grpc_pb.CreateNetworkV4Reply {
	rep := &resource_api_grpc_pb.CreateNetworkV4Reply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)
	glog.Info(req.Spec)

	var spec resource_model.NetworkV4Spec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ClientBadRequest
		return rep
	}
	if err = modelApi.validate.Struct(spec); err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ClientInvalidRequest
		return rep
	}

	// TODO Validate projectRole
	// TODO Validate cluster
	// TODO Validate spec
	// TODO Validate image

	var network resource_model.NetworkV4
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&network).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			return rep
		}

		network = resource_model.NetworkV4{
			Cluster:      spec.Cluster,
			Kind:         spec.Kind,
			Name:         spec.Name,
			Spec:         req.Spec,
			Status:       resource_model.StatusActive,
			StatusReason: fmt.Sprintf("CreateNetworkV4: user=%v, project=%v", req.UserName, req.ProjectName),
			Subnet:       spec.Spec.Subnet,
			StartIp:      spec.Spec.StartIp,
			EndIp:        spec.Spec.EndIp,
			Gateway:      spec.Spec.Gateway,
		}
		if err = db.Create(&network).Error; err != nil {
			rep.Err = err.Error()
			return rep
		}
	} else {
		rep.Err = fmt.Sprintf("Already Exists: cluster=%v, name=%v",
			spec.Cluster, spec.Name)
		return rep
	}

	networkPb, err := modelApi.convertNetworkV4(&network)
	if err != nil {
		rep.Err = err.Error()
		return rep
	}
	rep.Network = networkPb
	glog.Info("Completed CreateNetworkV4")
	return rep
}

func (modelApi *ResourceModelApi) UpdateNetworkV4(req *resource_api_grpc_pb.UpdateNetworkV4Request) *resource_api_grpc_pb.UpdateNetworkV4Reply {
	rep := &resource_api_grpc_pb.UpdateNetworkV4Reply{}
	var err error

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var spec resource_model.NetworkV4Spec
	if err = json.Unmarshal([]byte(req.Spec), &spec); err != nil {
		rep.Err = err.Error()
		return rep
	}

	var network resource_model.NetworkV4
	if err = db.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&network).Error; err != nil {
		rep.Err = err.Error()
		return rep
	}

	network.Spec = req.Spec
	network.Status = resource_model.StatusActive
	network.StatusReason = fmt.Sprintf("UpdateNetworkV4: user=%v, project=%v", req.UserName, req.ProjectName)
	network.Subnet = spec.Spec.Subnet
	network.StartIp = spec.Spec.StartIp
	network.EndIp = spec.Spec.EndIp
	network.Gateway = spec.Spec.Gateway
	db.Save(network)

	networkPb, err := modelApi.convertNetworkV4(&network)
	if err != nil {
		rep.Err = err.Error()
		return rep
	}
	rep.Network = networkPb
	glog.Info("Completed UpdateNetworkV4")
	return rep
}

func (modelApi *ResourceModelApi) DeleteNetworkV4(req *resource_api_grpc_pb.DeleteNetworkV4Request) (*resource_api_grpc_pb.DeleteNetworkV4Reply, error) {
	rep := &resource_api_grpc_pb.DeleteNetworkV4Reply{}
	var err error
	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		return nil, err
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	glog.Info("DEBUGlalalalala")
	glog.Info(req.Target)
	glog.Info(req.Cluster)

	var network resource_model.NetworkV4
	if err = db.Where("name = ?", req.Target).Delete(&network).Error; err != nil {
		return nil, err
	}

	networkPb, err := modelApi.convertNetworkV4(&network)
	if err != nil {
		return rep, err
	}
	rep.Network = networkPb
	glog.Info("Completed CreateNetworkV4")
	return rep, err
}

func (modelApi *ResourceModelApi) convertNetworkV4s(networks []resource_model.NetworkV4) []*resource_api_grpc_pb.NetworkV4 {
	pbNetworkV4s := make([]*resource_api_grpc_pb.NetworkV4, len(networks))
	for i, network := range networks {
		updatedAt, err := ptypes.TimestampProto(network.Model.UpdatedAt)
		createdAt, err := ptypes.TimestampProto(network.Model.CreatedAt)
		if err != nil {
			glog.Warningf("Invalid timestamp: %v", err)
			continue
		}

		pbNetworkV4s[i] = &resource_api_grpc_pb.NetworkV4{
			Cluster:      network.Cluster,
			Name:         network.Name,
			Kind:         network.Kind,
			Labels:       network.Labels,
			Status:       network.Status,
			StatusReason: network.StatusReason,
			UpdatedAt:    updatedAt,
			CreatedAt:    createdAt,
		}
	}

	return pbNetworkV4s
}

func (modelApi *ResourceModelApi) convertNetworkV4(network *resource_model.NetworkV4) (*resource_api_grpc_pb.NetworkV4, error) {
	updatedAt, err := ptypes.TimestampProto(network.Model.UpdatedAt)
	createdAt, err := ptypes.TimestampProto(network.Model.CreatedAt)
	if err != nil {
		return nil, err
	}

	networkPb := &resource_api_grpc_pb.NetworkV4{
		Cluster:      network.Cluster,
		Name:         network.Name,
		Kind:         network.Kind,
		Labels:       network.Labels,
		Status:       network.Status,
		StatusReason: network.StatusReason,
		UpdatedAt:    updatedAt,
		CreatedAt:    createdAt,
	}

	return networkPb, nil
}
