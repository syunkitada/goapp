package resource_model

import (
	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

const ImageKind = "Image"

type Image struct {
	gorm.Model
	Region       string `gorm:"not null;size:50;"`
	Name         string `gorm:"not null;size:255;"` // Name is unique in Region
	Kind         string `gorm:"not null;size:25;"`
	Labels       string `gorm:"not null;size:255;"`
	Description  string `gorm:"not null;size:255;"`
	Status       string `gorm:"not null;size:25;"`
	StatusReason string `gorm:"not null;size:50;"`
	Spec         string `gorm:"not null;size:5000;"`
}

type ImageSpec struct {
	Kind        string `validate:"required"`
	Region      string `validate:"required"`
	Name        string `validate:"required"`
	Description string
	Url         string
}

var ImageCmd map[string]index_model.Cmd = map[string]index_model.Cmd{
	"create_image": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: ImageKind,
		Help:    "create image",
	},
	"update_image": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeFile,
		ArgKind: ImageKind,
		Help:    "update image",
	},
	"get_images": index_model.Cmd{
		Arg:         index_model.ArgOptional,
		ArgType:     index_model.ArgTypeString,
		ArgKind:     ImageKind,
		Help:        "get images",
		TableHeader: []string{"Name", "Kind", "Region", "Status"},
	},
	"get_image": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: ImageKind,
		Help:    "get image",
	},
	"delete_image": index_model.Cmd{
		Arg:     index_model.ArgRequired,
		ArgType: index_model.ArgTypeString,
		ArgKind: ImageKind,
		Help:    "delete image",
	},
}

var ImagesTable = index_model.Table{
	Name:    "Images",
	Route:   "/Images",
	Kind:    "Table",
	DataKey: "Images",
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
			Link:           "Clusters/:datacenter/Resources/Images/Detail/:0/View",
			LinkParam:      "resource",
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
	Route:           "/Clusters/:datacenter/Resources/Images/Detail/:resource/:subkind",
	TabParam:        "subkind",
	GetQueries: []string{
		"GetImage",
		"GetImages", "GetImages"},
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
						"Server", "Pdu", "RackSpineRouter",
						"FloorLeafRouter", "FloorSpineRouter", "GatewayRouter",
					}},
			},
		},
	},
}
