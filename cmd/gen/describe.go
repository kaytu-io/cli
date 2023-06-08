package gen

import (
	"github.com/spf13/cobra"
)

var GetDescribeCmd = &cobra.Command{
	Use: "describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateDescribeCmd = &cobra.Command{
	Use: "describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteDescribeCmd = &cobra.Command{
	Use: "describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateDescribeCmd = &cobra.Command{
	Use: "describe",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetDescribeCmd.AddCommand(GetScheduleApiV0InsightTriggerCmd)

		GetDescribeCmd.AddCommand(PostScheduleApiV1DescribeResourceCmd)

		GetDescribeCmd.AddCommand(PutScheduleApiV1BenchmarkEvaluationTriggerCmd)

		GetDescribeCmd.AddCommand(PutScheduleApiV1ComplianceTriggerCmd)

		GetDescribeCmd.AddCommand(GetScheduleApiV0ComplianceSummarizerTriggerCmd)

		GetDescribeCmd.AddCommand(GetScheduleApiV0ComplianceTriggerCmd)

		GetDescribeCmd.AddCommand(GetScheduleApiV0DescribeTriggerCmd)

		GetDescribeCmd.AddCommand(GetScheduleApiV0SummarizeTriggerCmd)

		GetDescribeCmd.AddCommand(PutScheduleApiV1DescribeTriggerConnectionIdCmd)

		GetDescribeCmd.AddCommand(PutScheduleApiV1InsightEvaluationTriggerCmd)
}