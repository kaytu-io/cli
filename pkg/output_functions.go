package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"sort"
)

type outputFunction func(cmd *cobra.Command, commandName string, obj interface{}) error

var outputFunctions = map[string]outputFunction{
	"get-inventory-api-v2-analytics-spend-table": AnalyticsSpendTable,
}

func AnalyticsSpendTable(cmd *cobra.Command, commandName string, obj interface{}) error {
	var output []map[string]string
	dimension := *flags.ReadStringOptionalFlag(cmd, "Dimension")
	monthsMap := make(map[string]bool)
	for _, v := range obj.([]*models.GithubComKaytuIoKaytuEnginePkgInventoryAPISpendTableRow) {
		for t, _ := range v.CostValue {
			if !monthsMap[t] {
				monthsMap[t] = true
			}
		}
	}
	if dimension == "metric" {
		for _, v := range obj.([]*models.GithubComKaytuIoKaytuEnginePkgInventoryAPISpendTableRow) {
			o := map[string]string{
				"category": v.Category,
				"name":     v.DimensionName,
			}
			hasMonth := monthsMap
			for t, c := range v.CostValue {
				hasMonth[t] = false
				o[t] = fmt.Sprintf("%v", float64(int(c*100))/100)
			}
			for t, b := range hasMonth {
				if b {
					o[t] = ""
				}
			}
			output = append(output, o)
		}
	}
	var slicedOutput []map[string]string
	var rowsWidth []int
	screenWidth, _, _ := term.GetSize(0)
	for _, row := range output {
		width := len(row["category"]) + len(row["name"]) + 20
		slicedOutput = append(slicedOutput, map[string]string{
			"category": row["category"],
			"name":     row["name"],
		})
		rowsWidth = append(rowsWidth, width)
	}
	months := make([]string, 0, len(monthsMap))

	for m := range monthsMap {
		months = append(months, m)
	}
	sort.Strings(months)
	for _, month := range months {
		var maxWidth int
		for i, row := range output {
			rowsWidth[i] += max(len(month), len(row[month]))
			slicedOutput[i][month] = row[month]
			if rowsWidth[i] > maxWidth {
				maxWidth = rowsWidth[i]
			}
		}
		if maxWidth >= screenWidth-50 {
			bytes, err := json.MarshalIndent(slicedOutput, "", "  ")
			if err != nil {
				return fmt.Errorf("[output]: %v", err)
			}
			err = PrintTable(bytes, &map[string]int{"category": 1, "name": 2})
			if err != nil {
				return fmt.Errorf("[output]: %v", err)
			}
			slicedOutput = []map[string]string{}
			rowsWidth = []int{}
			for _, row := range output {
				width := len(row["category"]) + len(row["name"]) + 20
				slicedOutput = append(slicedOutput, map[string]string{
					"category": row["category"],
					"name":     row["name"],
				})
				rowsWidth = append(rowsWidth, width)
			}
		}
	}
	bytes, err := json.MarshalIndent(slicedOutput, "", "  ")
	if err != nil {
		return fmt.Errorf("[output]: %v", err)
	}
	return PrintTable(bytes, &map[string]int{"category": 1, "name": 2})
}
