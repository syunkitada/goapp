package resource_region_server

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "region-server",
	Short: "region-server",
	Long: `region-server
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		server := NewResourceRegionServer(&config.Conf)
		if err := server.Serv(); err != nil {
			glog.Fatal(err)
		}
	},
}
