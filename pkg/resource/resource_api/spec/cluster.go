package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type Cluster struct {
	Region       string `validate:"required"`
	Datacenter   string `validate:"required"`
	Name         string `validate:"required"`
	Kind         string `validate:"required"`
	Description  string
	DomainSuffix string `validate:"required"`
	Labels       string
	Warnings     int
	Criticals    int
	Nodes        int
	Instances    int
	Weight       int
	UpdatedAt    time.Time
	CreatedAt    time.Time
}

type GetCluster struct {
	Name string `validate:"required"`
}

type GetClusterData struct {
	Cluster Cluster
}

type GetClusters struct{}

type GetClustersData struct {
	Clusters []Cluster
}

type CreateCluster struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateClusterData struct{}

type UpdateCluster struct {
	Name         string
	Region       string
	Datacenter   string
	DomainSuffix string
	Token        string
	Project      string
	Kind         string
	Weight       int
	Endpoints    []string
}

type UpdateClusterData struct{}

type DeleteCluster struct {
	Name string `validate:"required"`
}

type DeleteClusterData struct{}

type DeleteClusters struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteClustersData struct{}

var ClustersTable = base_index_model.Table{
	Name:    "Clusters",
	Kind:    "Table",
	Route:   "",
	Subname: "cluster",
	DataKey: "Clusters",
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name:     "Name",
			IsSearch: true,
			Link:     "Clusters/:0/Resources/Computes",
			LinkKey:  "cluster",
			LinkSync: true,
			LinkDataQueries: []string{
				"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
		},
		base_index_model.TableColumn{Name: "Datacenter", IsSearch: true},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var VirtualAdminClustersTable = base_index_model.Table{
	Name:        "Clusters",
	Kind:        "Table",
	Route:       "",
	Subname:     "cluster",
	DataKey:     "Clusters",
	DataQueries: []string{"GetClusters"},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name:            "Name",
			IsSearch:        true,
			Align:           "left",
			Link:            "Regions/:Region/RegionResources/Clusters/:Cluster/Resources/Computes",
			LinkKey:         "Name",
			LinkParam:       "Cluster",
			LinkSync:        true,
			LinkDataQueries: []string{"GetComputes"},
		},
		base_index_model.TableColumn{Name: "Datacenter", IsSearch: true},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{Name: "Delete", Icon: "Delete",
			Kind:      "Form",
			DataKind:  "Region",
			SelectKey: "Name",
		},
	},
}
