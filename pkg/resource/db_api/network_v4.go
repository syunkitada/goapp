package db_api

import (
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils/ip_utils_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/consts"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
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
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
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
					Spec:         string(specBytes),
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
			var specBytes []byte
			if specBytes, err = json_utils.Marshal(val.Spec); err != nil {
				return
			}
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
					Spec:         string(specBytes),
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
		err = tx.Where("name = ? AND cluster = ?", input.Name, input.Cluster).Delete(&db_model.NetworkV4{}).Error
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

	isStaticNetworks := len(npspec.StaticNetworks) > 0
	netIds := []uint{}
	netPortMap := map[uint]map[string]bool{}
	netMacMap := map[uint]map[string]bool{}
	staticNetworkCountMap := map[string]int{}
	for _, network := range networks {
		// StaticNetworksが定義されてる場合は、networkの候補をStaticNetworksのみに絞る
		if isStaticNetworks {
			isStaticNetwork := false
			for _, snetwork := range npspec.StaticNetworks {
				if snetwork == network.Name {
					isStaticNetwork = true
					staticNetworkCountMap[snetwork] = 0
					break
				}
			}
			if !isStaticNetwork {
				continue
			}
		}
		netIds = append(netIds, network.ID)
		netPortMap[network.ID] = map[string]bool{}
		netMacMap[network.ID] = map[string]bool{}
	}

	// 候補となるnetworkのportをすべて取得し、利用可能なportを洗い出す
	var ports []db_model.NetworkV4Port
	if err = tx.Table("network_v4_ports").
		Select("network_v4_ports.*").
		Joins("INNER JOIN network_v4 ON network_v4.id = network_v4_ports.network_v4_id").
		Where("network_v4.id IN (?)", netIds).Find(&ports).Error; err != nil {
		return
	}

	for _, port := range ports {
		netPortMap[port.NetworkV4ID][port.Ip] = true
		netMacMap[port.NetworkV4ID][port.Mac] = true
	}

	nets := []spec.Network{}
	for _, net := range networks {
		var network *ip_utils_model.Network
		if network, err = ip_utils.ParseNetwork(net.Subnet, net.Gateway, net.StartIp, net.EndIp); err != nil {
			return
		}

		portMap, ok := netPortMap[net.ID]
		if !ok {
			continue
		}
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
		nets = append(nets, spec.Network{
			Id:           net.ID,
			Name:         net.Name,
			Subnet:       net.Subnet,
			Gateway:      net.Gateway,
			Kind:         net.Kind,
			Spec:         net.Spec,
			AvailableIps: availableIps,
		})
	}

	// すでにAssignされてるPortを取得し、新規にAssignする必要がある場合はAssignする
	var currentPorts []db_model.NetworkV4Port
	if err = tx.Where("resource_kind = ? AND resource_name = ?", kind, name).
		Find(&currentPorts).Error; err != nil {
		return
	}
	for _, port := range currentPorts {
		for _, net := range nets {
			if net.Id == port.NetworkV4ID {
				if count, ok := staticNetworkCountMap[net.Name]; ok {
					staticNetworkCountMap[net.Name] = count + 1
				}
				sports = append(sports, spec.PortSpec{
					NetworkID: net.Id,
					Version:   4,
					Subnet:    net.Subnet,
					Gateway:   net.Gateway,
					Ip:        port.Ip,
					Mac:       port.Mac,
				})
			} else {
				break
			}
		}
	}
	interfaces := npspec.Interfaces - len(currentPorts)
	if interfaces == 0 {
		return
	}

	if !isStaticNetworks {
		sort.Slice(nets, func(i, j int) bool {
			return len(nets[i].AvailableIps) < len(nets[j].AvailableIps)
		})
	} else {
		// StaticNetworksが定義されてる場合は、配列の定義順でNetworkをソートする
		sort.SliceStable(nets, func(i, j int) bool {
			// すでにAssign済みの場合は、Assign数が多いNetworkの優先度を下げる
			icount, iok := staticNetworkCountMap[nets[i].Name]
			jcount, jok := staticNetworkCountMap[nets[j].Name]
			if iok {
				if jok {
					return icount > jcount
				} else {
					return true
				}
			}

			ix := len(npspec.StaticNetworks)
			jx := ix
			for x, snet := range npspec.StaticNetworks {
				if snet == nets[i].Name {
					ix = x
				}
				if snet == nets[j].Name {
					jx = x
				}
			}
			return ix < jx
		})
	}

	netid := 0
	for i := 0; i < interfaces; i++ {
		var net spec.Network
		switch npspec.AssignPolicy {
		case consts.SchedulePolicyAffinity:
			if len(currentPorts) > 0 { // CurrentPortsが存在する場合は、Networkを既存に合わせる
				for _, n := range nets {
					if n.Id == currentPorts[0].NetworkV4ID {
						net = n
						break
					}
				}
			} else { // CurrentPortsが存在しない場合は、優先度の高いNetworkに固定する
				net = nets[0]
			}
		case consts.SchedulePolicyAntiAffinity:
			// 優先度の高い順にラウンドロビンで割り当てる
			net = nets[netid]
			if netid < len(nets)-1 {
				netid += 1
			} else {
				netid = 0
			}
		default:
			err = fmt.Errorf("Invalid AssignPolicy: %s", npspec.AssignPolicy)
			return
		}

		ip := net.AvailableIps[i]
		var mac string
		macMap, ok := netMacMap[net.Id]
		if !ok {
			err = fmt.Errorf("NotFound Mac: net.Id=%d", net.Id)
			return
		}
		mac, err = ip_utils.GenerateUniqueRandomMac(macMap, 100)
		if err != nil {
			return
		}

		port := db_model.NetworkV4Port{
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
			Kind:      net.Kind,
			Spec:      net.Spec,
		})
	}

	return
}
