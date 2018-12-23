package config

type MonitorConfig struct {
	AppDownTime     int
	ApiApp          AppConfig
	AlertManagerApp AppConfig
	AgentApp        AppConfig
	Database        DatabaseConfig
}
