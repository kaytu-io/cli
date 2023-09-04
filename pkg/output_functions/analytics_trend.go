package output_functions

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func AnalyticsTrend(cmd *cobra.Command, commandName string, obj interface{}) error {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()
	filter := cmd.Flags().Lookup("filter").Value.String()

	if filter != "" {
		fmt.Println(color.YellowString("No filter for this command!"))
	}

	switch typeOutput {
	case "table", "summary":
		return PrintTable(bytes, &map[string]int{"date": 1, "count": 2})

	case "list":
		return PrintList(bytes, &map[string]int{"date": 1, "count": 2})

	case "csv":
		return PrintCSV(bytes, &map[string]int{"date": 1, "count": 2})

	default:
		fmt.Println(string(bytes))
		return nil
	}
}
