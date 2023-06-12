package pkg

import (
	"fmt"
	"os"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/json"
)

type OutputType string

func PrintOutputForTypeArray(cmd *cobra.Command, obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("[printoutput] : %v", err)
	}

	typeOutput := cmd.Flags().Lookup("output-type").Value.String()
	if typeOutput == "json" {
		fmt.Println(string(bytes))
		return nil
	}

	var fields []map[string]interface{}
	var field map[string]interface{}
	err = json.Unmarshal(bytes, &fields)
	if err != nil {
		if err.Error() == "json: cannot unmarshal object into Go value of type []map[string]interface {}" {
			err = json.Unmarshal(bytes, &field)
		} else {
			return fmt.Errorf("[printoutput] : %v", err)
		}
	}
	printTable := table.NewWriter()
	printTable.SetOutputMirror(os.Stdout)

	var headers []interface{}
	var record []interface{}
	if len(fields) != 0 {
		for _, vl := range fields {
			for key, value := range vl {
				headers = append(headers, key)
				record = append(record, value)
			}
		}
	} else {
		for key, value := range field {
			headers = append(headers, key)
			record = append(record, value)
		}
	}
	printTable.AppendHeader(headers)
	printTable.AppendRows([]table.Row{record})
	printTable.AppendSeparator()
	printTable.Render()
	return nil
}
func PrintOutput(cmd *cobra.Command, obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("[printoutput] : %v", err)
	}

	typeOutput := cmd.Flags().Lookup("output-type").Value.String()
	if typeOutput == "json" {
		fmt.Println(string(bytes))
		return nil
	}

	var fields map[string]interface{}
	err = json.Unmarshal(bytes, &fields)
	if err != nil {
		return fmt.Errorf("[printoutput] : %v", err)
	}
	printTable := table.NewWriter()
	printTable.SetOutputMirror(os.Stdout)

	var headers []interface{}
	var record []interface{}
	for key, value := range fields {
		headers = append(headers, key)
		record = append(record, value)
	}
	printTable.AppendHeader(headers)
	printTable.AppendRows([]table.Row{record})
	printTable.AppendSeparator()
	printTable.Render()
	return nil
}
