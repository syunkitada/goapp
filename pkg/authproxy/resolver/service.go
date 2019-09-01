package resolver

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/spec"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) UpdateService(tctx *logger.TraceContext, db *gorm.DB, input *spec.UpdateService) (data *spec.UpdateServiceData, code uint8, err error) {
	return
}

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *spec.GetServiceIndex) (data *spec.GetServiceIndexData, code uint8, err error) {
	cmdMap := map[string]index_model.Cmd{}
	cmdMaps := []map[string]index_model.Cmd{
		spec.UserCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	code = base_const.CodeOk
	data = &spec.GetServiceIndexData{
		Index: index_model.Index{
			CmdMap: cmdMap,
		},
	}

	return
}
