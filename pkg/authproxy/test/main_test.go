package test

import (
	"testing"

	"github.com/golang/glog"
	"github.com/urfave/cli"

	"github.com/syunkitada/goapp/pkg/authproxy/core"
	"github.com/syunkitada/goapp/pkg/config"
)

var (
	Conf = &config.Conf
)

func TestMain(t *testing.T) {
	cli.VersionFlag = config.VersionFlag

	app := cli.NewApp()
	app.Name = "goapp-authproxy"
	app.Usage = "goapp-authproxy"
	app.Version = "0.0.1"
	app.Flags = append(config.CommonFlags, config.GlogFlags...)
	args := []string{app.Name, "--use-pwd", "--test-mode"}

	app.Action = func(c *cli.Context) error {
		config.Init(c)
		authproxy := core.NewAuthproxy(Conf)
		glog.Info(authproxy)
		responseIssueToken := authproxy.Auth.TestIssueToken(t)
		responseLogin := authproxy.Dashboard.TestLogin(t)
		responseGetState := authproxy.Dashboard.TestGetState(t, responseIssueToken)
		glog.Info(responseIssueToken)
		glog.Info(responseLogin)
		glog.Info(responseGetState)

		return nil
	}

	if err := app.Run(args); err != nil {
		glog.Error(err)
	}

	actual := "hello"
	expected := "hello"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}
