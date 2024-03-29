// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package list

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/keys"
	"github.com/spf13/cobra"
)

var ApikeysCmd = &cobra.Command{
	Use: "apikeys",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := keys.NewGetAuthAPIV1KeysParams()

		resp, err := client.Keys.GetAuthAPIV1Keys(req, auth)
		if err != nil {
			return fmt.Errorf("[GetAuthAPIV1Keys] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetAuthAPIV1Keys", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetAuthAPIV1Keys] : %v", err)
		}

		return nil

	},
}

func init() {

}
