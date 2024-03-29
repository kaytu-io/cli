// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package get

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/get/workspace"
)

var WorkspaceCmd = &cobra.Command{
	Use: "workspace",
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("invalid flags")
		}
		return cmd.Help()

	},
}

func init() {

	WorkspaceCmd.AddCommand(workspace.UsersCmd)

	WorkspaceCmd.AddCommand(workspace.BootstrapStatusCmd)

}
