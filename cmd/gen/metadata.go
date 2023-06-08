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
		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataServicesServiceNameCmd)

		GetMetadataCmd.AddCommand(GetMetadataApiV1MetadataKeyCmd)

		GetMetadataCmd.AddCommand(PostMetadataApiV1MetadataCmd)

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataResourcetypeCmd)

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataResourcetypeResourceTypeCmd)

		GetMetadataCmd.AddCommand(GetInventoryApiV2MetadataServicesCmd)
}