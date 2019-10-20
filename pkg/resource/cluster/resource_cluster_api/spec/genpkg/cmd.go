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
			"region,r": index_model.Flag{
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
	"get.nodes": index_model.Cmd{
		QueryName: "GetNodes",
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
	"sync.node": index_model.Cmd{
		QueryName: "SyncNode",
		FlagMap: map[string]index_model.Flag{
			"node,n": index_model.Flag{
				Required: false,
				FlagType: "base_spec.Node",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"report.node.task": index_model.Cmd{
		QueryName: "ReportNodeTask",
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
		"GetNodes": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"SyncNode": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"ReportNodeTask": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
