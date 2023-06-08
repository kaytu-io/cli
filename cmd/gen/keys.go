package gen

import (
	"github.com/spf13/cobra"
)

var GetKeysCmd = &cobra.Command{
	Use: "keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateKeysCmd = &cobra.Command{
	Use: "keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteKeysCmd = &cobra.Command{
	Use: "keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateKeysCmd = &cobra.Command{
	Use: "keys",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetKeysCmd.AddCommand(GetAuthApiV1KeysCmd)

		GetKeysCmd.AddCommand(PostAuthApiV1KeyCreateCmd)

		GetKeysCmd.AddCommand(PostAuthApiV1KeyIdActivateCmd)

		GetKeysCmd.AddCommand(PostAuthApiV1KeyIdSuspendCmd)

		GetKeysCmd.AddCommand(PostAuthApiV1KeyRoleCmd)

		GetKeysCmd.AddCommand(DeleteAuthApiV1KeyIdDeleteCmd)

		GetKeysCmd.AddCommand(GetAuthApiV1KeyIdCmd)
}