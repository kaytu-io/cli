package gen

import (
	"github.com/spf13/cobra"
)

var GetCostCmd = &cobra.Command{
	Use: "cost",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateCostCmd = &cobra.Command{
	Use: "cost",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteCostCmd = &cobra.Command{
	Use: "cost",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateCostCmd = &cobra.Command{
	Use: "cost",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetCostCmd.AddCommand(GetInventoryApiV1CostTopAccountsCmd)

		GetCostCmd.AddCommand(GetInventoryApiV1CostTopServicesCmd)
}