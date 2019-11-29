package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
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

var ClustersTable = index_model.Table{
	Name:    "Clusters",
	Kind:    "Table",
	Route:   "",
	Subname: "cluster",
	DataKey: "Clusters",
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name:      "Name",
			IsSearch:  true,
			Link:      "Clusters/:0/Resources/Computes",
			LinkParam: "cluster",
			LinkSync:  true,
			LinkGetQueries: []string{
				"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
		},
		index_model.TableColumn{Name: "Datacenter", IsSearch: true},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
