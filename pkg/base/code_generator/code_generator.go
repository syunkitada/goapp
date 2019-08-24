package code_generator

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
	"text/template"

	"github.com/syunkitada/goapp/pkg/base/base_model"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
)

func Generate(spec *base_model.Spec) {
	specType := reflect.TypeOf(spec.Meta)
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
	apiTemplatePath := filepath.Join(pkgPath, "templates", "genpkg", "api.go.tmpl")
	specPaths := append(splitedPwd[:pkgIndex], splitedPkg[2:]...)
	specPath := filepath.Join(specPaths...)
	pkgDir := filepath.Join(specPath, "genpkg")
	os_utils.MustMkdir(pkgDir, 0755)

	fmt.Println("DEBUG")
	fmt.Println(apiTemplatePath)
	fmt.Println(specPath)
	fmt.Println(pkgDir)
	for i, api := range spec.Apis {
		convertApi(&api)
		spec.Apis[i] = api
	}

	fmt.Println(spec.Apis)
	generateCodeFromTemplate(apiTemplatePath, pkgDir, "api.go", spec)
}

func generateCodeFromTemplate(templatePath string, pkgDir string, outputFile string, spec *base_model.Spec) {
	t := template.Must(template.ParseFiles(templatePath))
	filePath := filepath.Join(pkgDir, outputFile)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(f, spec); err != nil {
		log.Fatal(err)
	}
	cmd := "goimports -w " + filePath
	if out, err := exec.Command("sh", "-c", cmd).CombinedOutput(); err != nil {
		log.Fatalf("Failed cmd: %s, out=%s, err=%v", cmd, out, err)
	}
	fmt.Printf("Generated: %s\n", filePath)
}

func convertApi(api *base_model.Api) {
	queries := []base_model.Query{}
	for _, queryModel := range api.QueryModels {
		modelType := reflect.TypeOf(queryModel.Model)
		pkgPath := modelType.PkgPath()
		splitedPath := strings.Split(pkgPath, "/")
		pkgName := splitedPath[len(splitedPath)-1]
		name := modelType.Name()

		queries = append(queries, base_model.Query{
			RequiredAuth: queryModel.RequiredAuth,
			PkgPath:      pkgPath,
			PkgName:      pkgName,
			Name:         name,
		})
	}

	api.Queries = queries
}
