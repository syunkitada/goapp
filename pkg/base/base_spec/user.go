package base_spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type GetUser struct {
	Name string
}

type GetUserData struct {
	Name string
}

type GetAllUsers struct {
	Name string
}

type GetAllUsersData struct {
	Name string
}

type GetUsers struct {
	Name string
}

type GetUsersData struct {
	Name string
}

const UserKind = "User"

var UserCmd map[string]base_index_model.Cmd = map[string]base_index_model.Cmd{
	"create_user": base_index_model.Cmd{
		Arg:     base_const.ArgRequired,
		ArgType: base_const.ArgTypeFile,
		ArgKind: UserKind,
		Help:    "helptext",
	},
	"update_user": base_index_model.Cmd{
		Arg:     base_const.ArgRequired,
		ArgType: base_const.ArgTypeFile,
		ArgKind: UserKind,
		Help:    "helptext",
	},
	"get_users": base_index_model.Cmd{
		Arg:     base_const.ArgOptional,
		ArgType: base_const.ArgTypeString,
		ArgKind: UserKind,
		Help:    "helptext",
	},
	"get_user": base_index_model.Cmd{
		Arg:     base_const.ArgRequired,
		ArgType: base_const.ArgTypeString,
		ArgKind: UserKind,
		Help:    "helptext",
	},
	"delete_user": base_index_model.Cmd{
		Arg:     base_const.ArgRequired,
		ArgType: base_const.ArgTypeString,
		ArgKind: UserKind,
		Help:    "helptext",
	},
}
