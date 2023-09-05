package utils

import (
	"fmt"
	"strings"
)

func GenerateSetFieldsFromFlags(fv Field) (output string) {
	if len(fv.Children) > 0 {
		for _, param := range fv.Children {
			var line string
			if len(param.Children) > 0 {
				line = fmt.Sprintf(`req.Set{{.Name}}(%s)`, GenerateSetStruct(param))
			} else {
				switch param.Type {
				case "string":
					line = `req.Set{{.Name}}(flags.ReadStringFlag(cmd, "{{.Name}}"))`
				case "*string":
					line = `req.Set{{.Name}}(flags.ReadStringOptionalFlag(cmd, "{{.Name}}"))`
				case "int64":
					if param.Name == "StartTime" || param.Name == "EndTime" {
						line = `req.Set{{.Name}}(flags.ReadTimeFlag(cmd, "{{.Name}}"))`
					} else {
						line = `req.Set{{.Name}}(flags.ReadInt64Flag(cmd, "{{.Name}}"))`
					}
				case "*int64":
					if param.Name == "StartTime" || param.Name == "EndTime" {
						line = `req.Set{{.Name}}(flags.ReadTimeOptionalFlag(cmd, "{{.Name}}"))`
					} else {
						line = `req.Set{{.Name}}(flags.ReadInt64OptionalFlag(cmd, "{{.Name}}"))`
					}
				case "bool":
					line = `req.Set{{.Name}}(flags.ReadBooleanFlag(cmd, "{{.Name}}"))`
				case "*bool":
					line = `req.Set{{.Name}}(flags.ReadBooleanOptionalFlag(cmd, "{{.Name}}"))`
				case "[]string":
					line = `req.Set{{.Name}}(flags.ReadStringArrayFlag(cmd, "{{.Name}}"))`
				case "map[string]string":
					line = `req.Set{{.Name}}(flags.ReadMapStringFlag(cmd, "{{.Name}}"))`
				case "map[string][]string":
					line = `req.Set{{.Name}}(flags.ReadMapStringArrayFlag(cmd, "{{.Name}}"))`
				case "runtime.NamedReadCloser":
					line = `tfstateFile := flags.ReadStringFlag(cmd, "{{.Name}}")
		if tfstateFile == "" {
			dir, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
			}
			files, err := filepath.Glob(filepath.Join(dir, "*.tfstate"))
			if err != nil {
				return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
			}

			if len(files) == 0 {
				return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : No tfstate file found")
			}

			if len(files) > 1 {
				return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : Multiple .tfstate files found. Please specify one")
			}
			tfstateFile = files[0]
		}
		reader, err := os.Open(tfstateFile)
		{{.Name}} := runtime.NamedReader("{{.Name}}", reader)
		req.SetTerraformFile({{.Name}})`
				case "[]int64":
					line = `req.Set{{.Name}}(flags.ReadIntArrayFlag(cmd, "{{.Name}}"))`
				default:
					fmt.Println("Unknown type: " + param.Type)
					line = `req.Set{{.Name}}(v)`
				}
			}
			line = strings.ReplaceAll(line, "{{.Name}}", param.Name)
			output += line + "\n"
		}
	}

	return
}

func GenerateSetStruct(param Field) string {
	var children string
	for _, child := range param.Children {
		name := strings.ReplaceAll(child.Name, "-", "")
		name = strings.TrimPrefix(name, param.Name)
		children += fmt.Sprintf("%s: %s,\n", name, GenerateSetStruct(child))
	}

	resp := GenerateReadingFlag(param)
	if len(children) > 0 {
		if strings.HasPrefix(param.Type, "[]") {
			resp += `{
{
` + children + `
},
}`
		} else {
			resp += `{
` + children + `
}`
		}
	}
	return resp
}

func GenerateReadingFlag(param Field) string {
	switch param.Type {
	case "string":
		return fmt.Sprintf(`flags.ReadStringFlag(cmd, "%s")`, param.Name)
	case "*string":
		return fmt.Sprintf(`flags.ReadStringOptionalFlag(cmd, "%s")`, param.Name)
	case "int64":
		if param.Name == "StartTime" || param.Name == "EndTime" {
			return fmt.Sprintf(`flags.ReadTimeFlag(cmd, "%s")`, param.Name)
		} else {
			return fmt.Sprintf(`flags.ReadInt64Flag(cmd, "%s")`, param.Name)
		}
	case "*int64":
		if param.Name == "StartTime" || param.Name == "EndTime" {
			return fmt.Sprintf(`flags.ReadTimeOptionalFlag(cmd, "%s")`, param.Name)
		} else {
			return fmt.Sprintf(`flags.ReadInt64OptionalFlag(cmd, "%s")`, param.Name)
		}
	case "bool":
		return fmt.Sprintf(`flags.ReadBooleanFlag(cmd, "%s")`, param.Name)
	case "*bool":
		return fmt.Sprintf(`flags.ReadBooleanOptionalFlag(cmd, "%s")`, param.Name)
	case "[]string":
		return fmt.Sprintf(`flags.ReadStringArrayFlag(cmd, "%s")`, param.Name)
	case "map[string]string":
		return fmt.Sprintf(`flags.ReadMapStringFlag(cmd, "%s")`, param.Name)
	case "map[string][]string":
		return fmt.Sprintf(`flags.ReadMapStringArrayFlag(cmd, "%s")`, param.Name)
	case "[]int64":
		return fmt.Sprintf(`flags.ReadIntArrayFlag(cmd, "%s")`, param.Name)
	default:
		var ptype = param.Type
		if ptype[0] == '*' {
			ptype = "&" + ptype[1:]
		}

		if param.IsEnum {
			if strings.HasPrefix(ptype, "[]") {
				return fmt.Sprintf(`flags.ReadEnumArrayFlag[%s](cmd, "%s")`, ptype[2:], param.Name)
			}
			return fmt.Sprintf(`%s(flags.ReadStringFlag(cmd, "%s"))`, ptype, param.Name)
		} else {
			return ptype
		}
	}
}
