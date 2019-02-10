package resource_ctl

import (
	"github.com/spf13/cobra"
)

var (
	deleteCmdClusterFlag string
)

func init() {
	deleteNetworkCmd.Flags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")
	deleteNetworkCmd.MarkFlagRequired("cluster")

	DeleteCmd.AddCommand(deleteNetworkCmd)
	RootCmd.AddCommand(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Show resource",
	Long: `Show resource
	`,
}
