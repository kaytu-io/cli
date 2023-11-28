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

	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("benchmark-id", "", "BenchmarkID")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().Int64("count", 0, "Count")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("count")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().String("field", "", "Field")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.MarkFlagRequired("field")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("severities", nil, "Severities to filter by defaults to all severities except passed")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksSummaryCmd)

	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("time-at", 0, "timestamp for values in epoch seconds")

	ComplianceCmd.AddCommand(GetComplianceApiV1FindingsBenchmarkIdAccountsCmd)

	GetComplianceApiV1FindingsBenchmarkIdAccountsCmd.Flags().String("benchmark-id", "", "BenchmarkID")
	GetComplianceApiV1FindingsBenchmarkIdAccountsCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1FindingsBenchmarkIdAccountsCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1FindingsBenchmarkIdAccountsCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().Int64("time-at", 0, "timestamp for values in epoch seconds")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("end-time", "", "timestamp for end of the chart in epoch seconds")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("start-time", "", "timestamp for start of the chart in epoch seconds")

	ComplianceCmd.AddCommand(GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd)

	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().String("benchmark-id", "", "BenchmarkID")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().String("field", "", "Field")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.MarkFlagRequired("field")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")
	GetComplianceApiV1FindingsBenchmarkIdFieldCountCmd.Flags().StringArray("severities", nil, "Severities to filter by defaults to all severities except passed")

	ComplianceCmd.AddCommand(PostComplianceApiV1FindingsCmd)

	PostComplianceApiV1FindingsCmd.Flags().String("after-sort-key", "", "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-benchmark-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-connection-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("filters-connector", "", "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-policy-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-resource-collection", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-resource-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-resource-type-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-severity", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("filters-status", "", "")

	PostComplianceApiV1FindingsCmd.Flags().Int64("limit", 0, "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("limit")
	PostComplianceApiV1FindingsCmd.Flags().String("sort", "", "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("sort")

	ComplianceCmd.AddCommand(PostComplianceApiV1AiPolicyPolicyIdRemediationCmd)

	PostComplianceApiV1AiPolicyPolicyIdRemediationCmd.Flags().String("policy-id", "", "PolicyID")
	PostComplianceApiV1AiPolicyPolicyIdRemediationCmd.MarkFlagRequired("policy-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1QueriesSyncCmd)

}
