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

	Cpu ResourceMetricSystemCpuConfig
}

type ResourceMetricSystemCpuConfig struct {
	Enable            bool
	CheckProcsRunning ResourceMetricSystemCpuCheckProcsRunningConfig
	CheckProcsBlocked ResourceMetricSystemCpuCheckProcsBlockedConfig
}

type ResourceMetricSystemCpuCheckProcsRunningConfig struct {
	WarnRateLimit float64
	CritRateLimit float64
	Occurences    int
}

type ResourceMetricSystemCpuCheckProcsBlockedConfig struct {
	WarnRateLimit float64
	CritRateLimit float64
	Occurences    int
}

type ResourceProcCheckConfig struct {
	Cmd  string
	Name string
}
