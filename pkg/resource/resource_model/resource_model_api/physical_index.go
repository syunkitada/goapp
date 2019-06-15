package resource_model_api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_grpc_pb"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_utils"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
	"github.com/syunkitada/goapp/pkg/lib/codes"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/resource_model"
)

func (modelApi *ResourceModelApi) PhysicalAction(tctx *logger.TraceContext,
	req *authproxy_grpc_pb.ActionRequest, rep *authproxy_grpc_pb.ActionReply) {
	var err error
	var statusCode int64
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	data := map[string]interface{}{}
	response := authproxy_model.ActionResponse{
		Tctx: *req.Tctx,
	}

	var db *gorm.DB
	if db, err = modelApi.open(tctx); err != nil {
		authproxy_utils.MergeResponse(rep, &response, data, err, codes.RemoteDbError)
		return
	}
	defer func() {
		tmpErr := db.Close()
		if tmpErr != nil {
			logger.Error(tctx, tmpErr)
		}
	}()

	statusCode = codes.Unknown
	for _, query := range req.Queries {
		switch query.Kind {
		case "GetIndex":
			response.Index = *modelApi.getPhysicalIndex()
		case "GetDashboardIndex":
			response.Index = *modelApi.getPhysicalIndex()
			statusCode, err = modelApi.GetDatacenters(tctx, db, query, data)

		case "GetDatacenter":
			statusCode, err = modelApi.GetDatacenter(tctx, db, query, data)
		case "GetDatacenters":
			statusCode, err = modelApi.GetDatacenters(tctx, db, query, data)
		case "CreateDatacenter":
			statusCode, err = modelApi.CreateDatacenter(tctx, db, query)
		case "UpdateDatacenter":
			statusCode, err = modelApi.UpdateDatacenter(tctx, db, query)
		case "DeleteDatacenter":
			statusCode, err = modelApi.DeleteDatacenter(tctx, db, query)

		case "GetFloor":
			statusCode, err = modelApi.GetFloor(tctx, db, query, data)
		case "GetFloors":
			statusCode, err = modelApi.GetFloors(tctx, db, query, data)
		case "CreateFloor":
			statusCode, err = modelApi.CreateFloor(tctx, db, query)
		case "UpdateFloor":
			statusCode, err = modelApi.UpdateFloor(tctx, db, query)
		case "DeleteFloor":
			statusCode, err = modelApi.DeleteFloor(tctx, db, query)

		case "GetRack":
			statusCode, err = modelApi.GetRack(tctx, db, query, data)
		case "GetRacks":
			statusCode, err = modelApi.GetRacks(tctx, db, query, data)
		case "CreateRack":
			statusCode, err = modelApi.CreateRack(tctx, db, query)
		case "UpdateRack":
			statusCode, err = modelApi.UpdateRack(tctx, db, query)
		case "DeleteRack":
			statusCode, err = modelApi.DeleteRack(tctx, db, query)

		case "GetPhysicalResource":
			statusCode, err = modelApi.GetPhysicalResource(tctx, db, query, data)
		case "GetPhysicalResources":
			statusCode, err = modelApi.GetPhysicalResources(tctx, db, query, data)
		case "CreatePhysicalResource":
			statusCode, err = modelApi.CreatePhysicalResource(tctx, db, query)
		case "UpdatePhysicalResource":
			statusCode, err = modelApi.UpdatePhysicalResource(tctx, db, query)
		case "DeletePhysicalResource":
			statusCode, err = modelApi.DeletePhysicalResource(tctx, db, query)

		case "GetPhysicalModel":
			statusCode, err = modelApi.GetPhysicalModel(tctx, db, query, data)
		case "GetPhysicalModels":
			statusCode, err = modelApi.GetPhysicalModels(tctx, db, query, data)
		case "CreatePhysicalModel":
			statusCode, err = modelApi.CreatePhysicalModel(tctx, db, query)
		case "UpdatePhysicalModel":
			statusCode, err = modelApi.UpdatePhysicalModel(tctx, db, query)
		case "DeletePhysicalModel":
			statusCode, err = modelApi.DeletePhysicalModel(tctx, db, query)
		}

		if err != nil {
			break
		}
	}

	authproxy_utils.MergeResponse(rep, &response, data, err, statusCode)
}

func (modelApi *ResourceModelApi) getPhysicalIndex() *index_model.Index {
	cmdMap := map[string]index_model.Cmd{}
	cmdMaps := []map[string]index_model.Cmd{
		resource_model.DatacenterCmd,
		resource_model.RackCmd,
		resource_model.FloorCmd,
		resource_model.PhysicalModelCmd,
		resource_model.PhysicalResourceCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	return &index_model.Index{
		SyncDelay: 20000,
		CmdMap:    cmdMap,
		View: index_model.Panels{
			Name: "Root",
			Kind: "RoutePanels",
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
									Link:           "Datacenters/:datacenter/Resources/Resources/Detail/:0/View",
									LinkParam:      "resource",
									LinkSync:       false,
									LinkGetQueries: []string{"GetPhysicalResource"},
								},
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
								index_model.Action{
									Name:      "Delete",
									Icon:      "Delete",
									Kind:      "Form",
									DataKind:  "Rack",
									SelectKey: "Name",
								},
							},
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name: "Name", IsSearch: true,
									Link:           "Datacenters/:datacenter/Resources/Racks/Detail/:0/View",
									LinkParam:      "resource",
									LinkSync:       false,
									LinkGetQueries: []string{"GetRack"},
								},
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
								index_model.Action{
									Name:      "Delete",
									Icon:      "Delete",
									Kind:      "Form",
									DataKind:  "Floor",
									SelectKey: "Name",
								},
							},
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name: "Name", IsSearch: true,
									Link:           "Datacenters/:datacenter/Resources/Floors/Detail/:0/View",
									LinkParam:      "resource",
									LinkSync:       false,
									LinkGetQueries: []string{"GetFloor"},
								},
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
								index_model.Action{
									Name: "Delete", Icon: "Delete",
									Kind:      "Form",
									DataKind:  "PhysicalModel",
									SelectKey: "Name",
								},
							},
							ColumnActions: []index_model.Action{
								index_model.Action{Name: "Detail", Icon: "Detail"},
							},
							Columns: []index_model.TableColumn{
								index_model.TableColumn{
									Name:           "Name",
									IsSearch:       true,
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
					"Name":      "Resource",
					"Subname":   "resource",
					"Route":     "/Datacenters/:datacenter/Resources/:kind/Detail/:resource/:subkind",
					"Kind":      "RoutePanes",
					"PaneParam": "kind",
					"Panes": []interface{}{
						gin.H{
							"Name":            "Models",
							"Kind":            "RouteTabs",
							"RouteParamKey":   "kind",
							"RouteParamValue": "Models",
							"Route":           "/Datacenters/:datacenter/Resources/Models/Detail/:resource/:subkind",
							"TabParam":        "subkind",
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
						gin.H{
							"Name":            "Resources",
							"Kind":            "RouteTabs",
							"RouteParamKey":   "kind",
							"RouteParamValue": "Resources",
							"Route":           "/Datacenters/:datacenter/Resources/Resources/Detail/:resource/:subkind",
							"TabParam":        "subkind",
							"GetQueries": []string{
								"GetPhysicalResource",
								"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
							"ExpectedDataKeys": []string{"PhysicalResource"},
							"IsSync":           true,
							"Tabs": []interface{}{
								gin.H{
									"Name":    "View",
									"Route":   "/View",
									"Kind":    "View",
									"DataKey": "PhysicalResource",
									"Fields": []interface{}{
										gin.H{"Name": "Name", "Kind": "text"},
										gin.H{"Name": "Kind", "Kind": "select"},
									},
								},
								gin.H{
									"Name":         "Edit",
									"Route":        "/Edit",
									"Kind":         "Form",
									"DataKey":      "PhysicalResource",
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
				},
			},
		},
	}
}
