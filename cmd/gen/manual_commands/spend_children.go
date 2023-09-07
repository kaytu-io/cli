package manual_commands

import (
	"errors"
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/analytics"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/connections"
	"github.com/spf13/cobra"
)

var GetCostsCmd = &cobra.Command{
	Use:   "get-costs",
	Short: `Retrieving cost metrics with respect to specified filters. The API returns information such as the total cost and costs per each service based on the specified filters.`,
	Long:  `Retrieving cost metrics with respect to specified filters. The API returns information such as the total cost and costs per each service based on the specified filters.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			return fmt.Errorf("[get_costs] : %v", err)
		}
		view := flags.ReadStringFlag(cmd, "view")
		if view == "service" {
			return getByMetrics(cmd, client, auth)
		} else if view == "connection" {
			return getByConnections(cmd, client, auth)
		} else {
			return fmt.Errorf("[get_costs] : invalid view")
		}
	},
}

func getByMetrics(cmd *cobra.Command, client *client.KaytuServiceAPI, auth runtime.ClientAuthInfoWriter) error {
	req := analytics.NewGetInventoryAPIV2AnalyticsSpendMetricParams()

	req.SetConnectionGroup(flags.ReadStringOptionalFlag(cmd, "ConnectionGroup"))
	req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
	req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
	req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
	req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
	req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
	req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
	req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))

	resp, err := client.Analytics.GetInventoryAPIV2AnalyticsSpendMetric(req, auth)
	if err != nil {
		return fmt.Errorf("[get_costs] : %v", err)
	}

	err = pkg.PrintOutput(cmd, "get_costs", resp.GetPayload())
	if err != nil {
		return fmt.Errorf("[get_costs] : %v", err)
	}

	return nil
}

func getByConnections(cmd *cobra.Command, client *client.KaytuServiceAPI, auth runtime.ClientAuthInfoWriter) error {
	req := connections.NewGetOnboardAPIV1ConnectionsSummaryParams()

	req.SetConnectionID(flags.ReadStringArrayFlag(cmd, "ConnectionID"))
	req.SetConnector(flags.ReadStringArrayFlag(cmd, "Connector"))
	req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
	needResourceCount := false
	req.SetNeedResourceCount(&needResourceCount)
	req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
	req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
	req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
	req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))

	resp, err := client.Connections.GetOnboardAPIV1ConnectionsSummary(req, auth)
	if err != nil {
		return fmt.Errorf("[get_costs] : %v", err)
	}

	err = pkg.PrintOutput(cmd, "get_costs", resp.GetPayload())
	if err != nil {
		return fmt.Errorf("[get_costs] : %v", err)
	}

	return nil
}
