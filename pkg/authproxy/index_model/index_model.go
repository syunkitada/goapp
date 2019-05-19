package index_model

type Base struct {
	Name string
	Kind string
}

type Panels struct {
	Name      string
	Kind      string
	SyncDelay uint
	Panels    []interface{}
}
