package spec

import "github.com/syunkitada/goapp/pkg/base/base_model"

type Meta struct{}

var Spec = base_model.Spec{
	Meta: Meta{},
	Name: "Resource",
	Apis: []base_model.Api{
		// base_model.Api{
		// 	Name:            "Auth",
		// 	RequiredAuth:    false,
		// 	RequiredProject: false,
		// 	QueryModels: []base_model.QueryModel{
		// 		base_model.QueryModel{
		// 			Model: UpdateService{},
		// 		},
		// 		base_model.QueryModel{
		// 			Model: Login{},
		// 		},
		// 	},
		// },
		base_model.Api{
			Name:            "ResourcePhysical",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model: GetRegions{},
				},
			},
		},
		base_model.Api{
			Name:            "ResourceVirtual",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_model.QueryModel{
				base_model.QueryModel{
					Model: GetClusters{},
				},
			},
		},
	},
}
