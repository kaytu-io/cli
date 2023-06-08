package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/spf13/cobra"
)
var GetInventoryApiV1QueryCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_query",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		resp, err := client.SmartQuery.GetInventoryAPIV1Query(smart_query.NewGetInventoryAPIV1QueryParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		return nil
	},
}
var PostInventoryApiV1QueryQueryIdCmd = &cobra.Command{
	Use: "post_inventory_api_v_1_query_query_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		resp, err := client.SmartQuery.PostInventoryAPIV1QueryQueryID(smart_query.NewPostInventoryAPIV1QueryQueryIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV1QueryCountCmd = &cobra.Command{
	Use: "get_inventory_api_v_1_query_count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		resp, err := client.SmartQuery.GetInventoryAPIV1QueryCount(smart_query.NewGetInventoryAPIV1QueryCountParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		return nil
	},
}

