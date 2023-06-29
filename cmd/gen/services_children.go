package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/services"
	"github.com/spf13/cobra"
)

var GetInventoryApiV2ServicesSummaryCmd = &cobra.Command{
	Use: "services-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary] : %v", err)
		}

		req := services.NewGetInventoryAPIV2ServicesSummaryParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadStringOptionalFlag(cmd, "EndTime"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetTag(flags.ReadStringArrayFlag(cmd, "Tag"))

		resp, err := client.Services.GetInventoryAPIV2ServicesSummary(req, auth)
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
	Use: "service-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_2_services_summary_service_name] : %v", err)
		}

		req := services.NewGetInventoryAPIV2ServicesSummaryServiceNameParams()

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetConnectorID(flags.ReadStringArrayFlag(cmd, "ConnectorID"))
		req.SetEndTime(flags.ReadStringFlag(cmd, "EndTime"))
		req.SetServiceName(flags.ReadStringFlag(cmd, "ServiceName"))

		resp, err := client.Services.GetInventoryAPIV2ServicesSummaryServiceName(req, auth)
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
