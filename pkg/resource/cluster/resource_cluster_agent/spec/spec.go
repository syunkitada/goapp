package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"

	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Meta struct{}

var Spec = base_spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceClusterAgent",
	Kind: base_const.KindAgent,
	Apis: []base_spec_model.Api{
		base_spec_model.Api{
			Name:            "ResourceVirtualAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: api_spec.GetComputeConsole{}, Rep: api_spec.GetComputeConsoleData{}, Ws: true},
			},
		},
	},
}
