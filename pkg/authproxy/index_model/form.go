package index_model

type Form struct {
	Name         string
	Route        string
	Kind         string
	DataKey      string
	SubmitAction string
	Icon         string
	Fields       []Field
}
