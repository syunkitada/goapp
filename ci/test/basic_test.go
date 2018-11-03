package test

import (
	"testing"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
)

func TestBasic(t *testing.T) {
	var testCommand = &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			authproxy := core.NewAuthproxy(&config.Conf)
			glog.Info(authproxy)
			responseIssueToken := authproxy.Auth.TestIssueToken(t)
			responseLogin := authproxy.Dashboard.TestLogin(t)
			responseGetState := authproxy.Dashboard.TestGetState(t, responseIssueToken)
			glog.Info(responseIssueToken)
			glog.Info(responseLogin)
			glog.Info(responseGetState)
		},
	}

	cobra.OnInitialize(config.InitConfig)
	config.InitFlags(testCommand)

	if err := testCommand.Execute(); err != nil {
		t.Fatal(err)
	}
}
