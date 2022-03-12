package base_index_model

type Form struct {
	Name             string
	Route            string
	Kind             string
	DataKey          string
	SubmitAction     string
	Icon             string
	SubmitButtonName string
	DataQueries      []string
	SubmitQueries    []string
	Fields           []Field
}
