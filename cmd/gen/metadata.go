// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var MetadataCmd = &cobra.Command{
	Use: "metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	MetadataCmd.AddCommand(GetMetadataApiV1MetadataKeyCmd)
	GetMetadataApiV1MetadataKeyCmd.Flags().String("key", "", "Key")
	GetMetadataApiV1MetadataKeyCmd.MarkFlagRequired("key")

	MetadataCmd.AddCommand(PostMetadataApiV1MetadataCmd)
	PostMetadataApiV1MetadataCmd.Flags().String("key", "", "")
	PostMetadataApiV1MetadataCmd.MarkFlagRequired("key")
	PostMetadataApiV1MetadataCmd.Flags().String("value", "", "")
	PostMetadataApiV1MetadataCmd.MarkFlagRequired("value")

}
