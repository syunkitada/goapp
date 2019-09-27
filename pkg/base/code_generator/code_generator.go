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

	"github.com/syunkitada/goapp/pkg/base/base_model/spec_model"
	"github.com/syunkitada/goapp/pkg/lib/os_utils"
	"github.com/syunkitada/goapp/pkg/lib/str_utils"
)

func Generate(spec *spec_model.Spec) {
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
	specPaths := append(splitedPwd[:pkgIndex], splitedPkg[2:]...)
	specPath := filepath.Join(specPaths...)
	pkgDir := filepath.Join(specPath, "genpkg")
	os_utils.MustMkdir(pkgDir, 0755)

	fmt.Println("DEBUG")
	fmt.Println(specPath)
	fmt.Println(pkgDir)
	spec.QuerySet = map[string]spec_model.Query{}
	for i, api := range spec.Apis {
		convertApi(&api)
		spec.Apis[i] = api
		for _, query := range api.Queries {
			spec.QuerySet[query.Name] = query
		}
	}

	fmt.Println(spec.Apis)
	apiTemplatePath := filepath.Join(pkgPath, "templates", "genpkg", "api.go.tmpl")
	generateCodeFromTemplate(apiTemplatePath, pkgDir, "api.go", spec)

	cmdTemplatePath := filepath.Join(pkgPath, "templates", "genpkg", "cmd.go.tmpl")
	generateCodeFromTemplate(cmdTemplatePath, pkgDir, "cmd.go", spec)
}

func generateCodeFromTemplate(templatePath string, pkgDir string, outputFile string, spec interface{}) {
	t := template.Must(template.ParseFiles(templatePath))
	filePath := filepath.Join(pkgDir, outputFile)
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := t.Execute(f, spec); err != nil {
		log.Fatal(err)
	}
	if len(outputFile) == strings.LastIndex(outputFile, ".go")+3 {
		cmd := "goimports -w " + filePath
		if out, err := exec.Command("sh", "-c", cmd).CombinedOutput(); err != nil {
			log.Fatalf("Failed cmd: %s, out=%s, err=%v", cmd, out, err)
		}
	}
	fmt.Printf("Generated: %s\n", filePath)
}

func convertApi(api *spec_model.Api) {
	queries := []spec_model.Query{}
	for _, queryModel := range api.QueryModels {
		reqType := reflect.TypeOf(queryModel.Req)
		pkgPath := reqType.PkgPath()
		splitedPath := strings.Split(pkgPath, "/")
		pkgName := splitedPath[len(splitedPath)-1]
		name := reqType.Name()
		actionName, dataName := str_utils.SplitActionDataName(name)

		flags := []spec_model.Flag{}
		lenFields := reqType.NumField()
		for i := 0; i < lenFields; i++ {
			f := reqType.Field(i)
			required := false
			if validate, ok := f.Tag.Lookup("validate"); ok {
				splitedValidate := strings.Split(validate, ",")
				for _, tag := range splitedValidate {
					if tag == "required" {
						required = true
					}
				}
			}
			flagName := str_utils.ConvertToLowerFormat(f.Name)
			shortName := flagName[:1]
			if flag, ok := f.Tag.Lookup("flag"); ok {
				splitedFlag := strings.Split(flag, ",")
				flagName = splitedFlag[0]
				if len(splitedFlag) > 1 {
					shortName = splitedFlag[1]
				}
			}
			if shortName != "" {
				flagName += "," + shortName
			}

			flagKind, ok := f.Tag.Lookup("flagKind")
			if !ok {
				flagKind = ""
			}
			flagType := f.Type.String()
			flags = append(flags, spec_model.Flag{
				Name:     f.Name,
				FlagName: flagName,
				FlagType: flagType,
				FlagKind: flagKind,
				Required: required,
			})
		}

		repType := reflect.TypeOf(queryModel.Rep)
		repLenFields := repType.NumField()
		outputKind := ""
		outputFormat := ""
		for i := 0; i < repLenFields; i++ {
			f := repType.Field(i)
			switch f.Type.Kind() {
			case reflect.Slice:
				elem := f.Type.Elem()
				lenFields := elem.NumField()
				outputKind = "table"
				columns := []string{}
				for j := 0; j < lenFields; j++ {
					c := elem.Field(j)
					columns = append(columns, c.Name)
				}
				outputFormat = strings.Join(columns, ",")
			}
		}

		if !queryModel.RequiredAuth {
			queryModel.RequiredAuth = api.RequiredAuth
		}
		if !queryModel.RequiredProject {
			queryModel.RequiredProject = api.RequiredProject
		}

		queries = append(queries, spec_model.Query{
			RequiredAuth:    queryModel.RequiredAuth,
			RequiredProject: queryModel.RequiredProject,
			PkgPath:         pkgPath,
			PkgName:         pkgName,
			Name:            name,
			ActionName:      actionName,
			DataName:        dataName,
			CmdName:         str_utils.ConvertToLowerFormat(name),
			CmdOutputKind:   outputKind,
			CmdOutputFormat: outputFormat,
			Flags:           flags,
		})
	}

	api.Queries = queries
}
