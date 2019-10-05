package db_api

import (
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

	api.MustOpen()
	defer api.MustClose()

	if err = api.DB.AutoMigrate(&db_model.Region{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.Datacenter{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.Floor{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.Rack{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.PhysicalModel{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.PhysicalResource{}).Error; err != nil {
		return err
	}

	if err = api.DB.AutoMigrate(&db_model.Image{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.RegionService{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.NetworkV4{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.Cluster{}).Error; err != nil {
		return err
	}
	if err = api.DB.AutoMigrate(&db_model.Compute{}).Error; err != nil {
		return err
	}

	return nil
}
