// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var BenchmarksAssignmentCmd = &cobra.Command{
	Use: "benchmarks_assignment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	BenchmarksAssignmentCmd.AddCommand(PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)
	PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("benchmark-id", "", "")
	// PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.MarkFlagRequired("benchmark-id")
	PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("connection-id", "", "")
	// PostComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.MarkFlagRequired("connection-id")

	BenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd)
	GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd.Flags().String("benchmark-id", "", "")
	// GetComplianceApiV1AssignmentsBenchmarkBenchmarkIdCmd.MarkFlagRequired("benchmark-id")

	BenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsConnectionConnectionIdCmd)
	GetComplianceApiV1AssignmentsConnectionConnectionIdCmd.Flags().String("connection-id", "", "")
	// GetComplianceApiV1AssignmentsConnectionConnectionIdCmd.MarkFlagRequired("connection-id")

	BenchmarksAssignmentCmd.AddCommand(GetComplianceApiV1AssignmentsCmd)

	BenchmarksAssignmentCmd.AddCommand(DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd)
	DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("benchmark-id", "", "")
	// DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.MarkFlagRequired("benchmark-id")
	DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.Flags().String("connection-id", "", "")
	// DeleteComplianceApiV1AssignmentsBenchmarkIdConnectionConnectionIdCmd.MarkFlagRequired("connection-id")

}
