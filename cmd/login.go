package cmd

import (
	"fmt"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/auth0"
	apiclient "github.com/kaytu-io/cli-program/pkg/api/kaytu/client"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"time"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use: "login",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceCode, err := auth0.RequestDeviceCode()
		if err != nil {
			return fmt.Errorf("[login] : %v", err)
		}

		var accessToken string
		for {
			accessToken, err = auth0.AccessToken(deviceCode)
			if err != nil {
				time.Sleep(pkg.TimeSleep * time.Second)
				continue
			}
			break
		}

		workspaceName := cmd.Flags().Lookup("workspace-name").Value.String()
		if workspaceName == "" {
			client := apiclient.New(httptransport.New("app.kaytu.dev", "/keibi", []string{"https"}), strfmt.Default)
			bearerTokenAuth := httptransport.BearerToken(accessToken)
			opt := func(r *runtime.ClientOperation) {
				r.AuthInfo = bearerTokenAuth
			}
			resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), opt)
			if err != nil {
				return fmt.Errorf("[workspaces] : %v", err)
			}
			response := resp.GetPayload()

			var items []string
			for _, r := range response {
				items = append(items, r.Name)
			}
			prompt := promptui.Select{
				Label: "Select default workspace",
				Items: items,
			}
			_, result, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("[workspaces] : %v", err)
			}

			workspaceName = result
		}

		err = pkg.SetConfig(pkg.Config{
			AccessToken:      accessToken,
			DefaultWorkspace: workspaceName,
		})
		if err != nil {
			return fmt.Errorf("[login] : %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
