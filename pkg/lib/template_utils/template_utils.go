package template_utils

import (
	"os"
	"text/template"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func Template(tctx *logger.TraceContext, path string, perm os.FileMode, templatePath string, data interface{}) (err error) {
	var file *os.File
	file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return
	}
	defer func() {
		if tmpErr := file.Close(); tmpErr != nil {
			logger.Warningf(tctx, "Failed file.Close %s: %s", path, tmpErr.Error())
		}
	}()
	tmpl := template.Must(template.ParseFiles(templatePath))
	err = tmpl.Execute(file, data)
	return
}
