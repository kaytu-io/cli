package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2ServicesSummaryCmd = &cobra.Command{
	Use: "services-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV2ServicesSummaryParams()

		req.SetConnectionID(flags.ReadStringOptionalFlag("ConnectionID"))
		req.SetConnector(flags.ReadStringOptionalFlag("Connector"))
		req.SetEndTime(flags.ReadStringFlag("EndTime"))
		req.SetMinSpent(flags.ReadInt64OptionalFlag("MinSpent"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag("PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag("PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag("SortBy"))
		req.SetStartTime(flags.ReadStringFlag("StartTime"))
		req.SetTag(flags.ReadStringOptionalFlag("Tag"))

		resp, err := client.Benchmarks.GetInventoryAPIV2ServicesSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV2ServicesSummaryServiceNameCmd = &cobra.Command{
	Use: "services-summary-service-name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV2ServicesSummaryServiceNameParams()

		req.SetConnector(flags.ReadStringOptionalFlag("Connector"))
		req.SetConnectorID(flags.ReadStringOptionalFlag("ConnectorID"))
		req.SetEndTime(flags.ReadStringFlag("EndTime"))
		req.SetServiceName(flags.ReadStringFlag("ServiceName"))
		req.SetStartTime(flags.ReadStringFlag("StartTime"))

		resp, err := client.Benchmarks.GetInventoryAPIV2ServicesSummaryServiceName(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV1AccountsResourceCountCmd = &cobra.Command{
	Use: "accounts-resource-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_accounts_resource_count] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV1AccountsResourceCountParams()

		req.SetProvider(flags.ReadStringFlag("Provider"))
		req.SetSourceID(flags.ReadStringArrayFlag("SourceID"))

		resp, err := client.Benchmarks.GetInventoryAPIV1AccountsResourceCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_accounts_resource_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_accounts_resource_count] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV1ResourcesDistributionCmd = &cobra.Command{
	Use: "resources-distribution",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_distribution] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV1ResourcesDistributionParams()

		req.SetConnectionID(flags.ReadStringArrayFlag("ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag("Connector"))
		req.SetTimeWindow(flags.ReadStringFlag("TimeWindow"))

		resp, err := client.Benchmarks.GetInventoryAPIV1ResourcesDistribution(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_distribution] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_distribution] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV1ResourcesTopGrowingAccountsCmd = &cobra.Command{
	Use: "resources-top-growing-accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_growing_accounts] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV1ResourcesTopGrowingAccountsParams()

		req.SetCount(flags.ReadInt64Flag("Count"))
		req.SetProvider(flags.ReadStringFlag("Provider"))
		req.SetTimeWindow(flags.ReadStringFlag("TimeWindow"))

		resp, err := client.Benchmarks.GetInventoryAPIV1ResourcesTopGrowingAccounts(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_growing_accounts] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_growing_accounts] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV1ServicesDistributionCmd = &cobra.Command{
	Use: "services-distribution",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_services_distribution] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV1ServicesDistributionParams()

		req.SetProvider(flags.ReadStringFlag("Provider"))
		req.SetSourceID(flags.ReadStringArrayFlag("SourceID"))

		resp, err := client.Benchmarks.GetInventoryAPIV1ServicesDistribution(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_services_distribution] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_services_distribution] : %v", err)
		}

		return nil
	},
}
