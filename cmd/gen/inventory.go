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
	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagCmd)

	GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesAwsCmd)
	PostInventoryApiV1ResourcesAwsCmd.Flags().String("accept", "", "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().String("common", "", "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().StringArray("category", nil, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().StringArray("location", nil, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().StringArray("resource-type", nil, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().StringArray("service", nil, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().StringArray("source-id", nil, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().String("tags", "", "")

	PostInventoryApiV1ResourcesAwsCmd.Flags().Int64("no", 0, "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().Int64("size", 0, "")

	PostInventoryApiV1ResourcesAwsCmd.Flags().String("query", "", "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().String("direction", "", "")
	PostInventoryApiV1ResourcesAwsCmd.Flags().String("field", "", "")

	GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesAzureCmd)
	PostInventoryApiV1ResourcesAzureCmd.Flags().String("accept", "", "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().String("common", "", "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().StringArray("category", nil, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().StringArray("location", nil, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().StringArray("resource-type", nil, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().StringArray("service", nil, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().StringArray("source-id", nil, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().String("tags", "", "")

	PostInventoryApiV1ResourcesAzureCmd.Flags().Int64("no", 0, "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().Int64("size", 0, "")

	PostInventoryApiV1ResourcesAzureCmd.Flags().String("query", "", "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().String("direction", "", "")
	PostInventoryApiV1ResourcesAzureCmd.Flags().String("field", "", "")

	GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesFiltersCmd)
	PostInventoryApiV1ResourcesFiltersCmd.Flags().String("common", "", "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("category", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("connections", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("location", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("provider", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("resource-type", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("service", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().StringArray("tag-keys", nil, "")
	PostInventoryApiV1ResourcesFiltersCmd.Flags().String("tag-values", "", "")

	PostInventoryApiV1ResourcesFiltersCmd.Flags().String("query", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesRegionsCmd)
	GetInventoryApiV1ResourcesRegionsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().String("end-time", "", "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesCompositionKeyCmd)
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().String("key", "", "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().String("time", "", "")
	GetInventoryApiV2ResourcesCompositionKeyCmd.Flags().Int64("top", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricResourceTypeCmd)
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("resource-type", "", "")
	GetInventoryApiV2ResourcesMetricResourceTypeCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesTagKeyCmd)
	GetInventoryApiV2ServicesTagKeyCmd.Flags().String("key", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagKeyCmd)
	GetInventoryApiV2ResourcesTagKeyCmd.Flags().String("key", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTrendCmd)
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("datapoint-count", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("servicename", nil, "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ResourcesTrendCmd.Flags().StringArray("tag", nil, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesMetricServiceNameCmd)
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("service-name", "", "")
	GetInventoryApiV2ServicesMetricServiceNameCmd.Flags().String("start-time", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesTopRegionsCmd)
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().Int64("count", 0, "")

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

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesMetricCmd)
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("servicename", nil, "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ResourcesMetricCmd.Flags().StringArray("tag", nil, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ResourcesTagCmd)

	GetInventoryCmd.AddCommand(PostInventoryApiV1ResourcesCmd)
	PostInventoryApiV1ResourcesCmd.Flags().String("accept", "", "")
	PostInventoryApiV1ResourcesCmd.Flags().String("common", "", "")
	PostInventoryApiV1ResourcesCmd.Flags().StringArray("category", nil, "")
	PostInventoryApiV1ResourcesCmd.Flags().StringArray("location", nil, "")
	PostInventoryApiV1ResourcesCmd.Flags().StringArray("resource-type", nil, "")
	PostInventoryApiV1ResourcesCmd.Flags().StringArray("service", nil, "")
	PostInventoryApiV1ResourcesCmd.Flags().StringArray("source-id", nil, "")
	PostInventoryApiV1ResourcesCmd.Flags().String("tags", "", "")

	PostInventoryApiV1ResourcesCmd.Flags().Int64("no", 0, "")
	PostInventoryApiV1ResourcesCmd.Flags().Int64("size", 0, "")

	PostInventoryApiV1ResourcesCmd.Flags().String("query", "", "")
	PostInventoryApiV1ResourcesCmd.Flags().String("direction", "", "")
	PostInventoryApiV1ResourcesCmd.Flags().String("field", "", "")

	GetInventoryCmd.AddCommand(GetInventoryApiV1ResourcesCountCmd)

	GetInventoryCmd.AddCommand(GetInventoryApiV2CostCompositionCmd)
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2CostCompositionCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2CostCompositionCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2CostCompositionCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2CostCompositionCmd.Flags().Int64("top", 0, "")

	GetInventoryCmd.AddCommand(GetInventoryApiV2ServicesMetricCmd)
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().String("start-time", "", "")
	GetInventoryApiV2ServicesMetricCmd.Flags().StringArray("tag", nil, "")
}
