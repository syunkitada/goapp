package index_model

type Base struct {
	Name string
	Kind string
}

type Index struct {
	SyncDelay uint
	CmdMap    map[string]Cmd
	View      interface{}
}

type Cmd struct {
	Arg         string
	ArgType     string
	ArgKind     string
	FlagMap     map[string]Flag
	TableHeader []string
	Help        string
}

type Flag struct {
	Flag     string
	FlagType string
	Help     string
}

type Panels struct {
	Name   string
	Kind   string
	Panels []interface{}
}
