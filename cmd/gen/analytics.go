// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var SpendCmd = &cobra.Command{
	Use: "spend",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var AssetsCmd = &cobra.Command{
	Use: "assets",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsMetricCmd)

	GetInventoryApiV2AnalyticsMetricCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().String("end-time", "", "timestamp for resource count in epoch seconds")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().StringArray("metric-i-ds", nil, "Metric IDs")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().String("metric-type", "", "Metric type, default: assets")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().Int64("min-count", 0, "Minimum number of resources with this tag value, default 1")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().Int64("page-number", 0, "page number - default is 1")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().Int64("page-size", 0, "page size - default is 20")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().String("sort-by", "", "Sort by field - default is count")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().String("start-time", "", "timestamp for resource count change comparison in epoch seconds")
	GetInventoryApiV2AnalyticsMetricCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsCompositionKeyCmd)

	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().StringArray("connector", nil, "Connector types to filter by")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("end-time", "", "timestamp for resource count in epoch seconds")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("key", "", "Tag key")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.MarkFlagRequired("key")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("metric-type", "", "Metric type, default: assets")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("start-time", "", "timestamp for resource count change comparison in epoch seconds")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().Int64("top", 0, "How many top values to return default is 5")
	GetInventoryApiV2AnalyticsCompositionKeyCmd.MarkFlagRequired("top")

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsTrendCmd)

	GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("ids", nil, "")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().String("metric-type", "", "Metric type, default: assets")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
	GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsSpendTrendCmd)

	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("granularity", "", "Granularity of the table, default is daily")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().StringArray("metric-ids", nil, "")
	GetInventoryApiV2AnalyticsSpendTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsSpendCompositionCmd)

	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
	GetInventoryApiV2AnalyticsSpendCompositionCmd.Flags().Int64("top", 0, "How many top values to return default is 5")

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsSpendMetricsTrendCmd)

	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("granularity", "", "Granularity of the table, default is daily")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().StringArray("metric-ids", nil, "")
	GetInventoryApiV2AnalyticsSpendMetricsTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsTagCmd)

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsTagCmd)

	GetInventoryApiV2AnalyticsTagCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2AnalyticsTagCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsTagCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsTagCmd.Flags().String("end-time", "", "End time in unix timestamp format, default now")
	GetInventoryApiV2AnalyticsTagCmd.Flags().String("metric-type", "", "Metric type, default: assets")
	GetInventoryApiV2AnalyticsTagCmd.Flags().Int64("min-count", 0, "Minimum number of resources/spend with this tag value, default 1")
	GetInventoryApiV2AnalyticsTagCmd.Flags().String("start-time", "", "Start time in unix timestamp format, default now - 1 month")

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsCategoriesCmd)

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsCategoriesCmd)

	GetInventoryApiV2AnalyticsCategoriesCmd.Flags().String("metric-type", "", "Metric type, default: assets")

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsTableCmd)

	GetInventoryApiV2AnalyticsTableCmd.Flags().String("dimension", "", "Dimension of the table, default is metric")
	GetInventoryApiV2AnalyticsTableCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsTableCmd.Flags().String("granularity", "", "Granularity of the table, default is daily")
	GetInventoryApiV2AnalyticsTableCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

	AssetsCmd.AddCommand(GetInventoryApiV2AnalyticsMetricsListCmd)

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsMetricsListCmd)

	GetInventoryApiV2AnalyticsMetricsListCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetInventoryApiV2AnalyticsMetricsListCmd.Flags().String("metric-type", "", "Metric type, default: assets")

	//SpendCmd.AddCommand(GetInventoryApiV2AnalyticsSpendTableCmd)

	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().String("dimension", "", "Dimension of the table, default is metric")
	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().String("granularity", "", "Granularity of the table, default is daily")
	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().StringArray("metric-ids", nil, "")
	GetInventoryApiV2AnalyticsSpendTableCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")

}
