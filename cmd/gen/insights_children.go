package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/spf13/cobra"
)
var GetComplianceApiV1MetadataTagInsightCmd = &cobra.Command{
	Use: "get_compliance_api_v_1_metadata_tag_insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1MetadataTagInsight(insights.NewGetComplianceAPIV1MetadataTagInsightParams(), auth)
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
var GetComplianceApiV1InsightInsightIdCmd = &cobra.Command{
	Use: "get_compliance_api_v_1_insight_insight_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightID(insights.NewGetComplianceAPIV1InsightInsightIDParams(), auth)
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
var GetComplianceApiV1InsightInsightIdTrendCmd = &cobra.Command{
	Use: "get_compliance_api_v_1_insight_insight_id_trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightIDTrend(insights.NewGetComplianceAPIV1InsightInsightIDTrendParams(), auth)
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
var GetComplianceApiV1InsightCmd = &cobra.Command{
	Use: "get_compliance_api_v_1_insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1Insight(insights.NewGetComplianceAPIV1InsightParams(), auth)
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
	Use: "get_compliance_api_v_1_metadata_insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1MetadataInsight(insights.NewGetComplianceAPIV1MetadataInsightParams(), auth)
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
var GetComplianceApiV1MetadataTagInsightKeyCmd = &cobra.Command{
	Use: "get_compliance_api_v_1_metadata_tag_insight_key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight_key] : %v", err)
		}

		resp, err := client.Insights.GetComplianceAPIV1MetadataTagInsightKey(insights.NewGetComplianceAPIV1MetadataTagInsightKeyParams(), auth)
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


