package resource_model_api

import (
	"github.com/syunkitada/goapp/pkg/config"
)

type ResourceModelApi struct {
	Conf *config.Config
}

func NewResourceModelApi(conf *config.Config) *ResourceModelApi {
	modelApi := ResourceModelApi{
		Conf: conf,
	}

	return &modelApi
}
