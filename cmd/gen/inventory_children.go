package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/spf13/cobra"
)
var PostInventoryApiV1ResourcesAwsCmd = &cobra.Command{
	Use: "post_inventory_api_v_1_resources_aws",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_aws] : %v", err)
		}

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesAws(inventory.NewPostInventoryAPIV1ResourcesAwsParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_aws] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_aws] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV1ResourcesRegionsCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_resources_regions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_regions] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesRegions(inventory.NewGetInventoryAPIV1ResourcesRegionsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_regions] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_regions] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ResourcesTrendCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_resources_trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_trend] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTrend(inventory.NewGetInventoryAPIV2ResourcesTrendParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_trend] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ServicesTagKeyCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_services_tag_key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag_key] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ServicesTagKey(inventory.NewGetInventoryAPIV2ServicesTagKeyParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag_key] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ServicesTagCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_services_tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ServicesTag(inventory.NewGetInventoryAPIV2ServicesTagParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ResourcesCompositionKeyCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_resources_composition_key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_composition_key] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesCompositionKey(inventory.NewGetInventoryAPIV2ResourcesCompositionKeyParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_composition_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_composition_key] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ResourcesMetricCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_resources_metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesMetric(inventory.NewGetInventoryAPIV2ResourcesMetricParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ServicesCostTrendCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_services_cost_trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_cost_trend] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ServicesCostTrend(inventory.NewGetInventoryAPIV2ServicesCostTrendParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_cost_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_cost_trend] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ServicesMetricCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_services_metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ServicesMetric(inventory.NewGetInventoryAPIV2ServicesMetricParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ResourcesTagCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_resources_tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTag(inventory.NewGetInventoryAPIV2ResourcesTagParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ServicesCompositionKeyCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_services_composition_key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_composition_key] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ServicesCompositionKey(inventory.NewGetInventoryAPIV2ServicesCompositionKeyParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_composition_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_composition_key] : %v", err)
		}

		return nil
	},
}
var PostInventoryApiV1ResourcesAzureCmd = &cobra.Command{
	Use: "post_inventory_api_v_1_resources_azure",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_azure] : %v", err)
		}

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesAzure(inventory.NewPostInventoryAPIV1ResourcesAzureParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_azure] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_azure] : %v", err)
		}

		return nil
	},
}
var PostInventoryApiV1ResourcesFiltersCmd = &cobra.Command{
	Use: "post_inventory_api_v_1_resources_filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesFilters(inventory.NewPostInventoryAPIV1ResourcesFiltersParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

		return nil
	},
}
var PostInventoryApiV1ResourcesCmd = &cobra.Command{
	Use: "post_inventory_api_v_1_resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources] : %v", err)
		}

		resp, err := client.Inventory.PostInventoryAPIV1Resources(inventory.NewPostInventoryAPIV1ResourcesParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV1ResourcesCountCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_resources_count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_count] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesCount(inventory.NewGetInventoryAPIV1ResourcesCountParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_count] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV1ResourcesTopRegionsCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_resources_top_regions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesTopRegions(inventory.NewGetInventoryAPIV1ResourcesTopRegionsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ResourcesTagKeyCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_resources_tag_key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTagKey(inventory.NewGetInventoryAPIV2ResourcesTagKeyParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

		return nil
	},
}

