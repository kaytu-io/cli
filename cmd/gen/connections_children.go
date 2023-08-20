// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package gen

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/spf13/cobra"
)

var GetOnboardApiV1ConnectionsSummaryCmd = &cobra.Command{
	Use:   "connections-summary",
	Short: `Retrieving a list of connections summaries`,
	Long:  `Retrieving a list of connections summaries`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		req := connections.NewGetOnboardAPIV1ConnectionsSummaryParams()

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))
		req.SetHealthState(flags.ReadStringOptionalFlag(cmd, "HealthState"))
		req.SetLifecycleState(flags.ReadStringOptionalFlag(cmd, "LifecycleState"))
		req.SetNeedCost(flags.ReadBooleanOptionalFlag(cmd, "NeedCost"))
		req.SetNeedResourceCount(flags.ReadBooleanOptionalFlag(cmd, "NeedResourceCount"))
		req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
		req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
		req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummary(req, auth)
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "get-onboard-api-v1-connections-summary", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_onboard_api_v_1_connections_summary] : %v", err)
		}

		return nil
	},
}
