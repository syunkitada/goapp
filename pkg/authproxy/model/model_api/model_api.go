package model_api

import (
	"github.com/syunkitada/goapp/pkg/config"
)

type ModelApi struct {
	Conf *config.Config
}

func NewModelApi(conf *config.Config) *ModelApi {
	modelApi := ModelApi{
		Conf: conf,
	}

	return &modelApi
}
