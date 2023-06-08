package gen

import (
	"github.com/spf13/cobra"
)

var GetSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetSmartQueryCmd.AddCommand(GetInventoryApiV1QueryCmd)

		GetSmartQueryCmd.AddCommand(PostInventoryApiV1QueryQueryIdCmd)

		GetSmartQueryCmd.AddCommand(GetInventoryApiV1QueryCountCmd)
}