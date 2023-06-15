package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2ServicesSummaryServiceNameCmd = &cobra.Command{
	Use: "service-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV2ServicesSummaryServiceNameParams()

		req.SetConnector(flags.ReadStringOptionalFlag(cmd, "Connector"))
		req.SetConnectorID(flags.ReadStringOptionalFlag(cmd, "ConnectorID"))
		req.SetEndTime(flags.ReadStringFlag(cmd, "EndTime"))
		req.SetServiceName(flags.ReadStringFlag(cmd, "ServiceName"))

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

		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))
		req.SetSourceID(flags.ReadStringArrayFlag(cmd, "SourceID"))

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

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetTimeWindow(flags.ReadStringFlag(cmd, "TimeWindow"))

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
	Use: "top-growing-accounts-by-resources",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_growing_accounts] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV1ResourcesTopGrowingAccountsParams()

		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))
		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))
		req.SetTimeWindow(flags.ReadStringFlag(cmd, "TimeWindow"))

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

		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))
		req.SetSourceID(flags.ReadStringArrayFlag(cmd, "SourceID"))

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

var GetInventoryApiV2ServicesSummaryCmd = &cobra.Command{
	Use: "services-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		req := benchmarks.NewGetInventoryAPIV2ServicesSummaryParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

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
