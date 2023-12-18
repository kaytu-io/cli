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

var MetricCmd = &cobra.Command{
	Use: "metric",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := analytics.NewGetInventoryAPIV2AnalyticsMetricsMetricIDParams()

		req.SetMetricID(flags.ReadStringFlag(cmd, "MetricID"))

		resp, err := client.Analytics.GetInventoryAPIV2AnalyticsMetricsMetricID(req, auth)
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsMetricsMetricID] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetInventoryAPIV2AnalyticsMetricsMetricID", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsMetricsMetricID] : %v", err)
		}

		return nil

	},
}

func init() {

	MetricCmd.Flags().String("metric-id", "", "MetricID")

	MetricCmd.MarkFlagRequired("metric-id")

}