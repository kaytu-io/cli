package gen

import (
	"github.com/spf13/cobra"
)

var GetStackCmd = &cobra.Command{
	Use: "stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateStackCmd = &cobra.Command{
	Use: "stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteStackCmd = &cobra.Command{
	Use: "stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateStackCmd = &cobra.Command{
	Use: "stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GetStackCmd.AddCommand(DeleteScheduleApiV1StacksStackIdCmd)
	DeleteScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksFindingsJobIdCmd)
	GetScheduleApiV1StacksFindingsJobIdCmd.Flags().String("job-id", "", "")
	GetScheduleApiV1StacksFindingsJobIdCmd.Flags().Int64("no", 0, "")
	GetScheduleApiV1StacksFindingsJobIdCmd.Flags().Int64("size", 0, "")

	GetScheduleApiV1StacksFindingsJobIdCmd.Flags().String("direction", "", "")
	GetScheduleApiV1StacksFindingsJobIdCmd.Flags().String("field", "", "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksCmd)
	GetScheduleApiV1StacksCmd.Flags().StringArray("accoun-ids", nil, "")
	GetScheduleApiV1StacksCmd.Flags().StringArray("tag", nil, "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksResourceResourceIdCmd)
	GetScheduleApiV1StacksResourceResourceIdCmd.Flags().String("resource-id", "", "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdInsightCmd)
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("end-time", 0, "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("insight-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("stack-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("start-time", 0, "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdCmd)
	GetScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(PostScheduleApiV1StacksBenchmarkTriggerCmd)
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().StringArray("benchmarks", nil, "")
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(PostScheduleApiV1StacksCreateCmd)
	PostScheduleApiV1StacksCreateCmd.Flags().StringArray("resources", nil, "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("tags", "", "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("terrafrom-file", "", "")
}
