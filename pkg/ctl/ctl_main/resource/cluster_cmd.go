package resource

import (
	"fmt"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var getClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Show clusters",
	Long: `Show clusters
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetCluster(); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetCluster() error {
	var err error
	traceId := logger.NewTraceContext()
	startTime := logger.StartCtlTrace(traceId, appName)
	defer func() {
		logger.EndCtlTrace(traceId, appName, startTime, err)
	}()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Resource.CtlGetCluster(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetCluster.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Updated At", "Created At"})
	for _, cluster := range resp.Clusters {
		table.Append([]string{
			cluster.Name,
			fmt.Sprint(time.Unix(cluster.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(cluster.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}
