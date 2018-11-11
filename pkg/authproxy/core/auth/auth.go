package auth

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/config"
)

type Auth struct {
	conf              *config.Config
	authproxyModelApi *authproxy_model_api.AuthproxyModelApi
	token             *Token
}

func NewAuth(conf *config.Config, authproxyModelApi *authproxy_model_api.AuthproxyModelApi, token *Token) *Auth {
	auth := Auth{
		conf:              conf,
		authproxyModelApi: authproxyModelApi,
		token:             token,
	}
	return &auth
}
