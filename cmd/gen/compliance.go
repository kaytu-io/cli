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
	GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd.Flags().StringArray("severities", nil, "Severities to filter by")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksSummaryCmd)

	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksSummaryCmd.Flags().Int64("time-at", 0, "timestamp for values in epoch seconds")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdSummaryCmd.Flags().Int64("time-at", 0, "timestamp for values in epoch seconds")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdTreeCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdTreeCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdTreeCmd.MarkFlagRequired("benchmark-id")

	ComplianceCmd.AddCommand(GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd)

	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("benchmark-id", "", "Benchmark ID")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.MarkFlagRequired("benchmark-id")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connection-group", nil, "Connection groups to filter by ")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("end-time", "", "timestamp for end of the chart in epoch seconds")
	GetComplianceApiV1BenchmarksBenchmarkIdTrendCmd.Flags().String("start-time", "", "timestamp for start of the chart in epoch seconds")

	ComplianceCmd.AddCommand(PostComplianceApiV1FindingsCmd)

	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-benchmark-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-connection-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("filters-connector", "", "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-policy-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-resource-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-resource-type-id", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().StringArray("filters-severity", nil, "")
	PostComplianceApiV1FindingsCmd.Flags().String("filters-status", "", "")

	PostComplianceApiV1FindingsCmd.Flags().Int64("page-no", 0, "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("page-no")
	PostComplianceApiV1FindingsCmd.Flags().Int64("page-size", 0, "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("page-size")

	PostComplianceApiV1FindingsCmd.Flags().String("sorts-direction", "", "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("sorts-direction")
	PostComplianceApiV1FindingsCmd.Flags().String("sorts-field", "", "")
	PostComplianceApiV1FindingsCmd.MarkFlagRequired("sorts-field")

	ComplianceCmd.AddCommand(GetComplianceApiV1QueriesSyncCmd)

}
