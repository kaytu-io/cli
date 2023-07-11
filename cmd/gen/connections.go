// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"github.com/spf13/cobra"
)

var ConnectionsCmd = &cobra.Command{
	Use: "connections",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func init() {

	ConnectionsCmd.AddCommand(GetOnboardApiV1ConnectionsSummaryConnectionIdCmd)
	GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().String("connection-id", "", "ConnectionID")
	GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.MarkFlagRequired("connection-id")
	GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().Int64("end-time", 0, "end time in unix seconds")
	GetOnboardApiV1ConnectionsSummaryConnectionIdCmd.Flags().Int64("start-time", 0, "start time in unix seconds")

	ConnectionsCmd.AddCommand(GetOnboardApiV1ConnectionsSummaryCmd)
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().StringArray("connector", nil, "Connector")
	GetOnboardApiV1ConnectionsSummaryCmd.MarkFlagRequired("connector")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("end-time", 0, "end time in unix seconds")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("lifecycle-state", "", "lifecycle state filter")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("page-number", 0, "page number - default is 1")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("page-size", 0, "page size - default is 20")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().String("sort-by", "", "column to sort by - default is cost")
	GetOnboardApiV1ConnectionsSummaryCmd.Flags().Int64("start-time", 0, "start time in unix seconds")

}
