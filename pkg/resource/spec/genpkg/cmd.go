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
	"get.floor": index_model.Cmd{
		QueryName: "GetFloor",
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
	"get.floors": index_model.Cmd{
		QueryName:    "GetFloors",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
	},
	"create.floor": index_model.Cmd{
		QueryName: "CreateFloor",
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
	"update.floor": index_model.Cmd{
		QueryName: "UpdateFloor",
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
	"delete.floor": index_model.Cmd{
		QueryName: "DeleteFloor",
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
	"get.rack": index_model.Cmd{
		QueryName: "GetRack",
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
	"get.racks": index_model.Cmd{
		QueryName:    "GetRacks",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Floor,Unit",
	},
	"create.rack": index_model.Cmd{
		QueryName: "CreateRack",
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
	"update.rack": index_model.Cmd{
		QueryName: "UpdateRack",
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
	"delete.rack": index_model.Cmd{
		QueryName: "DeleteRack",
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
	"get.physical.model": index_model.Cmd{
		QueryName: "GetPhysicalModel",
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
	"get.physical.models": index_model.Cmd{
		QueryName:    "GetPhysicalModels",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
	},
	"create.physical.model": index_model.Cmd{
		QueryName: "CreatePhysicalModel",
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
	"update.physical.model": index_model.Cmd{
		QueryName: "UpdatePhysicalModel",
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
	"delete.physical.model": index_model.Cmd{
		QueryName: "DeletePhysicalModel",
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
	"get.physical.resource": index_model.Cmd{
		QueryName: "GetPhysicalResource",
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
	"get.physical.resources": index_model.Cmd{
		QueryName:    "GetPhysicalResources",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,Model,RackPosition,NetLinks,PowerLinks,Spec",
	},
	"create.physical.resource": index_model.Cmd{
		QueryName: "CreatePhysicalResource",
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
	"update.physical.resource": index_model.Cmd{
		QueryName: "UpdatePhysicalResource",
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
	"delete.physical.resource": index_model.Cmd{
		QueryName: "DeletePhysicalResource",
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
		"GetRegion":              base_model.QueryModel{},
		"GetRegions":             base_model.QueryModel{},
		"CreateRegion":           base_model.QueryModel{},
		"UpdateRegion":           base_model.QueryModel{},
		"DeleteRegion":           base_model.QueryModel{},
		"GetDatacenter":          base_model.QueryModel{},
		"GetDatacenters":         base_model.QueryModel{},
		"CreateDatacenter":       base_model.QueryModel{},
		"UpdateDatacenter":       base_model.QueryModel{},
		"DeleteDatacenter":       base_model.QueryModel{},
		"GetFloor":               base_model.QueryModel{},
		"GetFloors":              base_model.QueryModel{},
		"CreateFloor":            base_model.QueryModel{},
		"UpdateFloor":            base_model.QueryModel{},
		"DeleteFloor":            base_model.QueryModel{},
		"GetRack":                base_model.QueryModel{},
		"GetRacks":               base_model.QueryModel{},
		"CreateRack":             base_model.QueryModel{},
		"UpdateRack":             base_model.QueryModel{},
		"DeleteRack":             base_model.QueryModel{},
		"GetPhysicalModel":       base_model.QueryModel{},
		"GetPhysicalModels":      base_model.QueryModel{},
		"CreatePhysicalModel":    base_model.QueryModel{},
		"UpdatePhysicalModel":    base_model.QueryModel{},
		"DeletePhysicalModel":    base_model.QueryModel{},
		"GetPhysicalResource":    base_model.QueryModel{},
		"GetPhysicalResources":   base_model.QueryModel{},
		"CreatePhysicalResource": base_model.QueryModel{},
		"UpdatePhysicalResource": base_model.QueryModel{},
		"DeletePhysicalResource": base_model.QueryModel{},
	},
	"ResourceVirtual": map[string]base_model.QueryModel{
		"GetClusters": base_model.QueryModel{},
	},
}
