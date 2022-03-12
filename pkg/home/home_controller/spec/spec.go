package spec

import (
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec_model"
)

type Meta struct{}

var Spec = base_spec_model.Spec{
	Meta: Meta{},
	Name: "HomeApi",
	Kind: base_const.KindApi,
	Apis: []base_spec_model.Api{},
}
