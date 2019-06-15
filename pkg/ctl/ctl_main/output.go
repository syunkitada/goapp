package ctl_main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/syunkitada/goapp/pkg/authproxy/authproxy_model"
	"github.com/syunkitada/goapp/pkg/authproxy/index_model"
)

func (ctl *CtlMain) output(cmdInfo *index_model.Cmd, resp *authproxy_model.ActionResponse, flagMap map[string]interface{}) {
	fmt.Println(resp.Tctx.StatusCode, resp.Tctx.Err)
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
							r[i] = v.(string)
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
