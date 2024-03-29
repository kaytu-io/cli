// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package user

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/users"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var AccessCmd = &cobra.Command{
	Use: "access",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := users.NewPutAuthAPIV1UserRoleBindingParams()

		var config *models.GithubComKaytuIoKaytuEnginePkgAuthAPIPutRoleBindingRequest
		flags.ReadStructFlag(cmd, "Request", &config)
		req.SetRequest(config)

		resp, err := client.Users.PutAuthAPIV1UserRoleBinding(req, auth)
		if err != nil {
			return fmt.Errorf("[PutAuthAPIV1UserRoleBinding] : %v", err)
		}

		_ = resp

		return nil

	},
}

func init() {

	AccessCmd.Flags().String("request", "", "Request Body")

	AccessCmd.MarkFlagRequired("request")

}
