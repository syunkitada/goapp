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

type ResourceMetricsConfig struct {
	System ResourceMetricsSystemConfig
}

type ResourceMetricsSystemConfig struct {
	Enable       bool
	EnableLogin  bool
	EnableCpu    bool
	EnableMemory bool
	EnableProc   bool
	CacheLength  int

	Uptime       ResourceMetricsSystemUptimeConfig
	Login        ResourceMetricsSystemLoginConfig
	Proc         ResourceMetricsSystemProcConfig
	Cpu          ResourceMetricsSystemCpuConfig
	Mem          ResourceMetricsSystemMemConfig
	MemBuddyinfo ResourceMetricsSystemMemBuddyinfoConfig
	Disk         ResourceMetricsSystemDiskConfig
	DiskFs       ResourceMetricsSystemDiskFsConfig
	Net          ResourceMetricsSystemNetConfig
	NetDev       ResourceMetricsSystemNetDevConfig
}

type ResourceMetricsSystemUptimeConfig struct {
	Enable    bool
	CheckBoot ResourceMetricsSystemUptimeCheckBootConfig
}

type ResourceMetricsSystemUptimeCheckBootConfig struct {
	Occurences      int
	ReissueDuration int
	ReadinessSec    int64
}

type ResourceMetricsSystemLoginConfig struct {
	Enable     bool
	CheckLogin ResourceMetricsSystemLoginCheckLoginConfig
}

type ResourceMetricsSystemLoginCheckLoginConfig struct {
	Occurences      int
	ReissueDuration int
	WarnLoginSec    int64
	CritLoginSec    int64
}

type ResourceMetricsSystemProcConfig struct {
	Enable           bool
	CheckProcsStatus ResourceMetricsSystemProcCheckProcsStatusConfig
	CheckProcMap     map[string]ResourceMetricsSystemProcCheckProcConfig
}

type ResourceMetricsSystemProcCheckProcsStatusConfig struct {
	Occurences      int
	ReissueDuration int
}

type ResourceMetricsSystemProcCheckProcConfig struct {
	Occurences        int
	ReissueDuration   int
	Cmd               string
	Name              string
	WarnSchedWaitTime int64
	CritSchedWaitTime int64
}

type ResourceProcCheckConfig struct {
	Cmd  string
	Name string
}

type ResourceMetricsSystemCpuConfig struct {
	Enable            bool
	CheckProcsRunning ResourceMetricsSystemCpuCheckProcsRunningConfig
	CheckProcsBlocked ResourceMetricsSystemCpuCheckProcsBlockedConfig
}

type ResourceMetricsSystemCpuCheckProcsRunningConfig struct {
	WarnRateLimit   float64
	CritRateLimit   float64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricsSystemCpuCheckProcsBlockedConfig struct {
	WarnRateLimit   float64
	CritRateLimit   float64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricsSystemMemConfig struct {
	Enable         bool
	CheckAvailable ResourceMetricsSystemMemCheckAvailableConfig
	CheckPgscan    ResourceMetricsSystemMemCheckPgscanConfig
}

type ResourceMetricsSystemMemCheckAvailableConfig struct {
	WarnAvailableRatio float64
	Occurences         int
	ReissueDuration    int
}

type ResourceMetricsSystemMemCheckPgscanConfig struct {
	WarnPgscanDirect int64
	Occurences       int
	ReissueDuration  int
}

type ResourceMetricsSystemMemBuddyinfoConfig struct {
	Enable     bool
	CheckPages ResourceMetricsSystemMemBuddyinfoCheckPagesConfig
}

type ResourceMetricsSystemMemBuddyinfoCheckPagesConfig struct {
	WarnMinPages    int64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricsSystemDiskConfig struct {
	Enable       bool
	CheckIoDelay ResourceMetricsSystemCheckIoDelayConfig
}

type ResourceMetricsSystemCheckIoDelayConfig struct {
	Occurences        int
	ReissueDuration   int
	CritReadMsPerSec  int64
	WarnReadMsPerSec  int64
	CritWriteMsPerSec int64
	WarnWriteMsPerSec int64
	CritProgressIos   int64
	WarnProgressIos   int64
}

type ResourceMetricsSystemDiskFsConfig struct {
	Enable    bool
	CheckFree ResourceMetricsSystemDiskFsCheckFreeConfig
}

type ResourceMetricsSystemDiskFsCheckFreeConfig struct {
	WarnFreeRatio   float64
	CritFreeRatio   float64
	Occurences      int
	ReissueDuration int
}

type ResourceMetricsSystemNetConfig struct {
	Enable         bool
	CheckTcpErrors ResourceMetricsSystemNetCheckTcpErrorsConfig
}

type ResourceMetricsSystemNetCheckTcpErrorsConfig struct {
	Occurences             int
	ReissueDuration        int
	WarnOnPressures        bool
	CritOnPressures        bool
	WarnOnTcpAbortOnMemory bool
	CritOnTcpAbortOnMemory bool
	WarnOnListenDrops      bool
	CritOnListenDrops      bool
	WarnOnListenOverflows  bool
	CritOnListenOverflows  bool
}

type ResourceMetricsSystemNetDevConfig struct {
	Enable       bool
	StatFilters  []string
	CheckFilters []string
	CheckBytes   ResourceMetricsSystemNetDevCheckBytesConfig
	CheckErrors  ResourceMetricsSystemNetDevCheckErrorsConfig
}

type ResourceMetricsSystemNetDevCheckBytesConfig struct {
	Occurences      int
	ReissueDuration int
	WarnRatio       float64
	CritRatio       float64
}

type ResourceMetricsSystemNetDevCheckErrorsConfig struct {
	Occurences      int
	ReissueDuration int
	WarnErrors      int64
	CritErrors      int64
	WarnDrops       int64
	CritDrops       int64
}
