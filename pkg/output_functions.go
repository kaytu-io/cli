package pkg

import (
	output "github.com/kaytu-io/cli-program/pkg/output_functions"
	"github.com/spf13/cobra"
)

type outputFunction func(cmd *cobra.Command, commandName string, obj interface{}) error

var outputFunctions = map[string]outputFunction{
	"get-inventory-api-v2-analytics-spend-table": output.AnalyticsSpendTable,
}
