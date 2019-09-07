package genpkg

import (
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

var ResourcePhysicalCmdMap = map[string]index_model.Cmd{
	"get_regions": index_model.Cmd{
		QueryName:   "GetRegions",
		FlagMap:     map[string]index_model.Flag{},
		TableHeader: []string{},
	},
}
var ResourceVirtualCmdMap = map[string]index_model.Cmd{
	"get_clusters": index_model.Cmd{
		QueryName:   "GetClusters",
		FlagMap:     map[string]index_model.Flag{},
		TableHeader: []string{},
	},
}
