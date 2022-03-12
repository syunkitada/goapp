package base_index_model

type Tabs struct {
	Name             string
	SubNameParamKeys []string
	Kind             string
	Subname          string
	Route            string
	RouteParamKey    string
	RouteParamValue  string
	TabParam         string
	DataQueries      []string
	ExpectedDataKeys []string
	IsSync           bool
	Tabs             []interface{}
	Children         []interface{}
}
