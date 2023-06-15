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
		DeleteBenchmarksAssignmentCmd.AddCommand(DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)
DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("benchmark-id", "", "")
DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("connection-id", "", "")

		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd)
GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd.Flags().String("benchmark-id", "", "")

		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsConnectionConnectionIdCmd)
GetComplianceApiV1AssignmentsConnectionConnectionIdCmd.Flags().String("connection-id", "", "")

		GetBenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsCmd)

		GetBenchmarksAssignmentCmd.AddCommand(PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)
PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("benchmark-id", "", "")
PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("connection-id", "", "")
}