package utils

import (
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"strings"
)

func GenerateFlagDefinitions(tempName string, fv Field) (output string) {
	if len(fv.Children) > 0 {
		for _, param := range fv.Children {
			var line string
			if len(param.Children) > 0 {
				line = GenerateFlagDefinitions(tempName, param)
			} else {
				switch param.Type {
				case "string", "*string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "int64", "*int64":
					line = fmt.Sprintf(`%sCmd.Flags().Int64("%s", 0, "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "bool", "*bool":
					line = fmt.Sprintf(`%sCmd.Flags().Bool("%s", false, "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "[]string":
					line = fmt.Sprintf(`%sCmd.Flags().StringArray("%s", nil, "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "map[string]string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "map[string][]string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "%s")`, tempName, flags.Name(param.Name), param.Description)
				case "[]int64", "[]int":
					line = fmt.Sprintf(`%sCmd.Flags().Int64Slice("%s", nil, "%s")`, tempName, flags.Name(param.Name), param.Description)
				default:
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "%s")`, tempName, flags.Name(param.Name), param.Description)
				}

				if param.IsRequired {
					if !param.IsNestedRequirement {
						line += fmt.Sprintf("\n%[1]sCmd.MarkFlagRequired(\"%[2]s\")", tempName, flags.Name(param.Name))
					} else {
						if !strings.HasPrefix(param.Type, "*") && !strings.HasPrefix(param.Type, "[]") {
							line += fmt.Sprintf("\n%[1]sCmd.MarkFlagRequired(\"%[2]s\")", tempName, flags.Name(param.Name))
						}
					}
				}
			}
			line = strings.ReplaceAll(line, "{{.Name}}", param.Name)
			output += line + "\n"
		}
	}

	return
}
