package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connection"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2ConnectionsDataConnectionIdCmd = &cobra.Command{
	Use: "connections-data-connection-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data_connection_id] : %v", err)
		}

		req := connection.NewGetInventoryAPIV2ConnectionsDataConnectionIDParams()

		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connection.GetInventoryAPIV2ConnectionsDataConnectionID(req, auth)
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
	Use: "connections-data",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_connections_data] : %v", err)
		}

		req := connection.NewGetInventoryAPIV2ConnectionsDataParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connection.GetInventoryAPIV2ConnectionsData(req, auth)
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
