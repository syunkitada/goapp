package config

type MonitorConfig struct {
	AppDownTime     int
	ApiApp          MonitorApiAppConfig
	AlertManagerApp MonitorAlertManagerAppConfig
	AgentApp        MonitorAgentAppConfig
	Database        DatabaseConfig
}

type MonitorApiAppConfig struct {
	AppConfig
	IndexDatabaseMap map[string]MonitorDatabaseConfig
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

type MonitorDatabaseConfig struct {
	Connections []string
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
