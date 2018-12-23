package ctl_admin

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/monitor/monitor_model/monitor_model_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_model/resource_model_api"
)

type Ctl struct {
	Conf              *config.Config
	AuthproxyModelApi *authproxy_model_api.AuthproxyModelApi
	MonitorModelApi   *monitor_model_api.MonitorModelApi
	ResourceModelApi  *resource_model_api.ResourceModelApi
}

func NewCtl(conf *config.Config) *Ctl {
	ctl := Ctl{
		Conf:              conf,
		AuthproxyModelApi: authproxy_model_api.NewAuthproxyModelApi(conf),
		MonitorModelApi:   monitor_model_api.NewMonitorModelApi(conf),
		ResourceModelApi:  resource_model_api.NewResourceModelApi(conf, nil),
	}

	return &ctl
}
