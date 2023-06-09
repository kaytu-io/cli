package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/compliance"
	"github.com/spf13/cobra"
)

var GetComplianceApiV1BenchmarkBenchmarkIdTreeCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdTree",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_tree] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDTree(compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDTreeParams(), auth)
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

var GetComplianceApiV1BenchmarksCmd = &cobra.Command{
	Use: "Benchmarks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1Benchmarks(compliance.NewGetComplianceAPIV1BenchmarksParams(), auth)
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

var GetComplianceApiV1BenchmarksSummaryCmd = &cobra.Command{
	Use: "summarySummary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_summary] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksSummary(compliance.NewGetComplianceAPIV1BenchmarksSummaryParams(), auth)
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

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdSummary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummary(compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryParams(), auth)
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

var GetComplianceApiV1BenchmarksBenchmarkIdCmd = &cobra.Command{
	Use: "benchmarkBenchmarkId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksBenchmarkID(compliance.NewGetComplianceAPIV1BenchmarksBenchmarkIDParams(), auth)
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

var GetComplianceApiV1FindingsBenchmarkIdFieldTopCountCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdFieldTopCount",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_benchmark_id_field_top_count] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1FindingsBenchmarkIDFieldTopCount(compliance.NewGetComplianceAPIV1FindingsBenchmarkIDFieldTopCountParams(), auth)
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

var GetScheduleApiV1BenchmarkEvaluationsCmd = &cobra.Command{
	Use: "evaluationsEvaluations",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		resp, err := client.Compliance.GetScheduleAPIV1BenchmarkEvaluations(compliance.NewGetScheduleAPIV1BenchmarkEvaluationsParams(), auth)
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

var GetComplianceApiV1BenchmarksPoliciesPolicyIdCmd = &cobra.Command{
	Use: "policiesPoliciesPolicyId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_policies_policy_id] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksPoliciesPolicyID(compliance.NewGetComplianceAPIV1BenchmarksPoliciesPolicyIDParams(), auth)
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

var GetComplianceApiV1QueriesQueryIdCmd = &cobra.Command{
	Use: "queryQueryId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_queries_query_id] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1QueriesQueryID(compliance.NewGetComplianceAPIV1QueriesQueryIDParams(), auth)
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
	Use: "Findings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_findings] : %v", err)
		}

		resp, err := client.Compliance.PostComplianceAPIV1Findings(compliance.NewPostComplianceAPIV1FindingsParams(), auth)
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

var GetComplianceApiV1BenchmarkBenchmarkIdSummaryResultTrendCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdSummaryResultTrend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmark_benchmark_id_summary_result_trend] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrend(compliance.NewGetComplianceAPIV1BenchmarkBenchmarkIDSummaryResultTrendParams(), auth)
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

var GetComplianceApiV1BenchmarksBenchmarkIdPoliciesCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdPolicies",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_benchmarks_benchmark_id_policies] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1BenchmarksBenchmarkIDPolicies(compliance.NewGetComplianceAPIV1BenchmarksBenchmarkIDPoliciesParams(), auth)
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

var GetComplianceApiV1FindingsMetricsCmd = &cobra.Command{
	Use: "metricsMetrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_findings_metrics] : %v", err)
		}

		resp, err := client.Compliance.GetComplianceAPIV1FindingsMetrics(compliance.NewGetComplianceAPIV1FindingsMetricsParams(), auth)
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

var PostComplianceApiV1AlarmsTopCmd = &cobra.Command{
	Use: "topTop",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_alarms_top] : %v", err)
		}

		resp, err := client.Compliance.PostComplianceAPIV1AlarmsTop(compliance.NewPostComplianceAPIV1AlarmsTopParams(), auth)
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
