package os_utils

import (
	"os"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func Mkdir(path string, perm os.FileMode) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, perm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}
	return nil
}

func MustMkdir(path string, perm os.FileMode) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.Mkdir(path, perm); err != nil {
			logger.StdoutFatalf("Failed Mkdir: path=%s, err=%v", path, err)
		}
	} else if err != nil {
		logger.StdoutFatalf("Failed Mkdir: path=%s, err=%v", path, err)
	}
}

func PathExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
