package ctl

import (
	"github.com/spf13/cobra"

	authproxy_api_spec "github.com/syunkitada/goapp/pkg/authproxy/authproxy_api/spec"
	"github.com/syunkitada/goapp/pkg/base/code_generator"
)

var generateCodeCmd = &cobra.Command{
	Use:   "generate-code",
	Short: "generate-code",
	Run: func(cmd *cobra.Command, args []string) {
		code_generator.Generate(&authproxy_api_spec.Spec)
		code_generator.GenerateStatusCodes()
	},
}

func init() {
	RootCmd.AddCommand(generateCodeCmd)
}
