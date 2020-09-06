package config

type ResourceLogConfig struct {
	Path               string
	LogFormat          string
	MaxInitialReadSize int64
	CheckPrefix        string
	CheckMap           map[string]ResourceLogCheckConfig
}

type ResourceLogCheckConfig struct {
	Key             string
	Pattern         string
	Level           string
	ReissueDuration int
}

type ResourceMetricConfig struct {
	System ResourceMetricSystemConfig
}

type ResourceMetricSystemConfig struct {
	Enable       bool
	EnableLogin  bool
	EnableCpu    bool
	EnableMemory bool
	EnableProc   bool
	CacheLength  int
	ProcCheckMap map[string]ResourceProcCheckConfig

	Cpu          ResourceMetricSystemCpuConfig
	Mem          ResourceMetricSystemMemConfig
	MemBuddyinfo ResourceMetricSystemMemBuddyinfoConfig
	Disk         ResourceMetricSystemDiskConfig
	DiskFs       ResourceMetricSystemDiskFsConfig
}

type ResourceMetricSystemCpuConfig struct {
	Enable            bool
	CheckProcsRunning ResourceMetricSystemCpuCheckProcsRunningConfig
	CheckProcsBlocked ResourceMetricSystemCpuCheckProcsBlockedConfig
}

type ResourceMetricSystemCpuCheckProcsRunningConfig struct {
	WarnRateLimit   float64
	CritRateLimit   float64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricSystemCpuCheckProcsBlockedConfig struct {
	WarnRateLimit   float64
	CritRateLimit   float64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricSystemMemConfig struct {
	Enable         bool
	CheckAvailable ResourceMetricSystemMemCheckAvailableConfig
	CheckPgscan    ResourceMetricSystemMemCheckPgscanConfig
}

type ResourceMetricSystemMemCheckAvailableConfig struct {
	WarnAvailableRatio float64
	Occurences         int
	ReissueDuration    int
}

type ResourceMetricSystemMemCheckPgscanConfig struct {
	WarnPgscanDirect int64
	Occurences       int
	ReissueDuration  int
}

type ResourceMetricSystemMemBuddyinfoConfig struct {
	Enable     bool
	CheckPages ResourceMetricSystemMemBuddyinfoCheckPagesConfig
}

type ResourceMetricSystemMemBuddyinfoCheckPagesConfig struct {
	WarnMinPages    int64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricSystemDiskConfig struct {
	Enable bool
}

type ResourceMetricSystemDiskFsConfig struct {
	Enable    bool
	CheckFree ResourceMetricSystemDiskFsCheckFreeConfig
}

type ResourceMetricSystemDiskFsCheckFreeConfig struct {
	WarnFreeRatio   float64
	CritFreeRatio   float64
	Occurences      int
	ReissueDuration int
}

type ResourceProcCheckConfig struct {
	Cmd  string
	Name string
}
