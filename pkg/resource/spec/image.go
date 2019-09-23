package spec

import "github.com/syunkitada/goapp/pkg/authproxy/index_model"

type Image struct {
	Region       string
	Name         string
	Kind         string
	Labels       string
	Description  string
	Status       string
	StatusReason string
	Spec         interface{}
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
			Link:           "Regions/:Region/Resources/Images/Detail/:0/View",
			LinkParam:      "Name",
			LinkSync:       false,
			LinkGetQueries: []string{"GetImage"},
		},
		index_model.TableColumn{Name: "Kind"},
		index_model.TableColumn{Name: "UpdatedAt", Kind: "Time"},
		index_model.TableColumn{Name: "CreatedAt", Kind: "Time"},
	},
}
