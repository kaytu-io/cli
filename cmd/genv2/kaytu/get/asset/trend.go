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

var TrendCmd = &cobra.Command{
	Use: "trend",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := analytics.NewGetInventoryAPIV2AnalyticsTrendParams()

		req.SetConnectionGroup(flags.ReadStringArrayFlag(cmd, "ConnectionGroup"))

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))

		req.SetGranularity(flags.ReadStringOptionalFlag(cmd, "Granularity"))

		req.SetIds(flags.ReadStringArrayFlag(cmd, "Ids"))

		req.SetMetricType(flags.ReadStringOptionalFlag(cmd, "MetricType"))

		req.SetResourceCollection(flags.ReadStringArrayFlag(cmd, "ResourceCollection"))

		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Analytics.GetInventoryAPIV2AnalyticsTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsTrend] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetInventoryAPIV2AnalyticsTrend", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsTrend] : %v", err)
		}

		return nil

	},
}

func init() {

	TrendCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")

	TrendCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")

	TrendCmd.Flags().StringArray("connector", nil, "Connector type to filter by")

	TrendCmd.Flags().Int64("end-time", 0, "timestamp for end in epoch seconds")

	TrendCmd.Flags().String("granularity", "", "Granularity of the table, default is daily")

	TrendCmd.Flags().StringArray("ids", nil, "Metric IDs to filter by")

	TrendCmd.Flags().String("metric-type", "", "Metric type, default: assets")

	TrendCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")

	TrendCmd.Flags().Int64("start-time", 0, "timestamp for start in epoch seconds")

	TrendCmd.Flags().StringArray("tag", nil, "Key-Value tags in key=value format to filter by")

}
