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
		GetScheduleCmd.AddCommand(GetScheduleApiV1DescribeResourceJobsPendingCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1DescribeSourceJobsPendingCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1ResourceTypeProviderCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdJobsComplianceCmd)

		GetScheduleCmd.AddCommand(PostScheduleApiV1SourcesSourceIdJobsDescribeRefreshCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1ComplianceReportLastCompletedCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdJobsDescribeCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1SourcesSourceIdCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1SummarizeJobsPendingCmd)

		GetScheduleCmd.AddCommand(PostScheduleApiV1SourcesSourceIdJobsComplianceRefreshCmd)

		GetScheduleCmd.AddCommand(GetScheduleApiV1InsightJobsPendingCmd)
}