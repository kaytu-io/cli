// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var ResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	ResourceCmd.AddCommand(GetInventoryApiV2ResourcesMetricResourceTypeCmd)
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("connection-group", "", "Connection group to filter by - mutually exclusive with connectionId")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().Int64("end-time", 0, "timestamp for resource count in epoch seconds")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("resource-type", "", "ResourceType")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.MarkFlagRequired("resource-type")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().Int64("start-time", 0, "timestamp for resource count change comparison in epoch seconds")

}
