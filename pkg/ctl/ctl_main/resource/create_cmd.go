package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

var (
	createCmdFileFlag string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create resource",
	Long: `create resource
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := CreateResource(); err != nil {
			fmt.Println("Failed by following error")
			fmt.Println(err.Error())
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&createCmdFileFlag, "file", "f", "", "file (required)")
	createCmd.MarkFlagRequired("file")

	RootCmd.AddCommand(createCmd)
}

func CreateResource() error {
	traceId := logger.NewTraceId()
	startTime := logger.StartCtlTrace(traceId, appName)
	var err error

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		logger.EndCtlTrace(traceId, appName, startTime, err)
		return fmt.Errorf("Failed issue token: %v", err)
	}

	bytes, err := ioutil.ReadFile(createCmdFileFlag)
	if err != nil {
		logger.EndCtlTrace(traceId, appName, startTime, err)
		return fmt.Errorf("Failed read file: %v", err)
	}

	var resourceSpec resource_model.ResourceSpec
	if err = json.Unmarshal(bytes, &resourceSpec); err != nil {
		logger.EndCtlTrace(traceId, appName, startTime, err)
		return fmt.Errorf("Failed unmarshal file: %v", err)
	}

	switch resourceSpec.Kind {
	case resource_model.SpecNetworkV4:
		err = CreateNetworkV4(token.Token, string(bytes))
		logger.EndCtlTrace(traceId, appName, startTime, err)
		return err
	case resource_model.SpecCompute:
		resp, err := authproxy.Resource.CtlCreateCompute(token.Token, string(bytes))
		if err != nil {
			return err
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

	case resource_model.SpecContainer:
		glog.Info("Container")
	case resource_model.SpecImage:
		glog.Info("Image")
	case resource_model.SpecVolume:
		glog.Info("Volume")
	case resource_model.SpecLoadbalancer:
		glog.Info("Loadbalancer")
	}

	logger.EndCtlTrace(traceId, appName, startTime, err)
	return nil
}
