// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
package get

import (
	"errors"

	"github.com/spf13/cobra"

	"github.com/kaytu-io/cli-program/cmd/genv2/kaytu/get/query"
)

var QueryCmd = &cobra.Command{
	Use: "query",
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().ParseErrorsWhitelist.UnknownFlags {
			return errors.New("invalid flags")
		}
		return cmd.Help()

	},
}

func init() {

	QueryCmd.AddCommand(query.HistoryCmd)

}