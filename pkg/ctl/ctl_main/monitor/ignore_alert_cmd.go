package monitor

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var getIgnoreAlertCmd = &cobra.Command{
	Use:   "ignorealert",
	Short: "Show ignoreAlerts",
	Long: `Show ignoreAlerts
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetIgnoreAlert(); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteIgnoreAlertCmd = &cobra.Command{
	Use:   "ignorealert [ignorealert-name]",
	Short: "Show ignoreAlerts",
	Long: `Show ignoreAlerts
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeleteIgnoreAlert(args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetIgnoreAlert() error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken(tctx)
	if err != nil {
		return err
	}

	resp, err := authproxy.Monitor.CtlGetIgnoreAlert(token.Token, getCmdIndexFlag)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetIgnoreAlert.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Host", "Name", "Level", "User", "Reason", "Until", "Updated At", "Created At"})
	for _, ignoreAlert := range resp.IgnoreAlerts {
		table.Append([]string{
			ignoreAlert.Index,
			ignoreAlert.Host,
			ignoreAlert.Name,
			ignoreAlert.Level,
			ignoreAlert.User,
			ignoreAlert.Reason,
			"TODO",
			fmt.Sprint(time.Unix(ignoreAlert.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(ignoreAlert.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}

func CreateIgnoreAlert(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Monitor.CtlCreateIgnoreAlert(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Host", "Name", "Level", "User", "Reason", "Until", "Updated At", "Created At"})
	table.Append([]string{
		resp.IgnoreAlert.Index,
		resp.IgnoreAlert.Host,
		resp.IgnoreAlert.Name,
		resp.IgnoreAlert.Level,
		resp.IgnoreAlert.User,
		resp.IgnoreAlert.Reason,
		"TODO until",
		fmt.Sprint(time.Unix(resp.IgnoreAlert.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.IgnoreAlert.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func UpdateIgnoreAlert(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Monitor.CtlUpdateIgnoreAlert(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Index", "Host", "Name", "Level", "User", "Reason", "Until", "Updated At", "Created At"})
	table.Append([]string{
		resp.IgnoreAlert.Index,
		resp.IgnoreAlert.Host,
		resp.IgnoreAlert.Name,
		resp.IgnoreAlert.Level,
		resp.IgnoreAlert.User,
		resp.IgnoreAlert.Reason,
		"TODO until",
		fmt.Sprint(time.Unix(resp.IgnoreAlert.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.IgnoreAlert.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func DeleteIgnoreAlert(ignoreAlertId string) error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken(tctx)
	if err != nil {
		return err
	}

	var id uint64
	id, err = strconv.ParseUint(ignoreAlertId, 10, 64)
	if err != nil {
		return err
	}

	resp, err := authproxy.Monitor.CtlDeleteIgnoreAlert(token.Token, id)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	fmt.Println("Deleted")

	return nil
}
