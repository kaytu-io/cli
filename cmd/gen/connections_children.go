package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/spf13/cobra"
)

var GetOnboardApiV1ConnectionsSummaryConnectionIdCmd = &cobra.Command{
	Use: "summarySummaryConnectionId",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary_connection_id] : %v", err)
		}

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummaryConnectionID(connections.NewGetOnboardAPIV1ConnectionsSummaryConnectionIDParams(), auth)
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
	Use: "summarySummary",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummary(connections.NewGetOnboardAPIV1ConnectionsSummaryParams(), auth)
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
