package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
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
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
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
			"datacenter,d": index_model.Flag{
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
	"get.physical.resources": index_model.Cmd{
		QueryName:    "GetPhysicalResources",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,Spec",
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
			"datacenter,d": index_model.Flag{
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
}
var ResourceVirtualCmdMap = map[string]index_model.Cmd{
	"get.clusters": index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Weight",
	},
}

var ApiQueryMap = map[string]map[string]spec_model.QueryModel{
	"Auth": map[string]spec_model.QueryModel{
		"Login":         spec_model.QueryModel{},
		"UpdateService": spec_model.QueryModel{},
	},
	"ResourcePhysical": map[string]spec_model.QueryModel{
		"GetRegion":              spec_model.QueryModel{},
		"GetRegions":             spec_model.QueryModel{},
		"CreateRegion":           spec_model.QueryModel{},
		"UpdateRegion":           spec_model.QueryModel{},
		"DeleteRegion":           spec_model.QueryModel{},
		"GetDatacenter":          spec_model.QueryModel{},
		"GetDatacenters":         spec_model.QueryModel{},
		"CreateDatacenter":       spec_model.QueryModel{},
		"UpdateDatacenter":       spec_model.QueryModel{},
		"DeleteDatacenter":       spec_model.QueryModel{},
		"GetFloor":               spec_model.QueryModel{},
		"GetFloors":              spec_model.QueryModel{},
		"CreateFloor":            spec_model.QueryModel{},
		"UpdateFloor":            spec_model.QueryModel{},
		"DeleteFloor":            spec_model.QueryModel{},
		"GetRack":                spec_model.QueryModel{},
		"GetRacks":               spec_model.QueryModel{},
		"CreateRack":             spec_model.QueryModel{},
		"UpdateRack":             spec_model.QueryModel{},
		"DeleteRack":             spec_model.QueryModel{},
		"GetPhysicalModel":       spec_model.QueryModel{},
		"GetPhysicalModels":      spec_model.QueryModel{},
		"CreatePhysicalModel":    spec_model.QueryModel{},
		"UpdatePhysicalModel":    spec_model.QueryModel{},
		"DeletePhysicalModel":    spec_model.QueryModel{},
		"GetPhysicalResource":    spec_model.QueryModel{},
		"GetPhysicalResources":   spec_model.QueryModel{},
		"CreatePhysicalResource": spec_model.QueryModel{},
		"UpdatePhysicalResource": spec_model.QueryModel{},
		"DeletePhysicalResource": spec_model.QueryModel{},
	},
	"ResourceVirtual": map[string]spec_model.QueryModel{
		"GetClusters": spec_model.QueryModel{},
	},
}
