package gen

import "github.com/kaytu-io/cli-program/cmd/gen/manual_commands"

func init() {
	SpendCmd.AddCommand(manual_commands.GetCostsCmd)
	manual_commands.GetCostsCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	manual_commands.GetCostsCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	manual_commands.GetCostsCmd.Flags().StringArray("connector", nil, "Connector type to filter by")
	manual_commands.GetCostsCmd.Flags().String("end-time", "", "timestamp for end in epoch seconds")
	manual_commands.GetCostsCmd.Flags().Int64("page-number", 0, "page number - default is 1")
	manual_commands.GetCostsCmd.Flags().Int64("page-size", 0, "page size - default is 20")
	manual_commands.GetCostsCmd.Flags().String("sort-by", "", "Sort by field - default is cost")
	manual_commands.GetCostsCmd.Flags().String("start-time", "", "timestamp for start in epoch seconds")
	manual_commands.GetCostsCmd.Flags().String("view", "connection", "Whether to view by connection or service")
}
