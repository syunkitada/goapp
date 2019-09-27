package ctl_main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_model/index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
)

func (ctl *Ctl) index(args []string) error {
	var err error
	tctx := logger.NewCtlTraceContext(ctl.name)
	startTime := logger.StartTrace(tctx)
	defer func() { logger.EndTrace(tctx, startTime, err, 1) }()

	var ok bool
	var serviceName string
	if len(args) > 0 {
		serviceName = str_utils.ConvertToCamelFormat(args[0])
	} else {
		serviceName = ""
	}

	appUser := os.Getenv("APP_USER")
	appPassword := os.Getenv("APP_PASSWORD")
	appProject := os.Getenv("APP_PROJECT")
	ctl.client.SetProject(appProject)

	var loginData *base_spec.LoginData
	loginData, err = ctl.client.Login(tctx, &base_spec.Login{
		User:     appUser,
		Password: appPassword,
	})
	if err != nil {
		fmt.Printf("Failed Login: %v, %v\n", loginData, err)
		os.Exit(1)
	}

	if len(args) > 0 {
		if _, ok = loginData.Authority.ServiceMap[serviceName]; !ok {
			var project base_spec.ProjectService
			project, ok = loginData.Authority.ProjectServiceMap[appProject]
			if ok {
				_, ok = project.ServiceMap[serviceName]
			}
		}
	}

	if !ok {
		fmt.Printf("Usage: %s [SERVICE] [COMMAND] [FLAGS...]\n\n", ctl.name)

		fmt.Println("--- Available Services ---")
		snames := make([]string, 0, len(loginData.Authority.ServiceMap))
		for s := range loginData.Authority.ServiceMap {
			snames = append(snames, str_utils.ConvertToLowerFormat(s))
		}
		sort.Sort(sort.StringSlice(snames))
		for _, s := range snames {
			fmt.Println(s)
		}

		if project, ok := loginData.Authority.ProjectServiceMap[appProject]; ok {
			fmt.Println("\n--- Available Project Services ---")
			snames := make([]string, 0, len(loginData.Authority.ServiceMap))
			for s, _ := range project.ServiceMap {
				snames = append(snames, str_utils.ConvertToLowerFormat(s))
			}
			sort.Sort(sort.StringSlice(snames))
			for _, s := range snames {
				fmt.Println(s)
			}
		}
		return nil
	}

	// Get ServiceIndex, and exec cmd
	var getServiceIndexData *base_spec.GetServiceIndexData
	if getServiceIndexData, err = ctl.client.GetServiceIndex(tctx, &base_spec.GetServiceIndex{Name: serviceName}); err != nil {
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
	for query, cmd := range getServiceIndexData.Index.CmdMap {
		args := strings.Split(query, ".")
		helpQuery := strings.Join(args, " ")
		helpMsg := []string{helpQuery}
		flags := []string{}
		for f, flag := range cmd.FlagMap {
			sf := strings.Split(f, ",")
			if len(sf) == 2 {
				flags = append(flags, fmt.Sprintf("--%s, -%s [%s (%s)]", sf[0], sf[1], flag.FlagType, flag.FlagKind))
			} else {
				flags = append(flags, fmt.Sprintf("--%s [%s (%s)]", sf[0], flag.FlagType, flag.FlagKind))
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
				}
				break
			}
		}
	}

	params := map[string]interface{}{}
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
			if flag.Flag == base_const.ArgRequired {
				if !ok {
					cmdQuery = ""
					break
				}
			}
			if ok {
				switch cmdFlag.(type) {
				case string:
					if flag.FlagType == base_const.ArgTypeString {
						flagStr := cmdFlag.(string)
						switch flag.FlagKind {
						case "file":
							splitedFlag := strings.Split(flagStr, " ")
							data, err := json_utils.ReadFilesFromMultiPath(splitedFlag)
							if err != nil {
								return err
							}

							dataBytes, err := json_utils.Marshal(data)
							if err != nil {
								return err
							}
							params[key] = string(dataBytes)
						default:
							params[key] = cmdFlag.(string)
						}
					}
				case bool:
					if flag.FlagType == base_const.ArgTypeBool {
						params[key] = cmdFlag.(bool)
					}
				case int:
					if flag.FlagType == base_const.ArgTypeInt {
						params[key] = cmdFlag.(int)
					}
				}
			}
		}
	}

	_, isHelp := flagMap["help"]
	_, isShortHelp := shortFlagMap["h"]
	if cmdQuery == "" || isHelp || isShortHelp {
		fmt.Printf("Usage: %s [SERVICE] [COMMAND] [FLAGS...]\n\n", ctl.name)

		fmt.Printf("# Available Commands\n\n")

		sort.SliceStable(helpMsgs, func(i, j int) bool {
			return helpMsgs[i][0] < helpMsgs[j][0]
		})

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"COMMAND", "FLAGS"})
		table.SetBorder(false)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetRowLine(false)
		table.AppendBulk(helpMsgs)
		table.Render() // Send output

		return nil
	}

	if len(lastArgs) > 0 {
		specsBytes, err := json_utils.Marshal(lastArgs)
		if err != nil {
			return err
		}
		params["Args"] = string(specsBytes)
	}

	if ctl.baseConf.EnableDebug {
		fmt.Println("DEBUG params", params)
	}

	queries := []base_client.Query{
		base_client.Query{
			Name: str_utils.ConvertToCamelFormat(cmdQuery),
			Data: params,
		},
	}

	var resp Response
	if err = ctl.client.Request(tctx, serviceName, queries, &resp, true); err != nil {
		return err
	}
	ctl.output(&cmdInfo, &resp, flagMap, shortFlagMap)
	return nil
}

type Response struct {
	TraceId   string
	Code      uint8
	Error     string
	ResultMap map[string]Result
}

type Result struct {
	Data map[string]interface{}
}
