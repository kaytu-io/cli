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
		GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdInsightCmd)

		GetStackCmd.AddCommand(GetScheduleApiV1StacksStackIdCmd)

		GetStackCmd.AddCommand(PostScheduleApiV1StacksBenchmarkTriggerCmd)

		GetStackCmd.AddCommand(PostScheduleApiV1StacksCreateCmd)

		GetStackCmd.AddCommand(DeleteScheduleApiV1StacksStackIdCmd)

		GetStackCmd.AddCommand(GetScheduleApiV1StacksFindingsJobIdCmd)

		GetStackCmd.AddCommand(GetScheduleApiV1StacksCmd)

		GetStackCmd.AddCommand(GetScheduleApiV1StacksResourceResourceIdCmd)
}