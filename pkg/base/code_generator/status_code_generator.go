package code_generator

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/syunkitada/goapp/pkg/base/base_spec"
)

func GenerateStatusCodes() {
	specType := reflect.ValueOf(base_spec.CodeMap).Type()
	specPkgPath := specType.PkgPath()
	splitedPkg := strings.Split(specPkgPath, "/")

	pkgName := splitedPkg[2]

	pwd := os.Getenv("PWD")
	splitedPwd := strings.Split(pwd, "/")
	splitedPwd[0] = "/"
	pkgIndex := -1
	for i, dir := range splitedPwd {
		if dir == pkgName {
			pkgIndex = i
			break
		}
	}
	if pkgIndex == -1 {
		log.Fatal("Invalid PWD: you should be in pkg repository")
	}
	pkgPath := filepath.Join(splitedPwd[:pkgIndex+1]...)

	baseCodesTemplatePath := filepath.Join(pkgPath, "templates", "base_const", "codes.go.tmpl")
	baseCodesPkgDir := filepath.Join(pkgPath, "pkg", "base", "base_const")
	generateCodeFromTemplate(baseCodesTemplatePath, baseCodesPkgDir, "codes.go", base_spec.CodeMap)

	dashboardCodesTemplatePath := filepath.Join(pkgPath, "templates", "dashboard", "codes", "index.tsx.tmpl")
	dashboardCodesPkgDir := filepath.Join(pkgPath, "dashboard", "src", "lib", "codes")
	generateCodeFromTemplate(dashboardCodesTemplatePath, dashboardCodesPkgDir, "index.tsx", base_spec.CodeMap)
}
