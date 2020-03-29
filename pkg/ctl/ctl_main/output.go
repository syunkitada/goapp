package ctl_main

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"

	"github.com/syunkitada/goapp/pkg/base/base_index_model"
	"github.com/syunkitada/goapp/pkg/lib/json_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func (ctl *Ctl) output(cmdInfo *base_index_model.Cmd, resp *Response,
	flagMap map[string]interface{}, shortFlagMap map[string]interface{}) {
	outputFormat, ok := flagMap["out"]
	if !ok {
		outputFormat, ok = shortFlagMap["o"]
		if !ok {
			outputFormat = "table"
		}
	}

	switch outputFormat {
	case "json":
		outs := []string{}
		for _, queryData := range resp.ResultMap {
			dataBytes, err := json_utils.Marshal(queryData)
			if err != nil {
				logger.StdoutFatalf("Failed json marshal: %v", err)
			}
			outs = append(outs, string(dataBytes))
		}
		fmt.Println(strings.Join(outs, "\n"))

	case "yaml":
		outs := []string{}
		for _, queryData := range resp.ResultMap {
			dataBytes, err := json_utils.YamlMarshal(queryData)
			if err != nil {
				logger.StdoutFatalf("Failed json marshal: %v", err)
			}
			outs = append(outs, string(dataBytes))
		}
		fmt.Println(strings.Join(outs, "\n"))

	default:
		fmt.Printf("# Status: code=%d, traceId=%s\n", resp.Code, resp.TraceId)
		if resp.Error != "" {
			fmt.Printf("Error: %s\n", resp.Error)
			return
		}

		for query, result := range resp.ResultMap {
			fmt.Printf("# Status: query=%s, code=%d, err=%s\n", query, result.Code, result.Error)
			for key, data := range result.Data {
				fmt.Printf("# %s\n", key)

				switch cmdInfo.OutputKind {
				case "table":
					table := tablewriter.NewWriter(os.Stdout)
					tableHeader := strings.Split(cmdInfo.OutputFormat, ",")
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
					case map[string]interface{}:
						r := make([]string, len(tableHeader))
						for i, head := range tableHeader {
							if v, ok := data[head]; ok {
								r[i] = fmt.Sprint(v)
							} else {
								r[i] = "None"
							}
						}
						table.Append(r)
					}
					table.Render()
				default:
					fmt.Println(data)
				}
			}
		}
	}
}
