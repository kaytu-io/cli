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
	DeleteStackCmd.AddCommand(DeleteScheduleApiV1StacksStackIdCmd)
	DeleteScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksResourceCmd)
	GetScheduleApiV1StacksResourceCmd.Flags().String("resource-id", "", "")

	GetStackCmd.AddCommand(PostScheduleApiV1StacksBenchmarkTriggerCmd)
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().StringArray("benchmarks", nil, "")
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().String("stack-id", "", "")

	CreateStackCmd.AddCommand(PostScheduleApiV1StacksCreateCmd)
	PostScheduleApiV1StacksCreateCmd.Flags().String("config", "", "")
	PostScheduleApiV1StacksCreateCmd.Flags().StringArray("resources", nil, "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("tag", "", "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("terrafrom-file", "", "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksCmd)
	GetScheduleApiV1StacksCmd.Flags().StringArray("account-ids", nil, "")
	GetScheduleApiV1StacksCmd.Flags().StringArray("tag", nil, "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdInsightCmd)
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("end-time", 0, "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("insight-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("stack-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("start-time", 0, "")

	GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdCmd)
	GetScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(PostScheduleApiV1StacksDescriberTriggerCmd)
	PostScheduleApiV1StacksDescriberTriggerCmd.Flags().String("config", "", "")
	PostScheduleApiV1StacksDescriberTriggerCmd.Flags().String("stack-id", "", "")

	GetStackCmd.AddCommand(PostScheduleApiV1StacksStackIdFindingsCmd)
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().StringArray("benchmark-ids", nil, "")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().Int64("no", 0, "")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().Int64("size", 0, "")

	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("direction", "", "")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("field", "", "")

	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("stack-id", "", "")
}
