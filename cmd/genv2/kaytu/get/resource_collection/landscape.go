// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package resource_collection

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/resource_collection"
	"github.com/spf13/cobra"
)

var LandscapeCmd = &cobra.Command{
	Use: "landscape",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := resource_collection.NewGetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscapeParams()

		req.SetResourceCollectionID(flags.ReadStringFlag(cmd, "ResourceCollectionID"))

		resp, err := client.ResourceCollection.GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape(req, auth)
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetInventoryAPIV2ResourceCollectionResourceCollectionIDLandscape] : %v", err)
		}

		return nil

	},
}

func init() {

	LandscapeCmd.Flags().String("resource-collection-id", "", "Resource collection ID")

	LandscapeCmd.MarkFlagRequired("resource-collection-id")

}
