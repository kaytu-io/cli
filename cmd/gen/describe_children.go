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

var GetScheduleApiV0SummarizeTriggerCmd = &cobra.Command{
	Use: "trigger-summerize",
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
var PostScheduleApiV1DescribeResourceCmd = &cobra.Command{
	Use: "describe-resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_describe_resource] : %v", err)
		}

		req := describe.NewPostScheduleAPIV1DescribeResourceParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgDescribeAPIDescribeSingleResourceRequest{
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
	Use: "trigger-stack-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_insight_trigger] : %v", err)
		}

		req := describe.NewPostScheduleAPIV1StacksInsightTriggerParams()

		//req.SetRequest(v)

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

var PutScheduleApiV1InsightEvaluationTriggerCmd = &cobra.Command{
	Use: "trigger-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_insight_evaluation_trigger] : %v", err)
		}

		req := describe.NewPutScheduleAPIV1InsightEvaluationTriggerParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgDescribeAPITriggerInsightEvaluationRequest{
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

var GetScheduleApiV0ComplianceSummarizerTriggerCmd = &cobra.Command{
	Use: "trigger-compliance-summerizer",
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
	Use: "trigger-compliance",
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
var GetScheduleApiV0InsightTriggerCmd = &cobra.Command{
	Use: "trigger-insight",
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
var GetScheduleApiV1BenchmarkEvaluationsCmd = &cobra.Command{
	Use: "benchmark-evaluations",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		req := describe.NewGetScheduleAPIV1BenchmarkEvaluationsParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgDescribeAPIListBenchmarkEvaluationsRequest{
			BenchmarkID:       flags.ReadStringFlag(cmd, "BenchmarkID"),
			ConnectionID:      flags.ReadStringFlag(cmd, "ConnectionID"),
			Connector:         models.SourceType(flags.ReadStringFlag(cmd, "Connector")),
			EvaluatedAtAfter:  flags.ReadInt64Flag(cmd, "EvaluatedAtAfter"),
			EvaluatedAtBefore: flags.ReadInt64Flag(cmd, "EvaluatedAtBefore"),
		})

		resp, err := client.Describe.GetScheduleAPIV1BenchmarkEvaluations(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_benchmark_evaluations] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1InsightJobJobIdCmd = &cobra.Command{
	Use: "get-insight-job",
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

var PutScheduleApiV1BenchmarkEvaluationTriggerCmd = &cobra.Command{
	Use: "trigger-benchmark-evaluation",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[put_schedule_api_v_1_benchmark_evaluation_trigger] : %v", err)
		}

		req := describe.NewPutScheduleAPIV1BenchmarkEvaluationTriggerParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgDescribeAPITriggerBenchmarkEvaluationRequest{
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

var PutScheduleApiV1DescribeTriggerConnectionIdCmd = &cobra.Command{
	Use: "trigger-describe",
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
