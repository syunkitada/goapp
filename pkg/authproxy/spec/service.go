package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type UpdateService struct {
	Name string
}

type UpdateServiceData struct {
	Name string
}

type GetServiceIndex struct {
	Name string
}

type GetServiceIndexData struct {
	Index index_model.Index
}
