package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetComplianceApiV1InsightInsightIdCmd = &cobra.Command{
	Use: "insight-insight-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

        req := insights.NewGetComplianceAPIV1InsightInsightIDParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetEndTime(flags.ReadInt64OptionalFlag("EndTime"))
req.SetInsightID(flags.ReadStringFlag("InsightID"))
req.SetStartTime(flags.ReadInt64OptionalFlag("StartTime"))


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

var GetComplianceApiV1InsightInsightIdTrendCmd = &cobra.Command{
	Use: "insight-insight-id-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

        req := insights.NewGetComplianceAPIV1InsightInsightIDTrendParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetDatapointCount(flags.ReadInt64OptionalFlag("DatapointCount"))
req.SetEndTime(flags.ReadInt64OptionalFlag("EndTime"))
req.SetInsightID(flags.ReadStringFlag("InsightID"))
req.SetStartTime(flags.ReadInt64OptionalFlag("StartTime"))


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

var GetComplianceApiV1InsightCmd = &cobra.Command{
	Use: "insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

        req := insights.NewGetComplianceAPIV1InsightParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetEndTime(flags.ReadInt64OptionalFlag("EndTime"))
req.SetStartTime(flags.ReadInt64OptionalFlag("StartTime"))
req.SetTag(flags.ReadStringOptionalFlag("Tag"))


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
	Use: "metadata-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight] : %v", err)
		}

        req := insights.NewGetComplianceAPIV1MetadataInsightParams()

        req.SetConnector(flags.ReadStringArrayFlag("Connector"))


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

var GetComplianceApiV1MetadataTagInsightKeyCmd = &cobra.Command{
	Use: "metadata-tag-insight-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight_key] : %v", err)
		}

        req := insights.NewGetComplianceAPIV1MetadataTagInsightKeyParams()

        req.SetKey(flags.ReadStringFlag("Key"))


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

var GetComplianceApiV1MetadataTagInsightCmd = &cobra.Command{
	Use: "metadata-tag-insight",
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
