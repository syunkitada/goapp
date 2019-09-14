package ctl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_api"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap",
	Run: func(cmd *cobra.Command, args []string) {
		tctx := logger.NewTraceContext(baseConf.Host, "bootstrap")
		dbApi := db_api.New(&config.BaseConf, &config.MainConf)
		if err := dbApi.Bootstrap(tctx, false); err != nil {
			logger.Fatalf(tctx, "Failed Bootstrap: %v", err)
		}
		fmt.Println("Success Bootstrap")
		if err := dbApi.BootstrapResource(tctx, false); err != nil {
			logger.Fatalf(tctx, "Failed BootstrapResource: %v", err)
		}
		fmt.Println("Success BootstrapResource")
	},
}

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}
