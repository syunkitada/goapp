package exec_utils

import (
	"fmt"
	"regexp"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

var reDatabaseConnection = regexp.MustCompile(`^(\S+):(\S+)@.*\((\S+):(\S+)\)/(\S+)\?.*$`)

func CreateDatabase(tctx *logger.TraceContext, baseConf *base_config.Config, connection string, isRecreate bool) error {
	result := reDatabaseConnection.FindStringSubmatch(connection)
	if result == nil {
		return fmt.Errorf("Invalid connection: %v", connection)
	}

	user := result[1]
	password := result[2]
	host := result[3]
	port := result[4]
	database := result[5]

	mysqlCmd := fmt.Sprintf("mysql -u%s -p%s -h%s -P%s", user, password, host, port)
	if isRecreate {
		if !baseConf.EnableDevelop {
			return fmt.Errorf("Recreate database is not available on except develop mode")
		}

		if _, err := Shf(tctx, "%s -e 'drop database if exists %s'", mysqlCmd, database); err != nil {
			return err
		}
	}

	if _, err := Shf(tctx, "%s -e 'create database if not exists %s DEFAULT CHARACTER SET utf8'", mysqlCmd, database); err != nil {
		return err
	}

	return nil
}
