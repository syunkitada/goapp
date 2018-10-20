package resource

import (
	"os"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_client"
)

func init() {
	RootCmd.AddCommand(StatusCmd)
}

var StatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show resource api status",
	Long: `Show resource api status
                This is sample description1.
                This is sample description2.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := resource_api_client.NewResourceClient(&config.Conf)
		reply, err := client.Status()
		if err != nil {
			glog.Fatal(err)
		}

		data := [][]string{
			[]string{"clusterA", reply.Msg},
			[]string{"clusterB", reply.Msg},
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Cluster", "Status"})
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
	},
}
