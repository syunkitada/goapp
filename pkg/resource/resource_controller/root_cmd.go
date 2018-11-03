package resource_controller

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "controller",
	Short: "controller for management all resources",
	Long: `controller for management all resources
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceControllerServer(&config.Conf)
		if err := server.StartMainLoop(); err != nil {
			glog.Fatal(err)
		}

		if err := server.Serve(); err != nil {
			glog.Fatal(err)
		}
	},
}
