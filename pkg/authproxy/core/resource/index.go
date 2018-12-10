package resource

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
)

type Resource struct {
	host              string
	name              string
	conf              *config.Config
	resourceApiClient *resource_api_client.ResourceApiClient
}

func NewResource(conf *config.Config) *Resource {
	resource := Resource{
		host:              conf.Default.Host,
		name:              "authproxy:resource",
		conf:              conf,
		resourceApiClient: resource_api_client.NewResourceApiClient(conf),
	}
	return &resource
}
