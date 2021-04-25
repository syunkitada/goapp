package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
)

type RegionService struct {
	Region       string `validate:"required"`
	Name         string `validate:"required"` // Vip Domain
	Kind         string `validate:"required"`
	Status       string
	StatusReason string
	UpdatedAt    time.Time
	CreatedAt    time.Time
	Spec         interface{} `validate:"required"`
}

type RegionServiceComputeSpec struct {
	Kind           string             `validate:"required"`
	Image          string             `validate:"required"`
	SchedulePolicy SchedulePolicySpec `validate:"required"`
	NetworkPolicy  NetworkPolicySpec  `validate:"required"`
	Vcpus          uint               `validate:"required"`
	Memory         uint               `validate:"required"`
	Disk           uint               `validate:"required"`

	ImageSpec Image      // Auto Generated
	Name      string     // Auto Generated
	Cluster   string     // Auto Generated
	Ports     []PortSpec // Auto Generated
}

type SchedulePolicySpec struct {
	Replicas                           int `validate:"required"`
	ClusterFilters                     []string
	ClusterLabelFilters                []string
	NodeServiceFilters                 []string
	NodeServiceLabelFilters            []string
	NodeServiceLabelHardAffinities     []string
	NodeServiceLabelHardAntiAffinities []string
	NodeServiceLabelSoftAffinities     []string
	NodeServiceLabelSoftAntiAffinities []string
}

type NetworkPolicySpec struct {
	Version        int
	Interfaces     int
	AssignPolicy   string
	StaticNetworks []string
}

type PortSpec struct {
	NetworkID uint
	Version   int
	Subnet    string
	Gateway   string
	Ip        string
	Mac       string
	Kind      string
	Spec      string
}

type GetRegionService struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type GetRegionServiceData struct {
	RegionService RegionService
}

type GetRegionServices struct {
	Region string `validate:"required"`
}

type GetRegionServicesData struct {
	RegionServices []RegionService
}

type CreateRegionService struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateRegionServiceData struct{}

type UpdateRegionService struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateRegionServiceData struct{}

type DeleteRegionService struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type DeleteRegionServiceData struct{}

type DeleteRegionServices struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteRegionServicesData struct{}

var RegionServicesTable = base_index_model.Table{
	Name:        "RegionServices",
	Route:       "/RegionServices",
	Kind:        "Table",
	DataKey:     "RegionServices",
	DataQueries: []string{"GetRegionServices"},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "RegionService",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Align:           "left",
			Link:            "Regions/:Region/RegionResources/RegionServices/Detail/:0/View",
			LinkKeyMap:      map[string]string{"Name": "Name"},
			LinkSync:        false,
			LinkDataQueries: []string{"GetRegionService"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "Status"},
		base_index_model.TableColumn{Name: "StatusReason"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var RegionServicesDetail = base_index_model.Tabs{
	Name:            "RegionServices",
	Kind:            "RouteTabs",
	RouteParamKey:   "kind",
	RouteParamValue: "RegionServices",
	Route:           "/Regions/:Region/Resources/RegionServices/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	DataQueries: []string{
		"GetRegionService",
		"GetRegionServices", "GetImages"},
	ExpectedDataKeys: []string{"RegionService"},
	IsSync:           true,
	Tabs: []interface{}{
		base_index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "RegionService",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text"},
				base_index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		base_index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "RegionService",
			SubmitAction: "update image",
			Icon:         "Update",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text", Required: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				base_index_model.Field{Name: "Kind", Kind: "select", Required: true,
					Updatable: true,
					Options: []string{
						"Compute",
					}},
			},
		},
	},
}
