package template_utils

import (
	"os"
	"strings"
	"text/template"

	"github.com/syunkitada/goapp/pkg/lib/logger"
)

func subnetmask(subnet string) string {
	splitedSubnet := strings.Split(subnet, "/")
	if len(splitedSubnet) == 2 {
		return splitedSubnet[1]
	}
	return subnet
}

var templateFuncMap = template.FuncMap{
	"subnetmask": subnetmask,
}

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

func NewTemplate(tctx *logger.TraceContext, templatePath string) (tmpl *template.Template, err error) {
	splitedTemplatePath := strings.Split(templatePath, "/")
	tmpl, err = template.New(splitedTemplatePath[len(splitedTemplatePath)-1]).Funcs(templateFuncMap).ParseFiles(templatePath)
	return
}

func ExecTemplate(tctx *logger.TraceContext, tmpl *template.Template,
	path string, perm os.FileMode, data interface{}) (err error) {
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
	err = tmpl.Execute(file, data)
	return
}
