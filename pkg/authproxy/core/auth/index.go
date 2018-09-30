package auth

import (
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
	"github.com/syunkitada/goapp/pkg/config"
)

type Auth struct {
	Conf     *config.Config
	ModelApi *model_api.ModelApi
	Token    *Token
}

func NewAuth(conf *config.Config, modelApi *model_api.ModelApi, token *Token) *Auth {
	auth := Auth{
		Conf:     conf,
		ModelApi: modelApi,
		Token:    token,
	}
	return &auth
}
