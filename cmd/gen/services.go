package gen

import (
	"github.com/spf13/cobra"
)

var GetServicesCmd = &cobra.Command{
	Use: "services",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateServicesCmd = &cobra.Command{
	Use: "services",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteServicesCmd = &cobra.Command{
	Use: "services",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateServicesCmd = &cobra.Command{
	Use: "services",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {
	GetServicesCmd.AddCommand(GetInventoryApiV2ServicesSummaryCmd)
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("connection-id", nil, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().Int64("page-number", 0, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().Int64("page-size", 0, "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().String("sort-by", "", "")
	GetInventoryApiV2ServicesSummaryCmd.Flags().StringArray("tag", nil, "")

	GetServicesCmd.AddCommand(GetInventoryApiV2ServicesSummaryServiceNameCmd)
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().StringArray("connector", nil, "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().StringArray("connector-id", nil, "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("end-time", "", "")
	GetInventoryApiV2ServicesSummaryServiceNameCmd.Flags().String("service-name", "", "")
}
