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
	Driver      string
	Indexes     []string
	Connections []string
}

type MonitorApiAppConfig struct {
	AppConfig
}

type MonitorAlertManagerAppConfig struct {
	AppConfig
}

type MonitorAgentAppConfig struct {
	AppConfig
	Index                string
	FlushSpan            int
	LogReaderRefreshSpan int
	LogMap               map[string]MonitorLogConfig
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
