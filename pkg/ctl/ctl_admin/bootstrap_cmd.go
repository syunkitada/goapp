package ctl_admin

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/syunkitada/goapp/pkg/config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
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
	reDatabaseConnection = regexp.MustCompile(`^(\S+):(\S+)@.*\((\S+):(\S+)\)/(\S+)\?.*$`)
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
			logger.StdoutFatal(err)
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
			logger.StdoutFatal(err)
		}
	},
}

func (ctl *Ctl) Bootstrap(isRecreate bool) error {
	var err error
	tctx := logger.NewCtlTraceContext(ctl.Name)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	if err = ctl.createDatabases(isRecreate); err != nil {
		return err
	}

	if err = ctl.AuthproxyModelApi.Bootstrap(tctx); err != nil {
		return err
	}

	if err = ctl.ResourceModelApi.Bootstrap(tctx); err != nil {
		return err
	}

	for clusterName := range ctl.Conf.Resource.ClusterMap {
		clusterConf := *ctl.Conf
		clusterConf.Resource.Node.ClusterName = clusterName
		resourceClusterModelApi := resource_cluster_model_api.NewResourceClusterModelApi(&clusterConf)
		if err = resourceClusterModelApi.Bootstrap(tctx); err != nil {
			return err
		}
	}

	return nil
}

func (ctl *Ctl) createDatabases(isRecreate bool) error {
	databaseConnections := []string{
		ctl.Conf.Authproxy.Database.Connection,
		ctl.Conf.Resource.Database.Connection,
	}

	for _, cluster := range ctl.Conf.Resource.ClusterMap {
		fmt.Println("DEBUG connection", cluster.Database.Connection)
		databaseConnections = append(databaseConnections, cluster.Database.Connection)
	}

	for _, conn := range databaseConnections {
		if err := ctl.createDatabase(isRecreate, conn); err != nil {
			return err
		}
	}
	return nil
}

func (ctl *Ctl) createDatabase(isRecreate bool, connection string) error {
	conn, connErr := ctl.ParseDatabaseConnection(connection)
	if connErr != nil {
		return connErr
	}
	fmt.Printf("Creating database: %s\n", conn.database)

	if isRecreate {
		if !ctl.Conf.Default.EnableDevelop {
			return fmt.Errorf("Recreate database is not available on except develop mode")
		}
		if err := ctl.ExecMysql(conn, "drop database if exists "+conn.database); err != nil {
			return err
		}
	}

	if err := ctl.ExecMysql(conn, "create database if not exists "+conn.database+" DEFAULT CHARACTER SET utf8"); err != nil {
		return err
	}
	fmt.Printf("Created database: %s\n", conn.database)

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
	var err error

	if ctl.Conf.Default.EnableDatabaseLog {
		glog.Infof("Exec mysql -u%v -pxxx -h%v -P%v -e '%v'", conn.user, conn.host, conn.port, sql)
	}

	var out []byte
	if out, err = exec.Command(
		"mysql", "-u"+conn.user, "-p"+conn.password, "-h"+conn.host, "-P"+conn.port,
		"-e", sql).Output(); err != nil {
		return err
	}

	if ctl.Conf.Default.EnableDatabaseLog {
		glog.Infof("Exec mysql success: stdout=%v", string(out))
	}

	return nil
}
