package resource

import (
	// "fmt"
	"os"
	// "time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	createCmdClusterFlag string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create resource",
	Long: `create resource
	`,
}

var createComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "create compute",
	Long: `create compute
	`,
	Run: func(cmd *cobra.Command, args []string) {
		CreateCompute()
	},
}

func init() {
	createComputeCmd.Flags().StringVarP(&createCmdClusterFlag, "cluster", "c", "", "cluster (required)")
	createComputeCmd.MarkFlagRequired("cluster")

	createCmd.AddCommand(createComputeCmd)
	RootCmd.AddCommand(createCmd)
}

func CreateCompute() {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		glog.Fatal(err)
	}

	resp, err := authproxy.Resource.CtlCreateCompute(token.Token, createCmdClusterFlag)
	if err != nil {
		glog.Info(resp)
		glog.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name"})
	table.Append([]string{
		resp.Compute.Name,
	})
	table.Render()
}
