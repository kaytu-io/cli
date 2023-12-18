// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package list

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/list/accounts"
)

var AccountsCmd = &cobra.Command{
	Use: "accounts",
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("invalid flags")
		}
		return cmd.Help()

	},
}

func init() {

	AccountsCmd.AddCommand(accounts.SummaryCmd)

}