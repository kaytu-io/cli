package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/inventory"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetInventoryApiV2ResourcesTagKeyCmd = &cobra.Command{
	Use: "resources-tag-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_tag_key] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ResourcesTagKeyParams()

        req.SetKey(flags.ReadStringFlag("Key"))


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
	Use: "resources-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_trend] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ResourcesTrendParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetDatapointCount(flags.ReadStringOptionalFlag("DatapointCount"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetServicename(flags.ReadStringArrayFlag("Servicename"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))
req.SetTag(flags.ReadStringArrayFlag("Tag"))


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

var GetInventoryApiV2ServicesTagCmd = &cobra.Command{
	Use: "services-tag",
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

var GetInventoryApiV1ResourcesRegionsCmd = &cobra.Command{
	Use: "resources-regions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_regions] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV1ResourcesRegionsParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))


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

var GetInventoryApiV1ResourcesTopRegionsCmd = &cobra.Command{
	Use: "resources-top-regions",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_regions] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV1ResourcesTopRegionsParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetCount(flags.ReadInt64Flag("Count"))


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

var PostInventoryApiV1ResourcesAwsCmd = &cobra.Command{
	Use: "resources-aws",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_aws] : %v", err)
		}

        req := inventory.NewPostInventoryAPIV1ResourcesAwsParams()

        req.SetAccept(flags.ReadStringFlag("Accept"))
req.SetCommon(flags.ReadStringOptionalFlag("Common"))
req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
Category: flags.ReadStringArrayFlag("Category"),
Location: flags.ReadStringArrayFlag("Location"),
ResourceType: flags.ReadStringArrayFlag("ResourceType"),
Service: flags.ReadStringArrayFlag("Service"),
SourceID: flags.ReadStringArrayFlag("SourceID"),
Tags: flags.ReadMapStringFlag("Tags"),

},
Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
No: flags.ReadInt64Flag("No"),
Size: flags.ReadInt64Flag("Size"),

},
Query: flags.ReadStringFlag("Query"),
Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
{
Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag("Direction")),
Field: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag("Field")),

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
	Use: "resources-azure",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_azure] : %v", err)
		}

        req := inventory.NewPostInventoryAPIV1ResourcesAzureParams()

        req.SetAccept(flags.ReadStringFlag("Accept"))
req.SetCommon(flags.ReadStringOptionalFlag("Common"))
req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
Category: flags.ReadStringArrayFlag("Category"),
Location: flags.ReadStringArrayFlag("Location"),
ResourceType: flags.ReadStringArrayFlag("ResourceType"),
Service: flags.ReadStringArrayFlag("Service"),
SourceID: flags.ReadStringArrayFlag("SourceID"),
Tags: flags.ReadMapStringFlag("Tags"),

},
Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
No: flags.ReadInt64Flag("No"),
Size: flags.ReadInt64Flag("Size"),

},
Query: flags.ReadStringFlag("Query"),
Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
{
Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag("Direction")),
Field: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag("Field")),

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

var PostInventoryApiV1ResourcesCmd = &cobra.Command{
	Use: "resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources] : %v", err)
		}

        req := inventory.NewPostInventoryAPIV1ResourcesParams()

        req.SetAccept(flags.ReadStringFlag("Accept"))
req.SetCommon(flags.ReadStringOptionalFlag("Common"))
req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourcesRequest{
Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIFilters{
Category: flags.ReadStringArrayFlag("Category"),
Location: flags.ReadStringArrayFlag("Location"),
ResourceType: flags.ReadStringArrayFlag("ResourceType"),
Service: flags.ReadStringArrayFlag("Service"),
SourceID: flags.ReadStringArrayFlag("SourceID"),
Tags: flags.ReadMapStringFlag("Tags"),

},
Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
No: flags.ReadInt64Flag("No"),
Size: flags.ReadInt64Flag("Size"),

},
Query: flags.ReadStringFlag("Query"),
Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceSortItem{
{
Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag("Direction")),
Field: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISortFieldType(flags.ReadStringFlag("Field")),

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
	Use: "resources-count",
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

var GetInventoryApiV2ServicesCompositionKeyCmd = &cobra.Command{
	Use: "services-composition-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_composition_key] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ServicesCompositionKeyParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetKey(flags.ReadStringFlag("Key"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))
req.SetTop(flags.ReadInt64Flag("Top"))


		resp, err := client.Inventory.GetInventoryAPIV2ServicesCompositionKey(req, auth)
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

var GetInventoryApiV2ServicesMetricCmd = &cobra.Command{
	Use: "services-metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_metric] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ServicesMetricParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
req.SetSortBy(flags.ReadStringOptionalFlag("SortBy"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))
req.SetTag(flags.ReadStringArrayFlag("Tag"))


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

var GetInventoryApiV2ResourcesMetricCmd = &cobra.Command{
	Use: "resources-metric",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_metric] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ResourcesMetricParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
req.SetServicename(flags.ReadStringArrayFlag("Servicename"))
req.SetSortBy(flags.ReadStringOptionalFlag("SortBy"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))
req.SetTag(flags.ReadStringArrayFlag("Tag"))


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
	Use: "services-cost-trend",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_cost_trend] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ServicesCostTrendParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetDatapointCount(flags.ReadStringOptionalFlag("DatapointCount"))
req.SetEndTime(flags.ReadStringOptionalFlag("EndTime"))
req.SetStartTime(flags.ReadStringOptionalFlag("StartTime"))
req.SetTag(flags.ReadStringOptionalFlag("Tag"))


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

var GetInventoryApiV2ServicesTagKeyCmd = &cobra.Command{
	Use: "services-tag-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_tag_key] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ServicesTagKeyParams()

        req.SetKey(flags.ReadStringFlag("Key"))


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

var PostInventoryApiV1ResourcesFiltersCmd = &cobra.Command{
	Use: "resources-filters",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resources_filters] : %v", err)
		}

        req := inventory.NewPostInventoryAPIV1ResourcesFiltersParams()

        req.SetCommon(flags.ReadStringOptionalFlag("Common"))
req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetFiltersRequest{
Filters: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIResourceFilters{
Category: flags.ReadStringArrayFlag("Category"),
Connections: flags.ReadStringArrayFlag("Connections"),
Location: flags.ReadStringArrayFlag("Location"),
Provider: flags.ReadStringArrayFlag("Provider"),
ResourceType: flags.ReadStringArrayFlag("ResourceType"),
Service: flags.ReadStringArrayFlag("Service"),
TagKeys: flags.ReadStringArrayFlag("TagKeys"),
TagValues: flags.ReadMapStringArrayFlag("TagValues"),

},
Query: flags.ReadStringFlag("Query"),

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

var GetInventoryApiV2ResourcesCompositionKeyCmd = &cobra.Command{
	Use: "resources-composition-key",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_resources_composition_key] : %v", err)
		}

        req := inventory.NewGetInventoryAPIV2ResourcesCompositionKeyParams()

        req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
req.SetConnector(flags.ReadStringArrayFlag("Connector"))
req.SetKey(flags.ReadStringFlag("Key"))
req.SetTime(flags.ReadStringOptionalFlag("Time"))
req.SetTop(flags.ReadInt64Flag("Top"))


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

var GetInventoryApiV2ResourcesTagCmd = &cobra.Command{
	Use: "resources-tag",
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
