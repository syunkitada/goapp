package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetImage(tctx *logger.TraceContext, input *spec.GetImage) (data *spec.GetImageData, code uint8, err error) {
	var image *spec.Image
	if image, err = resolver.dbApi.GetImage(tctx, input); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetImageData{Image: *image}
	return
}

func (resolver *Resolver) GetImages(tctx *logger.TraceContext, input *spec.GetImages) (data *spec.GetImagesData, code uint8, err error) {
	var images []spec.Image
	if images, err = resolver.dbApi.GetImages(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetImagesData{Images: images}
	return
}

func (resolver *Resolver) CreateImage(tctx *logger.TraceContext, input *spec.CreateImage) (data *spec.CreateImageData, code uint8, err error) {
	var specs []spec.Image
	if specs, err = resolver.ConvertToImageSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateImages(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateImageData{}
	return
}

func (resolver *Resolver) UpdateImage(tctx *logger.TraceContext, input *spec.UpdateImage) (data *spec.UpdateImageData, code uint8, err error) {
	var specs []spec.Image
	if specs, err = resolver.ConvertToImageSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateImages(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateImageData{}
	return
}

func (resolver *Resolver) DeleteImage(tctx *logger.TraceContext, input *spec.DeleteImage) (data *spec.DeleteImageData, code uint8, err error) {
	if err = resolver.dbApi.DeleteImage(tctx, input); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteImageData{}
	return
}

func (resolver *Resolver) DeleteImages(tctx *logger.TraceContext, input *spec.DeleteImages) (data *spec.DeleteImagesData, code uint8, err error) {
	var specs []spec.Image
	if specs, err = resolver.ConvertToImageSpecs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteImages(tctx, specs); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteImagesData{}
	return
}

func (resolver *Resolver) ConvertToImageSpecs(specStr string) (data []spec.Image, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	specs := []spec.Image{}
	for _, base := range baseSpecs {
		if base.Kind != "Image" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.Image
		if err = json.Unmarshal(specBytes, &specData); err != nil {
			return
		}
		if err = resolver.Validate.Struct(&specData); err != nil {
			return
		}
		specs = append(specs, specData)
	}
	return
}
