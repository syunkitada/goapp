package config

type ResourceConfig struct {
	AppDownTime   int
	ApiApp        AppConfig
	ControllerApp AppConfig
	Database      DatabaseConfig
	Node          ResourceNodeConfig
	ClusterMap    map[string]ResourceClusterConfig
}

type ResourceClusterConfig struct {
	ApiApp        AppConfig
	ControllerApp AppConfig
	AgentApp      AppConfig
	Database      DatabaseConfig
}

type ResourceNodeConfig struct {
	ClusterName             string
	NetworkAvailabilityZone string
	NodeAvailabilityZone    string
	Compute                 ResourceComputeConfig
}

type ResourceComputeConfig struct {
	Driver  string // libvirt
	Libvirt ResourceLibvirtConfig
}

type ResourceLibvirtConfig struct {
	Enable             bool
	CpuMode            string // host-model
	CpuType            string // kvm, qemu
	MemoryType         string // local, hugepage
	DiskType           string // local
	NetworkType        string // local-linuxbridge
	NetworkVhostQueues int
}
