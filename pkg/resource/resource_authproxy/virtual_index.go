package resource_authproxy

import (
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

func (resource *Resource) getVirtualIndex() interface{} {
	return index_model.Panels{
		Name:      "Root",
		Kind:      "RoutePanels",
		SyncDelay: 20000,
		Panels: []interface{}{
			index_model.Table{
				Name:    "Clusters",
				Kind:    "Table",
				Route:   "",
				Subname: "cluster",
				DataKey: "Clusters",
				Columns: []index_model.TableColumn{
					index_model.TableColumn{
						Name:      "Name",
						IsSearch:  true,
						Link:      "Clusters/:0/Resources/Resources",
						LinkParam: "cluster",
						LinkSync:  true,
						LinkGetQueries: []string{
							"GetPhysicalResources", "GetRacks", "GetFloors", "GetPhysicalModels"},
					},
					index_model.TableColumn{Name: "Datacenter", IsSearch: true},
					index_model.TableColumn{Name: "UpdatedAt", Kind: "Time", Sort: "asc"},
					index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
				},
			},
		},
	}
}
