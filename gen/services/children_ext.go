package services

import (
	"errors"
	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/cli-program/gen/config"
	"github.com/kaytu-io/cli-program/gen/utils"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

func (g *Generator) ExtractChildren(root, servicePath string) error {
	hasPayload := map[string]bool{}

	childrenNames := map[string]interface{}{}
	err := filepath.Walk(servicePath, func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() || strings.HasSuffix(info.Name(), "_client.go") {
			return nil
		}

		name := info.Name()
		if strings.HasSuffix(name, "_parameters.go") {
			name = name[:len(name)-14]
			childrenNames[name] = struct{}{}
		} else if strings.HasSuffix(name, "_responses.go") {
			name = name[:len(name)-13]

			c, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			existPayload := strings.Contains(string(c), "GetPayload")
			hasPayload[name] = existPayload

			childrenNames[name] = struct{}{}
		} else {
			return errors.New("invalid filename: " + name)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for name := range childrenNames {
		apiName := strcase.ToCamel(name)
		apiName = strings.ReplaceAll(apiName, "ApiV", "APIV")
		apiName = strings.ReplaceAll(apiName, "Id", "ID")

		commandName := strings.ReplaceAll(name, "_", "-")
		commandName = config.UrlNames[commandName]
		if commandName == "" {
			return errors.New("url name not found for cmd: " + commandName)
		}

		fv := utils.ExtractFields(reflect.TypeOf(config.ParamModels[apiName]))
		output := utils.GenerateSetFieldsFromFlags(fv)
		outputFlags := utils.GenerateFlagDefinitions(strcase.ToCamel(name), fv)

		crud := "Get"
		if strings.HasPrefix(strings.ToLower(commandName), "create") ||
			strings.HasPrefix(strings.ToLower(commandName), "insert") {
			crud = "Create"
		}
		if strings.HasPrefix(strings.ToLower(commandName), "delete") ||
			strings.HasPrefix(strings.ToLower(commandName), "remove") {
			crud = "Delete"
		}
		if strings.HasPrefix(strings.ToLower(commandName), "update") ||
			strings.HasPrefix(strings.ToLower(commandName), "edit") {
			crud = "Update"
		}
		tmpl := Child{
			Parent:           g,
			CommandNameCamel: strcase.ToCamel(name),
			CommandNameSnake: strcase.ToSnake(name),
			CommandName:      commandName,
			APIName:          apiName,
			ParamString:      output,
			ParamFlags:       outputFlags,
			HasPayload:       hasPayload[name],
			CRUD:             crud,
		}
		g.Children = append(g.Children, tmpl)
	}
	//	for name, temp := range childrenMap {
	//		switch {
	//		case strings.Contains(temp.NameCamel, "Delete"):
	//			cmdReference += fmt.Sprintf(`
	//		Delete%sCmd.AddCommand(%sCmd)
	//`, strcase.ToCamel(serviceName), temp.NameCamel)
	//		case strings.Contains(temp.NameCamel, "Get"):
	//			cmdReference += fmt.Sprintf(`
	//		Get%sCmd.AddCommand(%sCmd)
	//`, strcase.ToCamel(serviceName), temp.NameCamel)
	//		case strings.Contains(temp.NameCamel, "Update"):
	//			cmdReference += fmt.Sprintf(`
	//		Update%sCmd.AddCommand(%sCmd)
	//`, strcase.ToCamel(serviceName), temp.NameCamel)
	//		case strings.Contains(temp.NameCamel, "Create"):
	//			cmdReference += fmt.Sprintf(`
	//		Create%sCmd.AddCommand(%sCmd)
	//`, strcase.ToCamel(serviceName), temp.NameCamel)
	//		default:
	//			cmdReference += fmt.Sprintf(`
	//		Get%sCmd.AddCommand(%sCmd)
	//`, strcase.ToCamel(serviceName), temp.NameCamel)
	//		}
	//
	//		fv := utils.ExtractFields(reflect.TypeOf(config.ParamModels[temp.APIName]))
	//		output := utils.GenerateFlagDefinitions(temp.NameCamel, fv)
	//		cmdReference += output
	//		//
	//		//for _, param := range temp.Params {
	//		//	cmdReference += fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")
	//		//`, temp.NameCamel, param.FlagName)
	//		//}
	//	}

	return nil
}
