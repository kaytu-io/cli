package gen

import (
	"github.com/spf13/cobra"
)

var GetInventoryCmd = &cobra.Command{
	Use: "inventory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateInventoryCmd = &cobra.Command{
	Use: "inventory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteInventoryCmd = &cobra.Command{
	Use: "inventory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateInventoryCmd = &cobra.Command{
	Use: "inventory",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesCompositionKeyCmd)

		GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesFiltersCmd)

		GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesCompositionKeyCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagKeyCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTrendCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesMetricCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagKeyCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesCountCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesCostTrendCmd)

		GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesAzureCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesRegionsCmd)

		GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesTopRegionsCmd)

		GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesAwsCmd)
}