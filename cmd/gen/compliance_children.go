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

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd = &cobra.Command{
	Use: "benchmark-benchmark-id-summary",
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

var GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd = &cobra.Command{
	Use: "benchmarks-benchmark-id-policies",
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

var GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd = &cobra.Command{
	Use: "benchmarks-policies-policy-id",
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
	Use: "findings-benchmark-id-field-top-count",
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
	Use: "queries-query-id",
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

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd = &cobra.Command{
	Use: "benchmark-benchmark-id-summary-result-trend",
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

var GetScheduleApiV1BenchmarkEvaluationsCmd = &cobra.Command{
	Use: "benchmark-evaluations",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		req := compliance.NewGetScheduleAPIV1BenchmarkEvaluationsParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIListBenchmarkEvaluationsRequest{
			BenchmarkID:       flags.ReadStringFlag(cmd, "BenchmarkID"),
			ConnectionID:      flags.ReadStringFlag(cmd, "ConnectionID"),
			Connector:         models.SourceType(flags.ReadStringFlag(cmd, "Connector")),
			EvaluatedAtAfter:  flags.ReadInt64Flag(cmd, "EvaluatedAtAfter"),
			EvaluatedAtBefore: flags.ReadInt64Flag(cmd, "EvaluatedAtBefore"),
		})

		resp, err := client.Compliance.GetScheduleAPIV1BenchmarkEvaluations(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd = &cobra.Command{
	Use: "benchmark-benchmark-id-tree",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_tree] : %v", err)
		}

		req := compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetStatus(flags.ReadStringFlag(cmd, "Status"))

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

var GetComplianceApiV1BenchmarksBenchmarkIdCmd = &cobra.Command{
	Use: "benchmarks-benchmark-id",
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

var GetComplianceApiV1BenchmarksCmd = &cobra.Command{
	Use: "benchmarks",
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

var PostComplianceApiV1AlarmsTopCmd = &cobra.Command{
	Use: "alarms-top",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_alarms_top] : %v", err)
		}

		req := compliance.NewPostComplianceAPIV1AlarmsTopParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetTopFieldRequest{
			Count: flags.ReadInt64Flag(cmd, "Count"),
			Field: models.GitlabComKeibiengineKeibiEnginePkgComplianceAPITopField(flags.ReadStringFlag(cmd, "Field")),
			Filters: &models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIFindingFilters{
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

var PostComplianceApiV1FindingsCmd = &cobra.Command{
	Use: "findings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_findings] : %v", err)
		}

		req := compliance.NewPostComplianceAPIV1FindingsParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIGetFindingsRequest{
			Filters: &models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIFindingFilters{
				BenchmarkID:    flags.ReadStringArrayFlag(cmd, "BenchmarkID"),
				ConnectionID:   flags.ReadStringArrayFlag(cmd, "ConnectionID"),
				Connector:      flags.ReadEnumArrayFlag[models.SourceType](cmd, "Connector"),
				PolicyID:       flags.ReadStringArrayFlag(cmd, "PolicyID"),
				ResourceID:     flags.ReadStringArrayFlag(cmd, "ResourceID"),
				ResourceTypeID: flags.ReadStringArrayFlag(cmd, "ResourceTypeID"),
				Severity:       flags.ReadStringArrayFlag(cmd, "Severity"),
				Status:         flags.ReadEnumArrayFlag[models.TypesComplianceResult](cmd, "Status"),
			},
			Page: &models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIFindingSortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GitlabComKeibiengineKeibiEnginePkgComplianceAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
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
