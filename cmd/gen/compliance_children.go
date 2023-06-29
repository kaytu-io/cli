package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var GetComplianceApiV1FindingsMetricsCmd = &cobra.Command{
	Use: "findings-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_metrics] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1FindingsMetricsParams()

		req.SetEnd(flags.ReadInt64OptionalFlag(cmd, "End"))
		req.SetStart(flags.ReadInt64OptionalFlag(cmd, "Start"))

		resp, err := client.Compliance.GetComplianceAPIV1FindingsMetrics(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_metrics] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_metrics] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1QueriesQueryIdCmd = &cobra.Command{
	Use: "get-query",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_queries_query_id] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1QueriesQueryIDParams()

		req.SetQueryID(flags.ReadStringFlag(cmd, "QueryID"))

		resp, err := client.Compliance.GetComplianceAPIV1QueriesQueryID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_queries_query_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_queries_query_id] : %v", err)
		}

		return nil
	},
}

var PostComplianceApiV1FindingsCmd = &cobra.Command{
	Use: "get-findings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_findings] : %v", err)
		}

		req := compliance.NewPostComplianceAPIV1FindingsParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetFindingsRequest{
			Filters: &models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters{
				BenchmarkID:    flags.ReadStringArrayFlag(cmd, "BenchmarkID"),
				ConnectionID:   flags.ReadStringArrayFlag(cmd, "ConnectionID"),
				Connector:      flags.ReadEnumArrayFlag[models.SourceType](cmd, "Connector"),
				PolicyID:       flags.ReadStringArrayFlag(cmd, "PolicyID"),
				ResourceID:     flags.ReadStringArrayFlag(cmd, "ResourceID"),
				ResourceTypeID: flags.ReadStringArrayFlag(cmd, "ResourceTypeID"),
				Severity:       flags.ReadStringArrayFlag(cmd, "Severity"),
				Status:         flags.ReadEnumArrayFlag[models.TypesComplianceResult](cmd, "Status"),
			},
			Page: &models.GithubComKaytuIoKaytuEnginePkgComplianceAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Sorts: []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingSortItem{
				{
					Direction: models.GithubComKaytuIoKaytuEnginePkgComplianceAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GithubComKaytuIoKaytuEnginePkgComplianceAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Compliance.PostComplianceAPIV1Findings(req, auth)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_findings] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_findings] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarksBenchmarkIdCmd = &cobra.Command{
	Use: "get-benchmark",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarksBenchmarkIDParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksBenchmarkID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd = &cobra.Command{
	Use: "benchmark-policies",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id_policies] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksBenchmarkIDPolicies(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id_policies] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id_policies] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarksSummaryCmd = &cobra.Command{
	Use: "benchmarks-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_summary] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarksSummaryParams()

		req.SetEnd(flags.ReadInt64Flag(cmd, "End"))
		req.SetStart(flags.ReadInt64Flag(cmd, "Start"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_summary] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_summary] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarksCmd = &cobra.Command{
	Use: "list-benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarksParams()

		resp, err := client.Compliance.GetComplianceAPIV1Benchmarks(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd = &cobra.Command{
	Use: "get-policy",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_policies_policy_id] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarksPoliciesPolicyIDParams()

		req.SetPolicyID(flags.ReadStringFlag(cmd, "PolicyID"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksPoliciesPolicyID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_policies_policy_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_policies_policy_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd = &cobra.Command{
	Use: "benchmark-findings-top-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_benchmark_id_field_top_count] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))
		req.SetField(flags.ReadStringFlag(cmd, "Field"))

		resp, err := client.Compliance.GetComplianceAPIV1FindingsBenchmarkIDFieldTopCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_benchmark_id_field_top_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_benchmark_id_field_top_count] : %v", err)
		}

		return nil
	},
}

var PostComplianceApiV1AlarmsTopCmd = &cobra.Command{
	Use: "top-alarms",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_alarms_top] : %v", err)
		}

		req := compliance.NewPostComplianceAPIV1AlarmsTopParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgComplianceAPIGetTopFieldRequest{
			Count: flags.ReadInt64Flag(cmd, "Count"),
			Field: models.GithubComKaytuIoKaytuEnginePkgComplianceAPITopField(flags.ReadStringFlag(cmd, "Field")),
			Filters: &models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingFilters{
				BenchmarkID:    flags.ReadStringArrayFlag(cmd, "BenchmarkID"),
				ConnectionID:   flags.ReadStringArrayFlag(cmd, "ConnectionID"),
				Connector:      flags.ReadEnumArrayFlag[models.SourceType](cmd, "Connector"),
				PolicyID:       flags.ReadStringArrayFlag(cmd, "PolicyID"),
				ResourceID:     flags.ReadStringArrayFlag(cmd, "ResourceID"),
				ResourceTypeID: flags.ReadStringArrayFlag(cmd, "ResourceTypeID"),
				Severity:       flags.ReadStringArrayFlag(cmd, "Severity"),
				Status:         flags.ReadEnumArrayFlag[models.TypesComplianceResult](cmd, "Status"),
			},
		})

		resp, err := client.Compliance.PostComplianceAPIV1AlarmsTop(req, auth)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_alarms_top] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_alarms_top] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd = &cobra.Command{
	Use: "benchmark-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd = &cobra.Command{
	Use: "benchmark-summary-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary_result_trend] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetEnd(flags.ReadInt64Flag(cmd, "End"))
		req.SetStart(flags.ReadInt64Flag(cmd, "Start"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary_result_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary_result_trend] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd = &cobra.Command{
	Use: "benchmark-tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_tree] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetStatus(flags.ReadStringArrayFlag(cmd, "Status"))

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDTree(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_tree] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_tree] : %v", err)
		}

		return nil
	},
}
