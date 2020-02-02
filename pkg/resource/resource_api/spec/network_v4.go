package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type NetworkV4 struct {
	Kind        string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Cluster     string `validate:"required"`
	Subnet      string `validate:"required"`
	StartIp     string `validate:"required"`
	EndIp       string `validate:"required"`
	Gateway     string `validate:"required"`
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
	Name   string `validate:"required"`
	Region string `validate:"required"`
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
}

var NetworkV4sTable = index_model.Table{
	Name:    "NetworkV4s",
	Route:   "/NetworkV4s",
	Kind:    "Table",
	DataKey: "NetworkV4s",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "NetworkV4",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Regions/:Region/Resources/NetworkV4s/Detail/:0/View",
			LinkKey:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetNetworkV4"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
