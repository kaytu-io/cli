package gen

import (
	"github.com/spf13/cobra"
)

var GetInsightsCmd = &cobra.Command{
	Use: "insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateInsightsCmd = &cobra.Command{
	Use: "insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteInsightsCmd = &cobra.Command{
	Use: "insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateInsightsCmd = &cobra.Command{
	Use: "insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GetInsightsCmd.AddCommand(GetInventoryApiV2InsightsInsightIdTrendCmd)
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().String("insight-id", "", "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetInventoryApiV2InsightsCmd)
	GetInventoryApiV2InsightsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().StringArray("insight-id", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().Int64("time", 0, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightGroupCmd)
	GetComplianceApiV1InsightGroupCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightGroupCmd.Flags().StringArray("connector", nil, "")
	GetComplianceApiV1InsightGroupCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightGroupCmd.Flags().Int64("start-time", 0, "")
	GetComplianceApiV1InsightGroupCmd.Flags().StringArray("tag", nil, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightCmd)
	GetComplianceApiV1InsightCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightCmd.Flags().StringArray("connector", nil, "")
	GetComplianceApiV1InsightCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightCmd.Flags().Int64("start-time", 0, "")
	GetComplianceApiV1InsightCmd.Flags().StringArray("tag", nil, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightCmd)
	GetComplianceApiV1MetadataInsightCmd.Flags().StringArray("connector", nil, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightCmd)

	GetInsightsCmd.AddCommand(GetInventoryApiV2InsightsJobJobIdCmd)
	GetInventoryApiV2InsightsJobJobIdCmd.Flags().String("job-id", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightGroupInsightGroupIdCmd)
	GetComplianceApiV1InsightGroupInsightGroupIdCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightGroupInsightGroupIdCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightGroupInsightGroupIdCmd.Flags().String("insight-group-id", "", "")
	GetComplianceApiV1InsightGroupInsightGroupIdCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd)
	GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd.Flags().Int64("datapoint-count", 0, "")
	GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd.Flags().String("insight-group-id", "", "")
	GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightInsightIdCmd)
	GetComplianceApiV1MetadataInsightInsightIdCmd.Flags().String("insight-id", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightKeyCmd)
	GetComplianceApiV1MetadataTagInsightKeyCmd.Flags().String("key", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdTrendCmd)
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("datapoint-count", 0, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().String("insight-id", "", "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdCmd)
	GetComplianceApiV1InsightInsightIdCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().String("insight-id", "", "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetInventoryApiV2InsightsInsightIdCmd)
	GetInventoryApiV2InsightsInsightIdCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsInsightIdCmd.Flags().String("insight-id", "", "")
	GetInventoryApiV2InsightsInsightIdCmd.Flags().Int64("time", 0, "")
}
