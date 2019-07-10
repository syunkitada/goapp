package config

type ResourceConfig struct {
	AppDownTime   int
	ApiApp        AppConfig
	ControllerApp ControllerAppConfig
	Database      DatabaseConfig
	ClusterMap    map[string]ResourceClusterConfig
	Node          ResourceNodeConfig
}

type ControllerAppConfig struct {
	AppConfig
	SyncResourceTimeout int
}

type ResourceClusterConfig struct {
	ApiApp        AppConfig
	ControllerApp ControllerAppConfig
	AgentApp      AppConfig
	Database      DatabaseConfig
}

type ResourceNodeConfig struct {
	ClusterName string
	Metrics     ResourceMetricsConfig
	Compute     ResourceComputeConfig
}

type ResourceComputeConfig struct {
	Enable  bool
	Driver  string // libvirt
	Libvirt ResourceLibvirtConfig
}

type ResourceLibvirtConfig struct {
	AvailableCpus      string
	CpuMode            string // host-model
	CpuType            string // kvm, qemu
	MemoryType         string // local, hugepage
	DiskType           string // local
	NetworkType        string // local-linuxbridge
	NetworkVhostQueues int
}

type ResourceMetricsConfig struct {
	System ResourceMetricsSystemConfig
}

type ResourceMetricsSystemConfig struct {
	Enable       bool
	EnableLogin  bool
	EnableCpu    bool
	EnableMemory bool
	CacheLength  int
}
