package output_functions

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"strings"
)

func CostsTrend(cmd *cobra.Command, commandName string, obj interface{}) error {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	data := string(bytes)
	data = strings.ReplaceAll(data, "count", "cost")
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()
	filter := cmd.Flags().Lookup("filter").Value.String()

	if filter != "" {
		fmt.Println(color.YellowString("No filter for this command!"))
	}

	switch typeOutput {
	case "table", "summary":
		return PrintTable([]byte(data), &map[string]int{"date": 1, "cost": 2})

	case "list":
		return PrintList([]byte(data), &map[string]int{"date": 1, "cost": 2})

	case "csv":
		return PrintCSV([]byte(data), &map[string]int{"date": 1, "cost": 2})

	default:
		fmt.Println(data)
		return nil
	}
}
