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
	GetResourceCmd.AddCommand(GetInventoryApiV2ResourcesRegionsCompositionCmd)
	GetInventoryApiV2ResourcesRegionsCompositionCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesRegionsCompositionCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesRegionsCompositionCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ResourcesRegionsCompositionCmd.Flags().Int64("start-time", 0, "")
	GetInventoryApiV2ResourcesRegionsCompositionCmd.Flags().Int64("top", 0, "")

	GetResourceCmd.AddCommand(GetInventoryApiV2ResourcesRegionsSummaryCmd)
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ResourcesRegionsSummaryCmd.Flags().Int64("start-time", 0, "")

	GetResourceCmd.AddCommand(GetInventoryApiV2ResourcesRegionsTrendCmd)
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().Int64("datapoint-count", 0, "")
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().Int64("end-time", 0, "")
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().StringArray("region", nil, "")
	GetInventoryApiV2ResourcesRegionsTrendCmd.Flags().Int64("start-time", 0, "")

	GetResourceCmd.AddCommand(PostInventoryApiV1ResourcesAwsCmd)
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

	GetResourceCmd.AddCommand(PostInventoryApiV1ResourcesCmd)
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

	GetResourceCmd.AddCommand(GetInventoryApiV1ResourcesCountCmd)

	GetResourceCmd.AddCommand(GetInventoryApiV1ResourcesRegionsCmd)
	GetInventoryApiV1ResourcesRegionsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().String("end-time", "", "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV1ResourcesRegionsCmd.Flags().String("start-time", "", "")

	GetResourceCmd.AddCommand(GetInventoryApiV1ResourcesTopRegionsCmd)
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV1ResourcesTopRegionsCmd.Flags().Int64("count", 0, "")

	GetResourceCmd.AddCommand(PostInventoryApiV1ResourceCmd)
	PostInventoryApiV1ResourceCmd.Flags().String("id", "", "")
	PostInventoryApiV1ResourceCmd.Flags().String("resource-type", "", "")

	GetResourceCmd.AddCommand(PostInventoryApiV1ResourcesAzureCmd)
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

	GetResourceCmd.AddCommand(PostInventoryApiV1ResourcesFiltersCmd)
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

}
