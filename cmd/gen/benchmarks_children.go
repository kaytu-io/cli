package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/benchmarks"
	"github.com/spf13/cobra"
)
var GetInventoryApiV1AccountsResourceCountCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_accounts_resource_count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_accounts_resource_count] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV1AccountsResourceCount(benchmarks.NewGetInventoryAPIV1AccountsResourceCountParams(), auth)
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
	Use: "get_inventory_api_v_1_resources_distribution",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_distribution] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV1ResourcesDistribution(benchmarks.NewGetInventoryAPIV1ResourcesDistributionParams(), auth)
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
	Use: "get_inventory_api_v_1_resources_top_growing_accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_resources_top_growing_accounts] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV1ResourcesTopGrowingAccounts(benchmarks.NewGetInventoryAPIV1ResourcesTopGrowingAccountsParams(), auth)
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
	Use: "get_inventory_api_v_1_services_distribution",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_services_distribution] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV1ServicesDistribution(benchmarks.NewGetInventoryAPIV1ServicesDistributionParams(), auth)
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
	Use: "get_inventory_api_v_2_services_summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV2ServicesSummary(benchmarks.NewGetInventoryAPIV2ServicesSummaryParams(), auth)
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
	Use: "get_inventory_api_v_2_services_summary_service_name",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		resp, err := client.Benchmarks.GetInventoryAPIV2ServicesSummaryServiceName(benchmarks.NewGetInventoryAPIV2ServicesSummaryServiceNameParams(), auth)
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
