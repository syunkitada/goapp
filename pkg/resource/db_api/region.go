package db_api

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (api *Api) GetRegions(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetRegions) (data *spec.GetRegionsData, err error) {
	var regions []spec.Region
	if err = db.Find(&regions).Error; err != nil {
		return
	}

	data = &spec.GetRegionsData{Regions: regions}
	return
}
