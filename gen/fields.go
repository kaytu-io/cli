package main

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
					line = `req.Set{{.Name}}(flags.ReadStringFlag("{{.Name}}"))`
				case "*string":
					line = `req.Set{{.Name}}(flags.ReadStringOptionalFlag("{{.Name}}"))`
				case "int64":
					line = `req.Set{{.Name}}(flags.ReadInt64Flag("{{.Name}}"))`
				case "*int64":
					line = `req.Set{{.Name}}(flags.ReadInt64OptionalFlag("{{.Name}}"))`
				case "bool":
					line = `req.Set{{.Name}}(flags.ReadBooleanFlag("{{.Name}}"))`
				case "*bool":
					line = `req.Set{{.Name}}(flags.ReadBooleanOptionalFlag("{{.Name}}"))`
				case "[]string":
					line = `req.Set{{.Name}}(flags.ReadStringArrayFlag("{{.Name}}"))`
				case "map[string]string":
					line = `req.Set{{.Name}}(flags.ReadMapStringFlag("{{.Name}}"))`
				case "map[string][]string":
					line = `req.Set{{.Name}}(flags.ReadMapStringArrayFlag("{{.Name}}"))`
				default:
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
		children += fmt.Sprintf("%s: %s,\n", child.Name, GenerateSetStruct(child))
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
		return fmt.Sprintf(`flags.ReadStringFlag("%s")`, param.Name)
	case "*string":
		return fmt.Sprintf(`flags.ReadStringOptionalFlag("%s")`, param.Name)
	case "int64":
		return fmt.Sprintf(`flags.ReadInt64Flag("%s")`, param.Name)
	case "*int64":
		return fmt.Sprintf(`flags.ReadInt64OptionalFlag("%s")`, param.Name)
	case "bool":
		return fmt.Sprintf(`flags.ReadBooleanFlag("%s")`, param.Name)
	case "*bool":
		return fmt.Sprintf(`flags.ReadBooleanOptionalFlag("%s")`, param.Name)
	case "[]string":
		return fmt.Sprintf(`flags.ReadStringArrayFlag("%s")`, param.Name)
	case "map[string]string":
		return fmt.Sprintf(`flags.ReadMapStringFlag("%s")`, param.Name)
	case "map[string][]string":
		return fmt.Sprintf(`flags.ReadMapStringArrayFlag("%s")`, param.Name)
	default:
		var ptype = param.Type
		if ptype[0] == '*' {
			ptype = "&" + ptype[1:]
		}

		if param.IsEnum {
			if strings.HasPrefix(ptype, "[]") {
				return fmt.Sprintf(`flags.ReadEnumArrayFlag[%s]("%s")`, ptype[2:], param.Name)
			}
			return fmt.Sprintf(`%s(flags.ReadStringFlag("%s"))`, ptype, param.Name)
		} else {
			return ptype
		}
	}
}
