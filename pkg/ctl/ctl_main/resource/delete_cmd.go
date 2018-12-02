package resource

import (
	"github.com/spf13/cobra"
)

var (
	deleteCmdClusterFlag string
)

func init() {
	DeleteCmd.PersistentFlags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")

	DeleteCmd.AddCommand(deleteNetworkV4Cmd)
	RootCmd.AddCommand(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Show resource",
	Long: `Show resource
	`,
}
