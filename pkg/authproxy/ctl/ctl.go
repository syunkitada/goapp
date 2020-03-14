package ctl

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/config"
	"github.com/syunkitada/goapp/pkg/base/base_config"
)

var baseConf base_config.Config

var RootCmd = &cobra.Command{
	Use: "ctl",
}

type Ctl struct {
	baseConf *base_config.Config
	mainConf *config.Config
}

func NewCtl(baseConf *base_config.Config, mainConf *config.Config) *Ctl {
	return &Ctl{
		baseConf: baseConf,
		mainConf: mainConf,
	}
}
