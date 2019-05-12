package resource_model_api

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_api/resource_api_grpc_pb"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) GetPhysicalIndex() interface{} {

	return index_model.Panels{
		Name:      "Root",
		Kind:      "RoutePanels",
		SyncDelay: 20000,
		Panels: []interface{}{
			index_model.Table{
				Name:    "Datacenters",
				Kind:    "Table",
				Route:   "",
				Subname: "datacenter",
				DataKey: "Datacenters",
				Columns: []index_model.TableColumn{
					index_model.TableColumn{
						Name:      "Name",
						IsSearch:  true,
						Link:      "Datacenters/:0/Resources/Resources",
						LinkParam: "datacenter",
						LinkSync:  true,
						LinkGetQueries: []string{
							"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
					},
					index_model.TableColumn{Name: "Region", IsSearch: true},
					index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
					index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
				},
			},
			index_model.Tabs{
				Name:             "Resources",
				Kind:             "RouteTabs",
				Subname:          "kind",
				Route:            "/Datacenters/:datacenter/Resources/:kind",
				TabParam:         "kind",
				GetQueries:       []string{"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
				ExpectedDataKeys: []string{"PhysicalResources", "Racks", "Floors", "PhysicalModels"},
				IsSync:           true,
				Tabs: []interface{}{
					index_model.Table{
						Name:    "Resources",
						Route:   "PhysicalResources",
						Kind:    "Table",
						DataKey: "PhysicalResources",
						Actions: []index_model.Action{
							index_model.Action{
								Name: "Create", Icon: "Create", Kind: "Form",
								DataKind: "PhysicalResource",
								Fields: []index_model.Field{
									index_model.Field{Name: "Name", Kind: "text",
										Require: true, Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$"},
									index_model.Field{Name: "Kind", Kind: "select", Require: true,
										Options: []string{
											"Server", "Pdu", "RackSpineRouter",
											"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
										}},
									index_model.Field{Name: "Rack", Kind: "select", Require: true,
										DataKey: "Racks"},
									index_model.Field{Name: "Model", Kind: "select", Require: true,
										DataKey: "PhysicalModels"},
								},
							},
						},
						SelectActions: []index_model.Action{
							index_model.Action{Name: "Delete", Icon: "Delete",
								Kind:      "Form",
								DataKind:  "PhysicalResource",
								SelectKey: "Name",
							},
						},
						ColumnActions: []index_model.Action{
							index_model.Action{Name: "Detail", Icon: "Detail"},
							index_model.Action{Name: "Update", Icon: "Update"},
						},
						Columns: []index_model.TableColumn{
							index_model.TableColumn{
								Name: "Name", IsSearch: true,
								Link:           "Datacenters/:datacenter/Resources/PhysicalResources/Detail/:0/View",
								LinkParam:      "resource",
								LinkSync:       false,
								LinkGetQueries: []string{"GetPhysicalResource"}},
							index_model.TableColumn{Name: "Kind"},
							index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "Action", Kind: "Action"},
						},
					},
					index_model.Table{
						Name:    "Racks",
						Route:   "/Racks",
						Kind:    "Table",
						DataKey: "Racks",
						SelectActions: []index_model.Action{
							index_model.Action{Name: "Delete", Icon: "Delete",
								Kind:      "Form",
								DataKind:  "Rack",
								SelectKey: "Name",
							},
						},
						Columns: []index_model.TableColumn{
							index_model.TableColumn{Name: "Name", IsSearch: true},
							index_model.TableColumn{Name: "Kind"},
							index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
						},
					},
					index_model.Table{
						Name:    "Floors",
						Route:   "/Floors",
						Kind:    "Table",
						DataKey: "Floors",
						SelectActions: []index_model.Action{
							index_model.Action{Name: "Delete", Icon: "Delete",
								Kind:      "Form",
								DataKind:  "Floor",
								SelectKey: "Name",
							},
						},
						Columns: []index_model.TableColumn{
							index_model.TableColumn{Name: "Name", IsSearch: true},
							index_model.TableColumn{Name: "Kind"},
							index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
						},
					},
					index_model.Table{
						Name:    "Models",
						Route:   "/Models",
						Kind:    "Table",
						DataKey: "PhysicalModels",
						Actions: []index_model.Action{
							index_model.Action{
								Name: "Create", Icon: "Create", Kind: "Form",
								DataKind: "PhysicalModel",
								Fields: []index_model.Field{
									index_model.Field{Name: "Name", Kind: "text", Require: true,
										Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
										RegExpMsg: "Please enter alphanumeric characters."},
									index_model.Field{Name: "Kind", Kind: "select", Require: true,
										Options: []string{
											"Server", "Pdu", "RackSpineRouter",
											"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
										}},
								},
							},
						},
						SelectActions: []index_model.Action{
							index_model.Action{Name: "Delete", Icon: "Delete",
								Kind:      "Form",
								DataKind:  "PhysicalModel",
								SelectKey: "Name",
							},
						},
						ColumnActions: []index_model.Action{
							index_model.Action{Name: "Detail", Icon: "Detail"},
						},
						Columns: []index_model.TableColumn{
							index_model.TableColumn{Name: "Name", IsSearch: true,
								Link:           "Datacenters/:datacenter/Resources/Models/Detail/:0/View",
								LinkParam:      "resource",
								LinkSync:       false,
								LinkGetQueries: []string{"GetPhysicalModel"}},
							index_model.TableColumn{Name: "Kind"},
							index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "desc"},
							index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							index_model.TableColumn{Name: "Action", Kind: "Action"},
						},
					},
				}, // Tabs
			},
			gin.H{
				"Name":     "Resource",
				"Subname":  "resource",
				"Route":    "/Datacenters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
				"TabParam": "subkind",
				"Kind":     "RouteTabs",
				"GetQueries": []string{
					"GetPhysicalModel",
					"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
				"ExpectedDataKeys": []string{"PhysicalModel"},
				"IsSync":           true,
				"Tabs": []interface{}{
					gin.H{
						"Name":    "View",
						"Route":   "/View",
						"Kind":    "View",
						"DataKey": "PhysicalModel",
						"Fields": []interface{}{
							gin.H{"Name": "Name", "Kind": "text"},
							gin.H{"Name": "Kind", "Kind": "select"},
						},
					},
					gin.H{
						"Name":         "Edit",
						"Route":        "/Edit",
						"Kind":         "Form",
						"DataKey":      "PhysicalModel",
						"SubmitAction": "Update",
						"Icon":         "Update",
						"Fields": []interface{}{
							gin.H{"Name": "Name", "Kind": "text", "Require": true,
								"Updatable": false,
								"Min":       5, "Max": 200, "RegExp": "^[0-9a-zA-Z]+$",
								"RegExpMsg": "Please enter alphanumeric characters."},
							gin.H{"Name": "Kind", "Kind": "select", "Require": true,
								"Updatable": true,
								"Options": []string{
									"Server", "Pdu", "RackSpineRouter",
									"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
								}},
						},
					},
				},
			},
		},
	}
}

func (modelApi *ResourceModelApi) convertDatacenters(tctx *logger.TraceContext, datacenters []resource_model.Datacenter) []*resource_api_grpc_pb.Datacenter {
	pbDatacenters := make([]*resource_api_grpc_pb.Datacenter, len(datacenters))
	for i, datacenter := range datacenters {
		updatedAt, err := ptypes.TimestampProto(datacenter.Model.UpdatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", datacenter.Model.UpdatedAt)
			continue
		}
		createdAt, err := ptypes.TimestampProto(datacenter.Model.CreatedAt)
		if err != nil {
			logger.Warningf(tctx, err,
				"Failed ptypes.TimestampProto: %v", datacenter.Model.CreatedAt)
			continue
		}

		pbDatacenters[i] = &resource_api_grpc_pb.Datacenter{
			Region:    datacenter.Region,
			Name:      datacenter.Name,
			Kind:      datacenter.Kind,
			UpdatedAt: updatedAt,
			CreatedAt: createdAt,
		}
	}

	return pbDatacenters
}
