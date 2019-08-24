package ctl

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap",
	Run: func(cmd *cobra.Command, args []string) {
		tctx := logger.NewTraceContext(baseConf.Host, "bootstrap")
		dbApi := db_api.New(&baseConf, &mainConf)
		if err := dbApi.Bootstrap(tctx, false); err != nil {
			logger.Fatalf(tctx, "Failed Bootstrap: %v", err)
		}
		fmt.Println("Success Bootstrap")
	},
}

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}
