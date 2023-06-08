package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/spf13/cobra"
)
var DeleteScheduleApiV1StacksStackIdCmd = &cobra.Command{
	Use: "delete_schedule_api_v_1_stacks_stack_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[delete_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		resp, err := client.Stack.DeleteScheduleAPIV1StacksStackID(stack.NewDeleteScheduleAPIV1StacksStackIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[delete_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[delete_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		return nil
	},
}
var GetScheduleApiV1StacksFindingsJobIdCmd = &cobra.Command{
	Use: "get_schedule_api_v_1_stacks_findings_job_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_findings_job_id] : %v", err)
		}

		resp, err := client.Stack.GetScheduleAPIV1StacksFindingsJobID(stack.NewGetScheduleAPIV1StacksFindingsJobIDParams(), auth)
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
	Use: "get_schedule_api_v_1_stacks",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks] : %v", err)
		}

		resp, err := client.Stack.GetScheduleAPIV1Stacks(stack.NewGetScheduleAPIV1StacksParams(), auth)
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
	Use: "get_schedule_api_v_1_stacks_resource_resource_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource_resource_id] : %v", err)
		}

		resp, err := client.Stack.GetScheduleAPIV1StacksResourceResourceID(stack.NewGetScheduleAPIV1StacksResourceResourceIDParams(), auth)
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
	Use: "get_schedule_api_v_1_stacks_stack_id_insight",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insight] : %v", err)
		}

		resp, err := client.Stack.GetScheduleAPIV1StacksStackIDInsight(stack.NewGetScheduleAPIV1StacksStackIDInsightParams(), auth)
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
	Use: "get_schedule_api_v_1_stacks_stack_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id] : %v", err)
		}

		resp, err := client.Stack.GetScheduleAPIV1StacksStackID(stack.NewGetScheduleAPIV1StacksStackIDParams(), auth)
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
	Use: "post_schedule_api_v_1_stacks_benchmark_trigger",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_benchmark_trigger] : %v", err)
		}

		resp, err := client.Stack.PostScheduleAPIV1StacksBenchmarkTrigger(stack.NewPostScheduleAPIV1StacksBenchmarkTriggerParams(), auth)
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
	Use: "post_schedule_api_v_1_stacks_create",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
		}

		resp, err := client.Stack.PostScheduleAPIV1StacksCreate(stack.NewPostScheduleAPIV1StacksCreateParams(), auth)
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
