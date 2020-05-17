package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
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
				base_spec_model.QueryModel{Req: base_spec.GetAllUsers{}, Rep: base_spec.GetAllUsersData{}},
				base_spec_model.QueryModel{Req: base_spec.GetUser{}, Rep: base_spec.GetUserData{}},
			},
		},
		base_spec_model.Api{
			Name:            "HomeProject",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: base_spec.GetUsers{}, Rep: base_spec.GetUsersData{}},
			},
		},
	},
}
