package compute_models

import "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"

type NetnsPort struct {
	Id                 int
	Name               string
	NetnsGateway       string
	NetnsIp            string
	VmIp               string
	VmMac              string
	Kind               string
	NetworkV4LocalSpec spec.NetworkV4LocalSpec
}
