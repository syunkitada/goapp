package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
)

type Meta struct{}

var Spec = spec_model.Spec{
	Meta: Meta{},
	Name: "ResourceClusterController",
	Kind: base_const.KindAgent,
	Apis: []spec_model.Api{},
}
