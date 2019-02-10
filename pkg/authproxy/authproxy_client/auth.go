package authproxy_client

import (
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type ResponseIssueToken struct {
	Token string
}

func (client *AuthproxyClient) IssueToken(tctx *logger.TraceContext) (*ResponseIssueToken, error) {
	var err error
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	req := authproxy_model.AuthRequest{
		Username: client.conf.Ctl.Username,
		Password: client.conf.Ctl.Password,
	}

	var resp ResponseIssueToken
	err = client.request(tctx, "token", &req, &resp)
	return &resp, err
}
