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
	"unicode"

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
	specPaths := append(splitedPwd[:pkgIndex], splitedPkg[2:]...)
	specPath := filepath.Join(specPaths...)
	pkgDir := filepath.Join(specPath, "genpkg")
	os_utils.MustMkdir(pkgDir, 0755)

	fmt.Println("DEBUG")
	fmt.Println(specPath)
	fmt.Println(pkgDir)
	for i, api := range spec.Apis {
		convertApi(&api)
		spec.Apis[i] = api
	}

	fmt.Println(spec.Apis)
	apiTemplatePath := filepath.Join(pkgPath, "templates", "genpkg", "api.go.tmpl")
	generateCodeFromTemplate(apiTemplatePath, pkgDir, "api.go", spec)

	cmdTemplatePath := filepath.Join(pkgPath, "templates", "genpkg", "cmd.go.tmpl")
	generateCodeFromTemplate(cmdTemplatePath, pkgDir, "cmd.go", spec)
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

		cmdRunes := []rune{}
		for i, r := range name {
			if i == 0 {
				cmdRunes = append(cmdRunes, unicode.ToLower(r))
				continue
			}
			if unicode.IsUpper(r) {
				cmdRunes = append(cmdRunes, '_', unicode.ToLower(r))
			} else {
				cmdRunes = append(cmdRunes, r)
			}
		}

		flags := []base_model.Flag{}
		lenFields := modelType.NumField()
		for i := 0; i < lenFields; i++ {
			f := modelType.Field(i)
			required := false
			if binding, ok := f.Tag.Lookup("binding"); ok && binding == "required" {
				required = true
			}
			flagName := strings.ToLower(f.Name)
			shortName := strings.ToLower(f.Name[:1])
			if flag, ok := f.Tag.Lookup("flag"); ok {
				splitedFlag := strings.Split(flag, ",")
				flagName = splitedFlag[0]
				if len(splitedFlag) > 1 {
					shortName = splitedFlag[1]
				}
			}

			flagKind, ok := f.Tag.Lookup("flagKind")
			if !ok {
				flagKind = ""
			}
			flagType := f.Type.String()
			flags = append(flags, base_model.Flag{
				Name:      f.Name,
				FlagName:  flagName,
				ShortName: shortName,
				FlagType:  flagType,
				FlagKind:  flagKind,
				Required:  required,
			})
		}

		queries = append(queries, base_model.Query{
			RequiredAuth: queryModel.RequiredAuth,
			PkgPath:      pkgPath,
			PkgName:      pkgName,
			Name:         name,
			CmdName:      string(cmdRunes),
			Flags:        flags,
		})
	}

	api.Queries = queries
}
