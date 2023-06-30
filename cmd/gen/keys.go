// ========== DO NOT EDIT THIS FILE! AUTO GENERATED!!! =========
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

	UpdateKeysCmd.AddCommand(PostAuthApiV1KeyRoleCmd)
	PostAuthApiV1KeyRoleCmd.Flags().Int64("id", 0, "")
	PostAuthApiV1KeyRoleCmd.MarkFlagRequired("id")
	PostAuthApiV1KeyRoleCmd.Flags().String("role-name", "", "")

	DeleteKeysCmd.AddCommand(DeleteAuthApiV1KeyIdDeleteCmd)
	DeleteAuthApiV1KeyIdDeleteCmd.Flags().String("id", "", "")
	DeleteAuthApiV1KeyIdDeleteCmd.MarkFlagRequired("id")

	GetKeysCmd.AddCommand(GetAuthApiV1KeyIdCmd)
	GetAuthApiV1KeyIdCmd.Flags().String("id", "", "")
	GetAuthApiV1KeyIdCmd.MarkFlagRequired("id")

	GetKeysCmd.AddCommand(GetAuthApiV1KeysCmd)

	CreateKeysCmd.AddCommand(PostAuthApiV1KeyCreateCmd)
	PostAuthApiV1KeyCreateCmd.Flags().String("name", "", "")
	PostAuthApiV1KeyCreateCmd.MarkFlagRequired("name")
	PostAuthApiV1KeyCreateCmd.Flags().String("role-name", "", "")

	GetKeysCmd.AddCommand(PostAuthApiV1KeyIdActivateCmd)
	PostAuthApiV1KeyIdActivateCmd.Flags().String("id", "", "")
	PostAuthApiV1KeyIdActivateCmd.MarkFlagRequired("id")

	GetKeysCmd.AddCommand(PostAuthApiV1KeyIdSuspendCmd)
	PostAuthApiV1KeyIdSuspendCmd.Flags().String("id", "", "")
	PostAuthApiV1KeyIdSuspendCmd.MarkFlagRequired("id")

}
