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
	GetBenchmarksCmd.AddCommand(GetInventoryApiV2ServicesSummaryServiceNameCmd)
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("connector", "", "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("connector-id", "", "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("service-name", "", "")

	GetBenchmarksCmd.AddCommand(GetInventoryApiV1AccountsResourceCountCmd)
	GetInventoryApiV1AccountsResourceCountCmd.Flags().String("provider", "", "")
	GetInventoryApiV1AccountsResourceCountCmd.Flags().StringArray("source-id", nil, "")

	GetBenchmarksCmd.AddCommand(GetInventoryApiV1ResourcesDistributionCmd)
	GetInventoryApiV1ResourcesDistributionCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV1ResourcesDistributionCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV1ResourcesDistributionCmd.Flags().String("time-window", "", "")

	GetBenchmarksCmd.AddCommand(GetInventoryApiV1ResourcesTopGrowingAccountsCmd)
	GetInventoryApiV1ResourcesTopGrowingAccountsCmd.Flags().Int64("count", 0, "")
	GetInventoryApiV1ResourcesTopGrowingAccountsCmd.Flags().String("provider", "", "")
	GetInventoryApiV1ResourcesTopGrowingAccountsCmd.Flags().String("time-window", "", "")

	GetBenchmarksCmd.AddCommand(GetInventoryApiV1ServicesDistributionCmd)
	GetInventoryApiV1ServicesDistributionCmd.Flags().String("provider", "", "")
	GetInventoryApiV1ServicesDistributionCmd.Flags().StringArray("source-id", nil, "")

	GetBenchmarksCmd.AddCommand(GetInventoryApiV2ServicesSummaryCmd)
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("tag", nil, "")
}
