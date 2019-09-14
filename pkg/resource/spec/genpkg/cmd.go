package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

var ResourcePhysicalCmdMap = map[string]index_model.Cmd{
	"get.region": index_model.Cmd{
		QueryName: "GetRegion",
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
	"get.regions": index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Name,Kind",
	},
	"create.region": index_model.Cmd{
		QueryName: "CreateRegion",
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
	"update.region": index_model.Cmd{
		QueryName: "UpdateRegion",
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
	"delete.region": index_model.Cmd{
		QueryName: "DeleteRegion",
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
	"get.datacenter": index_model.Cmd{
		QueryName: "GetDatacenter",
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
	"get.datacenters": index_model.Cmd{
		QueryName:    "GetDatacenters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix",
	},
	"create.datacenter": index_model.Cmd{
		QueryName: "CreateDatacenter",
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
	"update.datacenter": index_model.Cmd{
		QueryName: "UpdateDatacenter",
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
	"delete.datacenter": index_model.Cmd{
		QueryName: "DeleteDatacenter",
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
}
var ResourceVirtualCmdMap = map[string]index_model.Cmd{
	"get.clusters": index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Weight",
	},
}

var ApiQueryMap = map[string]map[string]base_model.QueryModel{
	"Auth": map[string]base_model.QueryModel{
		"Login":         base_model.QueryModel{},
		"UpdateService": base_model.QueryModel{},
	},
	"ResourcePhysical": map[string]base_model.QueryModel{
		"GetRegion":        base_model.QueryModel{},
		"GetRegions":       base_model.QueryModel{},
		"CreateRegion":     base_model.QueryModel{},
		"UpdateRegion":     base_model.QueryModel{},
		"DeleteRegion":     base_model.QueryModel{},
		"GetDatacenter":    base_model.QueryModel{},
		"GetDatacenters":   base_model.QueryModel{},
		"CreateDatacenter": base_model.QueryModel{},
		"UpdateDatacenter": base_model.QueryModel{},
		"DeleteDatacenter": base_model.QueryModel{},
	},
	"ResourceVirtual": map[string]base_model.QueryModel{
		"GetClusters": base_model.QueryModel{},
	},
}
