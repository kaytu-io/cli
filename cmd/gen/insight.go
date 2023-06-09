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
		GetInsightCmd.AddCommand(GetInventoryApiV2InsightsInsightIdTrendCmd)

		GetInsightCmd.AddCommand(GetInventoryApiV2InsightsCmd)

		GetInsightCmd.AddCommand(GetInventoryApiV2InsightsInsightIdCmd)
}