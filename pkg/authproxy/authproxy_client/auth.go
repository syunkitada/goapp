package authproxy_client

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type ResponseLogin struct {
	Token     string
	Name      string
	Authority *authproxy_model.UserAuthority
}

type ResponseGetIndex struct {
	Index index_model.Index
}

func (client *AuthproxyClient) Login(tctx *logger.TraceContext, serviceName string) (*ResponseLogin, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := authproxy_model.AuthRequest{
		Username: client.conf.Ctl.Username,
		Password: client.conf.Ctl.Password,
	}

	var resp ResponseLogin
	err = client.request(tctx, "auth/login", &req, &resp)
	return &resp, err
}

func (client *AuthproxyClient) GetIndex(tctx *logger.TraceContext, token string, serviceName string) (*ResponseGetIndex, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := authproxy_model.TokenAuthRequest{
		Token: token,
		Action: authproxy_model.ActionRequest{
			ProjectName: client.conf.Ctl.Project,
			ServiceName: serviceName,
			Queries: []authproxy_model.Query{
				authproxy_model.Query{Kind: "GetIndex"},
			},
		},
	}

	var resp ResponseGetIndex
	err = client.request(tctx, serviceName, &req, &resp)
	return &resp, err
}
