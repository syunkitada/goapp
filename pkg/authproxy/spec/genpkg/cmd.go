package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
)

var HomeCmdMap = map[string]index_model.Cmd{
	"get.all.users": index_model.Cmd{
		QueryName: "GetAllUsers",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
	"get.user": index_model.Cmd{
		QueryName: "GetUser",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
}
var HomeProjectCmdMap = map[string]index_model.Cmd{
	"get.users": index_model.Cmd{
		QueryName: "GetUsers",
		FlagMap: map[string]index_model.Flag{
			"name,n": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		OutputKind:   "",
		OutputFormat: "",
	},
}

var ApiQueryMap = map[string]map[string]spec_model.QueryModel{
	"Auth": map[string]spec_model.QueryModel{
		"Login":         spec_model.QueryModel{},
		"UpdateService": spec_model.QueryModel{},
	},
	"Home": map[string]spec_model.QueryModel{
		"GetAllUsers": spec_model.QueryModel{},
		"GetUser":     spec_model.QueryModel{},
	},
	"HomeProject": map[string]spec_model.QueryModel{
		"GetUsers": spec_model.QueryModel{},
	},
}
