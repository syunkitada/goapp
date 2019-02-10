package resource_ctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	appName = "ctl-resource"
)

var RootCmd = &cobra.Command{
	Use:   "resource",
	Short: "resource service",
	Long: `resource service
	`,
}

func PrintStackTrace(stackTrace []string) {
	fmt.Println("<StackTrace>")
	lenStackTrace := len(stackTrace)
	for i := lenStackTrace - 1; i >= 0; i = i - 1 {
		fmt.Printf("> %v\n", stackTrace[i])
	}
}
