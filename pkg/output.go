package pkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/iancoleman/strcase"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kaytu-io/cli-program/gen/config"
	"github.com/spf13/cobra"
	"github.com/teacat/jsonfilter"
	jsonmask "github.com/teambition/json-mask-go"
	"golang.org/x/term"
	"os"
	"sort"
	"strings"
)

func keepTwoDigits(v interface{}) interface{} {
	switch v.(type) {
	case []interface{}:
		for key, value := range v.([]interface{}) {
			v.([]interface{})[key] = keepTwoDigits(value)
		}
	case map[string]interface{}:
		for key, value := range v.(map[string]interface{}) {
			v.(map[string]interface{})[key] = keepTwoDigits(value)
		}
	case float64:
		v = fmt.Sprintf("%v", float64(int(v.(float64)*100))/100)
	case float32:
		v = fmt.Sprintf("%v", float32(int(v.(float64)*100))/100)
	}
	return v
}

func PrintOutput(cmd *cobra.Command, commandName string, obj interface{}) error {
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
		return PrintSummary(bytes, commandName)

	case "table":
		return PrintTable(bytes)

	case "list":
		return PrintList(bytes)

	case "csv":
		return PrintCSV(bytes)

	default:
		fmt.Println(string(bytes))
		return nil
	}
}

func PrintSummary(objJSON []byte, commandName string) error {
	if v, ok := config.UrlSummary[commandName]; ok {
		result, err := jsonmask.Mask(objJSON, v)
		if err != nil {
			return err
		}
		return PrintObject(result)
	}
	return PrintObject(objJSON)
}

func PrintObject(objJSON []byte) error {
	if objJSON[0] == '[' {
		return PrintTable(objJSON)
	}
	return PrintList(objJSON)
}

func PrintCSV(objJSON []byte) error {
	rows, err := extractRows(objJSON)
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

func PrintTable(objJSON []byte) error {
	rows, err := extractRows(objJSON)
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
	for _, row := range rows {
		for key := range row {
			exists := false
			for _, h := range headersOrdered {
				if h == key {
					exists = true
				}
			}
			if !exists {
				headersOrdered = append(headersOrdered, key)
			}
		}
	}
	sort.Slice(headersOrdered, func(i, j int) bool {
		return headersOrdered[i] < headersOrdered[j]
	})

	headersSize := 5
	var headers []interface{}
	for _, item := range headersOrdered {
		headers = append(headers, strcase.ToSnake(item))
		headersSize += len(strcase.ToSnake(item)) + 3
	}
	printTable.AppendHeader(headers)

	var maxRowSize int
	for _, item := range rows {
		rowSize := 5
		var row []interface{}
		for _, header := range headersOrdered {
			value := item[header]
			row = append(row, value)
			rowSize += len(value) + 3
		}
		if rowSize > maxRowSize {
			maxRowSize = rowSize
		}
		printTable.AppendRow(row)
	}
	if headersSize > maxRowSize {
		maxRowSize = headersSize
	}

	width, _, _ := term.GetSize(0)
	if maxRowSize > width {
		return PrintList(objJSON)
	}
	printTable.Render()
	return nil
}

func PrintList(objJSON []byte) error {
	rows, err := extractRows(objJSON)
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	if len(rows) == 0 {
		fmt.Println("no rows to show")
		return nil
	}

	var maxSize int
	var headersOrdered []string
	for _, row := range rows {
		for key := range row {
			exists := false
			for _, h := range headersOrdered {
				if h == key {
					exists = true
				}
			}
			if !exists {
				headersOrdered = append(headersOrdered, key)
				if len(key) > maxSize {
					maxSize = len(key)
				}
			}
		}
	}
	sort.Slice(headersOrdered, func(i, j int) bool {
		return headersOrdered[i] < headersOrdered[j]
	})

	for _, item := range rows {
		var row []interface{}
		for _, header := range headersOrdered {
			value := item[header]
			row = append(row, value)
			spaces := strings.Repeat(" ", maxSize-len(header)+2)
			if len(value) == 0 {
				continue
			}
			if value[0] == '[' {
				var strList []string
				if err = json.Unmarshal([]byte(value), &strList); err == nil {
					fmt.Println(fmt.Sprintf("%s %s%s", color.CyanString(strcase.ToCamel(header)+":"), spaces, strList))
					continue
				}
				fmt.Println(color.CyanString(strcase.ToCamel(header) + ":"))
				divider := fmt.Sprintf("\n%s\n", strings.Repeat("=", maxSize))
				fmt.Print(color.CyanString(divider))
				PrintTable([]byte(value))
				fmt.Print(color.CyanString(divider))
			} else {
				fmt.Println(fmt.Sprintf("%s %s%s", color.CyanString(strcase.ToCamel(header)+":"), spaces, value))
			}
		}
		divider := fmt.Sprintf("\n%s\n", strings.Repeat("-", maxSize+10))
		fmt.Println(color.CyanString(divider))
	}
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
			case float64:
				row[k] = fmt.Sprintf("%v", float64(int(v.(float64)*100))/100)
			case float32:
				row[k] = fmt.Sprintf("%v", float32(int(v.(float64)*100))/100)
			case int, int32, int64, string, bool:
				row[k] = fmt.Sprintf("%v", t)
			default:
				v = keepTwoDigits(v)
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
