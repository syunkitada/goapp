package resource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
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
		if err := UpdateResource(); err != nil {
			fmt.Println("Failed UpdateResource")
			for i, line := range strings.Split(err.Error(), "@@") {
				fmt.Printf("%v%v\n", strings.Repeat("  ", i), line)
			}
		}
	},
}

func init() {
	updateCmd.Flags().StringVarP(&updateCmdFileFlag, "file", "f", "", "file (required)")
	updateCmd.MarkFlagRequired("file")

	RootCmd.AddCommand(updateCmd)
}

func UpdateResource() error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err) }()

	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return fmt.Errorf("Failed issue token: %v", err)
	}

	bytes, err := ioutil.ReadFile(updateCmdFileFlag)
	if err != nil {
		return fmt.Errorf("Failed read file: %v", err)
	}

	var resourceSpec resource_model.ResourceSpec
	if err = json.Unmarshal(bytes, &resourceSpec); err != nil {
		return fmt.Errorf("Failed unmarshal file: %v", err)
	}

	switch resourceSpec.Kind {
	case resource_model.SpecNetworkV4:
		err = UpdateNetworkV4(token.Token, string(bytes))
		return err
	}

	return nil
}
