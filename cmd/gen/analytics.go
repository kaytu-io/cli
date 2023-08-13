// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var AnalyticsCmd = &cobra.Command{
        Use: "analytics",
        RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Help()
    },
}

func init() {
    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsCompositionKeyCmd)
        GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().StringArray("connector", nil, "Connector types to filter by")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("end-time", "", "timestamp for resource count in epoch seconds")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("key", "", "Tag key")
GetInventoryApiV2AnalyticsCompositionKeyCmd.MarkFlagRequired("key")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("metric-type", "", "Metric type, default: assets")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().String("start-time", "", "timestamp for resource count change comparison in epoch seconds")
GetInventoryApiV2AnalyticsCompositionKeyCmd.Flags().Int64("top", 0, "How many top values to return default is 5")
GetInventoryApiV2AnalyticsCompositionKeyCmd.MarkFlagRequired("top")

    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsMetricCmd)
        GetInventoryApiV2AnalyticsMetricCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
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

    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsRegionsSummaryCmd)
        GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().Int64("end-time", 0, "end time in unix seconds - default is one week ago")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().Int64("page-number", 0, "page number - default is 1")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().Int64("page-size", 0, "page size - default is 20")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().String("sort-by", "", "column to sort by - default is resource_count")
GetInventoryApiV2AnalyticsRegionsSummaryCmd.Flags().Int64("start-time", 0, "start time in unix seconds - default is now")

    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsTagCmd)
        GetInventoryApiV2AnalyticsTagCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
GetInventoryApiV2AnalyticsTagCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
GetInventoryApiV2AnalyticsTagCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
GetInventoryApiV2AnalyticsTagCmd.Flags().Int64("end-time", 0, "End time in unix timestamp format, default now")
GetInventoryApiV2AnalyticsTagCmd.Flags().String("metric-type", "", "Metric type, default: assets")
GetInventoryApiV2AnalyticsTagCmd.Flags().Int64("min-count", 0, "Minimum number of resources/spend with this tag value, default 1")
GetInventoryApiV2AnalyticsTagCmd.Flags().Int64("start-time", 0, "Start time in unix timestamp format, default now - 1 month")

    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsTrendCmd)
        GetInventoryApiV2AnalyticsTrendCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
GetInventoryApiV2AnalyticsTrendCmd.Flags().String("datapoint-count", "", "maximum number of datapoints to return, default is 30")
GetInventoryApiV2AnalyticsTrendCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("ids", nil, "")
GetInventoryApiV2AnalyticsTrendCmd.Flags().String("metric-type", "", "Metric type, default: assets")
GetInventoryApiV2AnalyticsTrendCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
GetInventoryApiV2AnalyticsTrendCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

    
        AnalyticsCmd.AddCommand(GetInventoryApiV2AnalyticsCategoriesCmd)
        GetInventoryApiV2AnalyticsCategoriesCmd.Flags().String("metric-type", "", "Metric type, default: assets")

    
}