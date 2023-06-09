package gen

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource"
	"github.com/spf13/cobra"
)

var PostInventoryApiV1ResourceCmd = &cobra.Command{
	Use: "Resource",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, auth, err := kaytu.GetKaytuAuthClient(cmd)
		if err != nil {
			return fmt.Errorf("[post_inventory_api_v_1_resource] : %v", err)
		}

		resp, err := client.Resource.PostInventoryAPIV1Resource(resource.NewPostInventoryAPIV1ResourceParams(), auth)
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
