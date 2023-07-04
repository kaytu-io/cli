// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var StackCmd = &cobra.Command{
	Use: "stack",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	StackCmd.AddCommand(PostScheduleApiV1StacksCreateCmd)
	PostScheduleApiV1StacksCreateCmd.Flags().String("config", "", "")
	PostScheduleApiV1StacksCreateCmd.Flags().StringArray("resources", nil, "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("tag", "", "")
	PostScheduleApiV1StacksCreateCmd.Flags().String("terrafrom-file", "", "")

	StackCmd.AddCommand(DeleteScheduleApiV1StacksStackIdCmd)
	DeleteScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")
	DeleteScheduleApiV1StacksStackIdCmd.MarkFlagRequired("stack-id")

	StackCmd.AddCommand(GetScheduleApiV1StacksStackIdCmd)
	GetScheduleApiV1StacksStackIdCmd.Flags().String("stack-id", "", "")
	GetScheduleApiV1StacksStackIdCmd.MarkFlagRequired("stack-id")

	StackCmd.AddCommand(GetScheduleApiV1StacksCmd)
	GetScheduleApiV1StacksCmd.Flags().StringArray("account-ids", nil, "")
	GetScheduleApiV1StacksCmd.Flags().StringArray("tag", nil, "")

	StackCmd.AddCommand(GetScheduleApiV1StacksResourceCmd)
	GetScheduleApiV1StacksResourceCmd.Flags().String("resource-id", "", "")
	GetScheduleApiV1StacksResourceCmd.MarkFlagRequired("resource-id")

	StackCmd.AddCommand(PostScheduleApiV1StacksStackIdFindingsCmd)
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().StringArray("benchmark-ids", nil, "")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().Int64("no", 0, "")
	PostScheduleApiV1StacksStackIdFindingsCmd.MarkFlagRequired("no")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().Int64("size", 0, "")
	PostScheduleApiV1StacksStackIdFindingsCmd.MarkFlagRequired("size")

	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("direction", "", "")
	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("field", "", "")

	PostScheduleApiV1StacksStackIdFindingsCmd.Flags().String("stack-id", "", "")
	PostScheduleApiV1StacksStackIdFindingsCmd.MarkFlagRequired("stack-id")

	StackCmd.AddCommand(GetScheduleApiV1StacksStackIdInsightCmd)
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("end-time", 0, "")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("insight-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.MarkFlagRequired("insight-id")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().String("stack-id", "", "")
	GetScheduleApiV1StacksStackIdInsightCmd.MarkFlagRequired("stack-id")
	GetScheduleApiV1StacksStackIdInsightCmd.Flags().Int64("start-time", 0, "")

	StackCmd.AddCommand(PostScheduleApiV1StacksBenchmarkTriggerCmd)
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().StringArray("benchmarks", nil, "")
	PostScheduleApiV1StacksBenchmarkTriggerCmd.Flags().String("stack-id", "", "")

	StackCmd.AddCommand(PostScheduleApiV1StacksDescriberTriggerCmd)
	PostScheduleApiV1StacksDescriberTriggerCmd.Flags().String("config", "", "")
	PostScheduleApiV1StacksDescriberTriggerCmd.Flags().String("stack-id", "", "")
	PostScheduleApiV1StacksDescriberTriggerCmd.MarkFlagRequired("stack-id")

}
