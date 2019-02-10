package resource_ctl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

var (
	updateCmdFileFlag string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "update resource",
	Long: `update resource
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := New(&config.Conf, nil)
		if err := ctl.UpdateResource(updateCmdFileFlag); err != nil {
			fmt.Println("Failed UpdateResource")
			for i, line := range strings.Split(err.Error(), "@@") {
				fmt.Printf("%v%v\n", strings.Repeat("  ", i), line)
			}
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateCmdFileFlag, "file", "f", "", "file (required)")
	if err := updateCmd.MarkFlagRequired("file"); err != nil {
		logger.StdoutFatal(err)
	}

	RootCmd.AddCommand(updateCmd)
}

func (ctl *ResourceCtl) UpdateResource(filePath string) error {
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
		err = ctl.UpdateNetwork(tctx, token, string(bytes))
		return err
	}

	return nil
}
