package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/cost"
	"github.com/spf13/cobra"
)

var GetInventoryApiV1CostTopServicesCmd = &cobra.Command{
	Use: "cost-top-services",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_services] : %v", err)
		}

		req := cost.NewGetInventoryAPIV1CostTopServicesParams()

		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))
		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))
		req.SetSourceID(flags.ReadStringFlag(cmd, "SourceID"))

		resp, err := client.Cost.GetInventoryAPIV1CostTopServices(req, auth)
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

var GetInventoryApiV1CostTopAccountsCmd = &cobra.Command{
	Use: "cost-top-accounts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_cost_top_accounts] : %v", err)
		}

		req := cost.NewGetInventoryAPIV1CostTopAccountsParams()

		req.SetCount(flags.ReadInt64Flag(cmd, "Count"))
		req.SetProvider(flags.ReadStringFlag(cmd, "Provider"))

		resp, err := client.Cost.GetInventoryAPIV1CostTopAccounts(req, auth)
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
