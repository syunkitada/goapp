package compute_models

type NetnsPort struct {
	Id             int
	Name           string
	ShareNetnsAddr string
	NetnsGateway   string
	NetnsAddr      string
	VmGateway      string
	VmIp           string
	VmAddr         string
	VmMac          string
}
