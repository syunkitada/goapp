package base_model

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
	Model        interface{}
}

type Query struct {
	RequiredAuth bool
	PkgPath      string
	PkgName      string
	Name         string
}

type Flag struct {
	Name      string
	FlagName  string
	ShortName string
	Type      string
	CobraType string
	Required  bool
}

type ServiceRouter struct {
	QueryMap  map[string]QueryModel
	Endpoints []string
}
