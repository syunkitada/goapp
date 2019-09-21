package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

type Meta struct{}

var Spec = spec_model.Spec{
	Meta: Meta{},
	Name: "Authproxy",
	Apis: []spec_model.Api{
		spec_model.Api{
			Name:            "Home",
			RequiredAuth:    true,
			RequiredProject: false,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: base_spec.GetAllUsers{}, Rep: base_spec.GetAllUsersData{}},
				spec_model.QueryModel{Req: base_spec.GetUser{}, Rep: base_spec.GetUserData{}},
			},
		},
		spec_model.Api{
			Name:            "HomeProject",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: base_spec.GetUsers{}, Rep: base_spec.GetUsersData{}},
			},
		},
	},
}
