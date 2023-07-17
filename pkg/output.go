package pkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
	"os"
	"sort"
)

func PrintOutput(cmd *cobra.Command, obj interface{}) error {
	typeOutput := cmd.Flags().Lookup("output-type").Value.String()

	switch typeOutput {
	case "summary":
		bytes, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return fmt.Errorf("[output] : %v", err)
		}
		fmt.Println(string(bytes))
		return nil

	case "table":
		return PrintTable(obj)

	case "csv":
		return PrintCSV(obj)

	default:
		bytes, err := json.MarshalIndent(obj, "", "  ")
		if err != nil {
			return fmt.Errorf("[output]: %v", err)
		}
		fmt.Println(string(bytes))
		return nil
	}
}

func PrintCSV(obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}

	rows, err := extractRows(bytes)
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	if len(rows) == 0 {
		fmt.Println("no rows to show")
		return nil
	}

	csvWriter := csv.NewWriter(os.Stdout)

	var headersOrdered []string
	for key := range rows[0] {
		headersOrdered = append(headersOrdered, key)
	}
	sort.Slice(headersOrdered, func(i, j int) bool {
		return headersOrdered[i] < headersOrdered[j]
	})

	var headers []string
	for _, item := range headersOrdered {
		headers = append(headers, strcase.ToSnake(item))
	}
	if err := csvWriter.Write(headers); err != nil {
		return err
	}

	for _, item := range rows {
		var row []string
		for _, header := range headersOrdered {

			value := item[header]
			row = append(row, value)
		}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}
	csvWriter.Flush()
	return nil
}
func PrintTable(obj interface{}) error {
	bytes, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}

	rows, err := extractRows(bytes)
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	if len(rows) == 0 {
		fmt.Println("no rows to show")
		return nil
	}

	printTable := table.NewWriter()
	printTable.SetOutputMirror(os.Stdout)

	var headersOrdered []string
	for key := range rows[0] {
		headersOrdered = append(headersOrdered, key)
	}
	sort.Slice(headersOrdered, func(i, j int) bool {
		return headersOrdered[i] < headersOrdered[j]
	})

	var headers []interface{}
	for _, item := range headersOrdered {
		headers = append(headers, strcase.ToSnake(item))
	}
	printTable.AppendHeader(headers)

	for _, item := range rows {
		var row []interface{}
		for _, header := range headersOrdered {

			value := item[header]
			row = append(row, value)
		}
		printTable.AppendRow(row)
	}
	printTable.Render()
	return nil
}

func extractRows(objJSON []byte) ([]map[string]string, error) {
	var rowsInterface []map[string]interface{}

	var fields map[string]interface{}
	if err := json.Unmarshal(objJSON, &fields); err != nil {
		rowsInterface, err = extractArray(objJSON)
		if err != nil {
			return nil, err
		}
	} else {
		rowsInterface = append(rowsInterface, fields)
	}

	var rows []map[string]string
	for _, item := range rowsInterface {
		row := map[string]string{}
		for k, v := range item {
			switch t := v.(type) {
			case int, int32, int64, float64, float32, string, bool:
				row[k] = fmt.Sprintf("%v", t)
			default:
				js, err := json.Marshal(v)
				if err != nil {
					return nil, err
				}
				row[k] = string(js)
			}
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func extractArray(objJSON []byte) ([]map[string]interface{}, error) {
	var fields []map[string]interface{}
	err := json.Unmarshal(objJSON, &fields)
	if err != nil {
		return nil, err
	}
	return fields, nil
}
