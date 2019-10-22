package os_utils

import (
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

type Route struct {
	Addr string
	Dev  string
	Via  string
}

func GetRouteMap(tctx *logger.TraceContext, netns string) (map[string]Route, error) {
	var out string
	var err error
	if netns == "" {
		out, err = exec_utils.Cmdf(tctx, "ip r")
		if err != nil {
			return nil, err
		}
	}

	routeMap := map[string]Route{}
	for _, line := range strings.Split(out, "\n") {
		if line != "" {
			splitedLine := strings.Split(line, " ")
			addr := ""
			dev := ""
			via := ""
			length := len(splitedLine)
			for i := 0; i < length; i++ {
				if i == 0 {
					addr = splitedLine[0]
				}
				if splitedLine[i] == "dev" {
					i += 1
					dev = splitedLine[i]
					continue
				}
				if splitedLine[i] == "via" {
					i += 1
					dev = splitedLine[i]
					continue
				}
			}
			routeMap[splitedLine[0]] = Route{
				Addr: addr,
				Dev:  dev,
				Via:  via,
			}
		}
	}

	return routeMap, nil
}
