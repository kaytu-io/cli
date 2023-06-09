package gen

import (
	"github.com/spf13/cobra"
)

var GetUsersCmd = &cobra.Command{
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var CreateUsersCmd = &cobra.Command{
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var DeleteUsersCmd = &cobra.Command{
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var UpdateUsersCmd = &cobra.Command{
	Use: "users",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}
func init() {
		GetUsersCmd.AddCommand(GetAuthApiV1WorkspaceRoleBindingsCmd)

		GetUsersCmd.AddCommand(DeleteAuthApiV1UserRoleBindingCmd)

		GetUsersCmd.AddCommand(GetAuthApiV1UserRoleBindingsCmd)

		GetUsersCmd.AddCommand(GetAuthApiV1UserUserIdCmd)

		GetUsersCmd.AddCommand(GetAuthApiV1UsersCmd)

		GetUsersCmd.AddCommand(DeleteAuthApiV1UserInviteCmd)

		GetUsersCmd.AddCommand(GetAuthApiV1UserUserIdWorkspaceMembershipCmd)

		GetUsersCmd.AddCommand(PostAuthApiV1UserInviteCmd)

		GetUsersCmd.AddCommand(PutAuthApiV1UserRoleBindingCmd)
}