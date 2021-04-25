// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

var ResourcePhysicalCmdMap = map[string]base_index_model.Cmd{
	"get.region": base_index_model.Cmd{
		QueryName: "GetRegion",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.datacenter": base_index_model.Cmd{
		QueryName: "GetDatacenter",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.datacenters": base_index_model.Cmd{
		QueryName:    "GetDatacenters",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.floor": base_index_model.Cmd{
		QueryName: "GetFloor",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
		Ws:           false,
	},
	"get.floors": base_index_model.Cmd{
		QueryName: "GetFloors",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
		Ws:           false,
	},
	"get.physical.model": base_index_model.Cmd{
		QueryName: "GetPhysicalModel",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
		Ws:           false,
	},
	"get.physical.models": base_index_model.Cmd{
		QueryName:    "GetPhysicalModels",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
		Ws:           false,
	},
	"get.physical.resource": base_index_model.Cmd{
		QueryName: "GetPhysicalResource",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
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
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.physical.resources": base_index_model.Cmd{
		QueryName: "GetPhysicalResources",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
}
var ResourcePhysicalAdminCmdMap = map[string]base_index_model.Cmd{
	"get.region": base_index_model.Cmd{
		QueryName: "GetRegion",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.region": base_index_model.Cmd{
		QueryName: "CreateRegion",
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
	"update.region": base_index_model.Cmd{
		QueryName: "UpdateRegion",
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
	"delete.region": base_index_model.Cmd{
		QueryName: "DeleteRegion",
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
	"delete.regions": base_index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"get.datacenter": base_index_model.Cmd{
		QueryName: "GetDatacenter",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.datacenters": base_index_model.Cmd{
		QueryName:    "GetDatacenters",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.datacenter": base_index_model.Cmd{
		QueryName: "CreateDatacenter",
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
	"update.datacenter": base_index_model.Cmd{
		QueryName: "UpdateDatacenter",
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
	"delete.datacenter": base_index_model.Cmd{
		QueryName: "DeleteDatacenter",
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
	"delete.datacenters": base_index_model.Cmd{
		QueryName: "DeleteDatacenters",
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
	"get.floor": base_index_model.Cmd{
		QueryName: "GetFloor",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
		Ws:           false,
	},
	"get.floors": base_index_model.Cmd{
		QueryName: "GetFloors",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
		Ws:           false,
	},
	"create.floor": base_index_model.Cmd{
		QueryName: "CreateFloor",
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
	"update.floor": base_index_model.Cmd{
		QueryName: "UpdateFloor",
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
	"delete.floor": base_index_model.Cmd{
		QueryName: "DeleteFloor",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
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
	"delete.floors": base_index_model.Cmd{
		QueryName: "DeleteFloors",
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
	"get.rack": base_index_model.Cmd{
		QueryName: "GetRack",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Floor,Unit",
		Ws:           false,
	},
	"get.racks": base_index_model.Cmd{
		QueryName: "GetRacks",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Floor,Unit",
		Ws:           false,
	},
	"create.rack": base_index_model.Cmd{
		QueryName: "CreateRack",
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
	"update.rack": base_index_model.Cmd{
		QueryName: "UpdateRack",
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
	"delete.rack": base_index_model.Cmd{
		QueryName: "DeleteRack",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
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
	"delete.racks": base_index_model.Cmd{
		QueryName: "DeleteRacks",
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
	"get.physical.model": base_index_model.Cmd{
		QueryName: "GetPhysicalModel",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
		Ws:           false,
	},
	"get.physical.models": base_index_model.Cmd{
		QueryName:    "GetPhysicalModels",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
		Ws:           false,
	},
	"create.physical.model": base_index_model.Cmd{
		QueryName: "CreatePhysicalModel",
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
	"update.physical.model": base_index_model.Cmd{
		QueryName: "UpdatePhysicalModel",
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
	"delete.physical.model": base_index_model.Cmd{
		QueryName: "DeletePhysicalModel",
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
	"delete.physical.models": base_index_model.Cmd{
		QueryName: "DeletePhysicalModels",
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
	"get.physical.resource": base_index_model.Cmd{
		QueryName: "GetPhysicalResource",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
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
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.physical.resources": base_index_model.Cmd{
		QueryName: "GetPhysicalResources",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"create.physical.resource": base_index_model.Cmd{
		QueryName: "CreatePhysicalResource",
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
	"update.physical.resource": base_index_model.Cmd{
		QueryName: "UpdatePhysicalResource",
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
	"delete.physical.resource": base_index_model.Cmd{
		QueryName: "DeletePhysicalResource",
		FlagMap: map[string]base_index_model.Flag{
			"datacenter,d": base_index_model.Flag{
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
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.physical.resources": base_index_model.Cmd{
		QueryName: "DeletePhysicalResources",
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
}
var ResourceVirtualAdminCmdMap = map[string]base_index_model.Cmd{
	"get.region": base_index_model.Cmd{
		QueryName: "GetRegion",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.region": base_index_model.Cmd{
		QueryName: "CreateRegion",
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
	"update.region": base_index_model.Cmd{
		QueryName: "UpdateRegion",
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
	"delete.region": base_index_model.Cmd{
		QueryName: "DeleteRegion",
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
	"delete.regions": base_index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"get.cluster": base_index_model.Cmd{
		QueryName: "GetCluster",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.cluster": base_index_model.Cmd{
		QueryName: "CreateCluster",
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
	"update.cluster": base_index_model.Cmd{
		QueryName: "UpdateCluster",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"domain.suffix,d": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"token,t": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"kind,k": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"weight,w": base_index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"endpoints,e": base_index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.cluster": base_index_model.Cmd{
		QueryName: "DeleteCluster",
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
	"delete.clusters": base_index_model.Cmd{
		QueryName: "DeleteClusters",
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
	"get.network.v4": base_index_model.Cmd{
		QueryName: "GetNetworkV4",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Cluster,Subnet,StartIp,EndIp,Gateway,Spec",
		Ws:           false,
	},
	"get.network.v4s": base_index_model.Cmd{
		QueryName: "GetNetworkV4s",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Cluster,Subnet,StartIp,EndIp,Gateway,Spec",
		Ws:           false,
	},
	"create.network.v4": base_index_model.Cmd{
		QueryName: "CreateNetworkV4",
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
	"update.network.v4": base_index_model.Cmd{
		QueryName: "UpdateNetworkV4",
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
	"delete.network.v4": base_index_model.Cmd{
		QueryName: "DeleteNetworkV4",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"cluster,c": base_index_model.Flag{
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
	"delete.network.v4s": base_index_model.Cmd{
		QueryName: "DeleteNetworkV4s",
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
	"get.image": base_index_model.Cmd{
		QueryName: "GetImage",
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
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.images": base_index_model.Cmd{
		QueryName: "GetImages",
		FlagMap: map[string]base_index_model.Flag{
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"create.image": base_index_model.Cmd{
		QueryName: "CreateImage",
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
	"update.image": base_index_model.Cmd{
		QueryName: "UpdateImage",
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
	"delete.image": base_index_model.Cmd{
		QueryName: "DeleteImage",
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
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.images": base_index_model.Cmd{
		QueryName: "DeleteImages",
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
	"get.region.service": base_index_model.Cmd{
		QueryName: "GetRegionService",
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
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.region.services": base_index_model.Cmd{
		QueryName: "GetRegionServices",
		FlagMap: map[string]base_index_model.Flag{
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"create.region.service": base_index_model.Cmd{
		QueryName: "CreateRegionService",
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
	"update.region.service": base_index_model.Cmd{
		QueryName: "UpdateRegionService",
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
	"delete.region.service": base_index_model.Cmd{
		QueryName: "DeleteRegionService",
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
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.region.services": base_index_model.Cmd{
		QueryName: "DeleteRegionServices",
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
		Kind:         "Terminal",
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
		Ws:           true,
	},
}
var ResourceVirtualCmdMap = map[string]base_index_model.Cmd{
	"get.region": base_index_model.Cmd{
		QueryName: "GetRegion",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.region": base_index_model.Cmd{
		QueryName: "CreateRegion",
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
	"update.region": base_index_model.Cmd{
		QueryName: "UpdateRegion",
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
	"delete.region": base_index_model.Cmd{
		QueryName: "DeleteRegion",
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
	"delete.regions": base_index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"get.cluster": base_index_model.Cmd{
		QueryName: "GetCluster",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
		Ws:           false,
	},
	"create.cluster": base_index_model.Cmd{
		QueryName: "CreateCluster",
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
	"update.cluster": base_index_model.Cmd{
		QueryName: "UpdateCluster",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"domain.suffix,d": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"token,t": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"kind,k": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"weight,w": base_index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"endpoints,e": base_index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.cluster": base_index_model.Cmd{
		QueryName: "DeleteCluster",
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
	"delete.clusters": base_index_model.Cmd{
		QueryName: "DeleteClusters",
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
	"get.image": base_index_model.Cmd{
		QueryName: "GetImage",
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
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.images": base_index_model.Cmd{
		QueryName: "GetImages",
		FlagMap: map[string]base_index_model.Flag{
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"create.image": base_index_model.Cmd{
		QueryName: "CreateImage",
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
	"update.image": base_index_model.Cmd{
		QueryName: "UpdateImage",
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
	"delete.image": base_index_model.Cmd{
		QueryName: "DeleteImage",
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
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.images": base_index_model.Cmd{
		QueryName: "DeleteImages",
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
	"get.region.service": base_index_model.Cmd{
		QueryName: "GetRegionService",
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
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"get.region.services": base_index_model.Cmd{
		QueryName: "GetRegionServices",
		FlagMap: map[string]base_index_model.Flag{
			"region,r": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
		Ws:           false,
	},
	"create.region.service": base_index_model.Cmd{
		QueryName: "CreateRegionService",
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
	"update.region.service": base_index_model.Cmd{
		QueryName: "UpdateRegionService",
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
	"delete.region.service": base_index_model.Cmd{
		QueryName: "DeleteRegionService",
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
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"delete.region.services": base_index_model.Cmd{
		QueryName: "DeleteRegionServices",
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
}
var ResourceMonitorCmdMap = map[string]base_index_model.Cmd{
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
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
			"target,t": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"time.duration,t": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"until.time,u": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
		Ws:           false,
	},
	"get.statistics": base_index_model.Cmd{
		QueryName: "GetStatistics",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
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
	"get.trace": base_index_model.Cmd{
		QueryName: "GetTrace",
		FlagMap: map[string]base_index_model.Flag{
			"cluster,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"trace.id,t": base_index_model.Flag{
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
	"ResourcePhysical": map[string]base_spec_model.QueryModel{
		"GetRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenter": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloor": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloors": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModel": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModels": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResource": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResources": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourcePhysicalAdmin": map[string]base_spec_model.QueryModel{
		"GetRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenter": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateDatacenter": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateDatacenter": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteDatacenter": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteDatacenters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloor": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloors": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateFloor": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateFloor": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteFloor": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteFloors": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRack": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRacks": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRack": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRack": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRack": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRacks": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModel": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModels": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreatePhysicalModel": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdatePhysicalModel": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalModel": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalModels": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResource": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResources": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreatePhysicalResource": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdatePhysicalResource": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalResource": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalResources": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceVirtualAdmin": map[string]base_spec_model.QueryModel{
		"GetRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetClusters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteClusters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodes": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNetworkV4": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNetworkV4s": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateNetworkV4": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateNetworkV4": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteNetworkV4": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteNetworkV4s": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImages": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImages": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetCompute": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetComputes": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetComputeConsole": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceVirtual": map[string]base_spec_model.QueryModel{
		"GetRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetClusters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCluster": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteClusters": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImages": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImage": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImages": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionService": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionServices": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceMonitor": map[string]base_spec_model.QueryModel{
		"GetClusters": base_spec_model.QueryModel{
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
		"GetStatistics": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogParams": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogs": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetTrace": base_spec_model.QueryModel{
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
