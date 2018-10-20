package resource

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
)

type Resource struct {
	Conf              *config.Config
	ResourceApiClient *resource_api_client.ResourceApiClient
}

func NewResource(conf *config.Config) *Resource {
	resource := Resource{
		Conf:              conf,
		ResourceApiClient: resource_api_client.NewResourceApiClient(conf),
	}
	return &resource
}
