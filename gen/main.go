package main

import (
	_ "embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/cli-program/cmd/flags"
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

	os.WriteFile(root+"/cmd/resources.go", []byte(`
package cmd

import (
	"github.com/kaytu-io/cli-program/cmd/gen"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var createCmd = &cobra.Command{
	Use: "create",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
var updateCmd = &cobra.Command{
	Use: "update",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	// ========== DO NOT EDIT THIS PART! AUTO GENERATED!!! =========
	// =============================================================

}
`), os.ModePerm)

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

		nameCommand := strings.ReplaceAll(strcase.ToSnake(name), "_", "-")
		r, err := regexp.Compile(`[a-zA-Z\-]+api-v-\d+-`)
		if err != nil {
			return err
		}
		c := r.ReplaceAll([]byte(nameCommand), []byte(""))
		nameCommand = string(c)

		fv := ExtractFields(reflect.TypeOf(paramModels[apiName]))
		output := GenerateSetFieldsFromFlags(fv)
		tmpl := ChildCmdTemplate{
			NameCamel:        strcase.ToCamel(name),
			NameSnake:        strcase.ToSnake(name),
			NameCommand:      nameCommand,
			APIName:          apiName,
			ServiceName:      strcase.ToCamel(serviceName),
			ServiceNameSnake: strcase.ToSnake(serviceName),
			ParamString:      output,
		}
		childrenMap[name] = tmpl
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

		switch {
		case strings.Contains(temp.NameCamel, "Delete"):
			cmdReference += fmt.Sprintf(`
		Delete%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
		case strings.Contains(temp.NameCamel, "Get"):
			cmdReference += fmt.Sprintf(`
		Get%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
		case strings.Contains(temp.NameCamel, "Update"):
			cmdReference += fmt.Sprintf(`
		Update%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
		case strings.Contains(temp.NameCamel, "Create"):
			cmdReference += fmt.Sprintf(`
		Create%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
		default:
			cmdReference += fmt.Sprintf(`
		Get%sCmd.AddCommand(%sCmd)
`, strcase.ToCamel(serviceName), temp.NameCamel)
		}

		fv := ExtractFields(reflect.TypeOf(paramModels[temp.APIName]))
		output := GenerateFlagDefinitions(temp.NameCamel, fv)
		cmdReference += output
		//
		//for _, param := range temp.Params {
		//	cmdReference += fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")
		//`, temp.NameCamel, param.FlagName)
		//}
	}
	cmdReference += "}"

	filename = fmt.Sprintf(root+"/cmd/gen/%s.go", serviceName)

	_, err = fservice.WriteString(cmdReference)
	if err != nil {
		return err
	}

	return nil
}

func GenerateFlagDefinitions(tempName string, fv Field) (output string) {
	if len(fv.Children) > 0 {
		for _, param := range fv.Children {
			var line string
			if len(param.Children) > 0 {
				line = GenerateFlagDefinitions(tempName, param)
			} else {
				switch param.Type {
				case "string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				case "*string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				case "int64":
					line = fmt.Sprintf(`%sCmd.Flags().Int64("%s", 0, "")`, tempName, flags.Name(param.Name))
				case "*int64":
					line = fmt.Sprintf(`%sCmd.Flags().Int64("%s", 0, "")`, tempName, flags.Name(param.Name))
				case "bool":
					line = fmt.Sprintf(`%sCmd.Flags().Bool("%s", false, "")`, tempName, flags.Name(param.Name))
				case "*bool":
					line = fmt.Sprintf(`%sCmd.Flags().Bool("%s", false, "")`, tempName, flags.Name(param.Name))
				case "[]string":
					line = fmt.Sprintf(`%sCmd.Flags().StringArray("%s", nil, "")`, tempName, flags.Name(param.Name))
				case "map[string]string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				case "map[string][]string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				default:
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				}
			}
			line = strings.ReplaceAll(line, "{{.Name}}", param.Name)
			output += line + "\n"
		}
	}

	return
}
