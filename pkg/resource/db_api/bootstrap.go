package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/db_model"
)

func (api *Api) BootstrapResource(tctx *logger.TraceContext, isRecreate bool) (err error) {
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 0) }()
	if err = exec_utils.CreateDatabase(tctx, api.baseConf, api.databaseConf.Connection, isRecreate); err != nil {
		return err
	}

	var db *gorm.DB
	db, err = api.Open(tctx)
	if err != nil {
		return err
	}
	defer api.Close(tctx, db)

	if err = db.AutoMigrate(&db_model.Region{}).Error; err != nil {
		return err
	}

	return nil
}
