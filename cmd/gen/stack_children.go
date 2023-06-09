package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var DeleteScheduleApiV1StacksStackIdCmd = &cobra.Command{
	Use: "stacks-stack-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		req := stack.NewDeleteScheduleAPIV1StacksStackIDParams()

		req.SetStackID(flags.ReadStringFlag(cmd, "StackID"))

		_, err = client.Stack.DeleteScheduleAPIV1StacksStackID(req, auth)
		if err != nil {
			return fmt.Errorf("[delete_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1StacksFindingsJobIdCmd = &cobra.Command{
	Use: "stacks-findings-job-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_findings_job_id] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksFindingsJobIDParams()

		req.SetJobID(flags.ReadStringFlag(cmd, "JobID"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIGetStackFindings{
			Page: &models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIFindingSortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgComplianceAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GitlabComKeibiengineKeibiEnginePkgComplianceAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Stack.GetScheduleAPIV1StacksFindingsJobID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_findings_job_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_findings_job_id] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksCmd = &cobra.Command{
	Use: "stacks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksParams()

		req.SetAccounIds(flags.ReadStringArrayFlag(cmd, "AccounIds"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Stack.GetScheduleAPIV1Stacks(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksResourceResourceIdCmd = &cobra.Command{
	Use: "stacks-resource-resource-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource_resource_id] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksResourceResourceIDParams()

		req.SetResourceID(flags.ReadStringFlag(cmd, "ResourceID"))

		resp, err := client.Stack.GetScheduleAPIV1StacksResourceResourceID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource_resource_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource_resource_id] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksStackIdInsightCmd = &cobra.Command{
	Use: "stacks-stack-id-insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insight] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksStackIDInsightParams()

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadStringFlag(cmd, "InsightID"))
		req.SetStackID(flags.ReadStringFlag(cmd, "StackID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Stack.GetScheduleAPIV1StacksStackIDInsight(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insight] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insight] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksStackIdCmd = &cobra.Command{
	Use: "stacks-stack-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksStackIDParams()

		req.SetStackID(flags.ReadStringFlag(cmd, "StackID"))

		resp, err := client.Stack.GetScheduleAPIV1StacksStackID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		return nil
	},
}

var PostScheduleApiV1StacksBenchmarkTriggerCmd = &cobra.Command{
	Use: "stacks-benchmark-trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_benchmark_trigger] : %v", err)
		}

		req := stack.NewPostScheduleAPIV1StacksBenchmarkTriggerParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgDescribeAPIStackBenchmarkRequest{
			Benchmarks: flags.ReadStringArrayFlag(cmd, "Benchmarks"),
			StackID:    flags.ReadStringOptionalFlag(cmd, "StackID"),
		})

		resp, err := client.Stack.PostScheduleAPIV1StacksBenchmarkTrigger(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_benchmark_trigger] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_benchmark_trigger] : %v", err)
		}

		return nil
	},
}

var PostScheduleApiV1StacksCreateCmd = &cobra.Command{
	Use: "stacks-create",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
		}

		req := stack.NewPostScheduleAPIV1StacksCreateParams()

		req.SetResources(flags.ReadStringArrayFlag(cmd, "Resources"))
		req.SetTags(flags.ReadStringOptionalFlag(cmd, "Tags"))
		//req.SetTerrafromFile(v)

		resp, err := client.Stack.PostScheduleAPIV1StacksCreate(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
		}

		return nil
	},
}
