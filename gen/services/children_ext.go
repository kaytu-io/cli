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
	"sort"
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
			return errors.New("url name not found for cmd: " + name)
		}

		fv := utils.ExtractFields(g.Swagger, reflect.TypeOf(config.ParamModels[apiName]).Name(), reflect.TypeOf(config.ParamModels[apiName]))
		output := utils.GenerateSetFieldsFromFlags(fv)
		outputFlags := utils.GenerateFlagDefinitions(strcase.ToCamel(name), fv)

		crud := "Get"
		if strings.HasPrefix(strings.ToLower(commandName), "create") ||
			strings.HasPrefix(strings.ToLower(commandName), "insert") ||
			strings.HasPrefix(strings.ToLower(commandName), "trigger") {
			crud = "Create"
			//commandName = strings.TrimPrefix(commandName, "create")
			//commandName = strings.TrimPrefix(commandName, "insert")
			//commandName = strings.TrimPrefix(commandName, "trigger")
		} else if strings.HasPrefix(strings.ToLower(commandName), "delete") ||
			strings.HasPrefix(strings.ToLower(commandName), "remove") {
			crud = "Delete"
			//commandName = strings.TrimPrefix(commandName, "remove")
			//commandName = strings.TrimPrefix(commandName, "delete")
		} else if strings.HasPrefix(strings.ToLower(commandName), "update") ||
			strings.HasPrefix(strings.ToLower(commandName), "edit") {
			crud = "Update"
			//commandName = strings.TrimPrefix(commandName, "update")
			//commandName = strings.TrimPrefix(commandName, "edit")
		} else if strings.HasPrefix(strings.ToLower(commandName), "list") {
			crud = "List"
			//commandName = strings.TrimPrefix(commandName, "list")
		} else {
			//commandName = strings.TrimPrefix(commandName, "get")
		}
		//commandName = strings.Trim(commandName, "-")

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

	sort.Slice(g.Children, func(i, j int) bool {
		return g.Children[i].CommandName < g.Children[j].CommandName
	})
	return nil
}
