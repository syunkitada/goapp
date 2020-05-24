// This code is auto generated.
// Don't modify this code.

package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

var HomeCmdMap = map[string]base_index_model.Cmd{
	"get.all.users": base_index_model.Cmd{
		QueryName: "GetAllUsers",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"get.user": base_index_model.Cmd{
		QueryName: "GetUser",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
	"update.user.password": base_index_model.Cmd{
		QueryName: "UpdateUserPassword",
		FlagMap: map[string]base_index_model.Flag{
			"current.password,c": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"new.password,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
			"new.password.confirm,n": base_index_model.Flag{
				Required: true,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
}
var HomeProjectCmdMap = map[string]base_index_model.Cmd{
	"get.users": base_index_model.Cmd{
		QueryName: "GetUsers",
		FlagMap: map[string]base_index_model.Flag{
			"name,n": base_index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		Kind:         "",
		OutputKind:   "",
		OutputFormat: "",
		Ws:           false,
	},
}

var ApiQueryMap = map[string]map[string]base_spec_model.QueryModel{
	"Auth": map[string]base_spec_model.QueryModel{
		"Login":         base_spec_model.QueryModel{},
		"UpdateService": base_spec_model.QueryModel{},
	},
	"Home": map[string]base_spec_model.QueryModel{
		"GetAllUsers": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: false,
		},
		"GetUser": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: false,
		},
		"UpdateUserPassword": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: false,
		},
	},
	"HomeProject": map[string]base_spec_model.QueryModel{
		"GetUsers": base_spec_model.QueryModel{
			RequiredAuth:    true,
			RequiredProject: true,
		},
	},
}
