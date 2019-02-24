package resource_ctl

import (
	"github.com/spf13/cobra"
)

var (
	deleteCmdClusterFlag string
)

func init() {
	deleteComputeCmd.Flags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")
	deleteComputeCmd.MarkFlagRequired("cluster")

	deleteNetworkCmd.Flags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")
	deleteNetworkCmd.MarkFlagRequired("cluster")

	deleteImageCmd.Flags().StringVarP(&deleteCmdClusterFlag, "cluster", "c", "", "Filtering by cluster")
	deleteImageCmd.MarkFlagRequired("cluster")

	DeleteCmd.AddCommand(deleteComputeCmd)
	DeleteCmd.AddCommand(deleteNetworkCmd)
	DeleteCmd.AddCommand(deleteImageCmd)
	RootCmd.AddCommand(DeleteCmd)
}

var DeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete resource",
	Long: `Delete resource
	`,
}
