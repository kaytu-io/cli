package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/spf13/cobra"
)

var GetScheduleApiV0ComplianceSummarizerTriggerCmd = &cobra.Command{
	Use: "summarizerSummarizerTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		_, err = client.Describe.GetScheduleAPIV0ComplianceSummarizerTrigger(describe.NewGetScheduleAPIV0ComplianceSummarizerTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0ComplianceTriggerCmd = &cobra.Command{
	Use: "triggerTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		_, err = client.Describe.GetScheduleAPIV0ComplianceTrigger(describe.NewGetScheduleAPIV0ComplianceTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0DescribeTriggerCmd = &cobra.Command{
	Use: "triggerTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		_, err = client.Describe.GetScheduleAPIV0DescribeTrigger(describe.NewGetScheduleAPIV0DescribeTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		return nil
	},
}
var PostScheduleApiV1DescribeResourceCmd = &cobra.Command{
	Use: "resourceResource",
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

var PutScheduleApiV1DescribeTriggerConnectionIdCmd = &cobra.Command{
	Use: "triggerTriggerConnectionId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		_, err = client.Describe.PutScheduleAPIV1DescribeTriggerConnectionID(describe.NewPutScheduleAPIV1DescribeTriggerConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0InsightTriggerCmd = &cobra.Command{
	Use: "triggerTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		_, err = client.Describe.GetScheduleAPIV0InsightTrigger(describe.NewGetScheduleAPIV0InsightTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0SummarizeTriggerCmd = &cobra.Command{
	Use: "triggerTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		_, err = client.Describe.GetScheduleAPIV0SummarizeTrigger(describe.NewGetScheduleAPIV0SummarizeTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		return nil
	},
}
var PutScheduleApiV1BenchmarkEvaluationTriggerCmd = &cobra.Command{
	Use: "evaluationEvaluationTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_benchmark_evaluation_trigger] : %v", err)
		}

		resp, err := client.Describe.PutScheduleAPIV1BenchmarkEvaluationTrigger(describe.NewPutScheduleAPIV1BenchmarkEvaluationTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_benchmark_evaluation_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_benchmark_evaluation_trigger] : %v", err)
		}

		return nil
	},
}

var PutScheduleApiV1ComplianceTriggerCmd = &cobra.Command{
	Use: "triggerTrigger",
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

var PutScheduleApiV1InsightEvaluationTriggerCmd = &cobra.Command{
	Use: "evaluationEvaluationTrigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_insight_evaluation_trigger] : %v", err)
		}

		resp, err := client.Describe.PutScheduleAPIV1InsightEvaluationTrigger(describe.NewPutScheduleAPIV1InsightEvaluationTriggerParams(), auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_insight_evaluation_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_insight_evaluation_trigger] : %v", err)
		}

		return nil
	},
}
