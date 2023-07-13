// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"fmt"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/stack"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var PostScheduleApiV1StacksCreateCmd = &cobra.Command{
	Use:   "create-stack",
	Short: `Create a stack by giving terraform statefile and additional resources`,
	Long: `Create a stack by giving terraform statefile and additional resources
Config structure for azure: {tenantId: string, objectId: string, secretId: string, clientId: string, clientSecret:string}
Config structure for aws: {accessKey: string, secretKey: string}`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_create] : %v", err)
		}

		req := stack.NewPostScheduleAPIV1StacksCreateParams()

		req.SetConfig(flags.ReadStringFlag(cmd, "Config"))
		req.SetTag(flags.ReadStringOptionalFlag(cmd, "Tag"))
		reader, err := os.Open(flags.ReadStringFlag(cmd, "TerraformFile"))
		TerraformFile := runtime.NamedReader("TerraformFile", reader)
		req.SetTerraformFile(TerraformFile)

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

var DeleteScheduleApiV1StacksStackIdCmd = &cobra.Command{
	Use:   "delete-stack",
	Short: `Delete a stack by ID`,
	Long:  `Delete a stack by ID`,
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

var GetScheduleApiV1StacksStackIdCmd = &cobra.Command{
	Use:   "get-stack",
	Short: `Get a stack details by ID`,
	Long:  `Get a stack details by ID`,
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

var GetScheduleApiV1StacksStackIdInsightCmd = &cobra.Command{
	Use:   "get-stack-insight",
	Short: `Get Insight results for a stack in the given time period`,
	Long:  `Get Insight results for a stack in the given time period`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insight] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksStackIDInsightParams()

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightID(flags.ReadInt64Flag(cmd, "InsightID"))
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

var PostScheduleApiV1StacksStackIdFindingsCmd = &cobra.Command{
	Use:   "list-stack-findings",
	Short: `Get all findings for a stack`,
	Long:  `Get all findings for a stack`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_stack_id_findings] : %v", err)
		}

		req := stack.NewPostScheduleAPIV1StacksStackIDFindingsParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgDescribeAPIGetStackFindings{
			BenchmarkIds: flags.ReadStringArrayFlag(cmd, "BenchmarkIds"),
			Page: models.GithubComKaytuIoKaytuEnginePkgComplianceAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Sorts: []*models.GithubComKaytuIoKaytuEnginePkgComplianceAPIFindingSortItem{
				{
					Direction: models.GithubComKaytuIoKaytuEnginePkgComplianceAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GithubComKaytuIoKaytuEnginePkgComplianceAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})
		req.SetStackID(flags.ReadStringFlag(cmd, "StackID"))

		resp, err := client.Stack.PostScheduleAPIV1StacksStackIDFindings(req, auth)
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_stack_id_findings] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_schedule_api_v_1_stacks_stack_id_findings] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksStackIdInsightsCmd = &cobra.Command{
	Use:   "list-stack-insights",
	Short: `Get all Insights results with the given filters`,
	Long:  `Get all Insights results with the given filters`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insights] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksStackIDInsightsParams()

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetInsightIds(flags.ReadIntArrayFlag(cmd, "InsightIds"))
		req.SetStackID(flags.ReadStringFlag(cmd, "StackID"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Stack.GetScheduleAPIV1StacksStackIDInsights(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insights] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_stack_id_insights] : %v", err)
		}

		return nil
	},
}

var GetScheduleApiV1StacksCmd = &cobra.Command{
	Use:   "list-stacks",
	Short: `Get list of stacks`,
	Long:  `Get list of stacks`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksParams()

		req.SetAccountIds(flags.ReadStringArrayFlag(cmd, "AccountIds"))
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

var GetScheduleApiV1StacksResourceCmd = &cobra.Command{
	Use:   "resource-stacks",
	Short: `Get list of all stacks containing a resource`,
	Long:  `Get list of all stacks containing a resource`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource] : %v", err)
		}

		req := stack.NewGetScheduleAPIV1StacksResourceParams()

		req.SetResourceID(flags.ReadStringFlag(cmd, "ResourceID"))

		resp, err := client.Stack.GetScheduleAPIV1StacksResource(req, auth)
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_schedule_api_v_1_stacks_resource] : %v", err)
		}

		return nil
	},
}
