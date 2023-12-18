// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package list

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/onboard"
	"github.com/spf13/cobra"
)

var ConnectorsCmd = &cobra.Command{
	Use: "connectors",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := onboard.NewGetOnboardAPIV1ConnectorParams()

		resp, err := client.Onboard.GetOnboardAPIV1Connector(req, auth)
		if err != nil {
			return fmt.Errorf("[GetOnboardAPIV1Connector] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "GetOnboardAPIV1Connector", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[GetOnboardAPIV1Connector] : %v", err)
		}

		return nil

	},
}

func init() {

}