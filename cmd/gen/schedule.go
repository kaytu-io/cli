package gen

import (
	"github.com/spf13/cobra"
)

var GetScheduleCmd = &cobra.Command{
	Use: "schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateScheduleCmd = &cobra.Command{
	Use: "schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteScheduleCmd = &cobra.Command{
	Use: "schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateScheduleCmd = &cobra.Command{
	Use: "schedule",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GetScheduleCmd.AddCommand(GetScheduleApiV1ResourceTypeProviderCmd)
	GetScheduleApiV1ResourceTypeProviderCmd.Flags().String("provider", "", "")

	GetScheduleCmd.AddCommand(PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd)
	PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd.Flags().String("source-id", "", "")

	GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesCmd)

	GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdJobsComplianceCmd)
	GetScheduleApiV1SourcesSourceIdJobsComplianceCmd.Flags().Int64("from", 0, "")
	GetScheduleApiV1SourcesSourceIdJobsComplianceCmd.Flags().String("source-id", "", "")
	GetScheduleApiV1SourcesSourceIdJobsComplianceCmd.Flags().Int64("to", 0, "")

	GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdJobsDescribeCmd)
	GetScheduleApiV1SourcesSourceIdJobsDescribeCmd.Flags().String("source-id", "", "")

	GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdCmd)
	GetScheduleApiV1SourcesSourceIdCmd.Flags().String("source-id", "", "")

	GetScheduleCmd.AddCommand(GetScheduleApiV1ComplianceReportLastCompletedCmd)

	GetScheduleCmd.AddCommand(GetScheduleApiV1DescribeResourceJobsPendingCmd)

	GetScheduleCmd.AddCommand(GetScheduleApiV1DescribeSourceJobsPendingCmd)

	GetScheduleCmd.AddCommand(GetScheduleApiV1InsightJobsPendingCmd)

	GetScheduleCmd.AddCommand(GetScheduleApiV1SummarizeJobsPendingCmd)

	GetScheduleCmd.AddCommand(PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd)
	PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd.Flags().String("source-id", "", "")
}
