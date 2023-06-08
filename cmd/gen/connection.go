package gen

import (
	"github.com/spf13/cobra"
)

var GetConnectionCmd = &cobra.Command{
	Use: "connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateConnectionCmd = &cobra.Command{
	Use: "connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteConnectionCmd = &cobra.Command{
	Use: "connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateConnectionCmd = &cobra.Command{
	Use: "connection",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetConnectionCmd.AddCommand(GetInventoryApiV2ConnectionsDataConnectionIdCmd)

		GetConnectionCmd.AddCommand(GetInventoryApiV2ConnectionsDataCmd)
}