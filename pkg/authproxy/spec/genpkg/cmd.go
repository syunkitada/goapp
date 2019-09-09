package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
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
		TableHeader: []string{},
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
		TableHeader: []string{},
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
		TableHeader: []string{},
	},
}

var ApiQueryMap = map[string]map[string]base_model.QueryModel{
	"Auth": map[string]base_model.QueryModel{
		"Login":         base_model.QueryModel{},
		"UpdateService": base_model.QueryModel{},
	},
	"Home": map[string]base_model.QueryModel{
		"GetAllUsers": base_model.QueryModel{},
		"GetUser":     base_model.QueryModel{},
	},
	"HomeProject": map[string]base_model.QueryModel{
		"GetUsers": base_model.QueryModel{},
	},
}
