package ctl_main

import (
	"fmt"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (ctl *CtlMain) Index(args []string) error {
	var err error
	tctx := logger.NewCtlTraceContext(ctl.name)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var token *authproxy_client.ResponseIssueToken
	// This should get userauthority
	if token, err = ctl.client.IssueToken(tctx); err != nil {
		return err
	}

	fmt.Println(token)
	return nil
}
