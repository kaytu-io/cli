package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insight"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2InsightsCmd = &cobra.Command{
	Use: "list-insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights] : %v", err)
		}

		req := insight.NewGetInventoryAPIV2InsightsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetInsightID(flags.ReadStringArrayFlag(cmd, "InsightID"))
		req.SetTime(flags.ReadInt64OptionalFlag(cmd, "Time"))

		resp, err := client.Insight.GetInventoryAPIV2Insights(req, auth)
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
	Use: "get-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id] : %v", err)
		}

		req := insight.NewGetInventoryAPIV2InsightsInsightIDParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetTime(flags.ReadInt64OptionalFlag(cmd, "Time"))

		resp, err := client.Insight.GetInventoryAPIV2InsightsInsightID(req, auth)
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
	Use: "get-insight-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id_trend] : %v", err)
		}

		req := insight.NewGetInventoryAPIV2InsightsInsightIDTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insight.GetInventoryAPIV2InsightsInsightIDTrend(req, auth)
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

var GetInventoryApiV2InsightsJobJobIdCmd = &cobra.Command{
	Use: "get-insight-job",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_job_job_id] : %v", err)
		}

		req := insight.NewGetInventoryAPIV2InsightsJobJobIDParams()

		req.SetJobID(flags.ReadStringFlag(cmd, "JobID"))

		resp, err := client.Insight.GetInventoryAPIV2InsightsJobJobID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_job_job_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_job_job_id] : %v", err)
		}

		return nil
	},
}
