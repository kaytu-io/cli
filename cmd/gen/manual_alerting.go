package gen

import "github.com/kaytu-io/cli-program/cmd/gen/manual_commands"

func init() {
	AlertingCmd.AddCommand(manual_commands.PUTAlertingApiV1RuleUpdateCmd)
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("rule-id", "", "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.MarkFlagRequired("rule-id")

	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().Int64("insight-id", 0, "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("benchmark-id", "", "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.MarkFlagsMutuallyExclusive("insight-id", "benchmark-id")

	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("connection-id", "", "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("connection-group", "", "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.MarkFlagsMutuallyExclusive("connection-id", "connection-group")

	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("connector", "", "")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("operator", "", "examples: \nvalue > 100 and value < 200 \nvalue > 100 and ( value < 1000 or value > 1200) \n( value > 10 and value < 100 ) or ( value > 130 and value < 200)")
	manual_commands.PUTAlertingApiV1RuleUpdateCmd.Flags().String("action-id", "", "")

	AlertingCmd.AddCommand(manual_commands.POSTAlertingApiV1RuleCreateCmd)

	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("insight-id", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("benchmark-id", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.MarkFlagsMutuallyExclusive("insight-id", "benchmark-id")

	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("connection-id", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("connection-group", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.MarkFlagsMutuallyExclusive("connection-id", "connection-group")

	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("connector", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("operator", "", "examples: \nvalue > 100 and value < 200 \nvalue > 100 and ( value < 1000 or value > 1200) \n( value > 10 and value < 100 ) or ( value > 130 and value < 200)")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.MarkFlagRequired("operator")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.Flags().String("action-id", "", "")
	manual_commands.POSTAlertingApiV1RuleCreateCmd.MarkFlagRequired("action-id")
}
