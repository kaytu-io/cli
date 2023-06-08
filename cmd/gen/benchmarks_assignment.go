package gen

import (
	"github.com/spf13/cobra"
)

var GetBenchmarksAssignmentCmd = &cobra.Command{
	Use: "benchmarks_assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateBenchmarksAssignmentCmd = &cobra.Command{
	Use: "benchmarks_assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteBenchmarksAssignmentCmd = &cobra.Command{
	Use: "benchmarks_assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateBenchmarksAssignmentCmd = &cobra.Command{
	Use: "benchmarks_assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsCmd)

		GetBenchmarksAssignmentCmd.AddCommand(PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)

		GetBenchmarksAssignmentCmd.AddCommand(DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)

		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd)

		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsConnectionConnectionIdCmd)
}