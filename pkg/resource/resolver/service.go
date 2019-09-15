package resolver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec/genpkg"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, db *gorm.DB, input *base_spec.GetServiceIndex) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourcePhysicalCmdMap,
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceIndexData{
			Index: index_model.Index{
				CmdMap: genpkg.ResourceVirtualCmdMap,
			},
		}
		code = base_const.CodeOk
	default:
		code = base_const.CodeClientNotFound
	}

	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext, db *gorm.DB,
	input *base_spec.GetServiceDashboardIndex) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	// TODO
	switch input.Name {
	case "ResourcePhysical":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: index_model.DashboardIndex{
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
										"get_physical-resources", "get_racks", "get_floors", "get_physical-models"},
								},
								index_model.TableColumn{Name: "Region", IsSearch: true},
								index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
								index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
							},
							SelectActions: []index_model.Action{
								index_model.Action{Name: "Delete", Icon: "Delete",
									Kind:      "Form",
									DataKind:  "Datacenter",
									SelectKey: "Name",
								},
							},
							Actions: []index_model.Action{
								index_model.Action{
									Name: "Create", Icon: "Create", Kind: "Form",
									DataKind: "Datacenter",
									Fields: []index_model.Field{
										index_model.Field{Name: "Name", Kind: "text", Require: true,
											Min: 5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
											RegExpMsg: "Please enter alphanumeric characters."},
										index_model.Field{Name: "Kind", Kind: "select", Require: true,
											Options: []string{
												"Private", "Share",
											}},
									},
								},
							},
						},
						gin.H{
							"Name":  "ResourcePhysical HOGE",
							"Kind":  "Msg",
							"Route": "",
						},
						gin.H{
							"Name":  "Piyo",
							"Kind":  "Msg",
							"Route": "/Piyo",
						},
					},
				},
			},
		}
		code = base_const.CodeOk
	case "ResourceVirtual":
		data = &base_spec.GetServiceDashboardIndexData{
			Index: index_model.DashboardIndex{
				View: index_model.Panels{
					Name: "Root",
					Kind: "RoutePanels",
					Panels: []interface{}{
						gin.H{
							"Name":  "ResourceVirtual HOGE",
							"Kind":  "Msg",
							"Route": "",
						},
						gin.H{
							"Name":  "Piyo",
							"Kind":  "Msg",
							"Route": "/Piyo",
						},
					},
				},
			},
		}
		code = base_const.CodeOk
	default:
		code = base_const.CodeClientNotFound
	}
	fmt.Println("DEBUG GetServiceDashboardIndex", input, data)

	return
}
