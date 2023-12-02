package cobra

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/kaytu-io/cli-program/codegen/api"
	"github.com/kaytu-io/cli-program/codegen/cmd"
	"github.com/kaytu-io/cli-program/codegen/swagger"
	"strings"

	//"github.com/kaytu-io/cli-program/gen/cmd"
	"os"
	"path/filepath"
	"text/template"
)

type Flag struct {
	Name         string
	Type         string
	DefaultValue string
	Usage        string
	Required     bool
	Persistence  bool
}

func (flag Flag) TypeFlagFunc() string {
	switch flag.Type {
	case "int64", "*int64":
		return "Int64"
	case "[]int64":
		return "Int64Slice"
	case "string", "*string":
		return "String"
	case "[]string":
		return "StringArray"
	case "bool", "*bool":
		return "Bool"
	default:
		return "String"
	}
}

func (flag Flag) FlagDefaultValue() string {
	switch flag.Type {
	case "int64", "*int64":
		return "0"
	case "[]int64":
		return "nil"
	case "string", "*string":
		return "\"\""
	case "[]string":
		return "nil"
	case "bool", "*bool":
		return "false"
	default:
		return "\"\""
	}
}

type Package struct {
	Name string
	Path string
}

type CobraCommand struct {
	Package  Package
	Command  cmd.Command
	Flags    []Flag
	Children []*CobraCommand
}

func FromCommand(c cmd.Command, swag *swagger.Swagger) CobraCommand {
	var flags []Flag
	for _, paramType := range c.API.ParamTypes {
		for _, subParam := range paramType.SubParams {
			name := strings.ReplaceAll(strcase.ToSnake(subParam.Name), "_", "-")
			model := swag.GetAPI(c.API.MethodName + "Params")
			required := false
			if strings.ToLower(subParam.Name) == "metricid" {
				fmt.Println(model)
			}
			for _, p := range model.Parameters {
				if strings.ToLower(strcase.ToCamel(p.Name)) == strings.ToLower(subParam.Name) {
					required = p.Required
				}
			}
			flags = append(flags, Flag{
				Name:         name,
				Type:         subParam.Type,
				DefaultValue: "", //TODO-Saleh
				Usage:        strings.ReplaceAll(subParam.Description, "\n", "\\n"),
				Required:     required,
				Persistence:  false,
			})
		}
	}
	return CobraCommand{
		Command:  c,
		Flags:    flags,
		Children: nil,
	}
}

func BuildCobraCommands(cmds []cmd.Command, swag *swagger.Swagger) CobraCommand {
	root := CobraCommand{
		Package: Package{
			Name: "genv2",
			Path: "github.com/kaytu-io/cli-program/cmd/genv2",
		},
		Command: cmd.Command{
			UsePath:          []string{"kaytu"},
			ShortDescription: "Kaytu CLI",
			LongDescription:  "",
			API:              api.API{},
		},
		Flags: []Flag{
			{
				Name:         "workspace-name",
				Type:         "string",
				DefaultValue: "",
				Usage:        "",
				Required:     false,
				Persistence:  true,
			},
			{
				Name:         "output-type",
				Type:         "string",
				DefaultValue: "\"summary\"",
				Usage:        "output type [summary, json, csv, list, table]",
				Required:     false,
				Persistence:  true,
			},
			{
				Name:         "filter",
				Type:         "string",
				DefaultValue: "",
				Usage:        "columns to display. syntax: https://github.com/teacat/jsonfilter#syntax",
				Required:     false,
				Persistence:  true,
			},
		},
		Children: nil,
	}

	for _, item := range cmds {
		addCommand(&root, item, swag)
	}
	return root
}

func addCommand(root *CobraCommand, cmdItem cmd.Command, swag *swagger.Swagger) {
	if len(cmdItem.UsePath) == 1 {
		child := FromCommand(cmdItem, swag)
		child.Package.Name = strcase.ToSnake(root.Command.UsePath[0])
		child.Package.Path = root.Package.Path + "/" + child.Package.Name
		root.Children = append(root.Children, &child)
		return
	}

	pathCmd := cmdItem.UsePath[0]
	cmdItem.UsePath = cmdItem.UsePath[1:]

	childExists := false
	for _, child := range root.Children {
		if child.Command.UsePath[0] == pathCmd {
			childExists = true
			addCommand(child, cmdItem, swag)
		}
	}

	if !childExists {
		pkgName := strcase.ToSnake(root.Command.UsePath[0])
		child := CobraCommand{
			Package: Package{
				Name: pkgName,
				Path: root.Package.Path + "/" + pkgName,
			},
			Command: cmd.Command{
				UsePath:          []string{pathCmd},
				ShortDescription: "", //TODO-Saleh populate description
				LongDescription:  "",
				API:              api.API{},
			},
			Flags:    nil,
			Children: nil,
		}
		addCommand(&child, cmdItem, swag)
		root.Children = append(root.Children, &child)
	}
}

func (cmd CobraCommand) CmdVarName() string {
	return strcase.ToCamel(cmd.Command.UsePath[0])
}
func (cmd CobraCommand) CmdUse() string {
	return cmd.Command.UsePath[0]
}
func (cmd CobraCommand) Filename() string {
	return strcase.ToSnake(cmd.Command.UsePath[0])
}

//go:embed run.template
var runTemplate string

func (cmd CobraCommand) RunImplementation() string {
	tmpl := template.New("CobraCommand")
	t, err := tmpl.Parse(runTemplate)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	err = t.Execute(buf, cmd)
	if err != nil {
		panic(err)
	}

	return buf.String()
}

//go:embed cmd.go.template
var cmdTemplate string

func GenerateCobraCommands(rootCmd CobraCommand, filePath string) {
	tmpl := template.New("CobraCommand")
	t, err := tmpl.Parse(cmdTemplate)
	if err != nil {
		panic(err)
	}

	path := filePath + "/" + rootCmd.Filename() + ".go"
	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	err = t.Execute(f, rootCmd)
	if err != nil {
		panic(err)
	}

	err = f.Close()
	if err != nil {
		panic(err)
	}

	for _, child := range rootCmd.Children {
		GenerateCobraCommands(*child, filePath+"/"+rootCmd.Filename())
	}
}
