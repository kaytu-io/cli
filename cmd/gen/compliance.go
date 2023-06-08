package gen

import (
	"github.com/spf13/cobra"
)

var GetComplianceCmd = &cobra.Command{
	Use: "compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateComplianceCmd = &cobra.Command{
	Use: "compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteComplianceCmd = &cobra.Command{
	Use: "compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateComplianceCmd = &cobra.Command{
	Use: "compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1QueriesQueryIdCmd)

		GetComplianceCmd.AddCommand(GetScheduleApiV1BenchmarkEvaluationsCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd)

		GetComplianceCmd.AddCommand(PostComplianceApiV1FindingsCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksSummaryCmd)

		GetComplianceCmd.AddCommand(PostComplianceApiV1AlarmsTopCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd)

		GetComplianceCmd.AddCommand(GetComplianceApiV1FindingsMetricsCmd)
}