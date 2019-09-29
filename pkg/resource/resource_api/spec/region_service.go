package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type RegionService struct {
	Region       string `validate:"required"`
	Name         string `validate:"required"` // Vip Domain
	Kind         string `validate:"required"`
	Status       string
	StatusReason string
	Cluster      string
	Spec         interface{} `validate:"required"`
}

type RegionComputeSpec struct {
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
	Replicas                    int `validate:"required"`
	ClusterFilters              []string
	ClusterLabelFilters         []string
	NodeFilters                 []string
	NodeLabelFilters            []string
	NodeLabelHardAffinities     []string
	NodeLabelHardAntiAffinities []string
	NodeLabelSoftAffinities     []string
	NodeLabelSoftAntiAffinities []string
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

var RegionServicesTable = index_model.Table{
	Name:    "RegionServices",
	Route:   "/RegionServices",
	Kind:    "Table",
	DataKey: "RegionServices",
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "RegionService",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Link:           "Regions/:Region/Resources/RegionServices/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetRegionService"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
