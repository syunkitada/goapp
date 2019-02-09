package monitor

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appName = "ctl-monitor"
)

var RootCmd = &cobra.Command{
	Use:   "monitor",
	Short: "monitor service",
	Long: `monitor service
	`,
}

func PrintStackTrace(stackTrace []string) {
	fmt.Println("<StackTrace>")
	lenStackTrace := len(stackTrace)
	for i := lenStackTrace - 1; i >= 0; i = i - 1 {
		fmt.Printf("> %v\n", stackTrace[i])
	}
}
