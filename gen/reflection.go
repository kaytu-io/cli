package main

import (
	"fmt"
	"reflect"
)

type Field struct {
	Name     string
	Type     string
	Children []Field
	IsEnum   bool
}

func ExtractFields(x reflect.Type) (resp Field) {
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
		fmt.Println(field.Name, fieldType.Name(), fieldType.Kind())
		if fieldType.Kind() == reflect.Struct {
			if !field.Anonymous {
				child.Children = append(child.Children, ExtractFields(fieldType).Children...)
			} else {
				resp.Children = append(resp.Children, ExtractFields(fieldType).Children...)
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
