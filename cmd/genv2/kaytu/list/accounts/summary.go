// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package accounts

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/spf13/cobra"
)

var SummaryCmd = &cobra.Command{
	Use: "summary",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := connections.NewGetOnboardAPIV1ConnectionsSummaryParams()

		req.SetConnectionGroups(flags.ReadStringArrayFlag(cmd, "ConnectionGroups"))

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))

		req.SetFilter(flags.ReadStringOptionalFlag(cmd, "Filter"))

		req.SetHealthState(flags.ReadStringOptionalFlag(cmd, "HealthState"))

		req.SetLifecycleState(flags.ReadStringOptionalFlag(cmd, "LifecycleState"))

		req.SetNeedCost(flags.ReadBooleanOptionalFlag(cmd, "NeedCost"))

		req.SetNeedResourceCount(flags.ReadBooleanOptionalFlag(cmd, "NeedResourceCount"))

		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))

		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))

		req.SetResourceCollection(flags.ReadStringArrayFlag(cmd, "ResourceCollection"))

		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))

		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[GetOnboardAPIV1ConnectionsSummary] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetOnboardAPIV1ConnectionsSummary", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetOnboardAPIV1ConnectionsSummary] : %v", err)
		}

		return nil

	},
}

func init() {

	SummaryCmd.Flags().StringArray("connection-groups", nil, "Connection Groups")

	SummaryCmd.Flags().StringArray("connection-id", nil, "Connection IDs")

	SummaryCmd.Flags().StringArray("connector", nil, "Connector")

	SummaryCmd.Flags().Int64("end-time", 0, "end time in unix seconds")

	SummaryCmd.Flags().String("filter", "", "Filter costs")

	SummaryCmd.Flags().String("health-state", "", "health state filter")

	SummaryCmd.Flags().String("lifecycle-state", "", "lifecycle state filter")

	SummaryCmd.Flags().Bool("need-cost", false, "for quicker inquiry send this parameter as false, default: true")

	SummaryCmd.Flags().Bool("need-resource-count", false, "for quicker inquiry send this parameter as false, default: true")

	SummaryCmd.Flags().Int64("page-number", 0, "page number - default is 1")

	SummaryCmd.Flags().Int64("page-size", 0, "page size - default is 20")

	SummaryCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")

	SummaryCmd.Flags().String("sort-by", "", "column to sort by - default is cost")

	SummaryCmd.Flags().Int64("start-time", 0, "start time in unix seconds")

}
