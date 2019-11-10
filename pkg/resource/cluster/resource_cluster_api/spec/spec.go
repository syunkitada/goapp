package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"

	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Meta struct{}

var Spec = spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceClusterApi",
	Kind: base_const.KindApi,
	Apis: []spec_model.Api{
		spec_model.Api{
			Name:            "ResourceVirtualAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []spec_model.QueryModel{
				spec_model.QueryModel{Req: api_spec.GetCompute{}, Rep: api_spec.GetComputeData{}},
				spec_model.QueryModel{Req: api_spec.GetComputes{}, Rep: api_spec.GetComputesData{}},
				spec_model.QueryModel{Req: api_spec.CreateCompute{}, Rep: api_spec.CreateComputeData{}},
				spec_model.QueryModel{Req: api_spec.UpdateCompute{}, Rep: api_spec.UpdateComputeData{}},
				spec_model.QueryModel{Req: api_spec.DeleteCompute{}, Rep: api_spec.DeleteComputeData{}},
				spec_model.QueryModel{Req: api_spec.DeleteComputes{}, Rep: api_spec.DeleteComputesData{}},

				spec_model.QueryModel{Req: api_spec.GetNodeServices{}, Rep: api_spec.GetNodeServicesData{}},
				spec_model.QueryModel{Req: api_spec.SyncNodeService{}, Rep: api_spec.SyncNodeServiceData{}},
				spec_model.QueryModel{Req: api_spec.ReportNodeServiceTask{}, Rep: api_spec.ReportNodeServiceTaskData{}},

				spec_model.QueryModel{Req: api_spec.ReportNode{}, Rep: api_spec.ReportNodeData{}},
			},
		},
	},
}
