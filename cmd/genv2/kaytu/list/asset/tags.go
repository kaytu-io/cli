// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package asset

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/analytics"
	"github.com/spf13/cobra"
)

var TagsCmd = &cobra.Command{
	Use: "tags",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := analytics.NewGetInventoryAPIV2AnalyticsTagParams()

		req.SetConnectionGroup(flags.ReadStringArrayFlag(cmd, "ConnectionGroup"))

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))

		req.SetMetricType(flags.ReadStringOptionalFlag(cmd, "MetricType"))

		req.SetMinCount(flags.ReadInt64OptionalFlag(cmd, "MinCount"))

		req.SetResourceCollection(flags.ReadStringArrayFlag(cmd, "ResourceCollection"))

		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Analytics.GetInventoryAPIV2AnalyticsTag(req, auth)
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsTag] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetInventoryAPIV2AnalyticsTag", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsTag] : %v", err)
		}

		return nil

	},
}

func init() {

	TagsCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")

	TagsCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")

	TagsCmd.Flags().StringArray("connector", nil, "Connector type to filter by")

	TagsCmd.Flags().Int64("end-time", 0, "End time in unix timestamp format, default now")

	TagsCmd.Flags().String("metric-type", "", "Metric type, default: assets")

	TagsCmd.Flags().Int64("min-count", 0, "Minimum number of resources/spend with this tag value, default 1")

	TagsCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")

	TagsCmd.Flags().Int64("start-time", 0, "Start time in unix timestamp format, default now - 1 month")

}
