package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks_assignment"
	"github.com/spf13/cobra"
)

var PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd = &cobra.Command{
	Use: "add-assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		req := benchmarks_assignment.NewPostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))

		resp, err := client.BenchmarksAssignment.PostComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID(req, auth)
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

var DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd = &cobra.Command{
	Use: "remove-assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		req := benchmarks_assignment.NewDeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionIDParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))
		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))

		_, err = client.BenchmarksAssignment.DeleteComplianceAPIV1AssignmentsBenchmarkIDConnectionConnectionID(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_compliance_api_v_1_assignments_benchmark_id_connection_connection_id] : %v", err)
		}

		return nil
	},
}
var GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd = &cobra.Command{
	Use: "benchmark-assignments",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_benchmark_benchmark_id] : %v", err)
		}

		req := benchmarks_assignment.NewGetComplianceAPIV1AssignmentsBenchmarkBenchmarkIDParams()

		req.SetBenchmarkID(flags.ReadStringFlag(cmd, "BenchmarkID"))

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1AssignmentsBenchmarkBenchmarkID(req, auth)
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
	Use: "connection-assignments",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments_connection_connection_id] : %v", err)
		}

		req := benchmarks_assignment.NewGetComplianceAPIV1AssignmentsConnectionConnectionIDParams()

		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1AssignmentsConnectionConnectionID(req, auth)
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
	Use: "list-assignments",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_assignments] : %v", err)
		}

		req := benchmarks_assignment.NewGetComplianceAPIV1AssignmentsParams()

		resp, err := client.BenchmarksAssignment.GetComplianceAPIV1Assignments(req, auth)
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
