package gen

import (
	"github.com/spf13/cobra"
)

var GetResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetResourceCmd.AddCommand(PostInventoryApiV1ResourceCmd)
}