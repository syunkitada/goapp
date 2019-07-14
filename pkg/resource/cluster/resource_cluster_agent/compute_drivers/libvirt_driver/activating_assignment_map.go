package libvirt_driver

import (
	"fmt"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (driver *LibvirtDriver) syncActivatingAssignmentMap(tctx *logger.TraceContext,
	assignmentMap map[uint]resource_model.ComputeAssignmentEx) error {
	fmt.Println("DEBUG driver", assignmentMap)
	out, err := exec_utils.Cmd(tctx, "brctl show")
	if err != nil {
		return err
	}
	bridgeInterfaceMap := map[string]map[string]bool{}
	for i, line := range strings.Split(out, "\n") {
		if i == 0 {
			continue
		}
		columns := strings.Split(line, "\t")
		if len(columns) == 0 {
			break
		}
		bridgeInterfaceMap[columns[0]] = map[string]bool{}
	}

	for _, assignment := range assignmentMap {
		fmt.Println("DEBUG assignment", assignment)
		for _, port := range assignment.ComputeSpec.Ports {
			bridgeName := fmt.Sprintf("br-compute%d", port.NetworkID)
			interfaceMap, ok := bridgeInterfaceMap[bridgeName]
			if !ok {
				if _, err = exec_utils.Cmdf(tctx, "brctl addbr %s", bridgeName); err != nil {
					return err
				}
			}

			if out, err = exec_utils.Cmdf(tctx, "ip addr show dev %s", bridgeName); err != nil {
				return err
			}
			gatewayAddress := fmt.Sprintf("%s/%s", port.Gateway, strings.Split(port.Subnet, "/")[1])
			if strings.Index(out, gatewayAddress) == -1 {
				if _, err = exec_utils.Cmdf(tctx, "ip addr add %s dev %s", gatewayAddress, bridgeName); err != nil {
					return err
				}
			}
			if strings.Index(out, "DOWN") > -1 {
				if _, err = exec_utils.Cmdf(tctx, "ip link set %s up", bridgeName); err != nil {
					return err
				}
			}
			fmt.Println("DEBUG interfaceMap", interfaceMap)
		}
	}

	// out, err := exec_utils.Cmdf(1, "brctl addbr %s", "test")
	return nil
}
