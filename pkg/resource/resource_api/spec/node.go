package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type Node struct {
	Name     string
	State    string
	Warnings int
	Errors   int
	Labels   string
}

type GetNodes struct {
	Cluster string `validate:"required"`
}

type GetNodesData struct {
	Nodes []Node
}

type GetNode struct {
	Cluster string `validate:"required"`
}

type GetNodeData struct {
	Node Node
}

var NodesTable = index_model.Table{
	Name:    "Nodes",
	Route:   "/Nodes",
	Kind:    "Table",
	DataKey: "Nodes",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Node",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Clusters/:Cluster/Resources/Nodes/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetNode"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var NodesDetail = index_model.Tabs{
	Name:            "Nodes",
	Kind:            "RouteTabs",
	RouteParamKey:   "Kind",
	RouteParamValue: "Nodes",
	Route:           "/Clusters/:Cluster/Resources/Nodes/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	GetQueries: []string{
		"GetNode",
	},
	ExpectedDataKeys: []string{"Node"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Node",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
	},
}
