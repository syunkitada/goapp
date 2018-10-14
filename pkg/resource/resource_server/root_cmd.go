package resource_server

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
)

var RootCmd = &cobra.Command{
	Use:   "server",
	Short: "server",
	Long: `server
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		resourceServer := NewResourceServer(&config.Conf)
		if err := resourceServer.Serv(); err != nil {
			glog.Fatal(err)
		}
	},
}
