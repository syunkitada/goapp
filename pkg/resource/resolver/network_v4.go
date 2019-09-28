package resolver

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/resource/spec"
)

func (resolver *Resolver) GetNetworkV4(tctx *logger.TraceContext, input *spec.GetNetworkV4, user *base_spec.UserAuthority) (data *spec.GetNetworkV4Data, code uint8, err error) {
	var rack *spec.NetworkV4
	if rack, err = resolver.dbApi.GetNetworkV4(tctx, input, user); err != nil {
		if gorm.IsRecordNotFoundError(err) {
			code = base_const.CodeOkNotFound
			return
		}
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNetworkV4Data{NetworkV4: *rack}
	return
}

func (resolver *Resolver) GetNetworkV4s(tctx *logger.TraceContext, input *spec.GetNetworkV4s, user *base_spec.UserAuthority) (data *spec.GetNetworkV4sData, code uint8, err error) {
	var racks []spec.NetworkV4
	if racks, err = resolver.dbApi.GetNetworkV4s(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOk
	data = &spec.GetNetworkV4sData{NetworkV4s: racks}
	return
}

func (resolver *Resolver) CreateNetworkV4(tctx *logger.TraceContext, input *spec.CreateNetworkV4, user *base_spec.UserAuthority) (data *spec.CreateNetworkV4Data, code uint8, err error) {
	var specs []spec.NetworkV4
	if specs, err = resolver.ConvertToNetworkV4Specs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.CreateNetworkV4s(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkCreated
	data = &spec.CreateNetworkV4Data{}
	return
}

func (resolver *Resolver) UpdateNetworkV4(tctx *logger.TraceContext, input *spec.UpdateNetworkV4, user *base_spec.UserAuthority) (data *spec.UpdateNetworkV4Data, code uint8, err error) {
	var specs []spec.NetworkV4
	if specs, err = resolver.ConvertToNetworkV4Specs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.UpdateNetworkV4s(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkUpdated
	data = &spec.UpdateNetworkV4Data{}
	return
}

func (resolver *Resolver) DeleteNetworkV4(tctx *logger.TraceContext, input *spec.DeleteNetworkV4, user *base_spec.UserAuthority) (data *spec.DeleteNetworkV4Data, code uint8, err error) {
	if err = resolver.dbApi.DeleteNetworkV4(tctx, input, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteNetworkV4Data{}
	return
}

func (resolver *Resolver) DeleteNetworkV4s(tctx *logger.TraceContext, input *spec.DeleteNetworkV4s, user *base_spec.UserAuthority) (data *spec.DeleteNetworkV4sData, code uint8, err error) {
	var specs []spec.NetworkV4
	if specs, err = resolver.ConvertToNetworkV4Specs(input.Spec); err != nil {
		code = base_const.CodeClientBadRequest
		return
	}
	if err = resolver.dbApi.DeleteNetworkV4s(tctx, specs, user); err != nil {
		code = base_const.CodeServerInternalError
		return
	}
	code = base_const.CodeOkDeleted
	data = &spec.DeleteNetworkV4sData{}
	return
}

func (resolver *Resolver) ConvertToNetworkV4Specs(specStr string) (specs []spec.NetworkV4, err error) {
	var baseSpecs []base_spec.Spec
	if err = json.Unmarshal([]byte(specStr), &baseSpecs); err != nil {
		return
	}

	for _, base := range baseSpecs {
		if base.Kind != "NetworkV4" {
			continue
		}
		var specBytes []byte
		if specBytes, err = json.Marshal(base.Spec); err != nil {
			return
		}
		var specData spec.NetworkV4
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
