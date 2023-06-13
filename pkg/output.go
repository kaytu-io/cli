package pkg

import (
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"k8s.io/apimachinery/pkg/util/json"
	"os"
	"sort"
)

type OutputType string

type Cell struct {
	Key   string
	Value interface{}
}

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

	var headers table.Row
	headerOrder := map[string]int{}
	var record []interface{}
	if len(fields) > 0 {
		firstRow := fields[0]
		idx := 0
		for key, _ := range firstRow {
			headers = append(headers, key)
			headerOrder[key] = idx
			idx++
		}
	}

	printTable.AppendHeader(headers)
	for _, vl := range fields {
		var cells []Cell
		for key, value := range vl {
			cells = append(cells, Cell{Key: key, Value: value})
		}
		sort.Slice(cells, func(i, j int) bool {
			return headerOrder[cells[i].Key] < headerOrder[cells[j].Key]
		})
		for _, v := range cells {
			record = append(record, v.Value)
		}

		printTable.AppendRow(record)
		record = nil
	}
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
