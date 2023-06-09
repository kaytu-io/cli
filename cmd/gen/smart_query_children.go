package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/smart_query"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var GetInventoryApiV1QueryCountCmd = &cobra.Command{
	Use: "query-count",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query_count] : %v", err)
		}

        req := smart_query.NewGetInventoryAPIV1QueryCountParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest{
Labels: flags.ReadStringArrayFlag("Labels"),
ProviderFilter: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISourceType(flags.ReadStringFlag("ProviderFilter")),
TitleFilter: flags.ReadStringFlag("TitleFilter"),

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
	Use: "query",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[get_inventory_api_v_1_query] : %v", err)
		}

        req := smart_query.NewGetInventoryAPIV1QueryParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIListQueryRequest{
Labels: flags.ReadStringArrayFlag("Labels"),
ProviderFilter: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISourceType(flags.ReadStringFlag("ProviderFilter")),
TitleFilter: flags.ReadStringFlag("TitleFilter"),

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
	Use: "query-query-id",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_query_query_id] : %v", err)
		}

        req := smart_query.NewPostInventoryAPIV1QueryQueryIDParams()

        req.SetAccept(flags.ReadStringFlag("Accept"))
req.SetQueryID(flags.ReadStringFlag("QueryID"))
req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIRunQueryRequest{
Page: &models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIPage{
No: flags.ReadInt64Flag("No"),
Size: flags.ReadInt64Flag("Size"),

},
Sorts: []*models.GitlabComKeibiengineKeibiEnginePkgInventoryAPISmartQuerySortItem{
{
Direction: models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIDirectionType(flags.ReadStringFlag("Direction")),
Field: flags.ReadStringFlag("Field"),

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
