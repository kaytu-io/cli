package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

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

var PostInventoryApiV1ResourcesAwsCmd = &cobra.Command{
	Use: "list-aws-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_aws] : %v", err)
		}

		req := inventory.NewPostInventoryAPIV1ResourcesAwsParams()

		req.SetAccept(flags.ReadStringFlag(cmd, "Accept"))
		req.SetCommon(flags.ReadStringOptionalFlag(cmd, "Common"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
			Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
				Category:     flags.ReadStringArrayFlag(cmd, "Category"),
				Location:     flags.ReadStringArrayFlag(cmd, "Location"),
				ResourceType: flags.ReadStringArrayFlag(cmd, "ResourceType"),
				Service:      flags.ReadStringArrayFlag(cmd, "Service"),
				SourceID:     flags.ReadStringArrayFlag(cmd, "SourceID"),
				Tags:         flags.ReadMapStringFlag(cmd, "Tags"),
			},
			Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Query: flags.ReadStringFlag(cmd, "Query"),
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesAws(req, auth)
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

var PostInventoryApiV1ResourcesAzureCmd = &cobra.Command{
	Use: "list-azure-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_azure] : %v", err)
		}

		req := inventory.NewPostInventoryAPIV1ResourcesAzureParams()

		req.SetAccept(flags.ReadStringFlag(cmd, "Accept"))
		req.SetCommon(flags.ReadStringOptionalFlag(cmd, "Common"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
			Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
				Category:     flags.ReadStringArrayFlag(cmd, "Category"),
				Location:     flags.ReadStringArrayFlag(cmd, "Location"),
				ResourceType: flags.ReadStringArrayFlag(cmd, "ResourceType"),
				Service:      flags.ReadStringArrayFlag(cmd, "Service"),
				SourceID:     flags.ReadStringArrayFlag(cmd, "SourceID"),
				Tags:         flags.ReadMapStringFlag(cmd, "Tags"),
			},
			Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Query: flags.ReadStringFlag(cmd, "Query"),
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesAzure(req, auth)
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
	Use: "list-resources-by-filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

		req := inventory.NewPostInventoryAPIV1ResourcesFiltersParams()

		req.SetCommon(flags.ReadStringOptionalFlag(cmd, "Common"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest{
			Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceFilters{
				Category:     flags.ReadStringArrayFlag(cmd, "Category"),
				Connections:  flags.ReadStringArrayFlag(cmd, "Connections"),
				Location:     flags.ReadStringArrayFlag(cmd, "Location"),
				Provider:     flags.ReadStringArrayFlag(cmd, "Provider"),
				ResourceType: flags.ReadStringArrayFlag(cmd, "ResourceType"),
				Service:      flags.ReadStringArrayFlag(cmd, "Service"),
				TagKeys:      flags.ReadStringArrayFlag(cmd, "TagKeys"),
				TagValues:    flags.ReadMapStringArrayFlag(cmd, "TagValues"),
			},
			Query: flags.ReadStringFlag(cmd, "Query"),
		})

		resp, err := client.Inventory.PostInventoryAPIV1ResourcesFilters(req, auth)
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

var GetInventoryApiV1ResourcesRegionsCmd = &cobra.Command{
	Use: "regions-by-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_regions] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV1ResourcesRegionsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesRegions(req, auth)
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
		req.SetKey(flags.ReadStringFlag(cmd, "Key"))
		req.SetTime(flags.ReadStringOptionalFlag(cmd, "Time"))
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

var GetInventoryApiV2ResourcesTagKeyCmd = &cobra.Command{
	Use: "get-resource-tag",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesTagKeyParams()

		req.SetKey(flags.ReadStringFlag(cmd, "Key"))

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

var GetInventoryApiV1ResourcesTopRegionsCmd = &cobra.Command{
	Use: "top-regions-by-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV1ResourcesTopRegionsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesTopRegions(req, auth)
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

var GetInventoryApiV2ResourcesTagCmd = &cobra.Command{
	Use: "list-resource-tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV2ResourcesTagParams()

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

var PostInventoryApiV1ResourcesCmd = &cobra.Command{
	Use: "list-all-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources] : %v", err)
		}

		req := inventory.NewPostInventoryAPIV1ResourcesParams()

		req.SetAccept(flags.ReadStringFlag(cmd, "Accept"))
		req.SetCommon(flags.ReadStringOptionalFlag(cmd, "Common"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
			Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
				Category:     flags.ReadStringArrayFlag(cmd, "Category"),
				Location:     flags.ReadStringArrayFlag(cmd, "Location"),
				ResourceType: flags.ReadStringArrayFlag(cmd, "ResourceType"),
				Service:      flags.ReadStringArrayFlag(cmd, "Service"),
				SourceID:     flags.ReadStringArrayFlag(cmd, "SourceID"),
				Tags:         flags.ReadMapStringFlag(cmd, "Tags"),
			},
			Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Query: flags.ReadStringFlag(cmd, "Query"),
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Inventory.PostInventoryAPIV1Resources(req, auth)
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
	Use: "all-resources-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_count] : %v", err)
		}

		req := inventory.NewGetInventoryAPIV1ResourcesCountParams()

		resp, err := client.Inventory.GetInventoryAPIV1ResourcesCount(req, auth)
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
		req.SetTop(flags.ReadInt64Flag(cmd, "Top"))

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
