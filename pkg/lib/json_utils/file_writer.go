package json_utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func WriteFile(tctx *logger.TraceContext, filePath string, data interface{}, perm os.FileMode) (err error) {
	var bytes []byte
	if bytes, err = json.Marshal(data); err != nil {
		return
	}
	err = ioutil.WriteFile(filePath, bytes, perm)
	return
}
