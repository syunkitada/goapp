package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const ClusterKind = "Cluster"

type Cluster struct {
	gorm.Model
	Datacenter   string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:50;unique_index;"`
	Kind         string `gorm:"not null;size:25;"`
	Description  string `gorm:"not null;size:200;"`
	DomainSuffix string `gorm:"not null;size:255;unique;"`
	Spec         string `gorm:"not null;size:1000;"`
}

type ClusterSpec struct {
	Kind         string `validate:"required"`
	Name         string `validate:"required"`
	Description  string
	Datacenter   string `validate:"required"`
	DomainSuffix string `validate:"required"`
	Spec         interface{}
}

var ClusterCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"CreateCluster": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: ClusterKind,
		Help:    "helptext",
	},
	"UpdateCluster": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: ClusterKind,
		Help:    "helptext",
	},
	"GetClusters": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     ClusterKind,
		Help:        "helptext",
		TableHeader: []string{"Name", "Kind", "Datacenter", "DomainSuffix"},
	},
	"GetCluster": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: ClusterKind,
		Help:    "helptext",
	},
	"DeleteCluster": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: ClusterKind,
		Help:    "helptext",
	},
}

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
