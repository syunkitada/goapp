package resolver

import (
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec/genpkg"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourcePhysicalCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourceVirtualCmdMap,
			},
		}
		code = base_const.CodeOk
	default:
		code = base_const.CodeClientNotFound
	}

	return
}
