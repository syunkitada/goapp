// This code is auto generated.
// Don't modify this code.

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
	"get.floor": index_model.Cmd{
		QueryName: "GetFloor",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.floors": index_model.Cmd{
		QueryName: "GetFloors",
		FlagMap: map[string]index_model.Flag{
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
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
		QueryName: "GetPhysicalResources",
		FlagMap: map[string]index_model.Flag{
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
	},
}
var ResourcePhysicalAdminCmdMap = map[string]index_model.Cmd{
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
	"delete.regions": index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"delete.datacenters": index_model.Cmd{
		QueryName: "DeleteDatacenters",
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
	"get.floor": index_model.Cmd{
		QueryName: "GetFloor",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.floors": index_model.Cmd{
		QueryName: "GetFloors",
		FlagMap: map[string]index_model.Flag{
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
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
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.floors": index_model.Cmd{
		QueryName: "DeleteFloors",
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
	"get.rack": index_model.Cmd{
		QueryName: "GetRack",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.racks": index_model.Cmd{
		QueryName: "GetRacks",
		FlagMap: map[string]index_model.Flag{
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
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
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.racks": index_model.Cmd{
		QueryName: "DeleteRacks",
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
	"delete.physical.models": index_model.Cmd{
		QueryName: "DeletePhysicalModels",
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
		QueryName: "GetPhysicalResources",
		FlagMap: map[string]index_model.Flag{
			"datacenter,d": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
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
	"delete.physical.resources": index_model.Cmd{
		QueryName: "DeletePhysicalResources",
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
}
var ResourceVirtualAdminCmdMap = map[string]index_model.Cmd{
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
	"delete.regions": index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"get.cluster": index_model.Cmd{
		QueryName: "GetCluster",
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
	"get.clusters": index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight",
	},
	"create.cluster": index_model.Cmd{
		QueryName: "CreateCluster",
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
	"update.cluster": index_model.Cmd{
		QueryName: "UpdateCluster",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"domain.suffix,d": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"token,t": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"kind,k": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"weight,w": index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"endpoints,e": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.cluster": index_model.Cmd{
		QueryName: "DeleteCluster",
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
	"delete.clusters": index_model.Cmd{
		QueryName: "DeleteClusters",
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
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Name,State,Warnings,Errors,Labels,MetricsGroups",
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
	"get.network.v4": index_model.Cmd{
		QueryName: "GetNetworkV4",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.network.v4s": index_model.Cmd{
		QueryName: "GetNetworkV4s",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Cluster,Subnet,StartIp,EndIp,Gateway",
	},
	"create.network.v4": index_model.Cmd{
		QueryName: "CreateNetworkV4",
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
	"update.network.v4": index_model.Cmd{
		QueryName: "UpdateNetworkV4",
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
	"delete.network.v4": index_model.Cmd{
		QueryName: "DeleteNetworkV4",
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
	"delete.network.v4s": index_model.Cmd{
		QueryName: "DeleteNetworkV4s",
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
	"get.image": index_model.Cmd{
		QueryName: "GetImage",
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
	"get.images": index_model.Cmd{
		QueryName: "GetImages",
		FlagMap: map[string]index_model.Flag{
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,Spec",
	},
	"create.image": index_model.Cmd{
		QueryName: "CreateImage",
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
	"update.image": index_model.Cmd{
		QueryName: "UpdateImage",
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
	"delete.image": index_model.Cmd{
		QueryName: "DeleteImage",
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
	"delete.images": index_model.Cmd{
		QueryName: "DeleteImages",
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
	"get.region.service": index_model.Cmd{
		QueryName: "GetRegionService",
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
	"get.region.services": index_model.Cmd{
		QueryName: "GetRegionServices",
		FlagMap: map[string]index_model.Flag{
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,Cluster,Spec",
	},
	"create.region.service": index_model.Cmd{
		QueryName: "CreateRegionService",
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
	"update.region.service": index_model.Cmd{
		QueryName: "UpdateRegionService",
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
	"delete.region.service": index_model.Cmd{
		QueryName: "DeleteRegionService",
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
	"delete.region.services": index_model.Cmd{
		QueryName: "DeleteRegionServices",
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
}
var ResourceVirtualCmdMap = map[string]index_model.Cmd{
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
	"delete.regions": index_model.Cmd{
		QueryName: "DeleteRegions",
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
	"get.cluster": index_model.Cmd{
		QueryName: "GetCluster",
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
	"get.clusters": index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight",
	},
	"create.cluster": index_model.Cmd{
		QueryName: "CreateCluster",
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
	"update.cluster": index_model.Cmd{
		QueryName: "UpdateCluster",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"region,r": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"datacenter,d": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"domain.suffix,d": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"token,t": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"project,p": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"kind,k": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"weight,w": index_model.Flag{
				Required: false,
				FlagType: "int",
				FlagKind: "",
			},
			"endpoints,e": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.cluster": index_model.Cmd{
		QueryName: "DeleteCluster",
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
	"delete.clusters": index_model.Cmd{
		QueryName: "DeleteClusters",
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
	"get.image": index_model.Cmd{
		QueryName: "GetImage",
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
	"get.images": index_model.Cmd{
		QueryName: "GetImages",
		FlagMap: map[string]index_model.Flag{
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,Spec",
	},
	"create.image": index_model.Cmd{
		QueryName: "CreateImage",
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
	"update.image": index_model.Cmd{
		QueryName: "UpdateImage",
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
	"delete.image": index_model.Cmd{
		QueryName: "DeleteImage",
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
	"delete.images": index_model.Cmd{
		QueryName: "DeleteImages",
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
	"get.region.service": index_model.Cmd{
		QueryName: "GetRegionService",
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
	"get.region.services": index_model.Cmd{
		QueryName: "GetRegionServices",
		FlagMap: map[string]index_model.Flag{
			"region,r": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,Cluster,Spec",
	},
	"create.region.service": index_model.Cmd{
		QueryName: "CreateRegionService",
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
	"update.region.service": index_model.Cmd{
		QueryName: "UpdateRegionService",
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
	"delete.region.service": index_model.Cmd{
		QueryName: "DeleteRegionService",
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
	"delete.region.services": index_model.Cmd{
		QueryName: "DeleteRegionServices",
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
}
var ResourceMonitorCmdMap = map[string]index_model.Cmd{
	"get.clusters": index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight",
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
		OutputFormat: "Name,State,Warnings,Errors,Labels,MetricsGroups",
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
	"get.alerts": index_model.Cmd{
		QueryName: "GetAlerts",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "table",
		OutputFormat: "Name,Time,Level,Handler,Msg,Tag",
	},
	"get.alert.rules": index_model.Cmd{
		QueryName: "GetAlertRules",
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
	"get.statistics": index_model.Cmd{
		QueryName: "GetStatistics",
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
	"get.logs": index_model.Cmd{
		QueryName: "GetLogs",
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
	"get.trace": index_model.Cmd{
		QueryName: "GetTrace",
		FlagMap: map[string]index_model.Flag{
			"cluster,c": index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"trace.id,t": index_model.Flag{
				Required: true,
				FlagType: "string",
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
	"ResourcePhysical": map[string]spec_model.QueryModel{
		"GetRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenter": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloor": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloors": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModel": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModels": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResource": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResources": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourcePhysicalAdmin": map[string]spec_model.QueryModel{
		"GetRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenter": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetDatacenters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateDatacenter": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateDatacenter": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteDatacenter": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteDatacenters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloor": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetFloors": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateFloor": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateFloor": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteFloor": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteFloors": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRack": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRacks": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRack": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRack": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRack": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRacks": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModel": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalModels": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreatePhysicalModel": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdatePhysicalModel": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalModel": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalModels": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResource": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetPhysicalResources": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreatePhysicalResource": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdatePhysicalResource": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalResource": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeletePhysicalResources": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceVirtualAdmin": map[string]spec_model.QueryModel{
		"GetRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetClusters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteClusters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodes": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNodeServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNetworkV4": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetNetworkV4s": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateNetworkV4": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateNetworkV4": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteNetworkV4": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteNetworkV4s": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImages": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImages": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceVirtual": map[string]spec_model.QueryModel{
		"GetRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegion": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegions": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetClusters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteCluster": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteClusters": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetImages": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImage": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteImages": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetRegionServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"CreateRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"UpdateRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionService": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"DeleteRegionServices": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
	"ResourceMonitor": map[string]spec_model.QueryModel{
		"GetClusters": spec_model.QueryModel{
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
		"GetAlerts": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetAlertRules": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetStatistics": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetLogs": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
		"GetTrace": spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
