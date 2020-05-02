package base_index_model

type View struct {
	Name            string
	Route           string
	Kind            string
	EnableWebSocket bool
	WebSocketKey    string
	WebSocketQuery  string
	DataQueries     []string
	DataKey         string
	Fields          []Field
	PanelsGroups    []interface{}
}
