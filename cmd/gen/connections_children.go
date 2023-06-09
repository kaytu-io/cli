package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/spf13/cobra"
)

var GetOnboardApiV1ConnectionsSummaryConnectionIdCmd = &cobra.Command{
	Use: "connections-summary-connection-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary_connection_id] : %v", err)
		}

		req := connections.NewGetOnboardAPIV1ConnectionsSummaryConnectionIDParams()

		req.SetConnectionID(flags.ReadStringFlag(cmd, "ConnectionID"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummaryConnectionID(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary_connection_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary_connection_id] : %v", err)
		}

		return nil
	},
}

var GetOnboardApiV1ConnectionsSummaryCmd = &cobra.Command{
	Use: "connections-summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		req := connections.NewGetOnboardAPIV1ConnectionsSummaryParams()

		req.SetConnectionID(flags.ReadStringOptionalFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetHealthState(flags.ReadStringOptionalFlag(cmd, "HealthState"))
		req.SetLifecycleState(flags.ReadStringOptionalFlag(cmd, "LifecycleState"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		return nil
	},
}
