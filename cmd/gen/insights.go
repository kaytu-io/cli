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
		GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightKeyCmd)

		GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataTagInsightCmd)

		GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdCmd)

		GetInsightsCmd.AddCommand(GetComplianceApiV1InsightInsightIdTrendCmd)

		GetInsightsCmd.AddCommand(GetComplianceApiV1InsightCmd)

		GetInsightsCmd.AddCommand(GetComplianceApiV1MetadataInsightCmd)
}