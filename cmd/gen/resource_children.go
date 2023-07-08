// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var PostInventoryApiV1ResourceCmd = &cobra.Command{
	Use: "create-resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		req := resource.NewPostInventoryAPIV1ResourceParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgInventoryAPIGetResourceRequest{
			ID:           flags.ReadStringOptionalFlag(cmd, "ID"),
			ResourceType: flags.ReadStringOptionalFlag(cmd, "ResourceType"),
		})

		resp, err := client.Resource.PostInventoryAPIV1Resource(req, auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesRegionsCompositionCmd = &cobra.Command{
	Use: "get-resource-composition-by-region",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_composition] : %v", err)
		}

		req := resource.NewGetInventoryAPIV2ResourcesRegionsCompositionParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))
		req.SetTop(flags.ReadInt64Flag(cmd, "Top"))

		resp, err := client.Resource.GetInventoryAPIV2ResourcesRegionsComposition(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_composition] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_composition] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesCountCmd = &cobra.Command{
	Use: "get-resource-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_count] : %v", err)
		}

		req := resource.NewGetInventoryAPIV2ResourcesCountParams()

		resp, err := client.Resource.GetInventoryAPIV2ResourcesCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_count] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesRegionsSummaryCmd = &cobra.Command{
	Use: "get-resource-summary-by-region",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_summary] : %v", err)
		}

		req := resource.NewGetInventoryAPIV2ResourcesRegionsSummaryParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Resource.GetInventoryAPIV2ResourcesRegionsSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_summary] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_summary] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ResourcesRegionsTrendCmd = &cobra.Command{
	Use: "get-resource-trend-by-region",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_trend] : %v", err)
		}

		req := resource.NewGetInventoryAPIV2ResourcesRegionsTrendParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetDatapointCount(flags.ReadInt64OptionalFlag(cmd, "DatapointCount"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetRegion(flags.ReadStringArrayFlag(cmd, "Region"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Resource.GetInventoryAPIV2ResourcesRegionsTrend(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_trend] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_regions_trend] : %v", err)
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

		req := resource.NewPostInventoryAPIV1ResourcesParams()

		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgInventoryAPIGetResourcesRequest{
			Filters: models.GithubComKaytuIoKaytuEnginePkgInventoryAPIFilters{
				ConnectionID: flags.ReadStringArrayFlag(cmd, "ConnectionID"),
				Connectors:   flags.ReadEnumArrayFlag[models.SourceType](cmd, "Connectors"),
				Location:     flags.ReadStringArrayFlag(cmd, "Location"),
				ResourceType: flags.ReadStringArrayFlag(cmd, "ResourceType"),
				Service:      flags.ReadStringArrayFlag(cmd, "Service"),
			},
			Page: &models.GithubComKaytuIoKaytuEnginePkgInventoryAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Query: flags.ReadStringFlag(cmd, "Query"),
			Sorts: []*models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceSortItem{
				{
					Direction: models.GithubComKaytuIoKaytuEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     models.GithubComKaytuIoKaytuEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag(cmd, "Field")),
				},
			},
		})

		resp, err := client.Resource.PostInventoryAPIV1Resources(req, auth)
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

var PostInventoryApiV1ResourcesFiltersCmd = &cobra.Command{
	Use: "list-resources-by-filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

		req := resource.NewPostInventoryAPIV1ResourcesFiltersParams()

		req.SetCommon(flags.ReadStringOptionalFlag(cmd, "Common"))
		req.SetRequest(&models.GithubComKaytuIoKaytuEnginePkgInventoryAPIGetFiltersRequest{
			Filters: models.GithubComKaytuIoKaytuEnginePkgInventoryAPIResourceFilters{
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

		resp, err := client.Resource.PostInventoryAPIV1ResourcesFilters(req, auth)
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

		req := resource.NewGetInventoryAPIV1ResourcesRegionsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetStartTime(flags.ReadStringOptionalFlag(cmd, "StartTime"))

		resp, err := client.Resource.GetInventoryAPIV1ResourcesRegions(req, auth)
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

var GetInventoryApiV1ResourcesTopRegionsCmd = &cobra.Command{
	Use: "top-regions-by-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

		req := resource.NewGetInventoryAPIV1ResourcesTopRegionsParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))

		resp, err := client.Resource.GetInventoryAPIV1ResourcesTopRegions(req, auth)
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
