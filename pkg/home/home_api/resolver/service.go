package resolver

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (resolver *Resolver) GetServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	cmdMap := map[string]base_index_model.Cmd{}
	cmdMaps := []map[string]base_index_model.Cmd{
		base_spec.UserCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	code = base_const.CodeOk
	data = &base_spec.GetServiceIndexData{
		Index: base_index_model.Index{
			CmdMap: cmdMap,
		},
	}

	return
}

func (resolver *Resolver) GetProjectServiceIndex(tctx *logger.TraceContext, input *base_spec.GetServiceIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceIndexData, code uint8, err error) {
	cmdMap := map[string]base_index_model.Cmd{}
	cmdMaps := []map[string]base_index_model.Cmd{
		base_spec.UserCmd,
	}
	for _, tmpCmdMap := range cmdMaps {
		for key, cmd := range tmpCmdMap {
			cmdMap[key] = cmd
		}
	}

	code = base_const.CodeOk
	data = &base_spec.GetServiceIndexData{
		Index: base_index_model.Index{
			CmdMap: cmdMap,
		},
	}

	return
}

func (resolver *Resolver) GetServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	fmt.Println("DEBUG GetService")
	switch input.Name {
	case "Home":
		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"User": user,
			},
			Index: base_index_model.DashboardIndex{
				DefaultRoute: map[string]interface{}{
					"Path": []string{"User", "View"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panes",
					Children: []interface{}{
						map[string]interface{}{
							"Name": "User",
							"Kind": "Tabs",
							"Children": []interface{}{
								map[string]interface{}{
									"Name":    "View",
									"Kind":    "View",
									"DataKey": "User",
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
													},
												},
											},
										},
									},
								},
								map[string]interface{}{
									"Name":             "PasswordSetting",
									"Kind":             "Form",
									"DataKey":          "User",
									"Icon":             "Update",
									"SubmitButtonName": "Change Password",
									"SubmitQueries":    []string{"UpdateUserPassword"},
									"Fields": []base_index_model.Field{
										base_index_model.Field{Name: "CurrentPassword", Kind: "Password", Updatable: true, Required: true},
										base_index_model.Field{Name: "NewPassword", Kind: "Password", Updatable: true, Required: true},
										base_index_model.Field{Name: "NewPasswordConfirm", Kind: "Password", Updatable: true, Required: true},
									},
								},
							},
						},
					},
				},
			},
		}

	default:
		code = base_const.CodeClientNotFound
	}

	return
}

func (resolver *Resolver) GetProjectServiceDashboardIndex(tctx *logger.TraceContext, input *base_spec.GetServiceDashboardIndex, user *base_spec.UserAuthority) (data *base_spec.GetServiceDashboardIndexData, code uint8, err error) {
	switch input.Name {
	case "HomeProject":
		data = &base_spec.GetServiceDashboardIndexData{
			Data: map[string]interface{}{
				"User": user,
			},
			Index: base_index_model.DashboardIndex{
				DefaultRoute: map[string]interface{}{
					"Path": []string{"User", "View"},
				},
				View: base_index_model.Panels{
					Name: "Root",
					Kind: "Panes",
					Children: []interface{}{
						base_index_model.Tabs{
							Name: "User",
							Kind: "Tabs",
							Children: []interface{}{
								base_index_model.View{
									Name:    "View",
									Kind:    "View",
									DataKey: "User",
									PanelsGroups: []interface{}{
										map[string]interface{}{
											"Name": "Detail",
											"Kind": "Cards",
											"Cards": []interface{}{
												map[string]interface{}{
													"Name": "Detail",
													"Kind": "Fields",
													"Fields": []base_index_model.Field{
														base_index_model.Field{Name: "Name"},
														base_index_model.Field{Name: "ProjectName"},
													},
												},
											},
										},
									},
								},
								base_index_model.Table{
									Name:        "Users",
									Kind:        "Table",
									DataQueries: []string{"GetProjectUsers"},
									DataKey:     "Users",
									Columns: []base_index_model.TableColumn{
										base_index_model.TableColumn{Name: "Name", IsSearch: true, Align: "left"},
										base_index_model.TableColumn{Name: "RoleName", Align: "left"},
									},
								},
							},
						},
					},
				},
			},
		}
	default:
		code = base_const.CodeClientNotFound
	}

	return
}
