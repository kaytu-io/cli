// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var GetSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateSmartQueryCmd = &cobra.Command{
	Use: "smart_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	GetSmartQueryCmd.AddCommand(GetInventoryApiV1QueryCmd)
	GetInventoryApiV1QueryCmd.Flags().StringArray("labels", nil, "")
	GetInventoryApiV1QueryCmd.Flags().String("provider-filter", "", "")
	GetInventoryApiV1QueryCmd.Flags().String("title-filter", "", "")
	GetInventoryApiV1QueryCmd.MarkFlagRequired("title-filter")

	GetSmartQueryCmd.AddCommand(PostInventoryApiV1QueryQueryIdCmd)
	PostInventoryApiV1QueryQueryIdCmd.Flags().String("accept", "", "")
	PostInventoryApiV1QueryQueryIdCmd.MarkFlagRequired("accept")
	PostInventoryApiV1QueryQueryIdCmd.Flags().String("query-id", "", "")
	PostInventoryApiV1QueryQueryIdCmd.MarkFlagRequired("query-id")
	PostInventoryApiV1QueryQueryIdCmd.Flags().Int64("no", 0, "")
	PostInventoryApiV1QueryQueryIdCmd.MarkFlagRequired("no")
	PostInventoryApiV1QueryQueryIdCmd.Flags().Int64("size", 0, "")
	PostInventoryApiV1QueryQueryIdCmd.MarkFlagRequired("size")

	PostInventoryApiV1QueryQueryIdCmd.Flags().String("direction", "", "")
	PostInventoryApiV1QueryQueryIdCmd.Flags().String("field", "", "")
	PostInventoryApiV1QueryQueryIdCmd.MarkFlagRequired("field")

	GetSmartQueryCmd.AddCommand(GetInventoryApiV1QueryCountCmd)
	GetInventoryApiV1QueryCountCmd.Flags().StringArray("labels", nil, "")
	GetInventoryApiV1QueryCountCmd.Flags().String("provider-filter", "", "")
	GetInventoryApiV1QueryCountCmd.Flags().String("title-filter", "", "")
	GetInventoryApiV1QueryCountCmd.MarkFlagRequired("title-filter")

}
