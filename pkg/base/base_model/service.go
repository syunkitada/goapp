package base_model

import "github.com/syunkitada/goapp/pkg/base/base_model/index_model"

type Spec struct {
	Meta interface{}
	Name string
	Apis []Api
}

type Api struct {
	Name            string
	Cmds            map[string]string
	RequiredAuth    bool
	RequiredProject bool
	RequiredService bool
	View            View
	Queries         []Query
	QueryModels     []QueryModel
}

type QueryModel struct {
	Cmd          string
	Help         string
	RequiredAuth bool
	ProjectRoles []string
	Roles        []string
	Req          interface{}
	Rep          interface{}
}

type Query struct {
	RequiredAuth    bool
	PkgPath         string
	PkgName         string
	Name            string
	CmdName         string
	CmdOutputKind   string
	CmdOutputFormat string
	Flags           []Flag
}

type Flag struct {
	Name      string
	FlagName  string
	ShortName string
	FlagType  string
	FlagKind  string
	CobraType string
	Required  bool
}

type ServiceRouter struct {
	QueryMap  map[string]QueryModel
	Endpoints []string
}

type GetServiceIndex struct {
	Name string
}

type GetServiceIndexData struct {
	Index index_model.Index
}
