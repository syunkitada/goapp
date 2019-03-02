package resource_ctl

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
)

type ResponseCreate struct {
	Tctx authproxy_grpc_pb.TraceContext
}

var (
	createCmdResourceTypeFlag string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create resource",
	Long: `create resource
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("require one args")
			return
		}

		ctl := New(&config.Conf, nil)
		if err := ctl.CreateResource(createCmdResourceTypeFlag, args); err != nil {
			fmt.Println("Failed by following error")
			fmt.Println(err.Error())
		}
	},
}

func init() {
	createCmd.PersistentFlags().StringVarP(&createCmdResourceTypeFlag, "type", "t", "virtual", "Type of resource")

	RootCmd.AddCommand(createCmd)
}

func (ctl *ResourceCtl) CreateResource(resourceType string, filePaths []string) error {
	var err error
	tctx := logger.NewCtlTraceContext(appName)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	data, err := json_utils.ReadFilesFromMultiPath(filePaths)
	if err != nil {
		return err
	}
	dataBytes, err := json_utils.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Println(string(dataBytes))

	req := resource_api_grpc_pb.ActionRequest{
		Spec: string(dataBytes),
	}

	var token *authproxy_client.ResponseIssueToken
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return err
	}

	switch resourceType {
	case "virtual":
		var resp ResponseCreate
		if err = ctl.client.Request(tctx, token, "CreateVirtualResource", &req, &resp); err != nil {
			return err
		}
		fmt.Println(resp)
	case "physical":
		var resp ResponseCreate
		if err = ctl.client.Request(tctx, token, "CreatePhysicalResource", &req, &resp); err != nil {
			return err
		}
		fmt.Println(resp)
	}

	return nil

	// bytes, err := ioutil.ReadFile(filePath)
	// if err != nil {
	// 	return fmt.Errorf("Failed read file: %v", err)
	// }

	// var resourceSpec resource_model.ResourceSpec
	// if err = json.Unmarshal(bytes, &resourceSpec); err != nil {
	// 	return fmt.Errorf("Failed unmarshal file: %v", err)
	// }

	// var token *authproxy_client.ResponseIssueToken
	// if token, err = ctl.client.IssueToken(tctx); err != nil {
	// 	return err
	// }

	// switch resourceSpec.Kind {
	// case resource_model.SpecNetworkV4:
	// 	err = ctl.CreateNetwork(tctx, token, string(bytes))
	// 	return err
	// case resource_model.SpecCompute:
	// 	err = ctl.CreateCompute(tctx, token, string(bytes))
	// 	return err
	// case resource_model.SpecImage:
	// 	err = ctl.CreateImage(tctx, token, string(bytes))
	// 	return err
	// default:
	// 	fmt.Printf("InvalidKind: %v\n", resourceSpec.Kind)
	// }

	// return nil
}
