// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/insights"
	"github.com/spf13/cobra"
)

var GetComplianceApiV1InsightInsightIdCmd = &cobra.Command{
	Use:   "get-insight",
	Short: `Retrieving the specified insight with ID. Provides details of the insight, including results during the specified time period for the specified connection.`,
	Long: `Retrieving the specified insight with ID. Provides details of the insight, including results during the specified time period for the specified connection.
Returns "all:provider" job results if connectionId is not defined.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightInsightIDParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-insight-insight-id", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataInsightInsightIdCmd = &cobra.Command{
	Use:   "get-insight-metadata",
	Short: `Retrieving insight metadata by id`,
	Long:  `Retrieving insight metadata by id`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataInsightInsightIDParams()

		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))

		resp, err := client.Insights.GetComplianceAPIV1MetadataInsightInsightID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-metadata-insight-insight-id", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_insight_insight_id] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightInsightIdTrendCmd = &cobra.Command{
	Use:   "get-insight-trend",
	Short: `Retrieving insight results datapoints for a specified connection during a specified time period.`,
	Long: `Retrieving insight results datapoints for a specified connection during a specified time period.
Returns "all:provider" job results if connectionId is not defined.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightInsightIDTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetDatapointCount(flags.ReadInt64OptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))

		resp, err := client.Insights.GetComplianceAPIV1InsightInsightIDTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-insight-insight-id-trend", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_insight_id_trend] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightGroupCmd = &cobra.Command{
	Use:   "list-insight-groups",
	Short: `Retrieving list of insight groups based on specified filters. The API provides details of insights, including results during the specified time period for the specified connection.`,
	Long: `Retrieving list of insight groups based on specified filters. The API provides details of insights, including results during the specified time period for the specified connection.
Returns "all:provider" job results if connectionId is not defined.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightGroupParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Insights.GetComplianceAPIV1InsightGroup(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-insight-group", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight_group] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1InsightCmd = &cobra.Command{
	Use:   "list-insights",
	Short: `Retrieving list of insights based on specified filters. Provides details of insights, including results during the specified time period for the specified connection.`,
	Long: `Retrieving list of insights based on specified filters. Provides details of insights, including results during the specified time period for the specified connection.
Returns "all:provider" job results if connectionId is not defined.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1InsightParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Insights.GetComplianceAPIV1Insight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-insight", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_insight] : %v", err)
		}

		return nil
	},
}

var GetComplianceApiV1MetadataTagInsightCmd = &cobra.Command{
	Use:   "list-insights-metadata-tags",
	Short: `Retrieving a list of insights tag keys with their possible values.`,
	Long:  `Retrieving a list of insights tag keys with their possible values.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		req := insights.NewGetComplianceAPIV1MetadataTagInsightParams()

		resp, err := client.Insights.GetComplianceAPIV1MetadataTagInsight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-compliance-api-v1-metadata-tag-insight", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_compliance_api_v_1_metadata_tag_insight] : %v", err)
		}

		return nil
	},
}
