package exec_utils

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func Cmd(tctx *logger.TraceContext, cmd string) (string, error) {
	cmd = fmt.Sprintf("timeout %d %s", tctx.GetTimeout(), cmd)
	cmds := strings.Split(cmd, " ")
	out, err := exec.Command(cmds[0], cmds[1:]...).CombinedOutput()
	if err != nil {
		logger.Warningf(tctx, "Failed Cmd: cmd=%s, out=%s, err=%s", cmd, string(out), err.Error())
	}
	return string(out), err
}

func Cmdf(tctx *logger.TraceContext, cmd string, args ...interface{}) (string, error) {
	return Cmd(tctx, fmt.Sprintf(cmd, args...))
}
