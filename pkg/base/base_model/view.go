package base_model

type View struct {
	Name    string
	Route   string
	Kind    string
	DataKey string
	Fields  []Field
}

type Form struct {
	Name         string
	Route        string
	Kind         string
	DataKey      string
	SubmitAction string
	Icon         string
	Fields       []Field
}

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

type Panels struct {
	Name   string
	Kind   string
	Panels []interface{}
}

type Table struct {
	Name          string
	Kind          string
	Route         string
	Subname       string
	DataKey       string
	Columns       []TableColumn
	Actions       []Action
	SelectActions []Action
	ColumnActions []Action
}

type TableColumn struct {
	Name           string
	Kind           string
	IsSearch       bool
	Link           string
	LinkParam      string
	LinkSync       bool
	LinkGetQueries []string
	Sort           string
}

type Action struct {
	Name      string
	Kind      string
	Icon      string
	DataKind  string
	SelectKey string
	Fields    []Field
}

type Field struct {
	Name      string
	Kind      string
	Require   bool
	Min       uint
	Max       uint
	RegExp    string
	RegExpMsg string
	Options   []string
	DataKey   string
	Updatable bool
}
