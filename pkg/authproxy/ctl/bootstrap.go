package ctl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/authproxy/db_api"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var bootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap",
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewCtl(&config.BaseConf, &config.MainConf)
		if tmpErr := ctl.Bootstrap(false); tmpErr != nil {
			logger.StdoutFatalf("Failed Bootstrap: %s\n", tmpErr.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(bootstrapCmd)
}

func (ctl *Ctl) Bootstrap(isRecreate bool) (err error) {
	tctx := logger.NewTraceContext(ctl.baseConf.Host, "bootstrap")
	dbApi := db_api.New(ctl.baseConf, ctl.mainConf)
	if err := dbApi.Bootstrap(tctx, isRecreate); err != nil {
		logger.Fatalf(tctx, "Failed Bootstrap: %v", err)
	}
	fmt.Println("Success Bootstrap")
	return
}
