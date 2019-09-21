package spec_model

type ServiceRouter struct {
	QueryMap  map[string]QueryModel
	Endpoints []string
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
	ActionName      string
	DataName        string
	CmdName         string
	CmdOutputKind   string
	CmdOutputFormat string
	Flags           []Flag
}
