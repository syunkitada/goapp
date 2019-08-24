package db_api

import (
	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) Bootstrap(tctx *logger.TraceContext, isRecreate bool) error {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 0) }()
	if err = exec_utils.CreateDatabase(tctx, api.baseConf, api.databaseConf.Connection, isRecreate); err != nil {
		return err
	}

	return nil
}
