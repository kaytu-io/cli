package pkg

import (
	"encoding/json"
	"fmt"
	output "github.com/kaytu-io/cli-program/pkg/output_functions"
	"github.com/spf13/cobra"
	"github.com/teacat/jsonfilter"
)

func PrintOutput(cmd *cobra.Command, commandName string, obj interface{}) error {
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()

	if function, ok := outputFunctions[commandName]; ok && typeOutput != "json" {
		return function(cmd, commandName, obj)
	} else {
		return PrintOutputDefault(cmd, commandName, obj)
	}
}

func PrintOutputDefault(cmd *cobra.Command, commandName string, obj interface{}) error {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()
	filter := cmd.Flags().Lookup("filter").Value.String()

	if filter != "" {
		bytes, err = jsonfilter.Filter(bytes, filter)
		if err != nil {
			return fmt.Errorf("[output]: %v", err)
		}
	}

	switch typeOutput {
	case "summary":
		return output.PrintSummary(bytes, commandName)

	case "table":
		return output.PrintTable(bytes, nil)

	case "list":
		return output.PrintList(bytes, nil)

	case "csv":
		return output.PrintCSV(bytes, nil)

	default:
		fmt.Println(string(bytes))
		return nil
	}
}
