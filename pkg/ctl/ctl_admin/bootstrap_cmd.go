package ctl_admin

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/resource/cluster/resource_cluster_model/resource_cluster_model_api"
)

type DatabaseConnection struct {
	user     string
	password string
	host     string
	port     string
	database string
}

var (
	reDatabaseConnection *regexp.Regexp = regexp.MustCompile(`^(\S+):(\S+)@.*\((\S+):(\S+)\)/(\S+)\?.*$`)
)

func init() {
	RootCmd.AddCommand(BootstrapCmd)
	RootCmd.AddCommand(RebootstrapCmd)
}

var BootstrapCmd = &cobra.Command{
	Use:   "bootstrap",
	Short: "bootstrap database",
	Long: `bootstrap database
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewCtl(&config.Conf)
		if err := ctl.Bootstrap(false); err != nil {
			glog.Fatal(err)
		}
	},
}

var RebootstrapCmd = &cobra.Command{
	Use:   "rebootstrap",
	Short: "rebootstrap database",
	Long: `rebootstrap database
	`,
	Run: func(cmd *cobra.Command, args []string) {
		ctl := NewCtl(&config.Conf)
		if err := ctl.Bootstrap(true); err != nil {
			glog.Fatal(err)
		}
	},
}

func (ctl *Ctl) Bootstrap(isRecreate bool) error {
	if err := ctl.CreateDatabases(isRecreate); err != nil {
		return err
	}

	if err := ctl.AuthproxyModelApi.Bootstrap(); err != nil {
		return err
	}

	if err := ctl.ResourceModelApi.Bootstrap(); err != nil {
		return err
	}

	for clusterName, _ := range ctl.Conf.Resource.ClusterMap {
		clusterConf := *ctl.Conf
		clusterConf.Resource.Cluster.Name = clusterName
		resourceClusterModelApi := resource_cluster_model_api.NewResourceClusterModelApi(&clusterConf)
		resourceClusterModelApi.Bootstrap()
	}

	return nil
}

func (ctl *Ctl) CreateDatabases(isRecreate bool) error {
	databaseConnections := []string{
		ctl.Conf.Authproxy.Database.Connection,
		ctl.Conf.Resource.Database.Connection,
	}

	for _, cluster := range ctl.Conf.Resource.ClusterMap {
		databaseConnections = append(databaseConnections, cluster.Database.Connection)
	}

	for _, conn := range databaseConnections {
		if err := ctl.CreateDatabase(isRecreate, conn); err != nil {
			return err
		}
	}
	return nil
}

func (ctl *Ctl) CreateDatabase(isRecreate bool, connection string) error {
	conn, connErr := ctl.ParseDatabaseConnection(connection)
	if connErr != nil {
		return connErr
	}

	if isRecreate {
		if !ctl.Conf.Default.EnableDevelop {
			return fmt.Errorf("Recreate database is not available on except develop mode")
		}
		if err := ctl.ExecMysql(conn, "drop database if exists "+conn.database); err != nil {
			return err
		}
	}

	if err := ctl.ExecMysql(conn, "create database if not exists "+conn.database); err != nil {
		return err
	}

	return nil
}

func (ctl *Ctl) ParseDatabaseConnection(connection string) (*DatabaseConnection, error) {
	result := reDatabaseConnection.FindStringSubmatch(connection)
	if result == nil {
		return nil, fmt.Errorf("Invalid connection: %v", connection)
	}

	return &DatabaseConnection{
		user:     result[1],
		password: result[2],
		host:     result[3],
		port:     result[4],
		database: result[5],
	}, nil
}

func (ctl *Ctl) ExecMysql(conn *DatabaseConnection, sql string) error {
	if ctl.Conf.Default.EnableDatabaseLog {
		glog.Infof("Exec mysql -u%v -pxxx -h%v -P%v -e '%v'", conn.user, conn.host, conn.port, sql)
	}
	if out, err := exec.Command(
		"mysql", "-u"+conn.user, "-p"+conn.password, "-h"+conn.host, "-P"+conn.port,
		"-e", sql).Output(); err != nil {
		return err
	} else {
		if ctl.Conf.Default.EnableDatabaseLog {
			glog.Infof("Exec mysql success: stdout=%v", out)
		}
	}

	return nil
}
