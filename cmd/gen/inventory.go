// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var InventoryCmd = &cobra.Command{
	Use: "inventory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	InventoryCmd.AddCommand(GetInventoryApiV2CostCompositionCmd)
	GetInventoryApiV2CostCompositionCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2CostCompositionCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2CostCompositionCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
	GetInventoryApiV2CostCompositionCmd.Flags().Int64("top", 0, "How many top values to return default is 5")

	InventoryCmd.AddCommand(GetInventoryApiV2CostMetricCmd)
	GetInventoryApiV2CostMetricCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2CostMetricCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2CostMetricCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2CostMetricCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2CostMetricCmd.Flags().Int64("page-number", 0, "page number - default is 1")
	GetInventoryApiV2CostMetricCmd.Flags().Int64("page-size", 0, "page size - default is 20")
	GetInventoryApiV2CostMetricCmd.Flags().String("sort-by", "", "Sort by field - default is cost")
	GetInventoryApiV2CostMetricCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2CostTrendCmd)
	GetInventoryApiV2CostTrendCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2CostTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2CostTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2CostTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
	GetInventoryApiV2CostTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2CostTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2AnalyticsSpendCompositionCmd)
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().Int64("top", 0, "How many top values to return default is 5")

	InventoryCmd.AddCommand(GetInventoryApiV2AnalyticsSpendMetricCmd)
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().Int64("page-number", 0, "page number - default is 1")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().Int64("page-size", 0, "page size - default is 20")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().String("sort-by", "", "Sort by field - default is cost")
	GetInventoryApiV2AnalyticsSpendMetricCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2AnalyticsSpendMetricsTrendCmd)
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("metric-ids", nil, "")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2AnalyticsSpendTrendCmd)
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("metric-ids", nil, "")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagCmd)
	GetInventoryApiV2ResourcesTagCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2ResourcesTagCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2ResourcesTagCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2ResourcesTagCmd.Flags().Int64("end-time", 0, "End time in unix timestamp format, default now")
	GetInventoryApiV2ResourcesTagCmd.Flags().Int64("min-count", 0, "Minimum number of resources with this tag value, default 1")

	InventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricResourceTypeCmd)
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("end-time", "", "timestamp for resource count in epoch seconds")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("resource-type", "", "ResourceType")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.MarkFlagRequired("resource-type")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("start-time", "", "timestamp for resource count change comparison in epoch seconds")

	InventoryCmd.AddCommand(GetInventoryApiV2ServicesCostTrendCmd)
	GetInventoryApiV2ServicesCostTrendCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().StringArray("services", nil, "Services to filter by")
	GetInventoryApiV2ServicesCostTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

}
