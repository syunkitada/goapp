package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

var ResourcePhysicalCmdMap = map[string]index_model.Cmd{
	"get.region": index_model.Cmd{
		QueryName: "GetRegion",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
	"get.regions": index_model.Cmd{
		QueryName:   "GetRegions",
		FlagMap:     map[string]index_model.Flag{},
		TableHeader: []string{},
	},
	"create.region": index_model.Cmd{
		QueryName: "CreateRegion",
		FlagMap: map[string]index_model.Flag{
			"spec": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
	"update.region": index_model.Cmd{
		QueryName: "UpdateRegion",
		FlagMap: map[string]index_model.Flag{
			"spec": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
	"delete.region": index_model.Cmd{
		QueryName: "DeleteRegion",
		FlagMap: map[string]index_model.Flag{
			"name": index_model.Flag{
				Required: false,
				FlagType: "string",
				FlagKind: "",
			},
		},
		TableHeader: []string{},
	},
}
var ResourceVirtualCmdMap = map[string]index_model.Cmd{
	"get.clusters": index_model.Cmd{
		QueryName:   "GetClusters",
		FlagMap:     map[string]index_model.Flag{},
		TableHeader: []string{},
	},
}

var ApiQueryMap = map[string]map[string]base_model.QueryModel{
	"Auth": map[string]base_model.QueryModel{
		"Login":         base_model.QueryModel{},
		"UpdateService": base_model.QueryModel{},
	},
	"ResourcePhysical": map[string]base_model.QueryModel{
		"GetRegion":    base_model.QueryModel{},
		"GetRegions":   base_model.QueryModel{},
		"CreateRegion": base_model.QueryModel{},
		"UpdateRegion": base_model.QueryModel{},
		"DeleteRegion": base_model.QueryModel{},
	},
	"ResourceVirtual": map[string]base_model.QueryModel{
		"GetClusters": base_model.QueryModel{},
	},
}
