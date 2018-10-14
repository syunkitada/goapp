package resource

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/region/resource_region_server"
	"github.com/syunkitada/goapp/pkg/resource/resource_server"
)

var rootCmd = &cobra.Command{}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	config.InitFlags(rootCmd)

	rootCmd.AddCommand(resource_server.RootCmd)
	rootCmd.AddCommand(resource_region_server.RootCmd)
}
