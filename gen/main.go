package main

import (
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed service.go.template
var serviceTemplate string

//go:embed child_cmd.go.template
var childCmdTemplate string

//go:embed children.go.template
var childrenTemplate string

func main() {
	root := "."
	_, err := os.Stat(root + "/pkg")
	if err != nil {
		root = ".."
	}
	os.Mkdir(root+"/cmd/gen", os.ModePerm)

	t := template.New("serviceTemplate")
	t, err = t.Parse(serviceTemplate)
	if err != nil {
		panic(err)
	}

	err = filepath.Walk(root+"/pkg/api/kaytu/client", func(path string, info fs.FileInfo, err error) error {
		if info == nil {
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		if strings.HasSuffix(path, "/kaytu/client") {
			return nil
		}

		filename := fmt.Sprintf(root+"/cmd/gen/%s.go", info.Name())
		os.RemoveAll(filename)

		f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
		if err != nil {
			return err
		}

		err = t.Execute(f, &ServiceTemplate{
			ServiceName:      strcase.ToCamel(info.Name()),
			ServiceNameSnake: info.Name(),
		})
		if err != nil {
			return err
		}

		rb, err := os.ReadFile(root + "/cmd/resources.go")
		if err != nil {
			return err
		}

		resourcesFile := string(rb)
		content := fmt.Sprintf(`// ========== DO NOT EDIT THIS PART! AUTO GENERATED!!! =========

	getCmd.AddCommand(gen.Get%[1]sCmd)
	updateCmd.AddCommand(gen.Update%[1]sCmd)
	deleteCmd.AddCommand(gen.Delete%[1]sCmd)
	createCmd.AddCommand(gen.Create%[1]sCmd)
`, strcase.ToCamel(info.Name()))
		resourcesFile = strings.ReplaceAll(resourcesFile, "// ========== DO NOT EDIT THIS PART! AUTO GENERATED!!! =========", content)

		err = os.WriteFile(root+"/cmd/resources.go", []byte(resourcesFile), os.ModePerm)
		if err != nil {
			return err
		}

		err = createChildren(root, info.Name(), path)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func createChildren(root, serviceName, servicePath string) error {
	t := template.New("childCmdTemplate")
	t, err := t.Parse(childCmdTemplate)
	if err != nil {
		panic(err)
	}

	filename := fmt.Sprintf(root+"/cmd/gen/%s_children.go", serviceName)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	tf := template.New("childrenTemplate")
	tf, err = tf.Parse(childrenTemplate)
	if err != nil {
		panic(err)
	}

	err = tf.Execute(f, &ChildCmdTemplate{ServiceNameSnake: strcase.ToSnake(serviceName)})
	if err != nil {
		panic(err)
	}

	childrenMap := map[string]ChildCmdTemplate{}

	err = filepath.Walk(servicePath, func(path string, info fs.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), "_client.go") {
			return nil
		}

		name := info.Name()
		idx := strings.LastIndex(name, ".")
		name = name[:idx]
		if strings.HasSuffix(name, "_parameters") {
			name = name[:len(name)-11]
		}
		if strings.HasSuffix(name, "_responses") {
			name = name[:len(name)-10]
		}

		apiName := strings.ReplaceAll(strcase.ToCamel(name), "ApiV", "APIV")
		apiName = strings.ReplaceAll(apiName, "Id", "ID")

		childrenMap[name] = ChildCmdTemplate{
			NameCamel:        strcase.ToCamel(name),
			NameSnake:        strcase.ToSnake(name),
			APIName:          apiName,
			ServiceName:      strcase.ToCamel(serviceName),
			ServiceNameSnake: strcase.ToSnake(serviceName),
		}
		return nil
	})
	if err != nil {
		return err
	}

	cmdReference := "func init() {"
	for _, temp := range childrenMap {
		err = t.Execute(f, &temp)
		if err != nil {
			return err
		}

		cmdReference += fmt.Sprintf(`
		Get%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
	}
	cmdReference += "}"

	filename = fmt.Sprintf(root+"/cmd/gen/%s.go", serviceName)
	f2, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = f2.WriteString(cmdReference)
	if err != nil {
		return err
	}

	return nil
}
