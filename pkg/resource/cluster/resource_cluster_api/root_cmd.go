package resource_cluster_api

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "cluster-api",
	Short: "This is api for controlle all resources",
	Long: `This is api for controlle all resources
	`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceClusterApiServer(&config.Conf)
		if err := server.StartMainLoop(); err != nil {
			glog.Fatal(err)
		}

		if err := server.Serve(); err != nil {
			glog.Fatal(err)
		}
	},
}
