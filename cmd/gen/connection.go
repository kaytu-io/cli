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
	GetConnectionCmd.AddCommand(GetInventoryApiV2ConnectionsDataCmd)
	GetInventoryApiV2ConnectionsDataCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ConnectionsDataCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ConnectionsDataCmd.Flags().Int64("start-time", 0, "")

	GetConnectionCmd.AddCommand(GetInventoryApiV2ConnectionsDataConnectionIdCmd)
	GetInventoryApiV2ConnectionsDataConnectionIdCmd.Flags().String("connection-id", "", "")
	GetInventoryApiV2ConnectionsDataConnectionIdCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ConnectionsDataConnectionIdCmd.Flags().Int64("start-time", 0, "")
}
