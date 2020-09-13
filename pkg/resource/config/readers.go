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
	Net          ResourceMetricSystemNetConfig
	NetDev       ResourceMetricSystemNetDevConfig
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

type ResourceMetricSystemNetConfig struct {
	Enable bool
}

type ResourceMetricSystemNetDevConfig struct {
	Enable       bool
	StatFilters  []string
	CheckFilters []string
	CheckBytes   ResourceMetricSystemNetDevCheckBytesConfig
	CheckErrors  ResourceMetricSystemNetDevCheckErrorsConfig
}

type ResourceMetricSystemNetDevCheckBytesConfig struct {
	Occurences      int
	ReissueDuration int
	WarnRatio       float64
	CritRatio       float64
}

type ResourceMetricSystemNetDevCheckErrorsConfig struct {
	Occurences      int
	ReissueDuration int
	WarnErrors      int64
	CritErrors      int64
	WarnDrops       int64
	CritDrops       int64
}

type ResourceProcCheckConfig struct {
	Cmd  string
	Name string
}
