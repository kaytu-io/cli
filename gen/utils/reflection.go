package utils

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/cli-program/gen/swagger"
	"reflect"
	"strings"
)

type Field struct {
	Name       string
	Type       string
	Children   []Field
	IsEnum     bool
	IsRequired bool
}

func fixParamName(name string) string {
	return strings.ReplaceAll(strcase.ToCamel(name), "Id", "ID")
}

func ExtractFields(swag *swagger.Swagger, parentType string, x reflect.Type) (resp Field) {
	for i := 0; i < x.NumField(); i++ {
		field := x.Field(i)
		fieldType := field.Type
		if field.Name == "Context" || field.Name == "HTTPClient" {
			continue
		}
		if !field.IsExported() {
			continue
		}

		for fieldType.Kind() == reflect.Slice {
			fieldType = fieldType.Elem()
		}

		for fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		child := Field{
			Name:     field.Name,
			Type:     field.Type.String(),
			Children: nil,
		}

		typeName := strings.TrimPrefix(strings.Trim(parentType, "*[]"), "models.")
		if strings.HasSuffix(typeName, "Params") {
			api := swag.GetAPI(typeName)
			for _, param := range api.Parameters {
				if fixParamName(param.Name) == child.Name {
					child.IsRequired = param.Required
				}
			}
			if api == nil {
				panic(fmt.Sprintf("failed to figure out api, %s", typeName))
			}
		} else {
			model := swag.GetModel(typeName)
			for _, param := range model.Parameters {
				if fixParamName(param.Name) == child.Name {
					child.IsRequired = param.Required
				}
			}
			if model == nil {
				panic(fmt.Sprintf("failed to figure out model, %s", typeName))
			}
		}

		if fieldType.Kind() == reflect.Struct {
			if !field.Anonymous {
				child.Children = append(child.Children, ExtractFields(swag, field.Type.String(), fieldType).Children...)
			} else {
				resp.Children = append(resp.Children, ExtractFields(swag, field.Type.String(), fieldType).Children...)
			}
		}

		if fieldType.Kind() != reflect.Struct && fieldType.Kind().String() != fieldType.Name() {
			child.IsEnum = true
		}

		if !field.Anonymous {
			resp.Children = append(resp.Children, child)
		}
	}
	return
}
