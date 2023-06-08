package cmd

import (
	"fmt"
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu"
	"github.com/kaytu-io/cli-program/pkg/api/kaytu/client/workspace"
	"github.com/spf13/cobra"
)

// workspacesCmd represents the workspaces command
var workspacesCmd = &cobra.Command{
	Use: "workspaces",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, opt, err := kaytu.GetKaytuClient(cmd)
		if err != nil {
			return fmt.Errorf("[workspaces] : %v", err)
		}

		resp, err := client.Workspace.GetWorkspaceAPIV1Workspaces(workspace.NewGetWorkspaceAPIV1WorkspacesParams(), opt)
		if err != nil {
			return fmt.Errorf("[workspaces] : %v", err)
		}

		err = pkg.PrintOutputForTypeArray(cmd, resp.GetPayload())
		if err != nil {
			return fmt.Errorf("[workspaces] : %v", err)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(workspacesCmd)
}
