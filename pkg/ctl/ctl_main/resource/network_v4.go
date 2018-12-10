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
)

var getNetworkV4Cmd = &cobra.Command{
	Use:   "networkv4",
	Short: "Show networks",
	Long: `Show networks
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := GetNetworkV4(); err != nil {
			glog.Fatal(err)
		}
	},
}

var deleteNetworkV4Cmd = &cobra.Command{
	Use:   "networkv4 [network-name]",
	Short: "Show networks",
	Long: `Show networks
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if err := DeleteNetworkV4(args[0]); err != nil {
			glog.Fatal(err)
		}
	},
}

func GetNetworkV4() error {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Resource.CtlGetNetworkV4(token.Token, getCmdClusterFlag, "%")
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	for _, network := range resp.Networks {
		table.Append([]string{
			network.Cluster,
			network.Name,
			network.Status,
			network.StatusReason,
			fmt.Sprint(time.Unix(network.UpdatedAt.Seconds, 0)),
			fmt.Sprint(time.Unix(network.CreatedAt.Seconds, 0)),
		})
	}
	table.Render()
	return nil
}

func CreateNetworkV4(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlCreateNetworkV4(token, spec)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Network.Cluster,
		resp.Network.Name,
		resp.Network.Status,
		resp.Network.StatusReason,
		fmt.Sprint(time.Unix(resp.Network.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Network.CreatedAt.Seconds, 0)),
	})
	table.Render()

	return nil
}

func UpdateNetworkV4(token string, spec string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	resp, err := authproxy.Resource.CtlUpdateNetworkV4(token, spec)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Network.Cluster,
		resp.Network.Name,
		resp.Network.Status,
		resp.Network.StatusReason,
		fmt.Sprint(time.Unix(resp.Network.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Network.CreatedAt.Seconds, 0)),
	})
	table.Render()
	return nil
}

func DeleteNetworkV4(networkName string) error {
	authproxy := core.NewAuthproxy(&config.Conf)
	token, err := authproxy.Auth.CtlIssueToken()
	if err != nil {
		return err
	}

	resp, err := authproxy.Resource.CtlDeleteNetworkV4(token.Token, deleteCmdClusterFlag, networkName)
	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Cluster", "Name", "Status", "Status Reason", "Updated At", "Created At"})
	table.Append([]string{
		resp.Network.Cluster,
		resp.Network.Name,
		resp.Network.Status,
		resp.Network.StatusReason,
		fmt.Sprint(time.Unix(resp.Network.UpdatedAt.Seconds, 0)),
		fmt.Sprint(time.Unix(resp.Network.CreatedAt.Seconds, 0)),
	})
	table.Render()
	return nil
}
