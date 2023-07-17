package cmd

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/auth0"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"time"
)

const RetrySleep = 3
const DefaultWorkspace = "keibi"

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use: "login",
	RunE: func(cmd *cobra.Command, args []string) error {
		deviceCode, err := auth0.RequestDeviceCode()
		if err != nil {
			return fmt.Errorf("[login-deviceCode]: %v", err)
		}

		var accessToken string
		for i := 0; i < 100; i++ {
			accessToken, err = auth0.AccessToken(deviceCode)
			if err != nil {
				time.Sleep(RetrySleep * time.Second)
				continue
			}
			break
		}
		if err != nil {
			return fmt.Errorf("[login-accessToken]: %v", err)
		}

		workspaceName := cmd.Flags().Lookup("workspace-name").Value.String()
		if workspaceName == "" {
			client, auth := kaytu.GetKaytuAuthClientWithConfig(DefaultWorkspace, accessToken)
			resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), auth)
			if err != nil {
				return fmt.Errorf("[login-workspaces]: %v", err)
			}
			response := resp.GetPayload()

			var items []string
			for _, r := range response {
				items = append(items, r.Name)
			}
			fmt.Println("\n")
			prompt := promptui.Select{
				Label: "Please select the default workspace",
				Items: items,
			}
			_, result, err := prompt.Run()
			if err != nil {
				return fmt.Errorf("[login-defaultWS]: %v", err)
			}

			workspaceName = result
		}

		err = pkg.SetConfig(pkg.Config{
			AccessToken:      accessToken,
			DefaultWorkspace: workspaceName,
		})
		if err != nil {
			return fmt.Errorf("[login-setConfig]: %v", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
