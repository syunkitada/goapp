package ctl_main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
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

	// TODO Login
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
		fmt.Printf("Usage: %s [SERVICE] [COMMAND] [OPTION] [FLAGS...]\n\n", ctl.name)

		fmt.Println("--- Available Services ---")
		snames := make([]string, 0, len(resp.Authority.ServiceMap))
		for s := range resp.Authority.ServiceMap {
			snames = append(snames, strings.ToLower(s))
		}
		sort.Sort(sort.StringSlice(snames))
		for _, s := range snames {
			fmt.Println(s)
		}

		if project, ok := resp.Authority.ProjectServiceMap[ctl.conf.Ctl.Project]; ok {
			fmt.Println("\n--- Available Project Services ---")
			snames := make([]string, 0, len(resp.Authority.ServiceMap))
			for s, _ := range project.ServiceMap {
				snames = append(snames, strings.ToLower(s))
			}
			sort.Sort(sort.StringSlice(snames))
			for _, s := range snames {
				fmt.Println(s)
			}
		}
		return nil
	}

	// Get ServiceIndex, and exec cmd
	var indexResp *authproxy_client.ResponseGetIndex
	if indexResp, err = ctl.client.GetIndex(tctx, resp.Token, serviceName); err != nil {
		return err
	}

	cmdArgs := []string{}
	flagMap := map[string]interface{}{}
	shortFlagMap := map[string]interface{}{}
	lastIndex := len(args) - 1
	isFlag := false
	for index, arg := range args {
		if strings.Index(arg, "--") == 0 {
			if len(arg) == 2 {
				cmdArgs = append(cmdArgs, strings.Join(args[index+1:], " "))
				break
			}

			// bool
			if index == lastIndex || strings.Index(args[index+1], "-") == 0 {
				flagMap[arg[2:]] = true
				continue
			}

			flagMap[arg[2:]] = args[index+1]
			isFlag = true
		} else if strings.Index(arg, "-") == 0 {
			if len(arg) == 1 {
				cmdArgs = append(cmdArgs, strings.Join(args[index+1:], " "))
				break
			}

			// bool
			if index == lastIndex || strings.Index(args[index+1], "-") == 0 {
				shortFlagMap[arg[1:]] = true
				continue
			}

			shortFlagMap[arg[1:]] = args[index+1]
			isFlag = true
		} else {
			if isFlag {
				isFlag = false
				continue
			}
			cmdArgs = append(cmdArgs, arg)
		}
	}

	cmdQuery := ""
	var cmdInfo index_model.Cmd
	lastArgs := []string{}
	helpMsgs := [][]string{}
	for query, cmd := range indexResp.Index.CmdMap {
		args := strings.Split(query, "_")
		helpQuery := strings.Join(args, " ")
		var helpMsg []string
		if cmd.Arg != "" {
			helpMsg = []string{helpQuery, fmt.Sprintf("type=%s,kind=%s (%s)", cmd.ArgType, cmd.ArgKind, cmd.Arg)}
		} else {
			helpMsg = []string{helpQuery, "", cmd.Help}
		}
		flags := []string{}
		for f, flag := range cmd.FlagMap {
			sf := strings.Split(f, ",")
			if len(sf) == 2 {
				flags = append(flags, fmt.Sprintf("--%s, -%s [%s (%s)]", sf[0], sf[1], flag.FlagType, flag.Flag))
			} else {
				flags = append(flags, fmt.Sprintf("--%s [%s (%s)]", sf[0], flag.FlagType, flag.Flag))
			}
		}

		sort.Sort(sort.StringSlice(flags))
		helpMsg = append(helpMsg, strings.Join(flags, "\n"))
		helpMsgs = append(helpMsgs, helpMsg)

		if len(cmdArgs) < 2 {
			continue
		}

		if len(cmdArgs) > len(args) {
			isMatch := true
			for i, arg := range args {
				if arg != cmdArgs[i+1] {
					isMatch = false
					break
				}
			}
			if isMatch {
				helpMsgs = [][]string{helpMsg}
				cmdQuery = query
				cmdInfo = cmd
				if len(cmdArgs)+1 > len(args) {
					if len(cmdArgs) > len(args)+1 {
						lastArgs = cmdArgs[len(args)+1:]
					}

					if cmd.Arg == "required" && len(lastArgs) == 0 {
						cmdQuery = ""
						break
					}
				}
				break
			}
		}
	}

	strParams := map[string]string{}
	boolParams := map[string]bool{}
	intParams := map[string]int{}
	if cmdInfo.FlagMap != nil {
		for key, flag := range cmdInfo.FlagMap {
			splitedKey := strings.Split(key, ",")
			key = splitedKey[0]
			var cmdFlag interface{}
			var ok bool
			if len(splitedKey) == 1 {
				cmdFlag, ok = flagMap[key]
			} else {
				cmdFlag, ok = flagMap[key]
				if !ok {
					cmdFlag, ok = shortFlagMap[splitedKey[1]]
				}
			}
			if flag.Flag == index_model.ArgRequired {
				if !ok {
					cmdQuery = ""
					break
				}
			}
			if ok {
				switch cmdFlag.(type) {
				case string:
					if flag.FlagType == index_model.ArgTypeString {
						strParams[key] = cmdFlag.(string)
					}
				case bool:
					if flag.FlagType == index_model.ArgTypeBool {
						boolParams[key] = cmdFlag.(bool)
					}
				case int:
					if flag.FlagType == index_model.ArgTypeInt {
						intParams[key] = cmdFlag.(int)
					}
				}
			}
		}
	}

	if ctl.conf.Default.EnableDebug {
		fmt.Println("DEBUG flagMap", flagMap)
		fmt.Println("DEBUG shortFlagMap", shortFlagMap)
		fmt.Println("DEBUG lastArg", lastArgs)
	}

	_, isHelp := flagMap["help"]
	_, isShortHelp := shortFlagMap["h"]
	if cmdQuery == "" || isHelp || isShortHelp {
		fmt.Printf("Usage: %s [SERVICE] [COMMAND] [OPTION] [FLAGS...]\n\n", ctl.name)

		fmt.Printf("# Available Commands\n\n")

		sort.SliceStable(helpMsgs, func(i, j int) bool {
			return helpMsgs[i][0] < helpMsgs[j][0]
		})

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"COMMAND", "OPTION", "FLAGS"})
		table.SetBorder(false)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetRowLine(false)
		table.AppendBulk(helpMsgs)
		table.Render() // Send output

		return nil
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
	} else if len(lastArgs) > 0 {
		specsBytes, err := json_utils.Marshal(lastArgs)
		if err != nil {
			return err
		}
		strParams["Args"] = string(specsBytes)
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
