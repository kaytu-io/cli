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
	GetInventoryCmd.AddCommand(GetInventoryApiV2CostCompositionCmd)
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2CostCompositionCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2CostCompositionCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2CostCompositionCmd.Flags().Int64("top", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricCmd)
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().Int64("min-count", 0, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("servicename", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("tag", nil, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTrendCmd)
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("datapoint-count", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("servicename", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("tag", nil, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagCmd)
	GetInventoryApiV2ResourcesTagCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesTagCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesTagCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ResourcesTagCmd.Flags().Int64("min-count", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesMetricCmd)
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("tag", nil, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesMetricServiceNameCmd)
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("service-name", "", "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2CostMetricCmd)
	GetInventoryApiV2CostMetricCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2CostMetricCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2CostMetricCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2CostMetricCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2CostMetricCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2CostMetricCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2CostMetricCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2CostTrendCmd)
	GetInventoryApiV2CostTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2CostTrendCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2CostTrendCmd.Flags().String("datapoint-count", "", "")
	GetInventoryApiV2CostTrendCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2CostTrendCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesCompositionKeyCmd)
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().String("key", "", "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().Int64("top", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricResourceTypeCmd)
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("resource-type", "", "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagKeyCmd)
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().String("key", "", "")
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().Int64("min-count", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagKeyCmd)
	GetInventoryApiV2ServicesTagKeyCmd.Flags().String("key", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagCmd)
}
