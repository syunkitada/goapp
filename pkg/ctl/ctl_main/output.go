package ctl_main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

func (ctl *CtlMain) outputCmdHelp(cmd string, cmdInfo index_model.Cmd) {
	cmdHelp := cmd
	if cmdInfo.Arg != "" {
		cmdHelp += fmt.Sprintf(" [%s:%s]", cmdInfo.ArgType, cmdInfo.Arg)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetBorder(false)
	table.SetRowLine(false)
	table.SetColumnSeparator("")
	table.SetCenterSeparator("")
	for key, flag := range cmdInfo.FlagMap {
		table.Append([]string{cmdHelp, cmdInfo.Help})
		table.Append([]string{fmt.Sprintf("--%s [%s:%s]", key, flag.FlagType, flag.Flag), flag.Help})
	}
	table.Render()

	// fmt.Printf("Invalid args: %s %s %v :%s\n", cmd, cmdInfo.Arg, cmdInfo.FlagMap, cmdInfo.Help)
}

func (ctl *CtlMain) output(cmdInfo *index_model.Cmd, resp *authproxy_model.ActionResponse, flagMap map[string]interface{}) {
	fmt.Printf("ResponseStatus: %d %s\n", resp.Tctx.StatusCode, resp.Tctx.Err)
	for key, data := range resp.Data {
		fmt.Printf("# %s\n", key)

		table := tablewriter.NewWriter(os.Stdout)
		tableHeader := cmdInfo.TableHeader
		table.SetHeader(tableHeader)

		switch data := data.(type) {
		case []interface{}:
			for _, raw := range data {
				switch raw := raw.(type) {
				case map[string]interface{}:
					r := make([]string, len(tableHeader))
					for i, head := range tableHeader {
						if v, ok := raw[head]; ok {
							r[i] = fmt.Sprint(v)
						} else {
							r[i] = "None"
						}
					}
					table.Append(r)
				}
			}
		}

		table.Render()
	}
}
