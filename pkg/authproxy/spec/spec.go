package spec

import "github.com/syunkitada/goapp/pkg/base/base_model"

type Meta struct{}

var Spec = base_model.Spec{
	Meta: Meta{},
	Name: "authproxy",
	Apis: []base_model.Api{
		base_model.Api{
			Name:            "service",
			RequiredService: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model:        UpdateService{},
					ProjectRoles: []string{"admin", "service"},
				},
			},
		},
		base_model.Api{
			Name: "home",
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model:        GetAllUsers{},
					ProjectRoles: []string{"admin"},
				},
				base_model.QueryModel{
					Model:        GetUser{},
					ProjectRoles: []string{"tenant", "admin"},
				},
			},
		},
		base_model.Api{
			Name:            "home.project",
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model:        GetUsers{},
					ProjectRoles: []string{"tenant"},
				},
			},
		},
	},
}
