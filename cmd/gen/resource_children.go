// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/spf13/cobra"
)

var PostAiApiV1GptRunCmd = &cobra.Command{
	Use:   "_",
	Short: ``,
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[post_ai_api_v_1_gpt_run] : %v", err)
		}

		req := resource.NewPostAiAPIV1GptRunParams()

		req.SetQuery(flags.ReadStringFlag(cmd, "Query"))

		resp, err := client.Resource.PostAiAPIV1GptRun(req, auth)
		if err != nil {
			return fmt.Errorf("[post_ai_api_v_1_gpt_run] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "post-ai-api-v1-gpt-run", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_ai_api_v_1_gpt_run] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesMetricResourceTypeCmd = &cobra.Command{
	Use:   "resource-metrics",
	Short: `Retrieving metrics for a specific resource type.`,
	Long:  `Retrieving metrics for a specific resource type.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		req := resource.NewGetInventoryAPIV2ResourcesMetricResourceTypeParams()

		req.SetConnectionGroup(flags.ReadStringOptionalFlag(cmd, "ConnectionGroup"))
		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
		req.SetResourceType(flags.ReadStringFlag(cmd, "ResourceType"))
		req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))

		resp, err := client.Resource.GetInventoryAPIV2ResourcesMetricResourceType(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-inventory-api-v2-resources-metric-resource-type", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		return nil
	},
}
