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
	GetUsersCmd.AddCommand(GetAuthApiV1UsersCmd)
	GetAuthApiV1UsersCmd.Flags().String("email", "", "")
	GetAuthApiV1UsersCmd.Flags().Bool("email-verified", false, "")
	GetAuthApiV1UsersCmd.Flags().String("role-name", "", "")

	DeleteUsersCmd.AddCommand(DeleteAuthApiV1UserRoleBindingCmd)
	DeleteAuthApiV1UserRoleBindingCmd.Flags().String("user-id", "", "")

	GetUsersCmd.AddCommand(GetAuthApiV1UserRoleBindingsCmd)

	GetUsersCmd.AddCommand(GetAuthApiV1UserUserIdCmd)
	GetAuthApiV1UserUserIdCmd.Flags().String("user-id", "", "")

	GetUsersCmd.AddCommand(GetAuthApiV1UserUserIdWorkspaceMembershipCmd)
	GetAuthApiV1UserUserIdWorkspaceMembershipCmd.Flags().String("user-id", "", "")

	DeleteUsersCmd.AddCommand(DeleteAuthApiV1UserInviteCmd)
	DeleteAuthApiV1UserInviteCmd.Flags().String("user-id", "", "")

	GetUsersCmd.AddCommand(GetAuthApiV1WorkspaceRoleBindingsCmd)

	GetUsersCmd.AddCommand(PostAuthApiV1UserInviteCmd)
	PostAuthApiV1UserInviteCmd.Flags().String("email", "", "")
	PostAuthApiV1UserInviteCmd.Flags().String("role-name", "", "")

	GetUsersCmd.AddCommand(PutAuthApiV1UserRoleBindingCmd)
	PutAuthApiV1UserRoleBindingCmd.Flags().String("role-name", "", "")
	PutAuthApiV1UserRoleBindingCmd.Flags().String("user-id", "", "")

}
