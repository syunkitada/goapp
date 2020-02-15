package base_spec_model

type Spec struct {
	Meta     interface{}
	Name     string
	Kind     string
	Apis     []Api
	QuerySet map[string]Query
}

type Api struct {
	Name            string
	Cmds            map[string]string
	RequiredAuth    bool
	RequiredProject bool
	Queries         []Query
	QueryModels     []QueryModel
	ViewModels      []interface{}
}
