package resource

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
)

var RootCmd = &cobra.Command{
	Use:   "resource",
	Short: "resource service",
	Long: `resource service
	`,
}

func init() {
	return
}

type ResourceCtl struct {
	conf              *config.Config
	resourceApiClient *resource_api_client.ResourceApiClient
}

func NewResourceCtl(conf *config.Config) *ResourceCtl {
	ctl := &ResourceCtl{
		conf:              conf,
		resourceApiClient: resource_api_client.NewResourceApiClient(conf),
	}
	return ctl
}
