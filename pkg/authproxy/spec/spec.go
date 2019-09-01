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
			Name:            "Auth",
			RequiredAuth:    false,
			RequiredProject: false,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model: base_spec.UpdateService{},
				},
				base_model.QueryModel{
					Model: base_spec.Login{},
				},
			},
		},
		base_model.Api{
			Name:            "Home",
			RequiredAuth:    true,
			RequiredProject: false,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model:        base_spec.GetAllUsers{},
					ProjectRoles: []string{"admin"},
				},
				base_model.QueryModel{
					Model:        base_spec.GetUser{},
					ProjectRoles: []string{"tenant", "admin"},
				},
			},
		},
		base_model.Api{
			Name:            "Home.Project",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model:        base_spec.GetUsers{},
					ProjectRoles: []string{"tenant"},
				},
			},
		},
	},
}
