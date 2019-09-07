package base_spec

import "github.com/syunkitada/goapp/pkg/base/base_model/index_model"

type UpdateService struct {
	Name         string
	Scope        string
	ProjectRoles []string
	Endpoints    []string
}

type UpdateServiceData struct {
	Name string
}

type GetServices struct{}

type GetServicesData struct {
	Services []Service
}

type Service struct {
	Name      string
	Scope     string
	Endpoints string
}

type GetServiceIndex struct {
	Name string
}

type GetServiceIndexData struct {
	Index index_model.Index
}
