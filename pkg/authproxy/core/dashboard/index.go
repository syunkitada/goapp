package dashboard

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model/authproxy_model_api"
	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
	"github.com/syunkitada/goapp/pkg/config"
)

type Dashboard struct {
	Conf              *config.Config
	Token             *auth.Token
	AuthproxyModelApi *authproxy_model_api.AuthproxyModelApi
}

func NewDashboard(conf *config.Config, authproxyModelApi *authproxy_model_api.AuthproxyModelApi, token *auth.Token) *Dashboard {
	dashboard := Dashboard{
		Conf:              conf,
		Token:             token,
		AuthproxyModelApi: authproxyModelApi,
	}
	return &dashboard
}
