package ctl_main

import (
	"fmt"
	"strings"

	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_client"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
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
		tmpServiceNames := []string{}
		splitedServiceName := strings.Split(serviceName, ".")
		for _, str := range splitedServiceName {
			tmpServiceNames = append(tmpServiceNames, strings.ToUpper(str[:1])+strings.ToLower(str[1:]))
		}
		serviceName = strings.Join(tmpServiceNames, ".")
	} else {
		serviceName = ""
	}

	var resp *authproxy_client.ResponseLogin
	// This should get userauthority
	if resp, err = ctl.client.Login(tctx, serviceName); err != nil {
		return err
	}

	if len(args) > 0 {
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

	var indexResp *authproxy_client.ResponseGetIndex
	if indexResp, err = ctl.client.GetIndex(tctx, resp.Token, serviceName); err != nil {
		return err
	}

	cmdArgs := []string{}
	flagMap := map[string]interface{}{}
	lastIndex := len(args) - 1
	isFlag := false
	for index, arg := range args {
		if strings.Index(arg, "--") == 0 {
			if len(arg) == 2 {
				cmdArgs = append(cmdArgs, strings.Join(args[index+1:], " "))
				break
			}

			// bool
			if index == lastIndex || strings.Index(args[index+1], "--") == 0 {
				flagMap[arg[2:]] = true
				continue
			}

			flagMap[arg[2:]] = args[index+1]
			isFlag = true
		} else {
			if isFlag {
				isFlag = false
				continue
			}
			cmdArgs = append(cmdArgs, arg)
		}
	}

	argsMap := map[string]index_model.Cmd{}
	cmdQuery := ""
	var cmdInfo index_model.Cmd
	argsStr := ""
	lastArgs := []string{}
	helpMsg := ""
	for query, cmd := range indexResp.Index.CmdMap {
		args := []rune{}
		for i, c := range query {
			if c >= 'A' && c <= 'Z' {
				c += 'a' - 'A'
				if i != 0 {
					args = append(args, ' ')
				}
			}
			args = append(args, c)
		}
		argsStr = string(args)
		argsMap[argsStr] = cmd

		if cmd.Arg != "" {
			helpMsg += fmt.Sprintf("%s [%s:%s]  :%s\n", argsStr, cmd.ArgType, cmd.Arg, cmd.Help)
		} else {
			helpMsg += fmt.Sprintf("%s  :%s\n", argsStr, cmd.Help)
		}

		if len(cmdArgs) < 2 {
			continue
		}

		splitedArgs := strings.Split(argsStr, " ")
		if len(cmdArgs)+1 >= len(splitedArgs) {
			isMatch := true
			for i, arg := range splitedArgs {
				if arg != cmdArgs[i+1] {
					isMatch = false
					break
				}
			}
			if isMatch {
				cmdQuery = query
				cmdInfo = cmd
				if len(cmdArgs)+1 > len(splitedArgs) {
					if len(cmdArgs) > len(splitedArgs)+1 {
						lastArgs = cmdArgs[len(splitedArgs)+1:]
					}

					if cmd.Arg == "required" && len(lastArgs) == 0 {
						ctl.outputCmdHelp(argsStr, cmd)
						return nil
					}
				}
				break
			}
		}
	}

	if cmdQuery == "" {
		fmt.Printf("# Available Commands\n-----------------------\n")
		fmt.Println(helpMsg)
		return nil
	}

	strParams := map[string]string{}
	if cmdInfo.FlagMap != nil {
		for key, flag := range cmdInfo.FlagMap {
			cmdFlag, ok := flagMap[key]
			if flag.Flag == index_model.ArgRequired {
				if !ok {
					ctl.outputCmdHelp(argsStr, cmdInfo)
					return nil
				}
			}
			if flag.FlagType == index_model.ArgTypeString {
				strParams[key] = cmdFlag.(string)
			}
		}
	}

	if cmdInfo.ArgType == "file" && len(lastArgs) > 0 {
		data, err := json_utils.ReadFilesFromMultiPath(lastArgs)
		if err != nil {
			return err
		}

		specs := []interface{}{}
		for _, d := range data {
			if _, ok := d["Kind"]; ok && d["Kind"].(string) == cmdInfo.ArgKind {
				if spec, ok := d["Spec"]; ok {
					specs = append(specs, spec)
				}
			}
		}

		specsBytes, err := json_utils.Marshal(specs)
		if err != nil {
			return err
		}

		queries := []authproxy_model.Query{
			authproxy_model.Query{
				Kind: cmdQuery,
				StrParams: map[string]string{
					"Specs": string(specsBytes),
				},
			},
		}

		var tmpResp *authproxy_model.ActionResponse
		if tmpResp, err = ctl.client.Action(tctx, resp.Token, serviceName, queries); err != nil {
			return err
		}

		ctl.output(&cmdInfo, tmpResp, flagMap)
		return nil
	}

	queries := []authproxy_model.Query{
		authproxy_model.Query{
			Kind:      cmdQuery,
			StrParams: strParams,
		},
	}

	var tmpResp *authproxy_model.ActionResponse
	if tmpResp, err = ctl.client.Action(tctx, resp.Token, serviceName, queries); err != nil {
		return err
	}
	ctl.output(&cmdInfo, tmpResp, flagMap)

	return nil
}
