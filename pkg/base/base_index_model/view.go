package base_index_model

type View struct {
	Name            string
	Route           string
	Kind            string
	EnableWebSocket bool
	DataQueries     []string
	DataKey         string
	Fields          []Field
	PanelsGroups    []interface{}
}