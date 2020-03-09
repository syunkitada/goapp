// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

var ResourceVirtualAdminCmdMap = map[string]base_index_model.Cmd{
	"get.compute": base_index_model.Cmd{
		QueryName: "GetCompute",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.computes": base_index_model.Cmd{
		QueryName: "GetComputes",
		FlagMap: map[string]base_index_model.Flag{
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.compute": base_index_model.Cmd{
		QueryName: "CreateCompute",
		FlagMap: map[string]base_index_model.Flag{
			"spec,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"update.compute": base_index_model.Cmd{
		QueryName: "UpdateCompute",
		FlagMap: map[string]base_index_model.Flag{
			"spec,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.compute": base_index_model.Cmd{
		QueryName: "DeleteCompute",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.computes": base_index_model.Cmd{
		QueryName: "DeleteComputes",
		FlagMap: map[string]base_index_model.Flag{
			"spec,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"get.compute.console": base_index_model.Cmd{
		QueryName: "GetComputeConsole",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
		Ws:           true,
	},
	"get.node.services": base_index_model.Cmd{
		QueryName: "GetNodeServices",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,Role,Status,StatusReason,State,StateReason,Token,Endpoints,Labels,Spec",
		Ws:           false,
	},
	"sync.node.service": base_index_model.Cmd{
		QueryName: "SyncNodeService",
		FlagMap: map[string]base_index_model.Flag{
			"node.service,n": base_index_model.Flag{
				Required: false,
				FlagType: "base_spec.NodeService",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "ComputeAssignments",
		Ws:           false,
	},
	"report.node.service.task": base_index_model.Cmd{
		QueryName: "ReportNodeServiceTask",
		FlagMap: map[string]base_index_model.Flag{
			"compute.assignment.reports,c": base_index_model.Flag{
				Required: false,
				FlagType: "[]spec.AssignmentReport",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"get.nodes": base_index_model.Cmd{
		QueryName: "GetNodes",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
		Ws:           false,
	},
	"get.node": base_index_model.Cmd{
		QueryName: "GetNode",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
		Ws:           false,
	},
	"get.node.metrics": base_index_model.Cmd{
		QueryName: "GetNodeMetrics",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"from.time,f": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"until.time,u": base_index_model.Flag{
				Required: false,
				FlagType: "*time.Time",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
		Ws:           false,
	},
	"report.node": base_index_model.Cmd{
		QueryName: "ReportNode",
		FlagMap: map[string]base_index_model.Flag{
			"project,p": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"state,s": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"warning,w": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"warnings,w": base_index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"error,e": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"errors,e": base_index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"timestate,t": base_index_model.Flag{
				Required: false,
				FlagType: "time.Time",
				FlagKind: "",
			},
			"logs,l": base_index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceLog",
				FlagKind: "",
			},
			"metrics,m": base_index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceMetric",
				FlagKind: "",
			},
			"events,e": base_index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceEvent",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"get.logs": base_index_model.Cmd{
		QueryName: "GetLogs",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"limit.logs,l": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"from.time,f": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"until.time,u": base_index_model.Flag{
				Required: false,
				FlagType: "time.Time",
				FlagKind: "",
			},
			"apps,a": base_index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
			"nodes,n": base_index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
			"trace.id,t": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"get.log.params": base_index_model.Cmd{
		QueryName: "GetLogParams",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "string",
		OutputFormat: "string",
		Ws:           false,
	},
	"get.events": base_index_model.Cmd{
		QueryName: "GetEvents",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Check,Level,Project,Node,Msg,ReissueDuration,Silenced,Time",
		Ws:           false,
	},
	"get.event.rule": base_index_model.Cmd{
		QueryName: "GetEventRule",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Project,Node,Name,Msg,Check,Level,Kind,Until,Spec",
		Ws:           false,
	},
	"get.event.rules": base_index_model.Cmd{
		QueryName: "GetEventRules",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Project,Node,Name,Msg,Check,Level,Kind,Until,Spec",
		Ws:           false,
	},
	"create.event.rules": base_index_model.Cmd{
		QueryName: "CreateEventRules",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"update.event.rules": base_index_model.Cmd{
		QueryName: "UpdateEventRules",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.event.rules": base_index_model.Cmd{
		QueryName: "DeleteEventRules",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
}

var ApiQueryMap = map[string]map[string]base_spec_model.QueryModel{
	"Auth": map[string]base_spec_model.QueryModel{
		"Login":         base_spec_model.QueryModel{},
		"UpdateService": base_spec_model.QueryModel{},
	},
	"ResourceVirtualAdmin": map[string]base_spec_model.QueryModel{
		"GetCompute": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetComputes": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCompute": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCompute": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCompute": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteComputes": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetComputeConsole": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"SyncNodeService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"ReportNodeServiceTask": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodes": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNode": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeMetrics": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"ReportNode": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogs": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogParams": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEvents": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEventRule": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEventRules": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateEventRules": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateEventRules": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteEventRules": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
