package gen

import (
	"github.com/spf13/cobra"
)

var GetRolesCmd = &cobra.Command{
	Use: "roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateRolesCmd = &cobra.Command{
	Use: "roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteRolesCmd = &cobra.Command{
	Use: "roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateRolesCmd = &cobra.Command{
	Use: "roles",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetRolesCmd.AddCommand(GetAuthApiV1RolesRoleNameCmd)

		GetRolesCmd.AddCommand(GetAuthApiV1RoleRoleNameKeysCmd)

		GetRolesCmd.AddCommand(GetAuthApiV1RoleRoleNameUsersCmd)

		GetRolesCmd.AddCommand(GetAuthApiV1RolesCmd)
}