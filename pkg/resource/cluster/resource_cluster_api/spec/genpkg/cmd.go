// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
)

var ResourceVirtualAdminCmdMap = map[string]index_model.Cmd{
	"get.compute": index_model.Cmd{
		QueryName: "GetCompute",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.computes": index_model.Cmd{
		QueryName: "GetComputes",
		FlagMap: map[string]index_model.Flag{
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk",
	},
	"create.compute": index_model.Cmd{
		QueryName: "CreateCompute",
		FlagMap: map[string]index_model.Flag{
			"spec,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"update.compute": index_model.Cmd{
		QueryName: "UpdateCompute",
		FlagMap: map[string]index_model.Flag{
			"spec,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.compute": index_model.Cmd{
		QueryName: "DeleteCompute",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.computes": index_model.Cmd{
		QueryName: "DeleteComputes",
		FlagMap: map[string]index_model.Flag{
			"spec,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.node.services": index_model.Cmd{
		QueryName: "GetNodeServices",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Name,Kind,Role,Status,StatusReason,State,StateReason,Labels,Spec",
	},
	"sync.node.service": index_model.Cmd{
		QueryName: "SyncNodeService",
		FlagMap: map[string]index_model.Flag{
			"node.service,n": index_model.Flag{
				Required: false,
				FlagType: "base_spec.NodeService",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"report.node.service.task": index_model.Cmd{
		QueryName: "ReportNodeServiceTask",
		FlagMap: map[string]index_model.Flag{
			"compute.assignment.reports,c": index_model.Flag{
				Required: false,
				FlagType: "[]spec.AssignmentReport",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.nodes": index_model.Cmd{
		QueryName: "GetNodes",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
	},
	"get.node": index_model.Cmd{
		QueryName: "GetNode",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.node.metrics": index_model.Cmd{
		QueryName: "GetNodeMetrics",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"from.time,f": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"until.time,u": index_model.Flag{
				Required: false,
				FlagType: "time.Time",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"report.node": index_model.Cmd{
		QueryName: "ReportNode",
		FlagMap: map[string]index_model.Flag{
			"project,p": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"state,s": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"warning,w": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"warnings,w": index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"error,e": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"errors,e": index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"timestate,t": index_model.Flag{
				Required: false,
				FlagType: "time.Time",
				FlagKind: "",
			},
			"logs,l": index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceLog",
				FlagKind: "",
			},
			"metrics,m": index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceMetric",
				FlagKind: "",
			},
			"events,e": index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceEvent",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.logs": index_model.Cmd{
		QueryName: "GetLogs",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"limit.logs,l": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"from.time,f": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"until.time,u": index_model.Flag{
				Required: false,
				FlagType: "time.Time",
				FlagKind: "",
			},
			"apps,a": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
			"nodes,n": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
			"trace.id,t": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.log.params": index_model.Cmd{
		QueryName: "GetLogParams",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "string",
		OutputFormat: "string",
	},
	"get.events": index_model.Cmd{
		QueryName: "GetEvents",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Check,Level,Project,Node,Msg,ReissueDuration,Silenced,Time",
	},
	"get.event.rule": index_model.Cmd{
		QueryName: "GetEventRule",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.event.rules": index_model.Cmd{
		QueryName: "GetEventRules",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Project,Node,Name,Msg,Check,Level,Kind,Until,Spec",
	},
	"create.event.rules": index_model.Cmd{
		QueryName: "CreateEventRules",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"update.event.rules": index_model.Cmd{
		QueryName: "UpdateEventRules",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.event.rules": index_model.Cmd{
		QueryName: "DeleteEventRules",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"specs,s": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "file",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
}

var ApiQueryMap = map[string]map[string]spec_model.QueryModel{
	"Auth": map[string]spec_model.QueryModel{
		"Login":         spec_model.QueryModel{},
		"UpdateService": spec_model.QueryModel{},
	},
	"ResourceVirtualAdmin": map[string]spec_model.QueryModel{
		"GetCompute": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetComputes": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCompute": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCompute": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCompute": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteComputes": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"SyncNodeService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"ReportNodeServiceTask": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodes": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNode": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeMetrics": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"ReportNode": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogs": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogParams": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEvents": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEventRule": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetEventRules": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateEventRules": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateEventRules": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteEventRules": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
