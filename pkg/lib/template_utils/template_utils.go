package template_utils

import (
	"os"
	"text/template"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func Template(tctx *logger.TraceContext, path string, perm os.FileMode, templatePath string, data interface{}) error {
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			logger.Warningf(tctx, "Failed file.Close %s: %s", path, err.Error())
		}
	}()
	tmpl := template.Must(template.ParseFiles(templatePath))
	tmpl.Execute(file, data)
	return nil
}
