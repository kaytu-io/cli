// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package spend

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/analytics"
	"github.com/spf13/cobra"
)

var CompositionCmd = &cobra.Command{
	Use: "composition",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := analytics.NewGetInventoryAPIV2AnalyticsSpendCompositionParams()

		req.SetConnectionGroup(flags.ReadStringArrayFlag(cmd, "ConnectionGroup"))

		req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))

		req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))

		req.SetEndTime(flags.ReadInt64OptionalFlag(cmd, "EndTime"))

		req.SetStartTime(flags.ReadInt64OptionalFlag(cmd, "StartTime"))

		req.SetTop(flags.ReadInt64OptionalFlag(cmd, "Top"))

		resp, err := client.Analytics.GetInventoryAPIV2AnalyticsSpendComposition(req, auth)
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsSpendComposition] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetInventoryAPIV2AnalyticsSpendComposition", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2AnalyticsSpendComposition] : %v", err)
		}

		return nil

	},
}

func init() {

	CompositionCmd.Flags().StringArray("connection-group", nil, "Connection group to filter by - mutually exclusive with connectionId")

	CompositionCmd.Flags().StringArray("connection-id", nil, "Connection IDs to filter by - mutually exclusive with connectionGroup")

	CompositionCmd.Flags().StringArray("connector", nil, "Connector type to filter by")

	CompositionCmd.Flags().Int64("end-time", 0, "timestamp for end in epoch seconds")

	CompositionCmd.Flags().Int64("start-time", 0, "timestamp for start in epoch seconds")

	CompositionCmd.Flags().Int64("top", 0, "How many top values to return default is 5")

}