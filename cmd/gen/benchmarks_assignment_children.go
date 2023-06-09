package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks_assignment"
	"github.com/spf13/cobra"
)

var DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdConnectionConnectionId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		_, err = client.BenchmarksAssignment.DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID(benchmarks_assignment.NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		return nil
	},
}
var GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd = &cobra.Command{
	Use: "benchmarkBenchmarkBenchmarkId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_benchmark_benchmark_id] : %v", err)
		}

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1AssignmentsBenchmarkBenchmarkID(benchmarks_assignment.NewGetComplianceAPIV1AssignmentsBenchmarkBenchmarkIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_benchmark_benchmark_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_benchmark_benchmark_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1AssignmentsConnectionConnectionIdCmd = &cobra.Command{
	Use: "connectionConnectionConnectionId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_connection_connection_id] : %v", err)
		}

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1AssignmentsConnectionConnectionID(benchmarks_assignment.NewGetComplianceAPIV1AssignmentsConnectionConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_connection_connection_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_connection_connection_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1AssignmentsCmd = &cobra.Command{
	Use: "Assignments",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments] : %v", err)
		}

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1Assignments(benchmarks_assignment.NewGetComplianceAPIV1AssignmentsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments] : %v", err)
		}

		return nil
	},
}

var PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd = &cobra.Command{
	Use: "benchmarkBenchmarkIdConnectionConnectionId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		resp, err := client.BenchmarksAssignment.PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID(benchmarks_assignment.NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		return nil
	},
}
