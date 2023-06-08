package gen

import (
	"github.com/spf13/cobra"
)

var GetBenchmarksCmd = &cobra.Command{
	Use: "benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateBenchmarksCmd = &cobra.Command{
	Use: "benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteBenchmarksCmd = &cobra.Command{
	Use: "benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateBenchmarksCmd = &cobra.Command{
	Use: "benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetBenchmarksCmd.AddCommand(GetInventoryApiV1ResourcesTopGrowingAccountsCmd)

		GetBenchmarksCmd.AddCommand(GetInventoryApiV1ServicesDistributionCmd)

		GetBenchmarksCmd.AddCommand(GetInventoryApiV2ServicesSummaryCmd)

		GetBenchmarksCmd.AddCommand(GetInventoryApiV2ServicesSummaryServiceNameCmd)

		GetBenchmarksCmd.AddCommand(GetInventoryApiV1AccountsResourceCountCmd)

		GetBenchmarksCmd.AddCommand(GetInventoryApiV1ResourcesDistributionCmd)
}