package monitor

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_api/monitor_api_client"
)

type Monitor struct {
	host             string
	name             string
	conf             *config.Config
	monitorApiClient *monitor_api_client.MonitorApiClient
}

func NewMonitor(conf *config.Config) *Monitor {
	monitor := Monitor{
		host:             conf.Default.Host,
		name:             "authproxy:monitor",
		conf:             conf,
		monitorApiClient: monitor_api_client.NewMonitorApiClient(conf),
	}
	return &monitor
}
