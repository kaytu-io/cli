// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2CostCompositionCmd = &cobra.Command{
	Use: "cost-composition",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_composition] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2CostCompositionParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))
		req.SetTop(flags.ReadInt64OptionalFlag(cmd, "Top"))

		resp, err := client.Inventory.GetInventoryAPIV2CostComposition(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_composition] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_composition] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2CostMetricCmd = &cobra.Command{
	Use: "cost-metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_metric] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2CostMetricParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV2CostMetric(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_metric] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_metric] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2CostTrendCmd = &cobra.Command{
	Use: "cost-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_trend] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2CostTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetDatapointCount(flags.ReadStringOptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV2CostTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_cost_trend] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesTagKeyCmd = &cobra.Command{
	Use: "get-resource-tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesTagKeyParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetKey(flags.ReadStringFlag(cmd, "Key"))
		req.SetMinCount(flags.ReadInt64OptionalFlag(cmd, "MinCount"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTagKey(req, auth)
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

var GetInventoryApiV2ServicesTagKeyCmd = &cobra.Command{
	Use: "get-service-tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag_key] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ServicesTagKeyParams()

		req.SetKey(flags.ReadStringFlag(cmd, "Key"))

		resp, err := client.Inventory.GetInventoryAPIV2ServicesTagKey(req, auth)
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

var GetInventoryApiV2ResourcesTagCmd = &cobra.Command{
	Use: "list-resource-tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesTagParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetMinCount(flags.ReadInt64OptionalFlag(cmd, "MinCount"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTag(req, auth)
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

var GetInventoryApiV2ServicesTagCmd = &cobra.Command{
	Use: "list-service-tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ServicesTagParams()

		resp, err := client.Inventory.GetInventoryAPIV2ServicesTag(req, auth)
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

var GetInventoryApiV2ResourcesMetricResourceTypeCmd = &cobra.Command{
	Use: "resource-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesMetricResourceTypeParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetResourceType(flags.ReadStringFlag(cmd, "ResourceType"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesMetricResourceType(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric_resource_type] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesTrendCmd = &cobra.Command{
	Use: "resource-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_trend] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetDatapointCount(flags.ReadStringOptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetServicename(flags.ReadStringArrayFlag(cmd, "Servicename"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesTrend(req, auth)
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

var GetInventoryApiV2ResourcesCompositionKeyCmd = &cobra.Command{
	Use: "resources-composition",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_composition_key] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesCompositionKeyParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetKey(flags.ReadStringFlag(cmd, "Key"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))
		req.SetTop(flags.ReadInt64Flag(cmd, "Top"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesCompositionKey(req, auth)
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
	Use: "resources-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesMetricParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetMinCount(flags.ReadInt64OptionalFlag(cmd, "MinCount"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetServicename(flags.ReadStringArrayFlag(cmd, "Servicename"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Inventory.GetInventoryAPIV2ResourcesMetric(req, auth)
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
	Use: "service-cost-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_cost_trend] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ServicesCostTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetDatapointCount(flags.ReadStringOptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetServices(flags.ReadStringArrayFlag(cmd, "Services"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV2ServicesCostTrend(req, auth)
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

var GetInventoryApiV2ServicesMetricServiceNameCmd = &cobra.Command{
	Use: "service-metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric_service_name] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ServicesMetricServiceNameParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetServiceName(flags.ReadStringFlag(cmd, "ServiceName"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV2ServicesMetricServiceName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric_service_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric_service_name] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ServicesMetricCmd = &cobra.Command{
	Use: "service-metrics",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ServicesMetricParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Inventory.GetInventoryAPIV2ServicesMetric(req, auth)
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
