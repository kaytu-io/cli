// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var ComplianceCmd = &cobra.Command{
	Use: "compliance",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	ComplianceCmd.AddCommand(GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd)
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().Int64("count", 0, "")
	// GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("count")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("count")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("field", "", "")
	// GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("field")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("field")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd)
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.MarkFlagRequired("benchmark-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd.MarkFlagRequired("benchmark-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().Int64("end", 0, "")
	// GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("end")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("end")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.Flags().Int64("start", 0, "")
	// GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("start")
	GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd.MarkFlagRequired("start")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd)
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.Flags().StringArray("status", nil, "")
	GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd.MarkFlagRequired("status")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksSummaryCmd)
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("end", 0, "")
	// GetComplianceApiV1BenchmarksSummaryCmd.MarkFlagRequired("end")
	GetComplianceApiV1BenchmarksSummaryCmd.MarkFlagRequired("end")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("start", 0, "")
	// GetComplianceApiV1BenchmarksSummaryCmd.MarkFlagRequired("start")
	GetComplianceApiV1BenchmarksSummaryCmd.MarkFlagRequired("start")

	ComplianceCmd.AddCommand(GetComplianceApiV1FindingsMetricsCmd)
	GetComplianceApiV1FindingsMetricsCmd.Flags().Int64("end", 0, "")
	GetComplianceApiV1FindingsMetricsCmd.Flags().Int64("start", 0, "")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdCmd)
	GetComplianceApiV1BenchmarksBenchmarkIdCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1BenchmarksBenchmarkIdCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdCmd.MarkFlagRequired("benchmark-id")

	ComplianceCmd.AddCommand(PostComplianceApiV1FindingsCmd)
	PostComplianceApiV1FindingsCmd.Flags().StringArray("benchmark-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("connection-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("connector", "", "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("policy-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("resource-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("resource-type-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("severity", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("status", "", "")

	PostComplianceApiV1FindingsCmd.Flags().Int64("no", 0, "")
	// PostComplianceApiV1FindingsCmd.MarkFlagRequired("no")
	PostComplianceApiV1FindingsCmd.Flags().Int64("size", 0, "")
	// PostComplianceApiV1FindingsCmd.MarkFlagRequired("size")

	PostComplianceApiV1FindingsCmd.Flags().String("direction", "", "")
	PostComplianceApiV1FindingsCmd.Flags().String("field", "", "")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd)
	GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd.Flags().String("policy-id", "", "")
	// GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd.MarkFlagRequired("policy-id")
	GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd.MarkFlagRequired("policy-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1QueriesQueryIdCmd)
	GetComplianceApiV1QueriesQueryIdCmd.Flags().String("query-id", "", "")
	// GetComplianceApiV1QueriesQueryIdCmd.MarkFlagRequired("query-id")
	GetComplianceApiV1QueriesQueryIdCmd.MarkFlagRequired("query-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksCmd)

	ComplianceCmd.AddCommand(GetComplianceApiV1QueriesSyncCmd)

	ComplianceCmd.AddCommand(PostComplianceApiV1AlarmsTopCmd)
	PostComplianceApiV1AlarmsTopCmd.Flags().Int64("count", 0, "")
	// PostComplianceApiV1AlarmsTopCmd.MarkFlagRequired("count")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("field", "", "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("benchmark-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("connection-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("connector", "", "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("policy-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("resource-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("resource-type-id", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().StringArray("severity", nil, "")
	PostComplianceApiV1AlarmsTopCmd.Flags().String("status", "", "")

}
