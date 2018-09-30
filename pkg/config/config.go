package config

import (
	"github.com/BurntSushi/toml"
	"github.com/golang/glog"
	"github.com/urfave/cli"
	"path/filepath"
)

var Conf Config

var CommonFlags = []cli.Flag{
	cli.StringFlag{Name: "config-dir", Value: "/etc/goapp", Usage: "config-dir"},
	cli.BoolFlag{Name: "use-pwd", Usage: "use PWD"},
	cli.BoolFlag{Name: "test-mode", Usage: "use test-mode"},
}

var VersionFlag = cli.BoolFlag{Name: "print-version, V", Usage: "print only the version"}

func Init(ctx *cli.Context) error {
	glogGangstaShim(ctx)

	newConfig := newConfig(ctx)
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
