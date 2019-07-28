package os_utils

import (
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func GetNetnsSet(tctx *logger.TraceContext) (map[string]bool, error) {
	out, err := exec_utils.Cmd(tctx, "ip netns")
	if err != nil {
		return nil, err
	}

	nsSet := map[string]bool{}
	for _, line := range strings.Split(out, "\n") {
		if line != "" {
			nsSet[strings.Split(line, " ")[0]] = true
		}
	}

	return nsSet, nil
}

func AddNetns(tctx *logger.TraceContext, netns string) error {
	_, err := exec_utils.Cmdf(tctx, "ip netns add %s", netns)
	return err
}
