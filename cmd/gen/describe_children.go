package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/describe"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var PutScheduleApiV1BenchmarkEvaluationTriggerCmd = &cobra.Command{
	Use: "benchmark-evaluation-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_benchmark_evaluation_trigger] : %v", err)
		}

		req := describe.NewPutScheduleAPIV1BenchmarkEvaluationTriggerParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPITriggerBenchmarkEvaluationRequest{
			BenchmarkID:  flags.ReadStringFlag(cmd, "BenchmarkID"),
			ConnectionID: flags.ReadStringFlag(cmd, "ConnectionID"),
			ResourceIDs:  flags.ReadStringArrayFlag(cmd, "ResourceIDs"),
		})

		resp, err := client.Describe.PutScheduleAPIV1BenchmarkEvaluationTrigger(req, auth)
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

var GetScheduleApiV0ComplianceSummarizerTriggerCmd = &cobra.Command{
	Use: "compliance-summarizer-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		req := describe.NewGetScheduleAPIV0ComplianceSummarizerTriggerParams()

		_, err = client.Describe.GetScheduleAPIV0ComplianceSummarizerTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_summarizer_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0ComplianceTriggerCmd = &cobra.Command{
	Use: "compliance-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		req := describe.NewGetScheduleAPIV0ComplianceTriggerParams()

		_, err = client.Describe.GetScheduleAPIV0ComplianceTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_compliance_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0DescribeTriggerCmd = &cobra.Command{
	Use: "describe-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		req := describe.NewGetScheduleAPIV0DescribeTriggerParams()

		_, err = client.Describe.GetScheduleAPIV0DescribeTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_describe_trigger] : %v", err)
		}

		return nil
	},
}
var PostScheduleApiV1DescribeResourceCmd = &cobra.Command{
	Use: "describe-resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_describe_resource] : %v", err)
		}

		req := describe.NewPostScheduleAPIV1DescribeResourceParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIDescribeSingleResourceRequest{
			AccessKey:        flags.ReadStringFlag(cmd, "AccessKey"),
			AccountID:        flags.ReadStringFlag(cmd, "AccountID"),
			AdditionalFields: flags.ReadMapStringFlag(cmd, "AdditionalFields"),
			Provider:         models.SourceType(flags.ReadStringFlag(cmd, "Provider")),
			ResourceType:     flags.ReadStringFlag(cmd, "ResourceType"),
			SecretKey:        flags.ReadStringFlag(cmd, "SecretKey"),
		})

		resp, err := client.Describe.PostScheduleAPIV1DescribeResource(req, auth)
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

var PostScheduleApiV1StacksInsightTriggerCmd = &cobra.Command{
	Use: "stacks-insight-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_insight_trigger] : %v", err)
		}

		req := describe.NewPostScheduleAPIV1StacksInsightTriggerParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIStackInsightRequest{
			Insights: flags.ReadIntArrayFlag(cmd, "Insights"),
			StackID:  flags.ReadStringOptionalFlag(cmd, "StackID"),
		})

		resp, err := client.Describe.PostScheduleAPIV1StacksInsightTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_insight_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_insight_trigger] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV0InsightTriggerCmd = &cobra.Command{
	Use: "insight-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		req := describe.NewGetScheduleAPIV0InsightTriggerParams()

		_, err = client.Describe.GetScheduleAPIV0InsightTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_insight_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV0SummarizeTriggerCmd = &cobra.Command{
	Use: "summarize-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		req := describe.NewGetScheduleAPIV0SummarizeTriggerParams()

		_, err = client.Describe.GetScheduleAPIV0SummarizeTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_0_summarize_trigger] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1InsightJobJobIdCmd = &cobra.Command{
	Use: "insight-job-job-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_job_job_id] : %v", err)
		}

		req := describe.NewGetScheduleAPIV1InsightJobJobIDParams()

		req.SetJobID(flags.ReadStringFlag(cmd, "JobID"))

		resp, err := client.Describe.GetScheduleAPIV1InsightJobJobID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_job_job_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_insight_job_job_id] : %v", err)
		}

		return nil
	},
}

var PutScheduleApiV1DescribeTriggerConnectionIdCmd = &cobra.Command{
	Use: "describe-trigger-connection-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		req := describe.NewPutScheduleAPIV1DescribeTriggerConnectionIDParams()

		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))

		_, err = client.Describe.PutScheduleAPIV1DescribeTriggerConnectionID(req, auth)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_describe_trigger_connection_id] : %v", err)
		}

		return nil
	},
}
var PutScheduleApiV1InsightEvaluationTriggerCmd = &cobra.Command{
	Use: "insight-evaluation-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_insight_evaluation_trigger] : %v", err)
		}

		req := describe.NewPutScheduleAPIV1InsightEvaluationTriggerParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPITriggerInsightEvaluationRequest{
			ConnectionID: flags.ReadStringFlag(cmd, "ConnectionID"),
			InsightID:    flags.ReadInt64Flag(cmd, "InsightID"),
			ResourceIDs:  flags.ReadStringArrayFlag(cmd, "ResourceIDs"),
		})

		resp, err := client.Describe.PutScheduleAPIV1InsightEvaluationTrigger(req, auth)
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
