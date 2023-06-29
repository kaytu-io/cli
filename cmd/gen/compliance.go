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
	GetComplianceCmd.AddCommand(GetComplianceApiV1FindingsMetricsCmd)
	GetComplianceApiV1FindingsMetricsCmd.Flags().Int64("end", 0, "")
	GetComplianceApiV1FindingsMetricsCmd.Flags().Int64("start", 0, "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1QueriesQueryIdCmd)
	GetComplianceApiV1QueriesQueryIdCmd.Flags().String("query-id", "", "")

	GetComplianceCmd.AddCommand(PostComplianceApiV1FindingsCmd)
	PostComplianceApiV1FindingsCmd.Flags().StringArray("benchmark-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("connection-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("connector", "", "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("policy-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("resource-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("resource-type-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("severity", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("status", "", "")

	PostComplianceApiV1FindingsCmd.Flags().Int64("no", 0, "")
	PostComplianceApiV1FindingsCmd.Flags().Int64("size", 0, "")

	PostComplianceApiV1FindingsCmd.Flags().String("direction", "", "")
	PostComplianceApiV1FindingsCmd.Flags().String("field", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdCmd)
	GetComplianceApiV1BenchmarksBenchmarkIdCmd.Flags().String("benchmark-id", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd)
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.Flags().String("benchmark-id", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksSummaryCmd)
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("end", 0, "")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("start", 0, "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksCmd)

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd)
	GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd.Flags().String("policy-id", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd)
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("benchmark-id", "", "")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().Int64("count", 0, "")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("field", "", "")

	GetComplianceCmd.AddCommand(PostComplianceApiV1AlarmsTopCmd)
	PostComplianceApiV1AlarmsTopCmd.Flags().Int64("count", 0, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("field", "", "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("benchmark-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("connection-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("connector", "", "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("policy-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("resource-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("resource-type-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("severity", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("status", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd.Flags().String("benchmark-id", "", "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().String("benchmark-id", "", "")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().Int64("end", 0, "")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().Int64("start", 0, "")

	GetComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.Flags().String("benchmark-id", "", "")
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.Flags().StringArray("status", nil, "")
}
