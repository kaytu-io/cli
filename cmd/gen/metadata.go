package gen

import (
	"github.com/spf13/cobra"
)

var GetMetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateMetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteMetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateMetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataResourcetypeCmd)
GetInventoryApiV2MetadataResourcetypeCmd.Flags().StringArray("connector", nil, "")
GetInventoryApiV2MetadataResourcetypeCmd.Flags().Int64("page-number", 0, "")
GetInventoryApiV2MetadataResourcetypeCmd.Flags().Int64("page-size", 0, "")
GetInventoryApiV2MetadataResourcetypeCmd.Flags().StringArray("service", nil, "")
GetInventoryApiV2MetadataResourcetypeCmd.Flags().StringArray("tag", nil, "")

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataResourcetypeResourceTypeCmd)
GetInventoryApiV2MetadataResourcetypeResourceTypeCmd.Flags().String("resource-type", "", "")

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataServicesCmd)
GetInventoryApiV2MetadataServicesCmd.Flags().StringArray("connector", nil, "")
GetInventoryApiV2MetadataServicesCmd.Flags().Bool("cost-support", false, "")
GetInventoryApiV2MetadataServicesCmd.Flags().Int64("page-number", 0, "")
GetInventoryApiV2MetadataServicesCmd.Flags().Int64("page-size", 0, "")
GetInventoryApiV2MetadataServicesCmd.Flags().StringArray("tag", nil, "")

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataServicesServiceNameCmd)
GetInventoryApiV2MetadataServicesServiceNameCmd.Flags().String("service-name", "", "")

		GetMetadataCmd.AddCommand(GetMetadataApiV1MetadataKeyCmd)
GetMetadataApiV1MetadataKeyCmd.Flags().String("key", "", "")

		GetMetadataCmd.AddCommand(PostMetadataApiV1MetadataCmd)
}