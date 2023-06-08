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

		GetConnectionsCmd.AddCommand(GetOnboardApiV1ConnectionsSummaryCmd)
}