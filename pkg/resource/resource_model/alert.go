package resource_model

type Alert struct {
	Name    string
	Time    string
	Level   string
	Handler string
	Msg     string
	Tag     map[string]string
}
