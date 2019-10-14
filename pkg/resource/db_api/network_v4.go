package db_api

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils/ip_utils_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (api *Api) GetNetworkV4(tctx *logger.TraceContext, input *spec.GetNetworkV4, user *base_spec.UserAuthority) (data *spec.NetworkV4, err error) {
	data = &spec.NetworkV4{}
	err = api.DB.Where("name = ? AND cluster = ? AND deleted_at IS NULL", input.Name, input.Cluster).
		First(data).Error
	return
}

func (api *Api) GetNetworkV4s(tctx *logger.TraceContext, input *spec.GetNetworkV4s, user *base_spec.UserAuthority) (data []spec.NetworkV4, err error) {
	err = api.DB.Where("cluster = ? AND deleted_at IS NULL", input.Cluster).Find(&data).Error
	return
}

func (api *Api) CreateNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			var tmp db_model.NetworkV4
			if err = tx.Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				First(&tmp).Error; err != nil {
				if !gorm.IsRecordNotFoundError(err) {
					return
				}
				tmp = db_model.NetworkV4{
					Name:         val.Name,
					Cluster:      val.Cluster,
					Kind:         val.Kind,
					Description:  val.Description,
					Status:       base_const.StatusActive,
					StatusReason: "CreateNetworkV4",
					Subnet:       val.Subnet,
					StartIp:      val.StartIp,
					EndIp:        val.EndIp,
					Gateway:      val.Gateway,
				}
				if err = tx.Create(&tmp).Error; err != nil {
					return
				}
			}
		}
		return
	})
	return
}

func (api *Api) UpdateNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Model(&db_model.NetworkV4{}).
				Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				Updates(&db_model.NetworkV4{
					Kind:         val.Kind,
					Description:  val.Description,
					Status:       base_const.StatusActive,
					StatusReason: "UpdateNetworkV4",
					Subnet:       val.Subnet,
					StartIp:      val.StartIp,
					EndIp:        val.EndIp,
					Gateway:      val.Gateway,
				}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) DeleteNetworkV4(tctx *logger.TraceContext, input *spec.DeleteNetworkV4, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		err = tx.Where("name = ? AND region = ?", input.Name, input.Region).Delete(&db_model.NetworkV4{}).Error
		return
	})
	return
}

func (api *Api) DeleteNetworkV4s(tctx *logger.TraceContext, input []spec.NetworkV4, user *base_spec.UserAuthority) (err error) {
	err = api.Transact(tctx, func(tx *gorm.DB) (err error) {
		for _, val := range input {
			if err = tx.Where("name = ? AND cluster = ?", val.Name, val.Cluster).
				Delete(&db_model.NetworkV4{}).Error; err != nil {
				return
			}
		}
		return
	})
	return
}

func (api *Api) AssignNetworkV4Port(tctx *logger.TraceContext, tx *gorm.DB,
	npspec *spec.NetworkPolicySpec, networks []db_model.NetworkV4,
	kind string, name string) (sports []spec.PortSpec, err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	netIds := []uint{}
	netPortMap := map[uint]map[string]bool{}
	netMacMap := map[uint]map[string]bool{}
	for _, network := range networks {
		netIds = append(netIds, network.ID)
		netPortMap[network.ID] = map[string]bool{}
	}

	var ports []db_model.NetworkV4Port
	if err = tx.Table("network_v4_ports").
		Select("network_v4.id, network_v4_ports.ip").
		Joins("INNER JOIN network_v4 ON network_v4.id = network_v4_ports.network_v4_id").
		Where("network_v4.id IN (?)", netIds).Find(&ports).Error; err != nil {
		return
	}

	for _, port := range ports {
		netPortMap[port.ID][port.Ip] = true
		netMacMap[port.ID][port.Mac] = true
	}

	nets := []resource_model.Network{}
	for _, net := range networks {
		var network *ip_utils_model.Network
		if network, err = ip_utils.ParseNetwork(net.Subnet, net.Gateway, net.StartIp, net.EndIp); err != nil {
			return
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
			Subnet:       net.Subnet,
			Gateway:      net.Gateway,
			AvailableIps: availableIps,
		})
	}

	for i := 0; i < npspec.Interfaces; i++ {
		switch npspec.AssignPolicy {
		case resource_model.SchedulePolicyAffinity:
			net := nets[0]
			ip := net.AvailableIps[i]
			var mac string
			macMap, _ := netMacMap[net.Id]
			mac, err = ip_utils.GenerateUniqueRandomMac(macMap, 100)
			if err != nil {
				return
			}

			port := resource_model.NetworkV4Port{
				NetworkV4ID:  net.Id,
				Ip:           ip,
				Mac:          mac,
				ResourceKind: kind,
				ResourceName: name,
			}
			if err = tx.Create(&port).Error; err != nil {
				return
			}
			sports = append(sports, spec.PortSpec{
				NetworkID: net.Id,
				Version:   4,
				Subnet:    net.Subnet,
				Gateway:   net.Gateway,
				Ip:        ip,
				Mac:       mac,
			})
		}
	}

	return
}
