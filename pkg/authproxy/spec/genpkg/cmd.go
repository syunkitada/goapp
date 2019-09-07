package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

var AuthCmdMap = map[string]index_model.Cmd{
	"update_service": index_model.Cmd{
		QueryName: "UpdateService",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"scope": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"projectroles": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
			"endpoints": index_model.Flag{
				Required: false,
				FlagType: "[]string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
	"login": index_model.Cmd{
		QueryName: "Login",
		FlagMap: map[string]index_model.Flag{
			"user": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
			"password": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
}
var HomeCmdMap = map[string]index_model.Cmd{
	"get_all_users": index_model.Cmd{
		QueryName: "GetAllUsers",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
	"get_user": index_model.Cmd{
		QueryName: "GetUser",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
}
var HomeProjectCmdMap = map[string]index_model.Cmd{
	"get_users": index_model.Cmd{
		QueryName: "GetUsers",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
}
