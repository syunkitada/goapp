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
	Compute ResourceComputeConfig
	LogMap  map[string]ResourceLogConfig
	Metric  ResourceMetricConfig
}

type ResourceComputeConfig struct {
	Enable               bool
	ConfirmRetryCount    int
	ConfirmRetryInterval int
	ConfigDir            string
	VarDir               string
	VmsDir               string
	ImagesDir            string

	VmNetnsGatewayStartIp string
	VmNetnsGatewayEndIp   string
	VmNetnsServiceIp      string
	VmNetnsStartIp        string
	VmNetnsEndIp          string

	Driver  string // libvirt
	Libvirt ResourceLibvirtConfig
}

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

type ResourceLibvirtConfig struct {
	AvailableCpus      string
	CpuMode            string // host-model
	CpuType            string // kvm, qemu
	MemoryType         string // local, hugepage
	DiskType           string // local
	NetworkType        string // local-linuxbridge
	NetworkVhostQueues int
}

type ResourceLogConfig struct {
	Path               string
	LogFormat          string
	MaxInitialReadSize int64
	AlertMap           map[string]ResourceLogAlertConfig
}

type ResourceLogAlertConfig struct {
	Key     string
	Pattern string
	Level   string
	Handler string
}

type ResourceMetricConfig struct {
	System ResourceMetricSystemConfig
}

type ResourceMetricSystemConfig struct {
	Enable       bool
	EnableLogin  bool
	EnableCpu    bool
	EnableMemory bool
	CacheLength  int
}

type TimeSeriesDatabaseConfig struct {
	Driver              string
	AlertDatabases      []string
	LogDatabases        []string
	MetricDatabases     []string
	PercistentDatabases []string
}

var BaseConf = base_config.Config{}

var MainConf = Config{
	Resource: ResourceConfig{
		Api: base_config.AppConfig{
			Name:                 "ResourceApi",
			NodeDownTimeDuration: 60,
		},
		Controller: ResourceControllerConfig{
			AppConfig: base_config.AppConfig{
				Name: "ReosurceController",
			},
			SyncRegionServiceTimeout: 10,
		},
	},
}
