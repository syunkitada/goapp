package monitor

import (
	"github.com/spf13/cobra"
)

var (
	getCmdIndexFlag string
)

func init() {
	GetCmd.PersistentFlags().StringVarP(&getCmdIndexFlag, "index", "i", "", "Filtering by index")

	GetCmd.AddCommand(getNodeCmd)
	GetCmd.AddCommand(getHostCmd)
	GetCmd.AddCommand(getLogCmd)
	RootCmd.AddCommand(GetCmd)
}

var GetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show monitor",
	Long: `Show monitor
	`,
}
