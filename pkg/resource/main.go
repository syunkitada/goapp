package resource

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_api"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_controller"
	"github.com/syunkitada/goapp/pkg/resource/resource_api"
	"github.com/syunkitada/goapp/pkg/resource/resource_controller"
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

	rootCmd.AddCommand(resource_api.RootCmd)
	rootCmd.AddCommand(resource_controller.RootCmd)
	rootCmd.AddCommand(resource_cluster_api.RootCmd)
	rootCmd.AddCommand(resource_cluster_controller.RootCmd)
}
