package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/spf13/cobra"
)
var PostScheduleApiV1DescribeResourceCmd = &cobra.Command{
	Use: "post_schedule_api_v_1_describe_resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_describe_resource] : %v", err)
		}

		resp, err := client.Describe.PostScheduleAPIV1DescribeResource(describe.NewPostScheduleAPIV1DescribeResourceParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_describe_resource] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_describe_resource] : %v", err)
		}

		return nil
	},
}
var PutScheduleApiV1ComplianceTriggerCmd = &cobra.Command{
	Use: "put_schedule_api_v_1_compliance_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_compliance_trigger] : %v", err)
		}

		resp, err := client.Describe.PutScheduleAPIV1ComplianceTrigger(describe.NewPutScheduleAPIV1ComplianceTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_compliance_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_compliance_trigger] : %v", err)
		}

		return nil
	},
}
var PutScheduleApiV1DescribeTriggerConnectionIdCmd = &cobra.Command{
	Use: "put_schedule_api_v_1_describe_trigger_connection_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		resp, err := client.Describe.PutScheduleAPIV1DescribeTriggerConnectionID(describe.NewPutScheduleAPIV1DescribeTriggerConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0ComplianceSummarizerTriggerCmd = &cobra.Command{
	Use: "get_schedule_api_v_0_compliance_summarizer_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		resp, err := client.Describe.GetScheduleAPIV0ComplianceSummarizerTrigger(describe.NewGetScheduleAPIV0ComplianceSummarizerTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0ComplianceTriggerCmd = &cobra.Command{
	Use: "get_schedule_api_v_0_compliance_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		resp, err := client.Describe.GetScheduleAPIV0ComplianceTrigger(describe.NewGetScheduleAPIV0ComplianceTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0DescribeTriggerCmd = &cobra.Command{
	Use: "get_schedule_api_v_0_describe_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		resp, err := client.Describe.GetScheduleAPIV0DescribeTrigger(describe.NewGetScheduleAPIV0DescribeTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0InsightTriggerCmd = &cobra.Command{
	Use: "get_schedule_api_v_0_insight_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		resp, err := client.Describe.GetScheduleAPIV0InsightTrigger(describe.NewGetScheduleAPIV0InsightTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0SummarizeTriggerCmd = &cobra.Command{
	Use: "get_schedule_api_v_0_summarize_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		resp, err := client.Describe.GetScheduleAPIV0SummarizeTrigger(describe.NewGetScheduleAPIV0SummarizeTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		return nil
	},
}

