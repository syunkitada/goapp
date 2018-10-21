package auth

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/config"
)

type Auth struct {
	Conf              *config.Config
	AuthproxyModelApi *authproxy_model_api.AuthproxyModelApi
	Token             *Token
}

func NewAuth(conf *config.Config, authproxyModelApi *authproxy_model_api.AuthproxyModelApi, token *Token) *Auth {
	auth := Auth{
		Conf:              conf,
		AuthproxyModelApi: authproxyModelApi,
		Token:             token,
	}
	return &auth
}
