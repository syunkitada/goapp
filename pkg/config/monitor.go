package config

type MonitorConfig struct {
	AppDownTime     int
	ProxyApp        AppConfig
	AlertManagerApp AppConfig
	AgentApp        AppConfig
	Database        DatabaseConfig
}
