package gen

import (
	"github.com/spf13/cobra"
)

var GetLocationCmd = &cobra.Command{
	Use: "location",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateLocationCmd = &cobra.Command{
	Use: "location",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteLocationCmd = &cobra.Command{
	Use: "location",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateLocationCmd = &cobra.Command{
	Use: "location",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetLocationCmd.AddCommand(GetInventoryApiV1LocationsConnectorCmd)
}