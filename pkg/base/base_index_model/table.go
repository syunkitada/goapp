package base_index_model

type Table struct {
	Name             string
	SubNameParamKeys []string
	Kind             string
	Route            string
	Subname          string
	DataKey          string
	DataQueries      []string
	DisablePaging    bool
	DisableToolbar   bool
	SearchForm       SearchForm
	Columns          []TableColumn
	Actions          []Action
	SelectActions    []Action
	ColumnActions    []Action
}

type SearchForm struct {
	Name   string
	Kind   string
	Inputs []TableInputField
}

type TableInputField struct {
	Name     string
	Type     string
	Multiple bool
	DataKey  string
	Data     interface{}
	Default  interface{}
}

type TableColumn struct {
	Name            string
	Kind            string
	IsSearch        bool
	Link            string
	LinkParam       string
	LinkKey         string
	LinkSync        bool
	LinkDataQueries []string // deprecated
	LinkPath        []string
	DataQueries     []string
	RowColoringMap  map[string]string
	FilterValues    []map[string]string
	Sort            string
	Icon            string
	Color           string
	InactiveColor   string
	Align           string
	View            interface{}
}

type Action struct {
	Name      string
	Kind      string
	Icon      string
	DataKind  string
	SelectKey string
	Query     string
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
	ListKey   string
	Updatable bool
	Fields    []Field
}
