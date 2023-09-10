package output_functions

import (
	"encoding/json"
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/spf13/cobra"
	"github.com/teacat/jsonfilter"
	"strings"
)

func GetCosts(cmd *cobra.Command, commandName string, obj interface{}) error {
	bytes, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()

	view := flags.ReadStringFlag(cmd, "view")
	if view == "service" {
		data := string(bytes)
		data = strings.ReplaceAll(data, "cost_dimension_name", "service_name")
		data = strings.ReplaceAll(data, "metrics", "services")
		data = strings.ReplaceAll(data, "\"connector\": [\n        \"AWS\"\n      ]", "\"connector\":\"AWS\"")
		data = strings.ReplaceAll(data, "\"connector\": [\n        \"Azure\"\n      ]", "\"connector\":\"Azure\"")
		bytes, err = jsonfilter.Filter([]byte(data), "services(total_cost,connector,service_name),total_count,total_cost")
		if err != nil {
			return fmt.Errorf("[output]: %v", err)
		}
		if typeOutput == "csv" {
			return PrintCSV(bytes, nil)
		}
		return PrintList(bytes, nil)
	} else if view == "connection" {
		data := string(bytes)
		data = strings.ReplaceAll(data, "providerConnectionName", "account_name")
		bytes, err = jsonfilter.Filter([]byte(data), "connections(cost,connector,resourceCount,account_name),connectionCount,totalCost")
		if err != nil {
			return fmt.Errorf("[output]: %v", err)
		}
		if typeOutput == "csv" {
			return PrintCSV(bytes, &map[string]int{"connections": 1, "connector": 1, "account_name": 2, "connectionCount": 3, "totalCost": 2})
		}
		return PrintList(bytes, &map[string]int{"connections": 1, "connector": 1, "account_name": 2, "connectionCount": 3, "totalCost": 2})
	}

	return PrintSummary(bytes, commandName)
}
