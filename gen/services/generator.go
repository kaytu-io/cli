package services

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

//go:embed service.go.template
var templateString string

//go:embed children.go.template
var childrenTemplateString string

type Generator struct {
	ServiceFileName  string
	ServiceNameCamel string
	ServiceNameSnake string

	Children []Child
}

type Child struct {
	Parent           *Generator
	CommandNameCamel string
	CommandNameSnake string
	CommandName      string
	APIName          string
	ParamString      string
	ParamFlags       string
	HasPayload       bool
	CRUD             string
}

func (g Generator) Generate(root string) error {
	t := template.New("")
	t, err := t.Parse(templateString)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(root+"/cmd/gen/%s.go", g.ServiceFileName)
	os.RemoveAll(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, g)
	if err != nil {
		return err
	}

	err = g.GenerateChildren(root)
	if err != nil {
		return err
	}

	return nil
}

func (g Generator) GenerateChildren(root string) error {
	t := template.New("")
	t, err := t.Parse(childrenTemplateString)
	if err != nil {
		return err
	}

	path := fmt.Sprintf(root+"/cmd/gen/%s_children.go", g.ServiceFileName)

	os.RemoveAll(path)
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	err = t.Execute(f, g)
	if err != nil {
		return err
	}
	return nil
}
