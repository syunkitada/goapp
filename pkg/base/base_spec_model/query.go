package base_spec_model

type ServiceRouter struct {
	Token     string
	QueryMap  map[string]QueryModel
	Endpoints []string
}

type QueryModel struct {
	RequiredAuth    bool
	RequiredProject bool
	Ws              bool
	Kind            string
	Cmd             string
	Help            string
	ProjectRoles    []string
	Roles           []string
	Req             interface{}
	Rep             interface{}
}

type Query struct {
	RequiredAuth    bool
	RequiredProject bool
	Ws              bool
	Kind            string
	PkgPath         string
	PkgName         string
	Name            string
	ActionName      string
	DataName        string
	CmdName         string
	CmdOutputKind   string
	CmdOutputFormat string
	Flags           []Flag
}
