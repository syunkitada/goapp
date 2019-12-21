package resource_model

type Event struct {
	Name    string
	Time    string
	Level   string
	Handler string
	Msg     string
	Tag     map[string]string
}
