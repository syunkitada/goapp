package base_spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

type UpdateService struct {
	Name            string
	Scope           string
	SyncRootCluster bool
	ProjectRoles    []string
	Endpoints       []string
	QueryMap        map[string]base_model.QueryModel
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
	Index index_model.Index
}

type GetServiceDashboardIndex struct {
	Name string
}

type GetServiceDashboardIndexData struct {
	Index index_model.Index
	Data  interface{}
}
