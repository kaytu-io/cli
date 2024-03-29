// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package kaytu

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/update"
)

var UpdateCmd = &cobra.Command{
	Use: "update",
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("invalid flags")
		}
		return cmd.Help()

	},
}

func init() {

	UpdateCmd.AddCommand(update.AccountCmd)

	UpdateCmd.AddCommand(update.CredentialCmd)

	UpdateCmd.AddCommand(update.UserCmd)

	UpdateCmd.AddCommand(update.WorkspaceCmd)

}
