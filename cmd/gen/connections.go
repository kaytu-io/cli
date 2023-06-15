package gen

import (
	"github.com/spf13/cobra"
)

var GetConnectionsCmd = &cobra.Command{
	Use: "connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateConnectionsCmd = &cobra.Command{
	Use: "connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteConnectionsCmd = &cobra.Command{
	Use: "connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateConnectionsCmd = &cobra.Command{
	Use: "connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetConnectionsCmd.AddCommand(GetOnboardApiV1ConnectionsSummaryConnectionIdCmd)
GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().String("connection-id", "", "")
GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().Int64("end-time", 0, "")
GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().Int64("start-time", 0, "")

		GetConnectionsCmd.AddCommand(GetOnboardApiV1ConnectionsSummaryCmd)
GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("connection-id", "", "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("connector", "", "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("end-time", 0, "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("health-state", "", "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("lifecycle-state", "", "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("page-number", 0, "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("page-size", 0, "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("sort-by", "", "")
GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("start-time", 0, "")
}