package ctl_main

import (
	"fmt"
	"strings"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (ctl *CtlMain) Index(args []string) error {
	var err error
	tctx := logger.NewCtlTraceContext(ctl.name)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var ok bool
	var serviceName string
	if len(args) > 0 {
		serviceName = args[0]
	} else {
		serviceName = ""
	}

	var resp *authproxy_client.ResponseLogin
	// This should get userauthority
	if resp, err = ctl.client.Login(tctx, serviceName); err != nil {
		return err
	}

	if len(args) > 0 {
		serviceName := args[0]
		tmpServiceNames := []string{}
		splitedServiceName := strings.Split(serviceName, ".")
		for _, str := range splitedServiceName {
			tmpServiceNames = append(tmpServiceNames, strings.ToUpper(str[:1])+strings.ToLower(str[1:]))
		}
		serviceName = strings.Join(tmpServiceNames, ".")
		if _, ok = resp.Authority.ServiceMap[serviceName]; !ok {
			var project authproxy_model.ProjectService
			project, ok = resp.Authority.ProjectServiceMap[ctl.conf.Ctl.Project]
			if ok {
				_, ok = project.ServiceMap[serviceName]
			}
		}
	}

	if !ok {
		fmt.Println("--- Available Services ---")
		for serviceName, _ := range resp.Authority.ServiceMap {
			fmt.Println(strings.ToLower(serviceName))
		}
		if project, ok := resp.Authority.ProjectServiceMap[ctl.conf.Ctl.Project]; ok {
			fmt.Println("--- Available Project Services ---")
			for serviceName, _ := range project.ServiceMap {
				fmt.Println(strings.ToLower(serviceName))
			}
		}
		return nil
	}

	// fmt.Println("ActionProjectService", resp.Authority.ActionProjectService)

	return nil
}
