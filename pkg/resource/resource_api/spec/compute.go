package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type Compute struct {
	Region        string
	Cluster       string
	RegionService string
	Name          string
	Kind          string
	Labels        string
	Status        string
	StatusReason  string
	Project       string
	Spec          interface{}
	LinkSpec      string
	IpAddrs       string
	Image         string
	Vcpus         uint
	Memory        uint
	Disk          uint
	UpdatedAt     time.Time
	CreatedAt     time.Time
}

type GetCompute struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type GetComputeData struct {
	Compute Compute
}

type GetComputeConsole struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type GetComputeConsoleData struct {
	Compute Compute
}

type WsComputeConsoleInput struct {
	Bytes []byte
}

type WsComputeConsoleOutput struct {
	Bytes []byte
}

type GetComputes struct {
	Region string `validate:"required"`
}

type GetComputesData struct {
	Computes []Compute
}

type CreateCompute struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateComputeData struct{}

type UpdateCompute struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateComputeData struct{}

type DeleteCompute struct {
	Name string `validate:"required"`
}

type DeleteComputeData struct{}

type DeleteComputes struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteComputesData struct{}

var ComputesTable = map[string]interface{}{
	"Name":        "Computes",
	"Kind":        "Pane",
	"DataQueries": []string{"GetComputes"},
	"Views": []interface{}{
		base_index_model.Table{
			Kind:    "Table",
			DataKey: "Computes",
			SelectActions: []base_index_model.Action{
				base_index_model.Action{
					Name:      "Delete",
					Icon:      "Delete",
					Kind:      "Form",
					DataKind:  "Compute",
					SelectKey: "Name",
				},
			},
			Columns: []base_index_model.TableColumn{
				base_index_model.TableColumn{
					Name: "Name", IsSearch: true,
					Align:      "left",
					LinkPath:   []string{"Regions", "RegionResources", "Clusters", "Resources", "Computes", "Compute", "View"},
					LinkKeyMap: map[string]string{"Name": "Name"},
				},
				base_index_model.TableColumn{Name: "RegionService"},
				base_index_model.TableColumn{Name: "Kind"},
				base_index_model.TableColumn{Name: "Status"},
				base_index_model.TableColumn{Name: "StatusReason"},
				base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
				base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
			},
		},
	},
	"Children": []interface{}{
		ComputesDetail,
	},
}

var ComputesDetail = base_index_model.Tabs{
	Name:   "Compute",
	Kind:   "Tabs",
	IsSync: true,
	Children: []interface{}{
		map[string]interface{}{
			"Name":        "View",
			"Kind":        "Box",
			"DataKey":     "Compute",
			"DataQueries": []string{"GetCompute"},
			"PanelsGroups": []interface{}{
				map[string]interface{}{
					"Name": "Detail",
					"Kind": "Cards",
					"Cards": []interface{}{
						map[string]interface{}{
							"Name": "Detail",
							"Kind": "Fields",
							"Fields": []base_index_model.Field{
								base_index_model.Field{Name: "Name"},
								base_index_model.Field{Name: "RegionService"},
								base_index_model.Field{Name: "Region"},
								base_index_model.Field{Name: "Cluster"},
								base_index_model.Field{Name: "Kind"},
								base_index_model.Field{Name: "Status"},
								base_index_model.Field{Name: "StatusReason"},
								base_index_model.Field{Name: "UpdatedAt"},
							},
						},
						map[string]interface{}{
							"Name":       "Spec",
							"Kind":       "Fields",
							"SubDataKey": "Spec",
							"Fields": []base_index_model.Field{
								base_index_model.Field{Name: "Kind"},
								base_index_model.Field{Name: "Image"},
								base_index_model.Field{Name: "Vcpus"},
								base_index_model.Field{Name: "Memory (MB)", DataKey: "Memory"},
								base_index_model.Field{Name: "Disk (GB)", DataKey: "Disk"},
								base_index_model.Field{Name: "Ips", DataKey: "Ports", Kind: "List",
									ListKey: "Ip",
									Fields: []base_index_model.Field{
										base_index_model.Field{Name: "Subnet"},
										base_index_model.Field{Name: "Gateway"},
									},
								},
							},
						},
					},
				},
			},
		},
		base_index_model.View{
			Name:            "Console",
			Kind:            "Console",
			DataKey:         "Compute",
			EnableWebSocket: true,
			WebSocketQuery:  "GetComputeConsole",
			WebSocketKey:    "ComputeConsole",
			WebSocketKind:   "Console",
		},
	},
}
