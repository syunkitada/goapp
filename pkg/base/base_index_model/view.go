package base_index_model

type View struct {
	Name            string
	Route           string
	Kind            string
	EnableWebSocket bool
	WebSocketKey    string
	WebSocketQuery  string
	WebSocketKind   string
	DataQueries     []string
	DataKey         string
	ViewParams      map[string]interface{}
	Fields          []Field
	PanelsGroups    []interface{}
}
