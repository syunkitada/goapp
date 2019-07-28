package compute_utils

import (
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

func InitNetns(tctx *logger.TraceContext, netns string, port compute_models.NetnsPort) error {
	bridgeName := netns + "-br"
	externalLink := netns + "-ex"
	internalLink := netns + "-in"
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

	return nil
}
