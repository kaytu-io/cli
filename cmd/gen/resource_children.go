package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/spf13/cobra"
	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
)

var PostInventoryApiV1ResourceCmd = &cobra.Command{
	Use: "resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

        req := resource.NewPostInventoryAPIV1ResourceParams()

        req.SetRequest(&models.GitlabComKeibiengineKeibiEnginePkgInventoryAPIGetResourceRequest{
ID: flags.ReadStringOptionalFlag(cmd, "ID"),
ResourceType: flags.ReadStringOptionalFlag(cmd, "ResourceType"),

})


		resp, err := client.Resource.PostInventoryAPIV1Resource(req, auth)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		return nil
	},
}
