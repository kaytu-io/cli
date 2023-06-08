package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connection"
	"github.com/spf13/cobra"
)
var GetInventoryApiV2ConnectionsDataConnectionIdCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_connections_data_connection_id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data_connection_id] : %v", err)
		}

		resp, err := client.Connection.GetInventoryAPIV2ConnectionsDataConnectionID(connection.NewGetInventoryAPIV2ConnectionsDataConnectionIDParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data_connection_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data_connection_id] : %v", err)
		}

		return nil
	},
}
var GetInventoryApiV2ConnectionsDataCmd = &cobra.Command{
	Use: "get_inventory_api_v_2_connections_data",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data] : %v", err)
		}

		resp, err := client.Connection.GetInventoryAPIV2ConnectionsData(connection.NewGetInventoryAPIV2ConnectionsDataParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data] : %v", err)
		}

		return nil
	},
}
