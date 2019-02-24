package resource_ctl

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
)

type ResourceCtl struct {
	ctl    *config.Config
	client *authproxy_client.AuthproxyClient
}

func New(ctl *config.Config, authproxy *core.Authproxy) *ResourceCtl {
	client := authproxy_client.New(&config.Conf, "Resource", authproxy)

	return &ResourceCtl{
		ctl:    ctl,
		client: client,
	}
}
