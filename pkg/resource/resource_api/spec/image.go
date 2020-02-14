package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
)

type Image struct {
	Region       string
	Name         string
	Kind         string
	Labels       string
	Description  string
	Status       string
	StatusReason string
	UpdatedAt    time.Time
	CreatedAt    time.Time
	Spec         interface{}
}

type ImageUrlSpec struct {
	Url string
}

type GetImage struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type GetImageData struct {
	Image Image
}

type GetImages struct {
	Region string `validate:"required"`
}

type GetImagesData struct {
	Images []Image
}

type CreateImage struct {
	Spec string `validate:"required" flagKind:"file"`
}

type CreateImageData struct{}

type UpdateImage struct {
	Spec string `validate:"required" flagKind:"file"`
}

type UpdateImageData struct{}

type DeleteImage struct {
	Name   string `validate:"required"`
	Region string `validate:"required"`
}

type DeleteImageData struct{}

type DeleteImages struct {
	Spec string `validate:"required" flagKind:"file"`
}

type DeleteImagesData struct{}

var ImagesTable = index_model.Table{
	Name:        "Images",
	Route:       "/Images",
	Kind:        "Table",
	DataKey:     "Images",
	DataQueries: []string{"GetImages"},
	SelectActions: []index_model.Action{
		index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Image",
			SelectKey: "Name",
		},
	},
	Columns: []index_model.TableColumn{
		index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Align:          "left",
			Link:           "Regions/:Region/RegionResources/Images/Detail/:0/View",
			LinkKey:        "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetImage"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var ImagesDetail = index_model.Tabs{
	Name:            "Images",
	Kind:            "RouteTabs",
	RouteParamKey:   "kind",
	RouteParamValue: "Images",
	Route:           "/Regions/:Region/Resources/Images/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	GetQueries: []string{
		"GetImage",
		"GetRegionServices", "GetImages"},
	ExpectedDataKeys: []string{"Image"},
	IsSync:           true,
	Tabs: []interface{}{
		index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Image",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text"},
				index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "Image",
			SubmitAction: "update image",
			Icon:         "Update",
			Fields: []index_model.Field{
				index_model.Field{Name: "Name", Kind: "text", Require: true,
					Updatable: false,
					Min:       5, Max: 200, RegExp: "^[0-9a-zA-Z]+$",
					RegExpMsg: "Please enter alphanumeric characters."},
				index_model.Field{Name: "Kind", Kind: "select", Require: true,
					Updatable: true,
					Options: []string{
						"Url",
					}},
			},
		},
	},
}
