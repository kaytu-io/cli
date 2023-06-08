package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/cost"
	"github.com/spf13/cobra"
)
var GetInventoryApiV1CostTopAccountsCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_cost_top_accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_accounts] : %v", err)
		}

		resp, err := client.Cost.GetInventoryAPIV1CostTopAccounts(cost.NewGetInventoryAPIV1CostTopAccountsParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_accounts] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_accounts] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV1CostTopServicesCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_cost_top_services",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_services] : %v", err)
		}

		resp, err := client.Cost.GetInventoryAPIV1CostTopServices(cost.NewGetInventoryAPIV1CostTopServicesParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_services] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_services] : %v", err)
		}

		return nil
	},
}

