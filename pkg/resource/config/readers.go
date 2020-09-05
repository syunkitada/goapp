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
	WarnFreeMb            int64
	WarnAvailableMb       int64
	IsWarnAvailableMbAuto bool
	Occurences            int
	ReissueDuration       int
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

type ResourceProcCheckConfig struct {
	Cmd  string
	Name string
}
