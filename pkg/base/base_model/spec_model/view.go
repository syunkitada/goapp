package spec_model

type Table struct {
	Name          string
	Data          interface{}
	Actions       []interface{}
	SelectActions []interface{}
	ColumnLinkMap map[string]Link
}

type Link struct {
	Target string
}

type Tab struct {
	Name string
	Tabs []interface{}
}
