package index_model

type Tabs struct {
	Name             string
	Kind             string
	Subname          string
	Route            string
	RouteParamKey    string
	RouteParamValue  string
	TabParam         string
	GetQueries       []string
	DataQueries      []string
	ExpectedDataKeys []string
	IsSync           bool
	Tabs             []interface{}
}
