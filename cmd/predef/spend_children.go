package predef

import (
	"errors"
	"fmt"
	"time"

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
	Short: `Retrieving costs by the give filters grouped by services or connections.`,
	Long:  `Retrieving costs by the give filters grouped by services or connections.`,
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

func init() {
	GetCostsCmd.Flags().String("start-time", fmt.Sprintf("%v", time.Now().AddDate(0, -1, 0).Format("2006-01-02")), "timestamp for start time")
	GetCostsCmd.Flags().String("end-time", fmt.Sprintf("%v", time.Now().AddDate(0, 0, -2).Format("2006-01-02")), "timestamp for end time")
	GetCostsCmd.Flags().Int64("page-number", 0, "page number - (Default 1)")
	GetCostsCmd.Flags().Int64("page-size", 0, "page size - (Default 20)")
	GetCostsCmd.Flags().String("sort-by", "", "Sort by field - (Default cost)")
	GetCostsCmd.Flags().String("view", "connection", "Whether to view by connection or service")
	GetCostsCmd.Flags().String("filter", "", "Filter in JSON format (Case Sensitive). Filters by matches:\n"+
		"\t- Match: JSON Format. Specifies a filter rule:\n"+
		"\t\t- Key: String. Specifies the filter dimension key. Valid values are [ConnectionID|Provider|ConnectionGroup|ConnectionName]\n"+
		"\t\t- Value: String Array. Specifies the filter value.\n"+
		"\t\t- MatchOption: String. Specifies the match option. Valid values are [EQUAL|STARTS_WITH|ENDS_WITH|CONTAINS]. (Default EQUAL).\n"+
		"\t\t              NOT the match by ~ character: [~EQUAL|~STARTS_WITH|~ENDS_WITH|~CONTAINS]\n"+
		"\t- AND: Array. Return results that match all Match objects.\n"+
		"\t- OR: Array. Return results that match either Match object.\n"+
		"Example:\n"+
		"\t {\"AND\":[{\"Match\":{\"Key\":\"ConnectionName\",\"Values\":[\"Software\"],\"MatchOption\":\"CONTAINS\"}},{\"Match\":{\"Key\":\"ConnectionName\",\"Values\":[\"T\"],\"MatchOption\":\"~STARTS_WITH\"}}]}\n"+
		"Give file by @ or file:// prefix")
}

func getByMetrics(cmd *cobra.Command, client *client.KaytuServiceAPI, auth runtime.ClientAuthInfoWriter) error {
	req := analytics.NewGetInventoryAPIV2AnalyticsSpendMetricParams()

	req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
	req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
	req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
	req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
	req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))
	req.SetFilter(flags.ReadStringOptionalFlag(cmd, "Filter"))

	endTime := flags.ReadTimeOptionalFlag(cmd, "EndTime")
	if endTime != nil {
		t := time.Unix(*endTime, 0).UTC()
		if (t.Year() == time.Now().Year() && t.Month() == time.Now().Month()) && (t.Day() == time.Now().Day() || t.Day() == (time.Now().Day()-1)) {
			return fmt.Errorf("[get_costs] : No accurate data for last two days")
		}
	}

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

	req.SetEndTime(flags.ReadTimeOptionalFlag(cmd, "EndTime"))
	needResourceCount := false
	req.SetNeedResourceCount(&needResourceCount)
	req.SetPageNumber(flags.ReadInt64OptionalFlag(cmd, "PageNumber"))
	req.SetPageSize(flags.ReadInt64OptionalFlag(cmd, "PageSize"))
	req.SetSortBy(flags.ReadStringOptionalFlag(cmd, "SortBy"))
	req.SetStartTime(flags.ReadTimeOptionalFlag(cmd, "StartTime"))
	req.SetFilter(flags.ReadStringOptionalFlag(cmd, "Filter"))

	endTime := flags.ReadTimeOptionalFlag(cmd, "EndTime")
	if endTime != nil {
		t := time.Unix(*endTime, 0).UTC()
		if (t.Year() == time.Now().Year() && t.Month() == time.Now().Month()) && (t.Day() == time.Now().Day() || t.Day() == (time.Now().Day()-1)) {
			return fmt.Errorf("[get_costs] : No accurate data for last two days")
		}
	}

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
