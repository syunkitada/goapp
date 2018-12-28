package resource_model_api

import (
	"encoding/json"
	"fmt"
	"net"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNetworkV4(tctx *logger.TraceContext, req *resource_api_grpc_pb.GetNetworkV4Request) *resource_api_grpc_pb.GetNetworkV4Reply {
	rep := &resource_api_grpc_pb.GetNetworkV4Reply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	var networks []resource_model.NetworkV4
	if err = db.Where("name like ?", req.Target).Find(&networks).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	rep.Networks = modelApi.convertNetworkV4s(tctx, networks)
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) CreateNetworkV4(req *resource_api_grpc_pb.CreateNetworkV4Request) *resource_api_grpc_pb.CreateNetworkV4Reply {
	rep := &resource_api_grpc_pb.CreateNetworkV4Reply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateNetworkV4Spec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	var network resource_model.NetworkV4
	tx := db.Begin()
	defer tx.Rollback()
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&network).Error; err != nil {
		if !gorm.IsRecordNotFoundError(err) {
			rep.Err = err.Error()
			rep.StatusCode = codes.RemoteDbError
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
		if err = tx.Create(&network).Error; err != nil {
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

	networkPb, err := modelApi.convertNetworkV4(&network)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Network = networkPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) UpdateNetworkV4(req *resource_api_grpc_pb.UpdateNetworkV4Request) *resource_api_grpc_pb.UpdateNetworkV4Reply {
	rep := &resource_api_grpc_pb.UpdateNetworkV4Reply{}

	db, err := gorm.Open("mysql", modelApi.conf.Resource.Database.Connection)
	defer db.Close()
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	db.LogMode(modelApi.conf.Default.EnableDatabaseLog)

	spec, statusCode, err := modelApi.validateNetworkV4Spec(db, req.Spec)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = statusCode
		return rep
	}

	tx := db.Begin()
	defer tx.Rollback()
	var network resource_model.NetworkV4
	if err = tx.Where("name = ? and cluster = ?", spec.Name, spec.Cluster).First(&network).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}

	network.Spec = req.Spec
	network.Status = resource_model.StatusActive
	network.StatusReason = fmt.Sprintf("UpdateNetworkV4: user=%v, project=%v", req.UserName, req.ProjectName)
	network.StartIp = spec.Spec.StartIp
	network.EndIp = spec.Spec.EndIp
	network.Gateway = spec.Spec.Gateway
	tx.Save(network)
	tx.Commit()

	networkPb, err := modelApi.convertNetworkV4(&network)
	if err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.ServerInternalError
		return rep
	}

	rep.Network = networkPb
	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) DeleteNetworkV4(req *resource_api_grpc_pb.DeleteNetworkV4Request) *resource_api_grpc_pb.DeleteNetworkV4Reply {
	rep := &resource_api_grpc_pb.DeleteNetworkV4Reply{}

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
	var network resource_model.NetworkV4
	if err = tx.Where("name = ?", req.Target).Delete(&network).Error; err != nil {
		rep.Err = err.Error()
		rep.StatusCode = codes.RemoteDbError
		return rep
	}
	tx.Commit()

	rep.StatusCode = codes.Ok
	return rep
}

func (modelApi *ResourceModelApi) convertNetworkV4s(tctx *logger.TraceContext, networks []resource_model.NetworkV4) []*resource_api_grpc_pb.NetworkV4 {
	pbNetworkV4s := make([]*resource_api_grpc_pb.NetworkV4, len(networks))
	for i, network := range networks {
		updatedAt, err := ptypes.TimestampProto(network.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err, "Failed ptypes.TimestampProto: %v", network.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(network.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err, "Failed ptypes.TimestampProto: %v", network.Model.CreatedAt)
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

func compareIp(ip1 net.IP, ip2 net.IP) int {
	if ok := ip1.Equal(ip2); ok {
		return 0
	}
	for i, ip := range ip1 {
		if ip > ip2[i] {
			return 1
		}
	}
	return -1
}

func (modelApi *ResourceModelApi) validateNetworkV4Spec(db *gorm.DB, specStr string) (resource_model.NetworkV4Spec, int64, error) {
	var spec resource_model.NetworkV4Spec
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
	case resource_model.SpecKindNetworkV4Local:
		startIp := net.ParseIP(spec.Spec.StartIp)
		if startIp == nil {
			errors = append(errors, "Invalid StartIp")
		}

		endIp := net.ParseIP(spec.Spec.EndIp)
		if endIp == nil {
			errors = append(errors, "Invalid StartIp")
		}

		gateway := net.ParseIP(spec.Spec.Gateway)
		if gateway == nil {
			errors = append(errors, "Invalid Gateway")
		}

		_, subnet, subnetErr := net.ParseCIDR(spec.Spec.Subnet)
		if subnetErr != nil {
			errors = append(errors, "Invalid Subnet")
		}

		if startIp != nil && endIp != nil && gateway != nil && subnetErr == nil {
			if !subnet.Contains(startIp) {
				errors = append(errors, "StartIp is not contained in Subnet")
			}
			if !subnet.Contains(endIp) {
				errors = append(errors, "StartIp is not contained in Subnet")
			}
			if !subnet.Contains(gateway) {
				errors = append(errors, "Gateway is not contained in Subnet")
			}
			if comp := compareIp(startIp, endIp); comp != -1 {
				errors = append(errors, "StartIp must be less than EndIp")
			}
		}

	default:
		errors = append(errors, fmt.Sprintf("Invalid kind: %v", spec.Spec.Kind))
	}

	if len(errors) > 0 {
		return spec, codes.ClientInvalidRequest, fmt.Errorf(strings.Join(errors, "\n"))
	}

	return spec, codes.Ok, nil
}
