package os_utils

import (
	"encoding/gob"
	"fmt"
	"os"
	"strings"

	"github.com/syunkitada/goapp/pkg/lib/exec_utils"
	"github.com/syunkitada/goapp/pkg/lib/logger"
)

const (
	FileTypeXz    = "xz"
	FileTypeQcow2 = "qcow2"
)

func DetectFileType(tctx *logger.TraceContext, filePath string) (fileType string, err error) {
	if fileType, err = exec_utils.Cmdf(tctx, "file %s", filePath); err != nil {
		return
	}
	splitedResult := strings.Split(fileType, " ")
	if len(splitedResult) == 0 {
		err = fmt.Errorf("Unknown File Type")
		return
	}
	switch splitedResult[1] {
	case "XZ":
		fileType = FileTypeXz
	case "QEMU":
		if splitedResult[4] == "(v2)" {
			fileType = FileTypeQcow2
		}
	default:
		fmt.Println("DEBUGT File")
		err = fmt.Errorf("Unknown File Type: %s", fileType)
	}

	return
}

func UnArchiveFile(tctx *logger.TraceContext, filePath string) (outputPath string, err error) {
	var fileType string
	if fileType, err = DetectFileType(tctx, filePath); err != nil {
		return
	}
	switch fileType {
	case FileTypeXz:
		if _, err = exec_utils.Cmdf(tctx, "file %s", filePath); err != nil {
			return
		}
		outputPath = strings.Split(filePath, ".xz")[0]
	default:
		outputPath = filePath
	}
	return
}

func SaveDataFileIfNotExist(tctx *logger.TraceContext, filePath string, data interface{}) (err error) {
	if _, tmpErr := os.Stat(filePath); tmpErr != nil {
		if os.IsNotExist(tmpErr) {
			f, tmpErr := os.Create(filePath)
			if tmpErr != nil {
				err = tmpErr
				return
			}
			defer f.Close()
			enc := gob.NewEncoder(f)

			if err = enc.Encode(data); err != nil {
				return
			}
		} else {
			err = tmpErr
		}
	}
	return
}

func LoadDataFile(tctx *logger.TraceContext, filePath string, data interface{}) (err error) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()

	dec := gob.NewDecoder(f)
	if err = dec.Decode(data); err != nil {
		return
	}
	return
}
