package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insight"
	"github.com/spf13/cobra"
)
var GetInventoryApiV2InsightsCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights] : %v", err)
		}

		resp, err := client.Insight.GetInventoryAPIV2Insights(insight.NewGetInventoryAPIV2InsightsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2InsightsInsightIdCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_insights_insight_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id] : %v", err)
		}

		resp, err := client.Insight.GetInventoryAPIV2InsightsInsightID(insight.NewGetInventoryAPIV2InsightsInsightIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2InsightsInsightIdTrendCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_insights_insight_id_trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id_trend] : %v", err)
		}

		resp, err := client.Insight.GetInventoryAPIV2InsightsInsightIDTrend(insight.NewGetInventoryAPIV2InsightsInsightIDTrendParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id_trend] : %v", err)
		}

		return nil
	},
}

}
