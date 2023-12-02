package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	output "github.com/kaytu-io/cli-program/pkg/output_functions"
	"github.com/spf13/cobra"
	"github.com/teacat/jsonfilter"
)

func PrintOutput(cmd *cobra.Command, commandName string, obj interface{}) error {
	var typeOutput string

	typeOutputFlag := cmd.Flags().Lookup("output-type")
	if typeOutputFlag != nil {
		typeOutput = typeOutputFlag.Value.String()
	}

	if function, ok := outputFunctions[commandName]; ok && typeOutput != "json" {
		return function(cmd, commandName, obj)
	} else {
		return PrintOutputDefault(cmd, commandName, obj)
	}
}

func PrintOutputDefault(cmd *cobra.Command, commandName string, obj interface{}) error {
	if obj == nil || (fmt.Sprintf("%v", obj) == "[]") {
		fmt.Println(color.CyanString("No Result"))
		return nil
	}

	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}

	var typeOutput, filter string

	typeOutputFlag := cmd.Flags().Lookup("output-type")
	if typeOutputFlag != nil {
		typeOutput = typeOutputFlag.Value.String()
	}
	filterFlag := cmd.Flags().Lookup("filter")
	if filterFlag != nil {
		filter = filterFlag.Value.String()
	}

	if filter != "" {
		bytes, err = jsonfilter.Filter(bytes, filter)
		if err != nil {
			return fmt.Errorf("[output]: %v", err)
		}
	}

	byteString := string(bytes[0])
	if byteString != "[" {
		if byteString != "{" {
			typeOutput = ""
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
