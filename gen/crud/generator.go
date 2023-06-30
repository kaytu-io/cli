package crud

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed crud.go.template
var templateString string

type Generator struct {
	Services []string
}

func (g Generator) Generate(root string) error {
	t := template.New("")
	t, err := t.Parse(templateString)
	if err != nil {
		return err
	}

	os.RemoveAll(root + "/cmd/crud.go")
	f, err := os.OpenFile(root+"/cmd/crud.go", os.O_CREATE|os.O_RDWR, os.ModePerm)
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
