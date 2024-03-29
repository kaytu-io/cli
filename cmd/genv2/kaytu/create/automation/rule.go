// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package automation

import (
	"errors"
	"fmt"

	"github.com/kaytu-io/cli-program/cmd/flags"
	"github.com/kaytu-io/cli-program/pkg"
	kaytuSDK "github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/alerting"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/models"
	"github.com/spf13/cobra"
)

var RuleCmd = &cobra.Command{
	Use: "rule",
	RunE: func(cmd *cobra.Command, args []string) error {

		client, auth, err := kaytuSDK.GetKaytuAuthClient(cmd)
		if err != nil {
			if errors.Is(err, pkg.ExpiredSession) {
				fmt.Println(err.Error())
				return nil
			}
			panic(err)
		}

		req := alerting.NewPostAlertingAPIV1RuleCreateParams()

		var config *models.GithubComKaytuIoKaytuEnginePkgAlertingAPICreateRuleRequest
		flags.ReadStructFlag(cmd, "Request", &config)
		req.SetRequest(config)

		resp, err := client.Alerting.PostAlertingAPIV1RuleCreate(req, auth)
		if err != nil {
			return fmt.Errorf("[PostAlertingAPIV1RuleCreate] : %v", err)
		}

		err = pkg.PrintOutput(cmd, "PostAlertingAPIV1RuleCreate", resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[PostAlertingAPIV1RuleCreate] : %v", err)
		}

		return nil

	},
}

func init() {

	RuleCmd.Flags().String("request", "", "Request Body")

	RuleCmd.MarkFlagRequired("request")

}
