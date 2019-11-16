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
		OutputFormat: "Name,State,Warnings,Errors,Labels",
	},
	"get.node": index_model.Cmd{
		QueryName: "GetNode",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
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
			"alerts,a": index_model.Flag{
				Required: false,
				FlagType: "[]spec.ResourceAlert",
				FlagKind: "",
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
		"ReportNode": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
