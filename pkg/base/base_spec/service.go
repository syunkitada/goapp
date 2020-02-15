package base_spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

type UpdateService struct {
	Name            string
	Token           string
	Scope           string
	SyncRootCluster bool
	ProjectRoles    []string
	Endpoints       []string
	QueryMap        map[string]base_spec_model.QueryModel
}

type UpdateServiceData struct {
	Name string
}

type GetServices struct{}

type GetServicesData struct {
	Services []Service
}

type Service struct {
	Name            string
	Token           string
	Scope           string
	Endpoints       string
	ProjectRoles    string
	QueryMap        string
	SyncRootCluster bool
}

type GetServiceIndex struct {
	Name string
}

type GetServiceIndexData struct {
	Index base_index_model.Index
}

type GetServiceDashboardIndex struct {
	Name string
}

type GetServiceDashboardIndexData struct {
	Index base_index_model.DashboardIndex
	Data  interface{}
}
