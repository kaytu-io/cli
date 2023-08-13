// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var InsightsCmd = &cobra.Command{
        Use: "insights",
        RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Help()
    },
}

func init() {
    
        InsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdCmd)
        GetComplianceApiV1InsightInsightIdCmd.Flags().StringArray("connection-id", nil, "filter the result by source id")
GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("end-time", 0, "unix seconds for the end time of the trend")
GetComplianceApiV1InsightInsightIdCmd.Flags().String("insight-id", "", "Insight ID")
GetComplianceApiV1InsightInsightIdCmd.MarkFlagRequired("insight-id")
GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("start-time", 0, "unix seconds for the start time of the trend")

    
        InsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightInsightIdCmd)
        GetComplianceApiV1MetadataInsightInsightIdCmd.Flags().String("insight-id", "", "Insight ID")
GetComplianceApiV1MetadataInsightInsightIdCmd.MarkFlagRequired("insight-id")

    
        InsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdTrendCmd)
        GetComplianceApiV1InsightInsightIdTrendCmd.Flags().StringArray("connection-id", nil, "filter the result by source id")
GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("datapoint-count", 0, "number of datapoints to return")
GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("end-time", 0, "unix seconds for the end time of the trend")
GetComplianceApiV1InsightInsightIdTrendCmd.Flags().String("insight-id", "", "Insight ID")
GetComplianceApiV1InsightInsightIdTrendCmd.MarkFlagRequired("insight-id")
GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("start-time", 0, "unix seconds for the start time of the trend")

    
        InsightsCmd.AddCommand(GetComplianceApiV1InsightGroupCmd)
        GetComplianceApiV1InsightGroupCmd.Flags().StringArray("connection-id", nil, "filter the result by source id")
GetComplianceApiV1InsightGroupCmd.Flags().StringArray("connector", nil, "filter insights by connector")
GetComplianceApiV1InsightGroupCmd.Flags().Int64("end-time", 0, "unix seconds for the end time of the trend")
GetComplianceApiV1InsightGroupCmd.Flags().Int64("start-time", 0, "unix seconds for the start time of the trend")
GetComplianceApiV1InsightGroupCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

    
        InsightsCmd.AddCommand(GetComplianceApiV1InsightCmd)
        GetComplianceApiV1InsightCmd.Flags().StringArray("connection-id", nil, "filter the result by source id")
GetComplianceApiV1InsightCmd.Flags().StringArray("connector", nil, "filter insights by connector")
GetComplianceApiV1InsightCmd.Flags().Int64("end-time", 0, "unix seconds for the end time of the trend")
GetComplianceApiV1InsightCmd.Flags().Int64("start-time", 0, "unix seconds for the start time of the trend")
GetComplianceApiV1InsightCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

    
        InsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightCmd)
        
    
}