package genpkg

import (
	"encoding/json"

	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_config"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_db_api"
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

type QueryResolver interface {
	GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (*base_spec.GetServiceIndexData, uint8, error)
	GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (*spec.GetRegionsData, uint8, error)
	GetClusters(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetClusters) (*spec.GetClustersData, uint8, error)
}

type QueryHandler struct {
	baseConf *base_config.Config
	appConf  *base_config.AppConfig
	dbApi    base_db_api.IApi
	resolver QueryResolver
}

func NewQueryHandler(baseConf *base_config.Config, appConf *base_config.AppConfig, dbApi base_db_api.IApi, resolver QueryResolver) *QueryHandler {
	return &QueryHandler{
		baseConf: baseConf,
		appConf:  appConf,
		dbApi:    dbApi,
		resolver: resolver,
	}
}

func (handler *QueryHandler) Exec(tctx *logger.TraceContext, req *base_model.Request, rep *base_model.Response) error {
	var err error
	for _, query := range req.Queries {
		switch query.Name {
		case "GetServiceIndex":
			var input base_spec.GetServiceIndex
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetServiceIndex(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetServiceIndex"] = data
			return err
		case "GetRegions":
			var input spec.GetRegions
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetRegions(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetRegions"] = data
			return err
		case "GetClusters":
			var input spec.GetClusters
			err = json.Unmarshal([]byte(query.Data), &input)
			if err != nil {
				return err
			}

			var db *gorm.DB
			if db, err = handler.dbApi.Open(tctx); err != nil {
				return err
			}
			defer handler.dbApi.Close(tctx, db)

			data, code, err := handler.resolver.GetClusters(tctx, db, &input)
			if err != nil {
				if code == 0 {
					code = base_const.CodeServerInternalError
				}
				rep.Error = err.Error()
			}
			rep.Code = code
			rep.Data["GetClusters"] = data
			return err
		}
	}
	return nil
}
