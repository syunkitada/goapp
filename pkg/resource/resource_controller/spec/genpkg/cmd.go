package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
)

var ApiQueryMap = map[string]map[string]spec_model.QueryModel{
	"Auth": map[string]spec_model.QueryModel{
		"Login":         spec_model.QueryModel{},
		"UpdateService": spec_model.QueryModel{},
	},
}
