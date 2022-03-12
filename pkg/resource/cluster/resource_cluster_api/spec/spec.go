package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"

	api_spec "github.com/syunkitada/goapp/pkg/resource/resource_api/spec"
)

type Meta struct{}

var Spec = base_spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceClusterApi",
	Kind: base_const.KindApi,
	Apis: []base_spec_model.Api{
		base_spec_model.Api{
			Name:            "ResourceVirtualAdmin",
			RequiredAuth:    true,
			RequiredProject: true,
			QueryModels: []base_spec_model.QueryModel{
				base_spec_model.QueryModel{Req: api_spec.GetCompute{}, Rep: api_spec.GetComputeData{}},
				base_spec_model.QueryModel{Req: api_spec.GetComputes{}, Rep: api_spec.GetComputesData{}},
				base_spec_model.QueryModel{Req: api_spec.CreateCompute{}, Rep: api_spec.CreateComputeData{}},
				base_spec_model.QueryModel{Req: api_spec.UpdateCompute{}, Rep: api_spec.UpdateComputeData{}},
				base_spec_model.QueryModel{Req: api_spec.DeleteCompute{}, Rep: api_spec.DeleteComputeData{}},
				base_spec_model.QueryModel{Req: api_spec.DeleteComputes{}, Rep: api_spec.DeleteComputesData{}},
				base_spec_model.QueryModel{Req: api_spec.GetComputeConsole{}, Rep: api_spec.GetComputeConsoleData{}, Ws: true},

				base_spec_model.QueryModel{Req: api_spec.GetNodeServices{}, Rep: api_spec.GetNodeServicesData{}},
				base_spec_model.QueryModel{Req: api_spec.SyncNodeService{}, Rep: api_spec.SyncNodeServiceData{}},
				base_spec_model.QueryModel{Req: api_spec.ReportNodeServiceTask{}, Rep: api_spec.ReportNodeServiceTaskData{}},

				base_spec_model.QueryModel{Req: api_spec.GetNodes{}, Rep: api_spec.GetNodesData{}},
				base_spec_model.QueryModel{Req: api_spec.GetNode{}, Rep: api_spec.GetNodeData{}},
				base_spec_model.QueryModel{Req: api_spec.GetNodeMetrics{}, Rep: api_spec.GetNodeMetricsData{}},
				base_spec_model.QueryModel{Req: api_spec.ReportNode{}, Rep: api_spec.ReportNodeData{}},

				base_spec_model.QueryModel{Req: api_spec.GetLogs{}, Rep: api_spec.GetLogsData{}},
				base_spec_model.QueryModel{Req: api_spec.GetLogParams{}, Rep: api_spec.GetLogParamsData{}},
				base_spec_model.QueryModel{Req: api_spec.GetEvents{}, Rep: api_spec.GetEventsData{}},

				base_spec_model.QueryModel{Req: api_spec.GetEventRule{}, Rep: api_spec.GetEventRuleData{}},
				base_spec_model.QueryModel{Req: api_spec.GetEventRules{}, Rep: api_spec.GetEventRulesData{}},
				base_spec_model.QueryModel{Req: api_spec.CreateEventRules{}, Rep: api_spec.CreateEventRulesData{}},
				base_spec_model.QueryModel{Req: api_spec.UpdateEventRules{}, Rep: api_spec.UpdateEventRulesData{}},
				base_spec_model.QueryModel{Req: api_spec.DeleteEventRules{}, Rep: api_spec.DeleteEventRulesData{}},
			},
		},
	},
}
