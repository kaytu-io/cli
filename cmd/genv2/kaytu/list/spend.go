// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package list

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/list/spend"
)

var SpendCmd = &cobra.Command{
	Use: "spend",
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("invalid flags")
		}
		return cmd.Help()

	},
}

func init() {

	SpendCmd.AddCommand(spend.CompositionCmd)

	SpendCmd.AddCommand(spend.MetricsCmd)

}
