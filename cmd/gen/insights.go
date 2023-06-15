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
	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdTrendCmd)
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("datapoint-count", 0, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().String("insight-id", "", "")
	GetComplianceApiV1InsightInsightIdTrendCmd.Flags().Int64("start-time", 0, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightCmd)
	GetComplianceApiV1InsightCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightCmd.Flags().StringArray("connector", nil, "")
	GetComplianceApiV1InsightCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightCmd.Flags().Int64("start-time", 0, "")
	GetComplianceApiV1InsightCmd.Flags().String("tag", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightInsightIdCmd)
	GetComplianceApiV1MetadataInsightInsightIdCmd.Flags().String("insight-id", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightCmd)
	GetComplianceApiV1MetadataInsightCmd.Flags().StringArray("connector", nil, "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightKeyCmd)
	GetComplianceApiV1MetadataTagInsightKeyCmd.Flags().String("key", "", "")

	GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightCmd)

	GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdCmd)
	GetComplianceApiV1InsightInsightIdCmd.Flags().StringArray("connection-id", nil, "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("end-time", 0, "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().String("insight-id", "", "")
	GetComplianceApiV1InsightInsightIdCmd.Flags().Int64("start-time", 0, "")
}
