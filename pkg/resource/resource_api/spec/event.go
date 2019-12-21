package spec

type ResourceEvent struct {
	Name    string
	Time    string
	Level   string
	Handler string
	Msg     string
	Tag     map[string]string
}

type GetEvents struct {
	Cluster string `validate:"required"`
}

type GetEventsData struct {
	Events []map[string]interface{}
}
