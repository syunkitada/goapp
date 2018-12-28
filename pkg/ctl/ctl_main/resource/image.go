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

var getImageCmd = &cobra.Command{
	Use:   "image",
	Short: "Show images",
	Long: `Show images
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetImage(); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteImageCmd = &cobra.Command{
	Use:   "image [image-name]",
	Short: "Show images",
	Long: `Show images
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeleteImage(args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetImage() error {
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

	resp, err := authproxy.Resource.CtlGetImage(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("GetImage.TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, image := range resp.Images {
		table.Append([]string{
			image.Cluster,
			image.Name,
			image.Status,
			image.StatusReason,
			fmt.Sprint(time.Unix(image.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(image.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()

	return nil
}

func CreateImage(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlCreateImage(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Image.Cluster,
		resp.Image.Name,
		resp.Image.Status,
		resp.Image.StatusReason,
		fmt.Sprint(time.Unix(resp.Image.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Image.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func UpdateImage(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlUpdateImage(token, spec)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Image.Cluster,
		resp.Image.Name,
		resp.Image.Status,
		resp.Image.StatusReason,
		fmt.Sprint(time.Unix(resp.Image.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Image.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func DeleteImage(imageName string) error {
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

	resp, err := authproxy.Resource.CtlDeleteImage(token.Token, deleteCmdClusterFlag, imageName)
	if err != nil {
		return err
	}
	if config.Conf.Default.EnableDebug {
		fmt.Printf("TraceID: %v\n", resp.TraceId)
	}

	fmt.Println("Deleted")

	return nil
}
