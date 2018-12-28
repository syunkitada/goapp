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

var getComputeCmd = &cobra.Command{
	Use:   "compute",
	Short: "Show computes",
	Long: `Show computes
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetCompute(); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteComputeCmd = &cobra.Command{
	Use:   "compute [compute-name]",
	Short: "Show computes",
	Long: `Show computes
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeleteCompute(args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetCompute() error {
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

	resp, err := authproxy.Resource.CtlGetCompute(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetCompute.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, compute := range resp.Computes {
		table.Append([]string{
			compute.Cluster,
			compute.Name,
			compute.Status,
			compute.StatusReason,
			fmt.Sprint(time.Unix(compute.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(compute.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}

func CreateCompute(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlCreateCompute(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Compute.Cluster,
		resp.Compute.Name,
		resp.Compute.Status,
		resp.Compute.StatusReason,
		fmt.Sprint(time.Unix(resp.Compute.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Compute.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func UpdateCompute(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlUpdateCompute(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Compute.Cluster,
		resp.Compute.Name,
		resp.Compute.Status,
		resp.Compute.StatusReason,
		fmt.Sprint(time.Unix(resp.Compute.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Compute.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func DeleteCompute(computeName string) error {
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

	resp, err := authproxy.Resource.CtlDeleteCompute(token.Token, deleteCmdClusterFlag, computeName)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	fmt.Println("Deleted")

	return nil
}
