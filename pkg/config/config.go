package config

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/urfave/cli"
	"os"
	"path/filepath"
)

var Conf Config

var CommonFlags = []cli.Flag{
	cli.StringFlag{Name: "config-dir", Value: "/etc/goapp", Usage: "config-dir"},
	cli.BoolFlag{Name: "use-testdata", Usage: "use testdata"},
}

var VersionFlag = cli.BoolFlag{Name: "print-version, V", Usage: "print only the version"}

func Init(ctx *cli.Context) error {
	glogGangstaShim(ctx)

	var configDir string
	if ctx.GlobalBool("use-testdata") {
		configDir = os.Getenv("PWD") + "/testdata"
	} else {
		configDir = ctx.GlobalString("config-dir")
	}

	newConfig := newConfig(ctx, configDir)
	_, err := toml.DecodeFile(newConfig.Default.ConfigFile, newConfig)
	if err != nil {
		glog.Errorf("Failed to decode file : %!s(MISSING)", err)
		return err
	}
	Conf = *newConfig

	glog.Infof("Loaded config-file(%v)", ctx.GlobalString("config-file"))
	return nil
}

func (conf *Config) Path(path string) string {
	return filepath.Join(conf.Default.ConfigDir, path)
}
