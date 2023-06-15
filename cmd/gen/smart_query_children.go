package gen

import (
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var GetInventoryApiV1QueryCountCmd = &cobra.Command{
	Use: "get-queries-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		req := smart_query.NewGetInventoryAPIV1QueryCountParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest{
			Labels:         flags.ReadStringArrayFlag(cmd, "Labels"),
			ProviderFilter: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISourceType(flags.ReadStringFlag(cmd, "ProviderFilter")),
			TitleFilter:    flags.ReadStringFlag(cmd, "TitleFilter"),
		})

		resp, err := client.SmartQuery.GetInventoryAPIV1QueryCount(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

		return nil
	},
}

var GetInventoryApiV1QueryCmd = &cobra.Command{
	Use: "get-queries",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		req := smart_query.NewGetInventoryAPIV1QueryParams()

		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest{
			Labels:         flags.ReadStringArrayFlag(cmd, "Labels"),
			ProviderFilter: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISourceType(flags.ReadStringFlag(cmd, "ProviderFilter")),
			TitleFilter:    flags.ReadStringFlag(cmd, "TitleFilter"),
		})

		resp, err := client.SmartQuery.GetInventoryAPIV1Query(req, auth)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

		return nil
	},
}

var PostInventoryApiV1QueryQueryIdCmd = &cobra.Command{
	Use: "get-query",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		req := smart_query.NewPostInventoryAPIV1QueryQueryIDParams()

		req.SetAccept(flags.ReadStringFlag(cmd, "Accept"))
		req.SetQueryID(flags.ReadStringFlag(cmd, "QueryID"))
		req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest{
			Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
				No:   flags.ReadInt64Flag(cmd, "No"),
				Size: flags.ReadInt64Flag(cmd, "Size"),
			},
			Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem{
				{
					Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag(cmd, "Direction")),
					Field:     flags.ReadStringFlag(cmd, "Field"),
				},
			},
		})

		resp, err := client.SmartQuery.PostInventoryAPIV1QueryQueryID(req, auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

		return nil
	},
}
