package db_api

import (
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
)

func (api *Api) BootstrapResource(tctx *logger.TraceContext, isRecreate bool) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 0) }()

	api.MustOpen()
	defer api.MustClose()

	if err = api.DB.AutoMigrate(&db_model.Compute{}).Error; err != nil {
		return
	}

	if err = api.DB.AutoMigrate(&db_model.NodeMeta{}).Error; err != nil {
		return
	}

	if err = api.DB.AutoMigrate(&db_model.ComputeAssignment{}).Error; err != nil {
		return
	}

	return
}
