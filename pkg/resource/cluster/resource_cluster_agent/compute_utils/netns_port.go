package compute_utils

import (
	"fmt"
	"strings"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_agent/compute_models"
)

func AssignNetnsPortId(assignedNetnsPortIds []bool) int {
	size := len(assignedNetnsPortIds)
	for i := 0; i < size; i++ {
		if !assignedNetnsPortIds[i] {
			assignedNetnsPortIds[i] = true
			return i
		}
	}
	return -1
}

func InitShareNetns(tctx *logger.TraceContext, conf *config.ResourceComputeConfig) error {
	shareNetns := "com-share"
	shareBr := "com-share-br"
	if _, err := exec_utils.Shf(tctx, "ip netns | grep %s || ip netns add %s", shareNetns, shareNetns); err != nil {
		return err
	}

	if _, err := exec_utils.Shf(tctx,
		"ip netns exec %s brctl show | grep %s || ip netns exec %s brctl addbr %s",
		shareNetns, shareBr, shareNetns, shareBr); err != nil {
		return err
	}

	if _, err := exec_utils.Shf(tctx,
		"ip netns exec %s ip addr show %s | grep %s || ip netns exec %s ip addr add %s dev %s",
		shareNetns, shareBr, conf.ShareNetnsGateway, shareNetns, conf.ShareNetnsGateway, shareBr); err != nil {
		return err
	}

	splitedSubnet := strings.Split(conf.ShareNetnsSubnet, "/")
	serviceAddr := fmt.Sprintf("%s/%s", conf.ShareNetnsHttpServiceIp, splitedSubnet[1])
	if _, err := exec_utils.Shf(tctx,
		"ip netns exec %s ip addr show %s | grep %s || ip netns exec %s ip addr add %s dev %s",
		shareNetns, shareBr, serviceAddr, shareNetns, serviceAddr, shareBr); err != nil {
		return err
	}

	if _, err := exec_utils.Shf(tctx,
		"ip netns exec %s ip link show %s | egrep UP|UNKNOWN || ip netns exec %s ip link set %s up",
		shareNetns, shareBr, shareNetns, shareBr); err != nil {
		return err
	}

	return nil
}

func InitNetns(tctx *logger.TraceContext, conf *config.ResourceComputeConfig, netns string, port compute_models.NetnsPort) error {
	bridgeName := netns + "-br"
	externalLink := netns + "-ex"
	internalLink := netns + "-in"
	shareExLink := netns + "-shex"
	shareInLink := netns + "-shin"

	if _, err := exec_utils.Cmdf(tctx, "brctl addbr %s", bridgeName); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip addr add dev %s %s", bridgeName, port.NetnsGateway); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link set %s up", bridgeName); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip route add %s via %s", port.VmIp, port.NetnsGateway); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link add %s type veth peer name %s", externalLink, internalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "brctl addif %s %s", bridgeName, externalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link set %s up", externalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link set %s netns %s", internalLink, netns); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip addr add dev %s %s", netns, internalLink, port.NetnsAddr); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip link set %s up", netns, internalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s brctl addbr vm-br", netns); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip addr add dev vm-br %s", netns, port.VmGateway); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip link set vm-br up", netns); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip link set %s up", netns, internalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s sysctl net.ipv4.ip_forward=1", netns); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s sysctl net.ipv4.conf.%s.proxy_arp=1", netns, internalLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip route add default via %s", netns, port.NetnsGateway); err != nil {
		return err
	}

	// Link share netns
	if _, err := exec_utils.Cmdf(tctx, "ip link add %s type veth peer name %s", shareExLink, shareInLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link set %s netns %s", shareInLink, netns); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip addr add dev %s %s",
		netns, shareInLink, port.ShareNetnsAddr); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip link set %s up", netns, shareInLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip link set %s netns %s", shareExLink, conf.ShareNetnsName); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s brctl addif %s %s",
		conf.ShareNetnsName, conf.ShareNetnsBridgeName, shareExLink); err != nil {
		return err
	}

	if _, err := exec_utils.Cmdf(tctx, "ip netns exec %s ip link set %s up", conf.ShareNetnsName, shareExLink); err != nil {
		return err
	}

	return nil
}
