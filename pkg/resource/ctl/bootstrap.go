package ctl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/config"
	"github.com/syunkitada/goapp/pkg/resource/db_api"

	cluster_db_api "github.com/syunkitada/goapp/pkg/resource/cluster/db_api"
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

		for clusterName, clusterConf := range config.MainConf.Resource.ClusterMap {
			fmt.Printf("Starting Bootstrap: %s\n", clusterName)
			clusterDbApi := cluster_db_api.New(&config.BaseConf, &clusterConf)
			if err := clusterDbApi.Bootstrap(tctx, false); err != nil {
				logger.Fatalf(tctx, "Failed Bootstrap: %v", err)
			}
			if err := clusterDbApi.BootstrapResource(tctx, false); err != nil {
				logger.Fatalf(tctx, "Failed Bootstrap: %v", err)
			}
			fmt.Printf("Success Bootstrap: %s\n", clusterName)
		}
	},
}

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}
