package ctl_main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/signal"
	"sort"
	"strings"
	"sync"
	"syscall"

	"github.com/gorilla/websocket"
	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/base/base_client"
	"github.com/syunkitada/goapp/pkg/base/base_const"
	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/base/base_spec"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
	"golang.org/x/crypto/ssh/terminal"
)

func (ctl *Ctl) index(args []string) (err error) {
	tctx := logger.NewTraceContext(ctl.baseConf.Host, ctl.name)
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
		sort.Strings(snames)
		for _, s := range snames {
			fmt.Println(s)
		}

		if project, ok := loginData.Authority.ProjectServiceMap[appProject]; ok {
			fmt.Println("\n--- Available Project Services ---")
			snames := make([]string, 0, len(loginData.Authority.ServiceMap))
			for s := range project.ServiceMap {
				snames = append(snames, str_utils.ConvertToLowerFormat(s))
			}
			sort.Strings(snames)
			for _, s := range snames {
				fmt.Println(s)
			}
		}
		return
	}

	// Get ServiceIndex, and exec cmd
	var getServiceIndexData *base_spec.GetServiceIndexData
	if getServiceIndexData, err = ctl.client.GetServiceIndex(tctx, &base_spec.GetServiceIndex{Name: serviceName}); err != nil {
		return
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
	var cmdInfo base_index_model.Cmd
	lastArgs := []string{}
	helpMsgs := [][]string{}
	for query, cmd := range getServiceIndexData.Index.CmdMap {
		helpQuery := query
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

		sort.Strings(flags)
		helpMsg = append(helpMsg, strings.Join(flags, "\n"))
		helpMsgs = append(helpMsgs, helpMsg)

		if len(cmdArgs) < 2 {
			continue
		}

		if len(cmdArgs) > 1 && cmdArgs[1] == query {
			if len(cmdArgs) > 2 {
				lastArgs = cmdArgs[2:]
			}

			cmdInfo = cmd
			helpMsgs = [][]string{helpMsg}
			cmdQuery = query
			break
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
			if flag.Required {
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
		return
	}

	if len(lastArgs) > 0 {
		var specsBytes []byte
		specsBytes, err = json_utils.Marshal(lastArgs)
		if err != nil {
			return
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
	if !cmdInfo.Ws {
		if err = ctl.client.Request(tctx, serviceName, queries, &resp, true); err != nil {
			return
		}
		ctl.output(&cmdInfo, &resp, flagMap, shortFlagMap)
	} else {
		var wsConn *websocket.Conn
		if wsConn, err = ctl.client.RequestWs(tctx, serviceName, queries, &resp, true); err != nil {
			return
		}
		switch cmdInfo.Kind {
		case "Terminal":
			err = startTerminal(tctx, &cmdInfo, wsConn)
		default:
			fmt.Printf("Invalid Kind: %s\n", cmdInfo.Kind)
		}
	}
	return
}

func startTerminal(tctx *logger.TraceContext, cmdInfo *base_index_model.Cmd, wsConn *websocket.Conn) (err error) {
	fmt.Println("Start Terminal: If you want to stop terminal, input '^].'")
	chMutex := sync.Mutex{}
	isDone := false
	doneCh := make(chan bool, 2)
	readCh := make(chan string, 10)
	sigCh := make(chan os.Signal, 1)

	// enter raw mode
	fd := int(os.Stdin.Fd())
	state, tmpErr := terminal.MakeRaw(fd)
	if tmpErr != nil {
		err = fmt.Errorf("Failed MakeRaw: err=%s", tmpErr.Error())
		return
	}

	signal.Notify(sigCh, syscall.SIGTERM)

	defer func() {
		chMutex.Lock()
		if tmpErr := wsConn.Close(); tmpErr != nil {
			logger.Warningf(tctx, "Failed wsConn.Close: err=%s", tmpErr.Error())
		} else {
			logger.Info(tctx, "Success wsConn.Close")
		}
		isDone = true
		chMutex.Unlock()
	}()

	var isPreExitKey bool
	stdin := bufio.NewReader(os.Stdin)
	go func() {
		for {
			ch, tmpErr := stdin.ReadByte()
			if tmpErr == io.EOF {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed ReadByte: EOF")
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}

			// ^]. でexitする
			if ch == 29 { // 29 == ^]
				isPreExitKey = true
			} else if isPreExitKey && ch == 46 { // 46 == .
				chMutex.Lock()
				if !isDone {
					doneCh <- true
				}
				chMutex.Unlock()
				return
			} else {
				isPreExitKey = false
			}

			input := TerminalInput{
				Bytes: []byte(string(ch)),
			}
			outputJson, tmpErr := json.Marshal(&input)
			if tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed json.Marshal: %s", tmpErr.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			tmpErr = wsConn.WriteMessage(websocket.TextMessage, outputJson)
			if tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed Write: %s", tmpErr.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
		}
	}()

	go func() {
		for {
			_, message, tmpErr := wsConn.ReadMessage()
			if tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed ReadMessage: %s", tmpErr.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			var output TerminalOutput
			if tmpErr := json.Unmarshal(message, &output); tmpErr != nil {
				chMutex.Lock()
				if !isDone {
					err = fmt.Errorf("Failed Unmarshal read message: %s", err.Error())
					doneCh <- true
				}
				chMutex.Unlock()
				return
			}
			readCh <- string(output.Bytes)
		}
	}()

	for {
		select {
		case ch := <-sigCh:
			if tmpErr := terminal.Restore(fd, state); tmpErr != nil {
				fmt.Printf("\nFailed terminal.Restore: err=%s\n", tmpErr.Error())
			}
			fmt.Printf("\nExit by %s\n", ch.String())
			return
		case <-doneCh:
			if tmpErr := terminal.Restore(fd, state); tmpErr != nil {
				fmt.Printf("\nFailed terminal.Restore: err=%s\n", tmpErr.Error())
			}
			fmt.Printf("\nExit by doneCh\n")
			return
		case str := <-readCh:
			fmt.Print(str)
		}
	}
}

type Response struct {
	TraceId   string
	Code      uint8
	Error     string
	ResultMap map[string]Result
}

type Result struct {
	Code  uint8
	Error string
	Data  map[string]interface{}
}

type TerminalInput struct {
	Bytes []byte
}

type TerminalOutput struct {
	Bytes []byte
}
