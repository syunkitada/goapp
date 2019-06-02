package index_model

type Base struct {
	Name string
	Kind string
}

type Index struct {
	SyncDelay uint
	Cmd       interface{}
	View      interface{}
}

type Panels struct {
	Name   string
	Kind   string
	Panels []interface{}
}
