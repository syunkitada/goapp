package authproxy

import (
	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:   "goapp-authproxy",
	Short: "goapp-authproxy",
	Long: `goapp-authproxy
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		authproxy := core.NewAuthproxy(&config.Conf)
		authproxy.Serv()
	},
}

func Main() {
	if err := rootCmd.Execute(); err != nil {
		glog.Fatal(err)
	}
}

func init() {
	cobra.OnInitialize(config.InitConfig)
	config.InitFlags(rootCmd)
}
