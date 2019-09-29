package ctl

import (
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/base/code_generator"
	resource_api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
	resource_controller_spec "github.com/syunkitada/goapp/pkg/resource/resource_controller/spec"
)

var generateCodeCmd = &cobra.Command{
	Use:   "generate-code",
	Short: "generate-code",
	Run: func(cmd *cobra.Command, args []string) {
		code_generator.Generate(&resource_api_spec.Spec)
		code_generator.Generate(&resource_controller_spec.Spec)
	},
}

func init() {
	RootCmd.AddCommand(generateCodeCmd)
}
