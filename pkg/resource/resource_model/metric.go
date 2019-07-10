package resource_model

type Metric struct {
	Name   string
	Time   string
	Tag    map[string]string
	Metric map[string]int64
}
