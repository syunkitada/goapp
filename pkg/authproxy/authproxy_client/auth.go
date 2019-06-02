package authproxy_client

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type ResponseLogin struct {
	Token     string
	Name      string
	Authority *authproxy_model.UserAuthority
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
