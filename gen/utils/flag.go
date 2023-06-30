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
				case "string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")
%[1]sCmd.MarkFlagRequired("%[2]s")`, tempName, flags.Name(param.Name))
				case "*string":
					line = fmt.Sprintf(`%sCmd.Flags().String("%s", "", "")`, tempName, flags.Name(param.Name))
				case "int64":
					line = fmt.Sprintf(`%sCmd.Flags().Int64("%s", 0, "")
%[1]sCmd.MarkFlagRequired("%[2]s")`, tempName, flags.Name(param.Name))
				case "*int64":
					line = fmt.Sprintf(`%sCmd.Flags().Int64("%s", 0, "")`, tempName, flags.Name(param.Name))
				case "bool":
					line = fmt.Sprintf(`%sCmd.Flags().Bool("%s", false, "")
%[1]sCmd.MarkFlagRequired("%[2]s")`, tempName, flags.Name(param.Name))
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
