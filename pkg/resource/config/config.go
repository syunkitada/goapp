package config

import (
	"net"
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_config"
)

type Config struct {
	Resource ResourceConfig
}

type ResourceConfig struct {
	Api         base_config.AppConfig
	Controller  ResourceControllerConfig
	ClusterName string
	ClusterMap  map[string]ResourceClusterConfig
}

type ResourceControllerConfig struct {
	base_config.AppConfig
	SyncRegionServiceTimeout int
}

type ResourceClusterConfig struct {
	Region             string
	Datacenter         string
	Kind               string
	Weight             int
	DomainSuffix       string
	TimeSeriesDatabase TimeSeriesDatabaseConfig
	Api                base_config.AppConfig
	Controller         base_config.AppConfig
	Agent              ResourceClusterAgentConfig
}

type ResourceClusterApiConfig struct {
	base_config.AppConfig
	RegionName string
}

type ResourceClusterAgentConfig struct {
	base_config.AppConfig
	Compute       ResourceComputeConfig
	ReportProject string
	LogMap        map[string]ResourceLogConfig
	Metrics       ResourceMetricsConfig
}

// ResourceComputeConfig is config for compute
type ResourceComputeConfig struct {
	Enable               bool
	ConfirmRetryCount    int
	ConfirmRetryInterval int

	// ConfigDir is directory for Compute config
	ConfigDir string

	// VarDir is directory for Compute data
	VarDir string
	// VmsDir is directory for VM data
	// Default is $VarDir/vms
	VmsDir string
	// Default is $VarDir/images
	ImagesDir string

	// VmNetnsGateway is Gateway in netns on host
	VmNetnsGatewayStartIp string
	VmNetnsGatewayEndIp   string

	// VmNetnsServiceIp is ServiceIp for VM in netns on host
	// NetnsService serve convenient services for VM
	VmNetnsServiceIp string

	// VmNetnsIp is assigned vm, this ip is available in netns on host
	// VmNetnsIp range is defined VmNetnsStartIp, and VmNetnsEndIp
	VmNetnsStartIp string
	VmNetnsEndIp   string

	// Driver is provider for VM.
	// Available providers are mock, qemu
	Driver string
}

// ResourceComputeExConfig is config for ComputeDriver
// This is auto generated from ResourceComputeConfig
type ResourceComputeExConfig struct {
	ResourceComputeConfig
	ConfirmRetryInterval  time.Duration
	VmNetnsGatewayStartIp net.IP
	VmNetnsGatewayEndIp   net.IP
	VmNetnsServiceIp      net.IP
	VmNetnsStartIp        net.IP
	VmNetnsEndIp          net.IP
	VmsDir                string
	ImagesDir             string
	UserdataTmpl          string
	VmServiceTmpl         string
	VmServiceShTmpl       string
	SystemdDir            string
}

// type ResourceLibvirtConfig struct {
// 	AvailableCpus      string
// 	CpuMode            string // host-model
// 	CpuType            string // kvm, qemu
// 	MemoryType         string // local, hugepage
// 	DiskType           string // local
// 	NetworkType        string // local-linuxbridge
// 	NetworkVhostQueues int
// }

type TimeSeriesDatabaseConfig struct {
	Driver          string
	EventDatabases  []string
	LogDatabases    []string
	MetricDatabases []string
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Resource: ResourceConfig{
		Api: base_config.AppConfig{
			Name:                        "ResourceApi",
			NodeServiceDownTimeDuration: 60,
		},
		Controller: ResourceControllerConfig{
			AppConfig: base_config.AppConfig{
				Name: "ReosurceController",
			},
			SyncRegionServiceTimeout: 10,
		},
	},
}
