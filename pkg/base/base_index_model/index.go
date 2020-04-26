package base_index_model

type DashboardIndex struct {
	SyncDelay    uint
	DefaultRoute map[string]interface{}
	View         interface{}
}

type Index struct {
	SyncDelay uint
	CmdMap    map[string]Cmd
	View      interface{}
}

type Cmd struct {
	QueryName    string
	Arg          string
	ArgType      string
	ArgKind      string
	FlagMap      map[string]Flag
	OutputKind   string
	OutputFormat string
	Help         string
	Ws           bool
	Kind         string
}

type Flag struct {
	Required bool
	Flag     string // depricated
	FlagType string // string, int, ...
	FlagKind string // empty or file
	Help     string
}

type Panels struct {
	Name     string
	Kind     string
	Panels   []interface{}
	Children []interface{}
}
