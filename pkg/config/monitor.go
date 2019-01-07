package config

type MonitorConfig struct {
	AppDownTime     int
	ApiApp          MonitorApiAppConfig
	AlertManagerApp MonitorAlertManagerAppConfig
	AgentApp        MonitorAgentAppConfig
	Database        DatabaseConfig
	Indexers        []MonitorIndexerConfig
}

type MonitorIndexerConfig struct {
	Driver          string
	Indexes         []string
	LogDatabases    []string
	MetricDatabases []string
}

type MonitorApiAppConfig struct {
	AppConfig
}

type MonitorAlertManagerAppConfig struct {
	AppConfig
}

type MonitorAgentAppConfig struct {
	AppConfig
	ReportIndex          string
	ReportProject        string
	ReportSpan           int
	LogReaderRefreshSpan int
	LogMap               map[string]MonitorLogConfig
	Metrics              MonitorMetricsConfig
}

type MonitorLogConfig struct {
	Path               string
	LogFormat          string
	MaxInitialReadSize int64
	AlertMap           map[string]MonitorLogAlertConfig
}

type MonitorLogAlertConfig struct {
	Key     string
	Pattern string
	Handler string
}

type MonitorMetricsConfig struct {
	System MonitorMetricsSystemConfig
}

type MonitorMetricsSystemConfig struct {
	Enable       bool
	EnableCpu    bool
	EnableMemory bool
}
