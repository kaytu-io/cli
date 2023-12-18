// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package insight

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
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

		req := insights.NewGetComplianceAPIV1InsightInsightIDTrendParams()

		req.SetConnectionGroup(flags.ReadStringArrayFlag(cmd, "ConnectionGroup"))

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))

		req.SetDatapointCount(flags.ReadInt64OptionalFlag(cmd, "DatapointCount"))

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))

		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))

		req.SetResourceCollection(flags.ReadStringArrayFlag(cmd, "ResourceCollection"))

		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightIDTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[GetComplianceAPIV1InsightInsightIDTrend] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetComplianceAPIV1InsightInsightIDTrend", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetComplianceAPIV1InsightInsightIDTrend] : %v", err)
		}

		return nil

	},
}

func init() {

	TrendCmd.Flags().StringArray("connection-group", nil, "filter the result by connection group")

	TrendCmd.Flags().StringArray("connection-id", nil, "filter the result by source id")

	TrendCmd.Flags().Int64("datapoint-count", 0, "number of datapoints to return")

	TrendCmd.Flags().Int64("end-time", 0, "unix seconds for the end time of the trend")

	TrendCmd.Flags().String("insight-id", "", "Insight ID")

	TrendCmd.MarkFlagRequired("insight-id")

	TrendCmd.Flags().StringArray("resource-collection", nil, "Resource collection IDs to filter by")

	TrendCmd.Flags().Int64("start-time", 0, "unix seconds for the start time of the trend")

}