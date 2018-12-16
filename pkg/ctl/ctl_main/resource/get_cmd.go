package resource

import (
	"github.com/spf13/cobra"
)

var (
	getCmdClusterFlag string
)

func init() {
	GetCmd.PersistentFlags().StringVarP(&getCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")

	GetCmd.AddCommand(getComputeCmd)
	GetCmd.AddCommand(getClusterCmd)
	GetCmd.AddCommand(getImageCmd)
	GetCmd.AddCommand(getNetworkV4Cmd)
	GetCmd.AddCommand(getNodeCmd)
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show resource",
	Long: `Show resource
	`,
}
