package resource_authproxy

import (
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
)

type Resource struct {
	host              string
	name              string
	path              string
	conf              *config.Config
	resourceApiClient *resource_api_client.ResourceApiClient
}

func New(conf *config.Config) *Resource {
	resource := Resource{
		host:              conf.Default.Host,
		name:              "authproxy:resource",
		path:              "/Resource",
		conf:              conf,
		resourceApiClient: resource_api_client.NewResourceApiClient(conf),
	}
	return &resource
}

func (resource *Resource) GetPath() string {
	return resource.path
}
