package spec

import "github.com/syunkitada/goapp/pkg/base/base_index_model"

type NetworkV4 struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Cluster     string `validate:"required"`
	Subnet      string `validate:"required"`
	StartIp     string `validate:"required"`
	EndIp       string `validate:"required"`
	Gateway     string `validate:"required"`
	Spec        interface{}
}

type NetworkV4LocalSpec struct {
	Resolvers []Resolver
	Nat       Nat
}

type Resolver struct {
	Resolver string
}

type Nat struct {
	Enable bool
	Ports  string
}

type GetNetworkV4 struct {
	Name    string `validate:"required"`
	Cluster string `validate:"required"`
}

type GetNetworkV4Data struct {
	NetworkV4 NetworkV4
}

type GetNetworkV4s struct {
	Cluster string `validate:"required"`
}

type GetNetworkV4sData struct {
	NetworkV4s []NetworkV4
}

type CreateNetworkV4 struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateNetworkV4Data struct{}

type UpdateNetworkV4 struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateNetworkV4Data struct{}

type DeleteNetworkV4 struct {
	Name    string `validate:"required"`
	Cluster string `validate:"required"`
}

type DeleteNetworkV4Data struct{}

type DeleteNetworkV4s struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteNetworkV4sData struct{}

type Network struct {
	Id           uint
	Name         string
	Subnet       string
	Gateway      string
	AvailableIps []string
	Kind         string
	Spec         string
}

var NetworkV4sTable = base_index_model.Table{
	Name:    "NetworkV4s",
	Route:   "/NetworkV4s",
	Kind:    "Table",
	DataKey: "NetworkV4s",
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "NetworkV4",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:            "Regions/:Region/Resources/NetworkV4s/Detail/:0/View",
			LinkKey:         "Name",
			LinkSync:        false,
			LinkDataQueries: []string{"GetNetworkV4"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
