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

//go:embed "child_empty_response_payload.go.template"
var childCmdEmptyResponsePayloadTemplate string

func main() {
	root := "."
	_, err := os.Stat(root + "/pkg")
	if err != nil {
		root = ".."
	}
	os.RemoveAll(root + "/cmd/gen")
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
		defer f.Close()

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

		err = createChildren(root, info.Name(), path, f)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		panic(err)
	}
}

func createChildren(root, serviceName, servicePath string, fservice *os.File) error {
	filename := fmt.Sprintf(root+"/cmd/gen/%s_children.go", serviceName)
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

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
	hasPayload := map[string]bool{}

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
		if strings.HasSuffix(info.Name(), "_responses.go") {
			c, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			existPayload := strings.Contains(string(c), "GetPayload")
			hasPayload[name] = existPayload
		}

		nameCommand := strings.Split(strcase.ToSnake(name), "_")
		newNameCommand := ""
		if len(nameCommand) >= 5 {
			for i := 5; i <= len(nameCommand)-1; i++ {
				if i == 6 {
					newNameCommand = strcase.ToSnake(nameCommand[i])
				}
				newNameCommand += strcase.ToCamel(nameCommand[i])
			}
		}
		childrenMap[name] = ChildCmdTemplate{
			NameCamel:        strcase.ToCamel(name),
			NameSnake:        strcase.ToSnake(name),
			NameCommand:      newNameCommand,
			APIName:          apiName,
			ServiceName:      strcase.ToCamel(serviceName),
			ServiceNameSnake: strcase.ToSnake(serviceName),
		}
		return nil
	})
	if err != nil {
		return err
	}

	withoutPayload := template.New("childCmdEmptyResponsePayloadTemplate")
	withoutPayload, err = withoutPayload.Parse(childCmdEmptyResponsePayloadTemplate)
	if err != nil {
		panic(err)
	}

	withPayload := template.New("childCmdTemplate")
	withPayload, err = withPayload.Parse(childCmdTemplate)
	if err != nil {
		panic(err)
	}

	cmdReference := "func init() {"
	for name, temp := range childrenMap {
		if hasPayload[name] {
			err = withPayload.Execute(f, &temp)
			if err != nil {
				return err
			}
		} else {
			err = withoutPayload.Execute(f, &temp)
			if err != nil {
				return err
			}
		}

		cmdReference += fmt.Sprintf(`
		Get%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
	}
	cmdReference += "}"

	filename = fmt.Sprintf(root+"/cmd/gen/%s.go", serviceName)

	_, err = fservice.WriteString(cmdReference)
	if err != nil {
		return err
	}

	return nil
}
