package gen

import (
	"github.com/spf13/cobra"
)

var GetInsightCmd = &cobra.Command{
	Use: "insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateInsightCmd = &cobra.Command{
	Use: "insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteInsightCmd = &cobra.Command{
	Use: "insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateInsightCmd = &cobra.Command{
	Use: "insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GetInsightCmd.AddCommand(GetInventoryApiV2InsightsCmd)
	GetInventoryApiV2InsightsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().StringArray("insight-id", nil, "")
	GetInventoryApiV2InsightsCmd.Flags().Int64("time", 0, "")

	GetInsightCmd.AddCommand(GetInventoryApiV2InsightsInsightIdCmd)
	GetInventoryApiV2InsightsInsightIdCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsInsightIdCmd.Flags().String("insight-id", "", "")
	GetInventoryApiV2InsightsInsightIdCmd.Flags().Int64("time", 0, "")

	GetInsightCmd.AddCommand(GetInventoryApiV2InsightsInsightIdTrendCmd)
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().String("insight-id", "", "")
	GetInventoryApiV2InsightsInsightIdTrendCmd.Flags().Int64("start-time", 0, "")

	GetInsightCmd.AddCommand(GetInventoryApiV2InsightsJobJobIdCmd)
	GetInventoryApiV2InsightsJobJobIdCmd.Flags().String("job-id", "", "")
}
