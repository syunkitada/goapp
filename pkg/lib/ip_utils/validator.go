package ip_utils

import (
	"net"

	"github.com/syunkitada/goapp/pkg/lib/error_utils"
	"github.com/syunkitada/goapp/pkg/lib/ip_utils/ip_utils_model"
)

func ParseNetwork(subnet string, gateway string, startIp string, endIp string) (*ip_utils_model.Network, error) {
	_, parsedSubnet, subnetErr := net.ParseCIDR(subnet)
	if subnetErr != nil {
		return nil, error_utils.NewInvalidDataError("subnet", subnet, "FailedParse")
	}

	parsedGateway := net.ParseIP(gateway)
	if parsedGateway == nil {
		return nil, error_utils.NewInvalidDataError("gateway", gateway, "FailedParse")
	}

	parsedStartIp := net.ParseIP(startIp)
	if parsedStartIp == nil {
		return nil, error_utils.NewInvalidDataError("startIp", startIp, "FailedParse")
	}

	parsedEndIp := net.ParseIP(endIp)
	if parsedEndIp == nil {
		return nil, error_utils.NewInvalidDataError("endIp", endIp, "FailedParse")
	}

	if !parsedSubnet.Contains(parsedStartIp) {
		return nil, error_utils.NewInvalidDataError("startIp", startIp, "startIp should be countained in subnet")
	}

	if !parsedSubnet.Contains(parsedEndIp) {
		return nil, error_utils.NewInvalidDataError("endIp", endIp, "endIp should be countained in subnet")
	}

	if !parsedSubnet.Contains(parsedGateway) {
		return nil, error_utils.NewInvalidDataError("gateway", gateway, "gateway should be countained in subnet")
	}
	// if compareIp(parsedStartIp, parsedEndIp) {
	// 	return nil, error_utils.NewInvalidDataError("endIp", endIp, "endIp should be bigger than startIp")
	// }
	return &ip_utils_model.Network{
		Subnet:  parsedSubnet,
		Gateway: parsedGateway,
		StartIp: parsedStartIp,
		EndIp:   parsedEndIp,
	}, nil
}
