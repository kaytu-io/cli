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
	GetDescribeCmd.AddCommand(GetScheduleApiV0SummarizeTriggerCmd)

	GetDescribeCmd.AddCommand(PostScheduleApiV1DescribeResourceCmd)
	PostScheduleApiV1DescribeResourceCmd.Flags().String("access-key", "", "")
	PostScheduleApiV1DescribeResourceCmd.Flags().String("account-id", "", "")
	PostScheduleApiV1DescribeResourceCmd.Flags().String("additional-fields", "", "")
	PostScheduleApiV1DescribeResourceCmd.Flags().String("provider", "", "")
	PostScheduleApiV1DescribeResourceCmd.Flags().String("resource-type", "", "")
	PostScheduleApiV1DescribeResourceCmd.Flags().String("secret-key", "", "")

	GetDescribeCmd.AddCommand(PostScheduleApiV1StacksInsightTriggerCmd)
	PostScheduleApiV1StacksInsightTriggerCmd.Flags().String("request", "", "")

	GetDescribeCmd.AddCommand(PutScheduleApiV1InsightEvaluationTriggerCmd)
	PutScheduleApiV1InsightEvaluationTriggerCmd.Flags().String("connection-id", "", "")
	PutScheduleApiV1InsightEvaluationTriggerCmd.Flags().Int64("insight-id", 0, "")
	PutScheduleApiV1InsightEvaluationTriggerCmd.Flags().StringArray("resource-i-ds", nil, "")

	GetDescribeCmd.AddCommand(GetScheduleApiV0ComplianceSummarizerTriggerCmd)

	GetDescribeCmd.AddCommand(GetScheduleApiV0ComplianceTriggerCmd)

	GetDescribeCmd.AddCommand(GetScheduleApiV0DescribeTriggerCmd)

	GetDescribeCmd.AddCommand(GetScheduleApiV0InsightTriggerCmd)

	GetDescribeCmd.AddCommand(GetScheduleApiV1BenchmarkEvaluationsCmd)
	GetScheduleApiV1BenchmarkEvaluationsCmd.Flags().String("benchmark-id", "", "")
	GetScheduleApiV1BenchmarkEvaluationsCmd.Flags().String("connection-id", "", "")
	GetScheduleApiV1BenchmarkEvaluationsCmd.Flags().String("connector", "", "")
	GetScheduleApiV1BenchmarkEvaluationsCmd.Flags().Int64("evaluated-at-after", 0, "")
	GetScheduleApiV1BenchmarkEvaluationsCmd.Flags().Int64("evaluated-at-before", 0, "")

	GetDescribeCmd.AddCommand(GetScheduleApiV1InsightJobJobIdCmd)
	GetScheduleApiV1InsightJobJobIdCmd.Flags().String("job-id", "", "")

	GetDescribeCmd.AddCommand(PutScheduleApiV1BenchmarkEvaluationTriggerCmd)
	PutScheduleApiV1BenchmarkEvaluationTriggerCmd.Flags().String("benchmark-id", "", "")
	PutScheduleApiV1BenchmarkEvaluationTriggerCmd.Flags().String("connection-id", "", "")
	PutScheduleApiV1BenchmarkEvaluationTriggerCmd.Flags().StringArray("resource-i-ds", nil, "")

	GetDescribeCmd.AddCommand(PutScheduleApiV1DescribeTriggerConnectionIdCmd)
	PutScheduleApiV1DescribeTriggerConnectionIdCmd.Flags().String("connection-id", "", "")
}
