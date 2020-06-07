package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

type Meta struct{}

var Spec = base_spec_model.Spec{
	Meta: Meta{},
	Name: "HomeApi",
	Kind: base_const.KindApi,
	Apis: []base_spec_model.Api{
		base_spec_model.Api{
			Name:            "Home",
			RequiredAuth:    true,
			RequiredProject: false,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: UpdateUserPassword{}, Rep: UpdateUserPasswordData{}},
			},
		},
		base_spec_model.Api{
			Name:            "HomeProject",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: GetProjectUsers{}, Rep: GetProjectUsersData{}},
			},
		},
	},
}
