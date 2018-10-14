package resource

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/ctl/ctl_main/resource/compute"
)

var RootCmd = &cobra.Command{
	Use:   "resource",
	Short: "resource service",
	Long: `resource service
                This is sample description1.
                This is sample description2.`,
}

func init() {
	RootCmd.AddCommand(compute.RootCmd)
}
