package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type Meta struct{}

var Spec = base_model.Spec{
	Meta: Meta{},
	Name: "Authproxy",
	Apis: []base_model.Api{
		base_model.Api{
			Name:            "Home",
			RequiredAuth:    true,
			RequiredProject: false,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{Req: base_spec.GetAllUsers{}, Rep: base_spec.GetAllUsersData{}},
				base_model.QueryModel{Req: base_spec.GetUser{}, Rep: base_spec.GetUserData{}},
			},
		},
		base_model.Api{
			Name:            "HomeProject",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{Req: base_spec.GetUsers{}, Rep: base_spec.GetUsersData{}},
			},
		},
	},
}
