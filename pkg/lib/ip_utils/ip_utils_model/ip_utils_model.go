package ip_utils_model

import "net"

type Network struct {
	Subnet  *net.IPNet
	Gateway net.IP
	StartIp net.IP
	EndIp   net.IP
}
