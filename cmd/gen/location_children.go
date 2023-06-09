package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/location"
	"github.com/spf13/cobra"
)

var GetInventoryApiV1LocationsConnectorCmd = &cobra.Command{
	Use: "connectorConnector",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_locations_connector] : %v", err)
		}

		resp, err := client.Location.GetInventoryAPIV1LocationsConnector(location.NewGetInventoryAPIV1LocationsConnectorParams(), auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_locations_connector] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_locations_connector] : %v", err)
		}

		return nil
	},
}
