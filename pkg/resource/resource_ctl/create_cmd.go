package resource_ctl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
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
		ctl := New(&config.Conf, nil)
		if err := ctl.CreateResource(createCmdFileFlag); err != nil {
			fmt.Println("Failed by following error")
			fmt.Println(err.Error())
		}
	},
}

func init() {
	createCmd.Flags().StringVarP(&createCmdFileFlag, "file", "f", "", "file (required)")
	if err := createCmd.MarkFlagRequired("file"); err != nil {
		logger.StdoutFatal(err)
	}

	RootCmd.AddCommand(createCmd)
}

func (ctl *ResourceCtl) CreateResource(filePath string) error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Failed read file: %v", err)
	}

	var resourceSpec resource_model.ResourceSpec
	if err = json.Unmarshal(bytes, &resourceSpec); err != nil {
		return fmt.Errorf("Failed unmarshal file: %v", err)
	}

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return err
	}

	switch resourceSpec.Kind {
	case resource_model.SpecNetworkV4:
		err = ctl.CreateNetwork(tctx, token, string(bytes))
		return err
	case resource_model.SpecCompute:
		err = ctl.CreateCompute(tctx, token, string(bytes))
		return err
	case resource_model.SpecContainer:
		glog.Info("Container")
	case resource_model.SpecImage:
		err = ctl.CreateImage(tctx, token, string(bytes))
		return err
	case resource_model.SpecVolume:
		glog.Info("Volume")
	case resource_model.SpecLoadbalancer:
		glog.Info("Loadbalancer")
	}

	return nil
}
