package ctl

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/base/base_config"
)

var baseConf base_config.Config

var RootCmd = &cobra.Command{
	Use: "ctl",
}
