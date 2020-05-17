package db_api

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (api *Api) BootstrapHome(tctx *logger.TraceContext, isRecreate bool) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 0) }()

	api.MustOpen()
	defer api.MustClose()

	return nil
}
