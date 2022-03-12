// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

var ResourceVirtualAdminCmdMap = map[string]base_index_model.Cmd{
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
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,IpAddrs,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
		Ws:           true,
	},
}

var ApiQueryMap = map[string]map[string]base_spec_model.QueryModel{
	"Auth": map[string]base_spec_model.QueryModel{
		"Login":         base_spec_model.QueryModel{},
		"UpdateService": base_spec_model.QueryModel{},
	},
	"ResourceVirtualAdmin": map[string]base_spec_model.QueryModel{
		"GetComputeConsole": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
