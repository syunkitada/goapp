package resource_model_api

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils/ip_utils_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetNetworkV4(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	resource, ok := query.StrParams["resource"]
	if !ok {
		return codes.ClientBadRequest, fmt.Errorf("resource is None")
	}

	var network resource_model.NetworkV4
	if err = db.Where(&resource_model.NetworkV4{
		Name: resource,
	}).First(&network).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["NetworkV4"] = network
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) GetNetworkV4s(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query, data map[string]interface{}) (int64, error) {
	var err error
	var networks []resource_model.NetworkV4
	if err = db.Find(&networks).Error; err != nil {
		return codes.RemoteDbError, err
	}
	data["NetworkV4s"] = networks
	return codes.OkRead, nil
}

func (modelApi *ResourceModelApi) CreateNetworkV4(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok {
		err = error_utils.NewInvalidRequestError("NotFound Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NetworkV4Spec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestError("Specs is empty")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		var data resource_model.NetworkV4
		if err = tx.Where("name = ?", spec.Name).First(&data).Error; err != nil {
			if !gorm.IsRecordNotFoundError(err) {
				return codes.RemoteDbError, err
			}

			data = resource_model.NetworkV4{
				Kind:         spec.Kind,
				Name:         spec.Name,
				Description:  spec.Description,
				Cluster:      spec.Cluster,
				Status:       resource_model.StatusActive,
				StatusReason: "CreateNetworkV4",
				Subnet:       spec.Subnet,
				StartIp:      spec.StartIp,
				EndIp:        spec.EndIp,
				Gateway:      spec.Gateway,
			}
			if err = tx.Create(&data).Error; err != nil {
				return codes.RemoteDbError, err
			}
		} else {
			err = error_utils.NewConflictAlreadyExistsError(spec.Name)
			return codes.ClientAlreadyExists, err
		}
	}

	tx.Commit()
	return codes.OkCreated, nil
}

func (modelApi *ResourceModelApi) UpdateNetworkV4(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NetworkV4Spec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	if len(specs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}
		network := &resource_model.NetworkV4{
			Kind:        spec.Kind,
			Description: spec.Description,
			StartIp:     spec.StartIp,
			EndIp:       spec.EndIp,
			Gateway:     spec.Gateway,
		}
		if err = tx.Model(network).Where("name = ?", spec.Name).Updates(network).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkUpdated, nil
}

func (modelApi *ResourceModelApi) DeleteNetworkV4(tctx *logger.TraceContext, db *gorm.DB,
	query *authproxy_grpc_pb.Query) (int64, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	tx := db.Begin()
	defer tx.Rollback()

	strSpecs, ok := query.StrParams["Specs"]
	if !ok || len(strSpecs) == 0 {
		err = error_utils.NewInvalidRequestEmptyError("Specs")
		return codes.ClientBadRequest, err
	}

	var specs []resource_model.NameSpec
	if err = json.Unmarshal([]byte(strSpecs), &specs); err != nil {
		return codes.ClientBadRequest, err
	}

	for _, spec := range specs {
		if err = modelApi.validate.Struct(&spec); err != nil {
			return codes.ClientBadRequest, err
		}

		if err = tx.Delete(&resource_model.NetworkV4{}, "name = ?", spec.Name).Error; err != nil {
			return codes.RemoteDbError, err
		}
	}

	tx.Commit()
	return codes.OkDeleted, nil
}

func (modelApi *ResourceModelApi) AssignNetworkV4Port(tctx *logger.TraceContext, tx *gorm.DB,
	spec *resource_model.NetworkSpec, networks []resource_model.NetworkV4) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	fmt.Println("DEBUG ASSIGN")
	netIds := []uint{}
	netPortMap := map[uint]map[string]bool{}
	netMacMap := map[uint]map[string]bool{}
	for _, network := range networks {
		netIds = append(netIds, network.ID)
		netPortMap[network.ID] = map[string]bool{}
	}

	var ports []resource_model.NetworkV4Port
	if err = tx.Raw("SELECT net.id, ports.ip FROM network_v4_ports as ports "+
		"INNER JOIN network_v4 as net ON net.id = ports.network_v4_id "+
		"WHERE net.id IN (?)", netIds).Scan(&ports).Error; err != nil {
		return err
	}

	for _, port := range ports {
		netPortMap[port.ID][port.Ip] = true
		netMacMap[port.ID][port.Mac] = true
	}

	nets := []resource_model.Network{}
	for _, net := range networks {
		var network *ip_utils_model.Network
		if network, err = ip_utils.ParseNetwork(net.Subnet, net.Gateway, net.StartIp, net.EndIp); err != nil {
			return err
		}

		portMap, _ := netPortMap[net.ID]
		availableIps := []string{}
		for {
			ipStr := network.StartIp.String()
			if _, ok := portMap[ipStr]; !ok {
				availableIps = append(availableIps, ipStr)
			}
			ip_utils.IncrementIp(network.StartIp)
			if ip_utils.CompareIp(network.StartIp, network.EndIp) == 0 {
				break
			}
		}
		nets = append(nets, resource_model.Network{
			Id:           net.ID,
			Name:         net.Name,
			AvailableIps: availableIps,
		})
	}

	// TODO AssignPort
	// TODO Support MultiPort
	// TODO Generate MAC
	net := nets[0]
	port := resource_model.NetworkV4Port{
		NetworkV4ID: net.Id,
		Ip:          net.AvailableIps[0],
	}
	if err = tx.Create(&port).Error; err != nil {
		return err
	}

	return nil
}

func (modelApi *ResourceModelApi) RegisterRecord(tctx *logger.TraceContext, db *gorm.DB, compute *resource_model.Compute) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	// TODO
	// register a record
	// implrment dns service

	return nil
}
