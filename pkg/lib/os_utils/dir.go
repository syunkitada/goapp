package os_utils

import (
	"os"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func MustMkdir(path string, perm os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, perm); err != nil {
			logger.StdoutFatal(err)
		}
	} else if err != nil {
		logger.StdoutFatalf("Failed Mkdir: path=%s, err=%s", path, err.Error())
	}
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
