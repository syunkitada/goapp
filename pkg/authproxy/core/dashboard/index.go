package dashboard

import (
	"github.com/syunkitada/goapp/pkg/authproxy/core/auth"
	"github.com/syunkitada/goapp/pkg/authproxy/model/model_api"
	"github.com/syunkitada/goapp/pkg/config"
)

type Dashboard struct {
	Conf     *config.Config
	Token    *auth.Token
	ModelApi *model_api.ModelApi
}

func NewDashboard(conf *config.Config, modelApi *model_api.ModelApi, token *auth.Token) *Dashboard {
	dashboard := Dashboard{
		Conf:     conf,
		Token:    token,
		ModelApi: modelApi,
	}
	return &dashboard
}
