package predef

import (
	"github.com/kaytu-io/cli-program/pkg"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use: "logout",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := pkg.RemoveConfig()
		if err != nil {
			return err
		}
		return nil
	},
}
