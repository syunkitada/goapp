package resource

import (
	"github.com/spf13/cobra"
)

var (
	deleteCmdClusterFlag string
)

func init() {
	deleteNetworkV4Cmd.Flags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")
	deleteNetworkV4Cmd.MarkFlagRequired("cluster")

	DeleteCmd.AddCommand(deleteNetworkV4Cmd)
	RootCmd.AddCommand(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Show resource",
	Long: `Show resource
	`,
}
