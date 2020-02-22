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
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
	},
	"get.datacenters": base_index_model.Cmd{
		QueryName:    "GetDatacenters",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
	},
	"get.physical.models": base_index_model.Cmd{
		QueryName:    "GetPhysicalModels",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
	},
	"get.datacenters": base_index_model.Cmd{
		QueryName:    "GetDatacenters",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Region,DomainSuffix,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Zone,Floor",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Floor,Unit",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Floor,Unit",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
	},
	"get.physical.models": base_index_model.Cmd{
		QueryName:    "GetPhysicalModels",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Unit,Description,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Datacenter,Cluster,Rack,PhysicalModel,RackPosition,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
	},
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Name,Kind,Role,Status,StatusReason,State,StateReason,Labels,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Cluster,Subnet,StartIp,EndIp,Gateway,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Kind,Name,Description,Cluster,Subnet,StartIp,EndIp,Gateway,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
	},
	"delete.network.v4": base_index_model.Cmd{
		QueryName: "DeleteNetworkV4",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Region,Cluster,RegionService,Name,Kind,Labels,Status,StatusReason,Project,Spec,LinkSpec,Image,Vcpus,Memory,Disk,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
	},
	"get.regions": base_index_model.Cmd{
		QueryName:    "GetRegions",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Name,Kind,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
	},
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Labels,Description,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Region,Name,Kind,Status,StatusReason,UpdatedAt,CreatedAt,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
	},
}
var ResourceMonitorCmdMap = map[string]base_index_model.Cmd{
	"get.clusters": base_index_model.Cmd{
		QueryName:    "GetClusters",
		FlagMap:      map[string]base_index_model.Flag{},
		OutputKind:   "table",
		OutputFormat: "Region,Datacenter,Name,Kind,Description,DomainSuffix,Labels,Warnings,Criticals,Nodes,Instances,Weight,UpdatedAt,CreatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
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
		OutputKind:   "table",
		OutputFormat: "Name,DisabledServices,ActiveServices,CriticalServices,DisabledServicesData,ActiveServicesData,CriticalServicesData,SuccessEvents,CriticalEvents,WarningEvents,SilencedEvents,SuccessEventsData,CriticalEventsData,WarningEventsData,SilencedEventsData,MetricsGroups,Labels,UpdatedAt",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "string",
		OutputFormat: "string",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "table",
		OutputFormat: "Check,Level,Project,Node,Msg,ReissueDuration,Silenced,Time",
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
		OutputKind:   "table",
		OutputFormat: "Project,Node,Name,Msg,Check,Level,Kind,Until,Spec",
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
		OutputKind:   "table",
		OutputFormat: "Project,Node,Name,Msg,Check,Level,Kind,Until,Spec",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
		OutputKind:   "",
		OutputFormat: "",
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
