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

	Uptime       ResourceMetricSystemUptimeConfig
	Login        ResourceMetricSystemLoginConfig
	Proc         ResourceMetricSystemProcConfig
	Cpu          ResourceMetricSystemCpuConfig
	Mem          ResourceMetricSystemMemConfig
	MemBuddyinfo ResourceMetricSystemMemBuddyinfoConfig
	Disk         ResourceMetricSystemDiskConfig
	DiskFs       ResourceMetricSystemDiskFsConfig
	Net          ResourceMetricSystemNetConfig
	NetDev       ResourceMetricSystemNetDevConfig
}

type ResourceMetricSystemUptimeConfig struct {
	Enable    bool
	CheckBoot ResourceMetricSystemUptimeCheckBootConfig
}

type ResourceMetricSystemUptimeCheckBootConfig struct {
	Occurences      int
	ReissueDuration int
	ReadinessSec    int64
}

type ResourceMetricSystemLoginConfig struct {
	Enable     bool
	CheckLogin ResourceMetricSystemLoginCheckLoginConfig
}

type ResourceMetricSystemLoginCheckLoginConfig struct {
	Occurences      int
	ReissueDuration int
	WarnLoginSec    int64
	CritLoginSec    int64
}

type ResourceMetricSystemProcConfig struct {
	Enable           bool
	CheckProcsStatus ResourceMetricSystemProcCheckProcsStatusConfig
	CheckProcMap     map[string]ResourceMetricSystemProcCheckProcConfig
}

type ResourceMetricSystemProcCheckProcsStatusConfig struct {
	Occurences      int
	ReissueDuration int
}

type ResourceMetricSystemProcCheckProcConfig struct {
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
	Enable       bool
	CheckIoDelay ResourceMetricSystemCheckIoDelayConfig
}

type ResourceMetricSystemCheckIoDelayConfig struct {
	Occurences        int
	ReissueDuration   int
	CritReadMsPerSec  int64
	WarnReadMsPerSec  int64
	CritWriteMsPerSec int64
	WarnWriteMsPerSec int64
	CritProgressIos   int64
	WarnProgressIos   int64
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
	Enable         bool
	CheckTcpErrors ResourceMetricSystemNetCheckTcpErrorsConfig
}

type ResourceMetricSystemNetCheckTcpErrorsConfig struct {
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
