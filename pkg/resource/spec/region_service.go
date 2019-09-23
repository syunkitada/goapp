package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type RegionService struct {
	Region       string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:63;"` // Vip Domain
	Project      string `gorm:"not null;size:63;"`
	Kind         string `gorm:"not null;size:25;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:100000;"`
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
