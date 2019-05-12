package index_model

type Tabs struct {
	Name             string
	Kind             string
	Subname          string
	Route            string
	TabParam         string
	GetQueries       []string
	ExpectedDataKeys []string
	IsSync           bool
	Tabs             []interface{}
}
