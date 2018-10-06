package resource

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_client"
)

type Resource struct {
	Conf           *config.Config
	ResourceClient resource_client.ResourceClient
}

func NewResource(conf *config.Config) *Resource {
	resource := Resource{
		Conf:           conf,
		ResourceClient: resource_client.NewResourceClient(conf),
	}
	return &resource
}
