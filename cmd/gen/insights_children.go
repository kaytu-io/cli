package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2InsightsInsightIdTrendCmd = &cobra.Command{
	Use: "old-get-insight-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id_trend] : %v", err)
		}

		req := insights.NewGetInventoryAPIV2InsightsInsightIDTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetInventoryAPIV2InsightsInsightIDTrend(req, auth)
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

var GetInventoryApiV2InsightsCmd = &cobra.Command{
	Use: "old-list-insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights] : %v", err)
		}

		req := insights.NewGetInventoryAPIV2InsightsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetInsightID(flags.ReadStringArrayFlag(cmd, "InsightID"))
		req.SetTime(flags.ReadInt64OptionalFlag(cmd, "Time"))

		resp, err := client.Insights.GetInventoryAPIV2Insights(req, auth)
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

var GetComplianceApiV1InsightGroupCmd = &cobra.Command{
	Use: "list-insight-groups",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightGroupParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Insights.GetComplianceAPIV1InsightGroup(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightCmd = &cobra.Command{
	Use: "list-insights",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Insights.GetComplianceAPIV1Insight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataInsightCmd = &cobra.Command{
	Use: "list-insights-metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataInsightParams()

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		resp, err := client.Insights.GetComplianceAPIV1MetadataInsight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataTagInsightCmd = &cobra.Command{
	Use: "list-insights-metadata-tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataTagInsightParams()

		resp, err := client.Insights.GetComplianceAPIV1MetadataTagInsight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2InsightsJobJobIdCmd = &cobra.Command{
	Use: "old-get-insight-job",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_job_job_id] : %v", err)
		}

		req := insights.NewGetInventoryAPIV2InsightsJobJobIDParams()

		req.SetJobID(flags.ReadStringFlag(cmd, "JobID"))

		resp, err := client.Insights.GetInventoryAPIV2InsightsJobJobID(req, auth)
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

var GetComplianceApiV1InsightGroupInsightGroupIdCmd = &cobra.Command{
	Use: "get-insight-group",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightGroupInsightGroupIDParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightGroupID(flags.ReadStringFlag(cmd, "InsightGroupID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightGroupInsightGroupID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightGroupInsightGroupIdTrendCmd = &cobra.Command{
	Use: "get-insight-group-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id_trend] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightGroupInsightGroupIDTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetDatapointCount(flags.ReadInt64OptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightGroupID(flags.ReadStringFlag(cmd, "InsightGroupID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightGroupInsightGroupIDTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group_insight_group_id_trend] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataInsightInsightIdCmd = &cobra.Command{
	Use: "get-insight-metadata",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataInsightInsightIDParams()

		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))

		resp, err := client.Insights.GetComplianceAPIV1MetadataInsightInsightID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataTagInsightKeyCmd = &cobra.Command{
	Use: "list-insights-metadata-tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight_key] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataTagInsightKeyParams()

		req.SetKey(flags.ReadStringFlag(cmd, "Key"))

		resp, err := client.Insights.GetComplianceAPIV1MetadataTagInsightKey(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight_key] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightInsightIdTrendCmd = &cobra.Command{
	Use: "get-insight-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightInsightIDTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetDatapointCount(flags.ReadInt64OptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightIDTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightInsightIdCmd = &cobra.Command{
	Use: "get-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightInsightIDParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2InsightsInsightIdCmd = &cobra.Command{
	Use: "old-get-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_insights_insight_id] : %v", err)
		}

		req := insights.NewGetInventoryAPIV2InsightsInsightIDParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetTime(flags.ReadInt64OptionalFlag(cmd, "Time"))

		resp, err := client.Insights.GetInventoryAPIV2InsightsInsightID(req, auth)
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
