package services

import (
	"errors"
	"fmt"
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

		parentType := reflect.TypeOf(config.ParamModels[apiName]).Name()
		typeName := strings.TrimPrefix(strings.Trim(parentType, "*[]"), "models.")
		api := g.Swagger.GetAPI(typeName)
		var apiDescription string
		if api != nil {
			apiDescription = strings.ReplaceAll(api.Description, "`", "'")
		} else {
			fmt.Println("failed to get api for:", typeName)
		}

		commandName := strings.ReplaceAll(name, "_", "-")
		commandName = config.UrlNames[commandName]
		if commandName == "" {
			return errors.New("url name not found for cmd: " + name)
		}

		fv := utils.ExtractFields(g.Swagger, parentType, false, reflect.TypeOf(config.ParamModels[apiName]))
		output := utils.GenerateSetFieldsFromFlags(fv)
		outputFlags := utils.GenerateFlagDefinitions(strcase.ToCamel(name), fv)

		crud := "Get"
		if strings.HasPrefix(strings.ToLower(commandName), "create") ||
			strings.HasPrefix(strings.ToLower(commandName), "insert") ||
			strings.HasPrefix(strings.ToLower(commandName), "trigger") {
			crud = "Create"
		} else if strings.HasPrefix(strings.ToLower(commandName), "delete") ||
			strings.HasPrefix(strings.ToLower(commandName), "remove") {
			crud = "Delete"
		} else if strings.HasPrefix(strings.ToLower(commandName), "update") ||
			strings.HasPrefix(strings.ToLower(commandName), "edit") {
			crud = "Update"
		} else if strings.HasPrefix(strings.ToLower(commandName), "list") {
			crud = "List"
		} else {
		}

		tmpl := Child{
			Parent:           g,
			CommandNameCamel: strcase.ToCamel(name),
			CommandNameSnake: strcase.ToSnake(name),
			CommandName:      commandName,
			APIName:          apiName,
			APIDescription:   apiDescription,
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
