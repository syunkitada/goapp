package spec

import (
	"time"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
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

var ImagesTable = base_index_model.Table{
	Name:        "Images",
	Route:       "/Images",
	Kind:        "Table",
	DataKey:     "Images",
	DataQueries: []string{"GetImages"},
	SelectActions: []base_index_model.Action{
		base_index_model.Action{
			Name:      "Delete",
			Icon:      "Delete",
			Kind:      "Form",
			DataKind:  "Image",
			SelectKey: "Name",
		},
	},
	Columns: []base_index_model.TableColumn{
		base_index_model.TableColumn{
			Name: "Name", IsSearch: true,
			Align:           "left",
			Link:            "Regions/:Region/RegionResources/Images/Detail/:0/View",
			LinkKeyMap:      map[string]string{"Name": "Name"},
			LinkSync:        false,
			LinkDataQueries: []string{"GetImage"},
		},
		base_index_model.TableColumn{Name: "Kind"},
		base_index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		base_index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}

var ImagesDetail = base_index_model.Tabs{
	Name:            "Images",
	Kind:            "RouteTabs",
	RouteParamKey:   "kind",
	RouteParamValue: "Images",
	Route:           "/Regions/:Region/Resources/Images/Detail/:Name/:Subkind",
	TabParam:        "Subkind",
	DataQueries: []string{
		"GetImage",
		"GetRegionServices", "GetImages"},
	ExpectedDataKeys: []string{"Image"},
	IsSync:           true,
	Tabs: []interface{}{
		base_index_model.View{
			Name:    "View",
			Route:   "/View",
			Kind:    "View",
			DataKey: "Image",
			Fields: []base_index_model.Field{
				base_index_model.Field{Name: "Name", Kind: "text"},
				base_index_model.Field{Name: "Kind", Kind: "select"},
			},
		},
		base_index_model.Form{
			Name:         "Edit",
			Route:        "/Edit",
			Kind:         "Form",
			DataKey:      "Image",
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
						"Url",
					}},
			},
		},
	},
}
