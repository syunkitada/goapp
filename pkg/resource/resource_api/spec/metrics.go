package spec

type MetricsGroup struct {
	Name         string
	MetricsGroup []Metrics
}

type Metrics struct {
	Name   string
	Keys   []string
	Values []map[string]interface{}
}
